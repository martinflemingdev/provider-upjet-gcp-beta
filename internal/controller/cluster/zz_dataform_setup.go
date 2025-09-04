// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	repository "github.com/upbound/provider-gcp-beta/internal/controller/cluster/dataform/repository"
	repositoryiammember "github.com/upbound/provider-gcp-beta/internal/controller/cluster/dataform/repositoryiammember"
	repositoryreleaseconfig "github.com/upbound/provider-gcp-beta/internal/controller/cluster/dataform/repositoryreleaseconfig"
	repositoryworkflowconfig "github.com/upbound/provider-gcp-beta/internal/controller/cluster/dataform/repositoryworkflowconfig"
)

// Setup_dataform creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dataform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		repository.Setup,
		repositoryiammember.Setup,
		repositoryreleaseconfig.Setup,
		repositoryworkflowconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated_dataform creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated_dataform(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		repository.SetupGated,
		repositoryiammember.SetupGated,
		repositoryreleaseconfig.SetupGated,
		repositoryworkflowconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
