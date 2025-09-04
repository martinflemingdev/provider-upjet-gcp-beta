// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	projectserviceidentity "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/cloudplatform/projectserviceidentity"
	serviceaccount "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/cloudplatform/serviceaccount"
	healthcheck "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/compute/healthcheck"
	regionbackendservice "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/compute/regionbackendservice"
	regionsecuritypolicy "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/compute/regionsecuritypolicy"
	cluster "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/container/cluster"
	nodepool "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/container/nodepool"
	repository "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/dataform/repository"
	repositoryiammember "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/dataform/repositoryiammember"
	repositoryreleaseconfig "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/dataform/repositoryreleaseconfig"
	repositoryworkflowconfig "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/dataform/repositoryworkflowconfig"
	servertlspolicy "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/networksecurity/servertlspolicy"
	providerconfig "github.com/upbound/provider-gcp-beta/internal/controller/namespaced/providerconfig"
)

// Setup_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectserviceidentity.Setup,
		serviceaccount.Setup,
		healthcheck.Setup,
		regionbackendservice.Setup,
		regionsecuritypolicy.Setup,
		cluster.Setup,
		nodepool.Setup,
		repository.Setup,
		repositoryiammember.Setup,
		repositoryreleaseconfig.Setup,
		repositoryworkflowconfig.Setup,
		servertlspolicy.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_monolith creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_monolith(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		projectserviceidentity.SetupGated,
		serviceaccount.SetupGated,
		healthcheck.SetupGated,
		regionbackendservice.SetupGated,
		regionsecuritypolicy.SetupGated,
		cluster.SetupGated,
		nodepool.SetupGated,
		repository.SetupGated,
		repositoryiammember.SetupGated,
		repositoryreleaseconfig.SetupGated,
		repositoryworkflowconfig.SetupGated,
		servertlspolicy.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
