/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/MacroPower/provider-authentik/apis/provider/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Application.
func (mg *Application) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ProtocolProvider),
		Extract:      resource.ExtractParamPath("id", true),
		Reference:    mg.Spec.ForProvider.ProtocolProviderRef,
		Selector:     mg.Spec.ForProvider.ProtocolProviderSelector,
		To: reference.To{
			List:    &v1alpha1.ProxyList{},
			Managed: &v1alpha1.Proxy{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProtocolProvider")
	}
	mg.Spec.ForProvider.ProtocolProvider = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProtocolProviderRef = rsp.ResolvedReference

	return nil
}