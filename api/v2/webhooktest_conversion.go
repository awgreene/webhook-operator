package v2

import (
	v1 "github.com/awgreene/webhook-operator/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

var _ conversion.Convertible = &WebhookTest{}

// ConvertTo converts this ConversionTest to the Hub version (v1).
func (src *WebhookTest) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.WebhookTest)

	// Object Meta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Mutate = src.Spec.Conversion.Mutate
	dst.Spec.Valid = src.Spec.Conversion.Valid

	return nil
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *WebhookTest) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.WebhookTest)

	// Object Meta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.Conversion.Valid = src.Spec.Valid
	dst.Spec.Conversion.Mutate = src.Spec.Mutate

	return nil
}
