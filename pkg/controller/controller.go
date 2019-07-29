package controller

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/kerberushq/kerberus/pkg/controller/crdrequest"
	namespacectrl "github.com/kerberushq/kerberus/pkg/controller/namespace"
	configv1alpha1client "github.com/kerberushq/kerberus/pkg/generated/config/clientset/versioned/typed/config/v1alpha1"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error

// AddToManager adds all Controllers to the Manager
func AddToManager(ctx context.Context, log *logrus.Entry, m manager.Manager) error {
	// Get Config client for controllers
	configClient, err := configv1alpha1client.NewForConfig(m.GetConfig())
	if err != nil {
		log.Fatalf("Failed to initialize configv1alpha1client client with %s", err)
	}

	configList, err := configClient.KerberusConfigs("kerberus").List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to initialize kerberus with config %s", err)
		os.Exit(1)
	}
	if len(configList.Items) > 1 {
		log.Fatal("Found multiple config object for kerberus. Only one config instance is allowed")
		os.Exit(1)
	}
	config := configList.Items[0]

	if config.Spec.CRDRequest.Enable {
		if err := crdrequest.Add(ctx, log, m); err != nil {
			return err
		}
	}

	if config.Spec.RBACConfig.Enable {
		if err := namespacectrl.Add(ctx, log, m); err != nil {
			return err
		}
	}

	return nil
}
