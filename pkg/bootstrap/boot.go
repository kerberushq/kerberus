package bootstrap

//go:generate go get github.com/go-bindata/go-bindata/go-bindata
//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix data data/...
//go:generate gofmt -s -l -w bindata.go

import (
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// boostrap created required CRD objects onto the cluster
type boostrap struct {
	kubeClientSet *kubernetes.Clientset
	apiExtClient  apiextensionsclientset.Interface

	log *logrus.Entry
}

// newReconciler returns a new reconcile.Reconciler
func New(log *logrus.Entry, mgr manager.Manager) (*boostrap, error) {

	apiExtClient, err := apiextensionsclientset.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatalf("Failed to initialize apiextensions client with %s", err)
		return nil, err
	}

	kubeClientSet, err := kubernetes.NewForConfig(mgr.GetConfig())
	if err != nil {
		log.Fatalf("Failed to initialize kubeclient client with %s", err)
		return nil, err
	}

	return &boostrap{
		kubeClientSet: kubeClientSet,
		apiExtClient:  apiExtClient,

		log: log,
	}, nil
}

func (b *boostrap) Initialize() error {
	if os.Getenv("BOOTSTRAP") == strings.ToLower("true") {
		b.log.Info("auto-boostrap enabled")

		for _, file := range []string{
			"crds/crd_v1alpha1_crdconfig_crd.yaml",
			"crds/crd_v1alpha1_crdrequest_crd.yaml",
		} {
			var crdConfig apiextensionsv1beta1.CustomResourceDefinition
			data, err := Asset(file)
			if err != nil {
				return err
			}

			err = yaml.Unmarshal(data, &crdConfig)
			if err != nil {
				return err
			}

			_, err = b.apiExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(&crdConfig)
			if errors.IsAlreadyExists(err) {
				b.log.Debugf("bootstrap skip with error %s", err)
				continue
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}
