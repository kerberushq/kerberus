package retry

import (
	"k8s.io/apimachinery/pkg/util/wait"
)

func RetryOnError(backoff wait.Backoff, errorFunc func(error) bool, fn func() error) error {
	var lastConflictErr error
	err := wait.ExponentialBackoff(backoff, func() (bool, error) {
		err := fn()
		switch {
		case err == nil:
			return true, nil
		case errorFunc(err):
			lastConflictErr = err
			return false, nil
		default:
			return false, err
		}
	})
	if err == wait.ErrWaitTimeout {
		err = lastConflictErr
	}
	return err
}
