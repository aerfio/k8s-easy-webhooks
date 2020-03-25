package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
	"knative.dev/pkg/webhook/resourcesemantics"
)

var _ resourcesemantics.GenericCRD = (*Tester)(nil)

func (tst *Tester) Validate(context.Context) *apis.FieldError {
	if tst.Spec.Data != nil && *tst.Spec.Data > 12 {
		return apis.ErrGeneric("Incorrect values", ".Spec.Data")
	}
	return nil
}

func (tst *Tester) SetDefaults(context.Context) {
	if tst.Spec.Foo == "" {
		tst.Spec.Foo = "defaulted"
	}
}
