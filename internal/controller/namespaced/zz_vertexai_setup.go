// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	dataset "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/vertexai/dataset"
	endpoint "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/vertexai/endpoint"
	endpointiammember "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/vertexai/endpointiammember"
	featurestore "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/vertexai/featurestore"
	featurestoreentitytype "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/vertexai/featurestoreentitytype"
	tensorboard "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/vertexai/tensorboard"
)

// Setup_vertexai creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_vertexai(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.Setup,
		endpoint.Setup,
		endpointiammember.Setup,
		featurestore.Setup,
		featurestoreentitytype.Setup,
		tensorboard.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_vertexai creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_vertexai(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		dataset.SetupGated,
		endpoint.SetupGated,
		endpointiammember.SetupGated,
		featurestore.SetupGated,
		featurestoreentitytype.SetupGated,
		tensorboard.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
