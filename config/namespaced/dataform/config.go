// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package dataform

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_dataform_repository", func(r *config.Resource) {
		r.MarkAsRequired("region")
	})
	// IAMMember needs full resource path
	p.AddResourceConfigurator("google_dataform_repository_iam_member", func(r *config.Resource) {
		r.References["repository"] = config.Reference{
			TerraformName: "google_dataform_repository",
			Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(\"id\", true)",
		}
	})
	// target GA MR type for serviceAccountSelector compatibility
	p.AddResourceConfigurator("google_dataform_repository_workflow_config", func(r *config.Resource) {
		r.References["invocation_config.service_account"] = config.Reference{
			Type:                "github.com/upbound/provider-gcp/apis/cluster/cloudplatform/v1beta1.ServiceAccount",
			Extractor:           "github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(\"email\", false)",
			RefFieldName:        "ServiceAccountGaRef",
			SelectorFieldName:   "ServiceAccountGaSelector",
		}
	})
}
