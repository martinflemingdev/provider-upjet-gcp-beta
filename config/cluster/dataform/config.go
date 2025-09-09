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
    // // generates Ref and Selector fields, needs just the repo name, fails on full path
    // p.AddResourceConfigurator("google_dataform_repository_release_config", func(r *config.Resource) {
    //     r.References["repository"] = config.Reference{
    //         TerraformName: "google_dataform_repository",
    //         Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
    //     }
    // })
    // generates Ref and Selector fields, needs just the repo name, fails on full path
	p.AddResourceConfigurator("google_dataform_repository_workflow_config", func(r *config.Resource) {
		// r.References["repository"] = config.Reference{
		// 	TerraformName: "google_dataform_repository",
		// 	Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractResourceID()",
		// }
        // Enables cross-API group ServiceAccount references (gcp-beta ↔ gcp GA providers)
        r.References["invocation_config.service_account"] = config.Reference{
            TerraformName: "google_service_account",
            Extractor:     "github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath(\"email\", false)",
        }
    })
}