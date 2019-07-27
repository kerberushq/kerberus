package controller

import (
	"context"

	"github.com/sirupsen/logrus"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/kerberushq/kerberus/pkg/controller/crdrequest"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager) error

// AddToManager adds all Controllers to the Manager
func AddToManager(ctx context.Context, log *logrus.Entry, m manager.Manager) error {
	if err := crdrequest.Add(ctx, log, m); err != nil {
		return err
	}
	return nil
}
