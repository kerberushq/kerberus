package crdrequest

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
	apiextensionsapi "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	crdv1alpha1 "github.com/kerberushq/kerberus/pkg/apis/crd/v1alpha1"
	crdv1alpha1client "github.com/kerberushq/kerberus/pkg/generated/clientset/versioned/typed/crd/v1alpha1"
	"github.com/kerberushq/kerberus/pkg/util/sort"
)

// Add creates a new crdrequest Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(ctx context.Context, log *logrus.Entry, mgr manager.Manager) error {
	return add(ctx, log, mgr, newReconciler(log, mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(log *logrus.Entry, mgr manager.Manager) reconcile.Reconciler {
	apiExtClient, err := apiextensionsclientset.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatal("Failed to initialize apiextensions client with %s", err)
	}

	kubeClientSet, err := kubernetes.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatalf("Failed to initialize kubeclient client with %s", err)
	}

	crdClient, err := crdv1alpha1client.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatal("Failed to initialize crdv1alpha1client client with %s", err)
	}

	return &Reconcilecrdrequest{
		kubeClientSet: kubeClientSet,
		crdClient:     crdClient,
		apiExtClient:  apiExtClient,

		scheme: mgr.GetScheme(),
		log:    log,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(ctx context.Context, log *logrus.Entry, mgr manager.Manager, r reconcile.Reconciler) error {
	err := crdv1alpha1.AddToScheme(mgr.GetScheme())
	if err != nil {
		return err
	}

	c, err := controller.New("crdrequest-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &apiextensionsapi.CustomResourceDefinition{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &crdv1alpha1.CRDRequest{},
	})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource crdrequest and requeue the owner Namespace
	return c.Watch(&source.Kind{Type: &crdv1alpha1.CRDRequest{}}, &handler.EnqueueRequestForObject{})
}

// blank assignment to verify that Reconcilecrdrequest implements reconcile.Reconciler
var _ reconcile.Reconciler = &Reconcilecrdrequest{}

// Reconcilecrdrequest reconciles a crdrequest object
type Reconcilecrdrequest struct {
	kubeClientSet *kubernetes.Clientset
	crdClient     crdv1alpha1client.CrdV1alpha1Interface
	apiExtClient  apiextensionsclientset.Interface

	scheme *runtime.Scheme
	log    *logrus.Entry
}

// Reconcile reads that state of the cluster for a crdrequest object and makes changes based on the state read
// and what is in the crdrequest.Spec
func (r *Reconcilecrdrequest) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// HACK: Because we manage non-namespaced objects with namespaced objects
	// sometimes we don't get namespace in the reconcile.Request. This is due to fact
	// that owner field does not contain namespace of the owning object.
	// Because of this we will get all namespaces of all potentially duplicate
	// objects and just reconcile them all.
	// https://github.com/kubernetes-sigs/controller-runtime/issues/543
	reconcileList := []reconcile.Request{}
	if len(request.Namespace) == 0 {
		objects, err := r.crdClient.CRDRequests(metav1.NamespaceAll).List(metav1.ListOptions{})
		if err != nil {
			return reconcile.Result{Requeue: true}, err
		}
		for _, req := range objects.Items {
			reconcileList = append(reconcileList, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      req.Name,
					Namespace: req.Namespace,
				},
			})
		}

		for _, req := range reconcileList {
			return r.Reconcile(req)
		}
	} else {
		r.log = r.log.WithField("Request.Namespace", request.Namespace).WithField("Request.Name", request.Name)
		r.log.Info("Reconciling")
		return r.reconcoleCRDRequest(request)
	}
	return reconcile.Result{Requeue: false}, nil
}

func (r *Reconcilecrdrequest) reconcoleCRDRequest(request reconcile.Request) (reconcile.Result, error) {
	instance, err := r.crdClient.CRDRequests(request.Namespace).Get(request.Name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// reconcilers will delete objects
			return reconcile.Result{Requeue: false}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{Requeue: false}, err
	}

	if len(instance.Status.Conditions) > 0 && instance.Status.Conditions[0].State == crdv1alpha1.StateFailed {
		return reconcile.Result{Requeue: false}, err
	}

	admit, err := r.admitCRDRequest(instance)
	// do not admit
	if !admit {
		r.updateStatus(instance, false, crdv1alpha1.StateFailed, err.Error())
		return reconcile.Result{Requeue: false}, nil
	}
	if err != nil {
		r.updateStatus(instance, false, crdv1alpha1.StateFailed, err.Error())
		r.log.Errorf("error [%s] with admin status %t not allowed", err.Error(), admit)
		return reconcile.Result{Requeue: false}, nil
	}

	for _, crd := range instance.Spec.CRDList {
		r.log.Debugf("create %s", crd.GetName())
		err := r.createOrUpdateCRDRequest(instance)
		if err != nil {
			r.updateStatus(instance, false, crdv1alpha1.StateFailed, err.Error())
			r.log.Error(err)
			return reconcile.Result{Requeue: false}, err
		}
	}

	return reconcile.Result{Requeue: false}, nil
}

func (r *Reconcilecrdrequest) createOrUpdateCRDRequest(instance *crdv1alpha1.CRDRequest) error {
	for _, crd := range instance.Spec.CRDList {
		// if CRDRequest is owned we add owners annotation so when we clean
		// them if CRDRequest is deleted CRD will be too. Else we leave them around
		if instance.Spec.Owned {
			r.log.Debugf("crdrequest %s is owned", instance.Name)
			if err := controllerutil.SetControllerReference(instance, &crd, r.scheme); err != nil {
				return err
			}
		}
		existing, err := r.apiExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crd.GetName(), metav1.GetOptions{})
		// no errors - already exist so update
		if err == nil {
			// nil certain fields for diff check
			existing.Spec.Names.ListKind = ""
			existing.Spec.Versions = nil
			existing.Spec.AdditionalPrinterColumns = nil
			existing.Spec.Conversion = nil

			if diff.ObjectReflectDiff(existing.Spec, crd.Spec) == "<no diffs>" {
				continue
			}
			_, err = r.apiExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Update(&crd)
			if err != nil {
				return err
			}
		}
		_, err = r.apiExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&crd)
		if err != nil {
			r.log.Error(err)
			return err
		}
	}

	return nil
}

func (r *Reconcilecrdrequest) admitCRDRequest(instance *crdv1alpha1.CRDRequest) (bool, error) {
	var denyAll bool
	var allowAll bool

	configList, err := r.crdClient.CRDConfigs("").List(metav1.ListOptions{})
	if err != nil {
		return false, err
	}

	var allow []string
	var deny []string
	state := make(map[string]bool)
	for _, list := range configList.Items {
		for _, rule := range list.Spec.Allow {
			denyAll = list.Spec.DenyAll
			allow = append(allow, rule)
		}
		for _, rule := range list.Spec.Deny {
			allowAll = list.Spec.AllowAll
			deny = append(deny, rule)
		}
	}
	if (denyAll && allowAll) || (!denyAll && !allowAll) {
		allowAll = false
		denyAll = true
	}

	allow = sort.Unique(allow)
	deny = sort.Unique(deny)

	r.log.Debugf("allowAll %t, denyAll %t", allowAll, denyAll)
	r.log.Debugf("allow list %s", allow)
	r.log.Debugf("deny list %s", deny)
	for _, crd := range instance.Spec.CRDList {
		// if allowAll - confirm with deny list
		// state should be all false to admit

		if allowAll {
			for _, rx := range deny {
				r.log.Debugf("matching %s with %s", rx, crd.GetName())
				found, _ := regexp.MatchString(rx, crd.GetName())
				if found {
					state[crd.GetName()] = found
				}
			}
		}

		// of denyAll - confirm with allow list
		// state should be all true to admit
		if denyAll {
			r.log.Debug(crd.GetName())
			state[crd.GetName()] = false
			for _, rx := range allow {
				r.log.Debugf("matching %s with %s", rx, crd.GetName())
				found, _ := regexp.MatchString(rx, crd.GetName())
				if found {
					state[crd.GetName()] = found
				}
			}
		}
	}

	// check state and admit/deny
	if allowAll {
		for name, crd := range state {
			if crd {
				return false, fmt.Errorf("object %s was denied", name)
			}
		}
		return true, nil
	}

	if denyAll {
		for name, crd := range state {
			if !crd {
				return false, fmt.Errorf("object %s was not in allow list", name)
			}
		}
		return true, nil
	}

	return false, fmt.Errorf("error while matching rules")
}

func (r *Reconcilecrdrequest) updateStatus(instance *crdv1alpha1.CRDRequest, status bool, state crdv1alpha1.State, message string) error {
	if len(instance.Status.Conditions) > 5 {
		instance.Status.Conditions = instance.Status.Conditions[1:]
	}
	condition := crdv1alpha1.Condition{
		LastTransitionTime: metav1.NewTime(time.Now()),
		Status:             status,
		State:              state,
		Message:            message,
	}
	instance.Status.Conditions = append(instance.Status.Conditions, condition)
	_, err := r.crdClient.CRDRequests(instance.Namespace).UpdateStatus(instance)
	return err
}
