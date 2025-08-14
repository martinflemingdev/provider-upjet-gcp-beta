// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package vertexai

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_vertex_ai_index", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_vertex_ai_tensorboard", func(r *config.Resource) {
		config.MarkAsRequired(r.TerraformResource, "region")
	})
	p.AddResourceConfigurator("google_vertex_ai_endpoint_iam_member", func(r *config.Resource) {
		// 1) make region explicit to avoid the resolver error (same as above) https://github.com/hashicorp/terraform-provider-google/issues/21026
		config.MarkAsRequired(r.TerraformResource, "region")

		// 2) wire up references so endpointRef/endpointSelector are generated
		r.References["endpoint"] = config.Reference{
			TerraformName: "google_vertex_ai_endpoint",
			// If you still hit issues after adding region, uncomment this so
			// endpointRef resolves to the full resource name from observation:
			// Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractParamPath(\"id\", true)",
		}
	})
}