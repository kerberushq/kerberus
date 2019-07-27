package bootstrap

//go:generate go get github.com/go-bindata/go-bindata/go-bindata
//go:generate go-bindata -nometadata -pkg $GOPACKAGE -prefix data data/...
//go:generate gofmt -s -l -w bindata.go

import (
	"github.com/sirupsen/logrus"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
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
		log.Fatal("Failed to initialize apiextensions client with %s", err)
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

}
