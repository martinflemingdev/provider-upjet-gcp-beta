// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/upbound/provider-gcp-beta/config/namespaced/compute"
	"github.com/upbound/provider-gcp-beta/config/namespaced/container"
	"github.com/upbound/provider-gcp-beta/config/namespaced/dataform"
	"github.com/upbound/provider-gcp-beta/config/namespaced/networksecurity"
	"github.com/upbound/provider-gcp-beta/hack"
)

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced(_ context.Context, sdkProvider *schema.Provider, generationProvider bool) (*ujconfig.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		ujconfig.WithDefaultResourceOptions(
			groupOverrides(),
			externalNameConfig(),
			defaultVersion(),
			resourceConfigurator(),
			descriptionOverrides(),
		),
		ujconfig.WithRootGroup("gcp-beta.m.upbound.io"),
		ujconfig.WithShortName("gcp-beta"),
		// Comment out the following line to generate all resources.
		ujconfig.WithIncludeList(resourceList(cliReconciledExternalNameConfigs)),
		ujconfig.WithTerraformPluginSDKIncludeList(resourceList(terraformPluginSDKExternalNameConfigs)),
		ujconfig.WithReferenceInjectors([]ujconfig.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithSkipList(skipList),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithMainTemplate(hack.MainTemplate),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
	)

	registerTerraformConversions(pc)
	for _, configure := range []func(provider *ujconfig.Provider){
		container.Configure,
		networksecurity.Configure,
		compute.Configure,
		dataform.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc, nil
}

func registerTerraformConversions(pc *ujconfig.Provider) {
	for n, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		// the controller will be reconciling on the CRD API version
		// with the converted API (with embedded objects in place of
		// singleton lists), so we need the appropriate Terraform
		// converter in this case.
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}
		pc.Resources[n] = r
	}
}
