package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
	"knative.dev/pkg/webhook/resourcesemantics"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var _ resourcesemantics.GenericCRD = (*Tester)(nil)

func (tst *Tester) Validate(context.Context) *apis.FieldError {
	logf.Log.Info("lel validates")
	return nil
}

func (tst *Tester) SetDefaults(context.Context) {
	logf.Log.Info("lel defaults")
}
