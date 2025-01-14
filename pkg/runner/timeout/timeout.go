package timeout

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	DefaultApplyTimeout   = 5 * time.Second
	DefaultAssertTimeout  = 30 * time.Second
	DefaultErrorTimeout   = 30 * time.Second
	DefaultDeleteTimeout  = 15 * time.Second
	DefaultCleanupTimeout = 30 * time.Second
	DefaultExecTimeout    = 5 * time.Second
)

func Get(fallback time.Duration, config *metav1.Duration, test *metav1.Duration, step *metav1.Duration, operation *metav1.Duration) time.Duration {
	if operation != nil {
		return operation.Duration
	}
	if step != nil {
		return step.Duration
	}
	if test != nil {
		return test.Duration
	}
	if config != nil {
		return config.Duration
	}
	return fallback
}

func Context(fallback time.Duration, config *metav1.Duration, test *metav1.Duration, step *metav1.Duration, operation *metav1.Duration) (context.Context, context.CancelFunc) {
	timeout := Get(fallback, config, test, step, operation)
	return context.WithTimeout(context.Background(), timeout)
}
