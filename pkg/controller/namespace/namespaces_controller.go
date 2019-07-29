package namespaces

import (
	"context"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	//apiextensionsapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	configv1alpha1 "github.com/kerberushq/kerberus/pkg/apis/config/v1alpha1"
	rbacv1alpha1 "github.com/kerberushq/kerberus/pkg/apis/rbac/v1alpha1"
	rbacv1alpha1client "github.com/kerberushq/kerberus/pkg/generated/rbac/clientset/versioned/typed/rbac/v1alpha1"
)

// Add creates a new NamespacePermission Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(ctx context.Context, log *logrus.Entry, config configv1alpha1.Config, mgr manager.Manager) error {
	return add(ctx, log, mgr, newReconciler(log, config, mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(log *logrus.Entry, config configv1alpha1.Config, mgr manager.Manager) reconcile.Reconciler {
	apiExtClient, err := apiextensionsclientset.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatal("Failed to initialize apiextensions client with %s", err)
	}

	kubeClientSet, err := kubernetes.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatalf("Failed to initialize kubeclient client with %s", err)
	}

	rbacClient, err := rbacv1alpha1client.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatal("Failed to initialize crdv1alpha1client client with %s", err)
	}

	return &ReconcileNamespace{
		kubeClientSet: kubeClientSet,
		rbacClient:    rbacClient,
		apiExtClient:  apiExtClient,

		config: config,
		scheme: mgr.GetScheme(),
		log:    log,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(ctx context.Context, log *logrus.Entry, mgr manager.Manager, r reconcile.Reconciler) error {
	err := rbacv1alpha1.AddToScheme(mgr.GetScheme())
	if err != nil {
		return err
	}

	c, err := controller.New("namespace-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource RoleBinding and requeue the owner Namespace
	err = c.Watch(&source.Kind{Type: &corev1.Namespace{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource RoleBinding and requeue the owner Namespace
	return c.Watch(&source.Kind{Type: &rbacv1.RoleBinding{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &corev1.Namespace{},
	})
}

// blank assignment to verify that ReconcileNamespace implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNamespace{}

// ReconcileNamespace reconciles a NamespacePermission object
type ReconcileNamespace struct {
	kubeClientSet *kubernetes.Clientset
	rbacClient    rbacv1alpha1client.RbacV1alpha1Interface
	apiExtClient  apiextensionsclientset.Interface

	config configv1alpha1.Config
	scheme *runtime.Scheme
	log    *logrus.Entry
}

// Reconcile reads that state of the cluster for a NamespacePermission object and makes changes based on the state read
// and what is in the NamespacePermission.Spec
func (r *ReconcileNamespace) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	r.log = r.log.WithField("Request.Namespace", request.Namespace).WithField("Request.Name", request.Name)
	r.log.Info("Reconciling Namespaces")

	instance, err := r.kubeClientSet.CoreV1().Namespaces().Get(request.Name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{Requeue: false}, err
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{Requeue: true}, err
	}

	// list all NamespacePermission objects and apply them into the namespaces
	// TODO: define only controller namespace
	namespacePermissions, err := r.rbacClient.NamespacePermissions(metav1.NamespaceAll).List(metav1.ListOptions{})
	if errors.IsNotFound(err) {
		return reconcile.Result{Requeue: false}, err
	} else if err != nil {
		return reconcile.Result{Requeue: true}, err
	}

	for _, namespacePermission := range namespacePermissions.Items {
		namespacePermissionCopy := namespacePermission.DeepCopy()
		// requst.Name is the namespace name
		roleBinding := newRoleBinding(namespacePermissionCopy, request.Name)
		if err := controllerutil.SetControllerReference(instance, roleBinding, r.scheme); err != nil {
			return reconcile.Result{Requeue: true}, err
		}
		err := r.createOrUpdateRolebinding(roleBinding)
		if err != nil {
			return reconcile.Result{Requeue: true}, err
		}
	}

	return reconcile.Result{Requeue: false}, nil
}

// TODO: Reuse code. this is duplicate
func (r *ReconcileNamespace) createOrUpdateRolebinding(rolebinding *rbacv1.RoleBinding) error {
	existing, err := r.kubeClientSet.RbacV1().RoleBindings(rolebinding.Namespace).Get(rolebinding.Name, metav1.GetOptions{})
	if err == nil {
		// compare key fields to know if we need update
		if diff.ObjectReflectDiff(existing.Subjects, rolebinding.Subjects) == "<no diffs>" &&
			diff.ObjectReflectDiff(existing.RoleRef, rolebinding.RoleRef) == "<no diffs>" {
			return nil
		}

		_, err = r.kubeClientSet.RbacV1().RoleBindings(rolebinding.Namespace).Update(rolebinding)
		if err != nil {
			// update does not work in all cases due to immutable fields
			err := r.kubeClientSet.RbacV1().RoleBindings(rolebinding.Namespace).Delete(rolebinding.Name, &metav1.DeleteOptions{})
			if err != nil {
				return err
			}
			_, err = r.kubeClientSet.RbacV1().RoleBindings(rolebinding.Namespace).Create(rolebinding)
			if err != nil {
				return err
			}
		}
		return nil
	}
	_, err = r.kubeClientSet.RbacV1().RoleBindings(rolebinding.Namespace).Create(rolebinding)
	if err != nil {
		r.log.Error(err.Error())
		return err
	}
	return nil
}

// newRoleBinding returns a namespace level role-binding
func newRoleBinding(cr *rbacv1alpha1.NamespacePermission, namespace string) *rbacv1.RoleBinding {
	labels := map[string]string{
		"config": cr.Name,
	}
	return &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: namespace,
			Labels:    labels,
		},
		Subjects: cr.Spec.Subjects,
		RoleRef:  cr.Spec.RoleRef,
	}
}
