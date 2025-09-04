// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/upbound/provider-gcp-beta/config/cluster/common"
)

// terraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the no-fork
// architecture for this provider.
var terraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{

	// cloudplatform
	//
	// This resource does not support import.
	"google_project_service_identity": config.IdentifierFromProvider,
	// Service accounts can be imported using their URI, e.g. projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com
	"google_service_account": config.TemplatedStringAsIdentifier("account_id", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/serviceAccounts/{{ .external_name }}@{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}.iam.gserviceaccount.com"),

	// container
	//
	// Imported by using the following format: projects/my-gcp-project/locations/us-east1-a/clusters/my-cluster
	"google_container_cluster": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/clusters/{{ .external_name }}"),
	// Imported by using the following format: my-gcp-project/us-east1-a/my-cluster/main-pool
	"google_container_node_pool": config.TemplatedStringAsIdentifier("name", "{{ .setup.configuration.project }}/{{ .parameters.location }}/{{ .parameters.cluster }}/{{ .external_name }}"),

	// networksecurity
	//
	// The resource can be imported using projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{name}}
	"google_network_security_server_tls_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ .setup.configuration.project }}/locations/{{ .parameters.location }}/serverTlsPolicies/{{ .external_name }}"),

	// compute
	//
	// Imported by using projects/{{project}}/global/healthChecks/{{name}}
	"google_compute_health_check": config.TemplatedStringAsIdentifier("name", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/global/healthChecks/{{ .external_name }}"),
	// Imported by using projects/{{project}}/regions/{{region}}/securityPolicies/{{name}}
	"google_compute_region_security_policy": config.TemplatedStringAsIdentifier("name", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/regions/{{ .parameters.region }}/securityPolicies/{{ .external_name }}"),
	// Imported by using projects/{{project}}/regions/{{region}}/backendServices/{{name}}
	"google_compute_region_backend_service": config.TemplatedStringAsIdentifier("name", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/regions/{{ .parameters.region }}/backendServices/{{ .external_name }}"),

	// dataform
	//
	// Imported by using projects/{{project}}/locations/{{region}}/repositories/{{name}}
	"google_dataform_repository": config.TemplatedStringAsIdentifier("name", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/locations/{{ .parameters.region }}/repositories/{{ .external_name }}"),
	// Imported by using the following format: projects/{{project}}/locations/{{region}}/repositories/{{repository}} roles/viewer user:jane@example.com
	"google_dataform_repository_iam_member": config.IdentifierFromProvider,
	// Imported by using projects/{{project}}/locations/{{region}}/repositories/{{repository}}/releaseConfigs/{{name}}
	"google_dataform_repository_release_config": config.TemplatedStringAsIdentifier("name", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/locations/{{ .parameters.region }}/repositories/{{ .parameters.repository }}/releaseConfigs/{{ .external_name }}"),
	// Imported by using projects/{{project}}/locations/{{region}}/repositories/{{repository}}/workflowConfigs/{{name}}
	"google_dataform_repository_workflow_config": config.TemplatedStringAsIdentifier("name", "projects/{{ if .parameters.project }}{{ .parameters.project }}{{ else }}{{ .setup.configuration.project }}{{ end }}/locations/{{ .parameters.region }}/repositories/{{ .parameters.repository }}/workflowConfigs/{{ .external_name }}"),
}

// cliReconciledExternalNameConfigs contains all external name configurations
// belonging to Terraform resources to be reconciled under the CLI-based
// architecture for this provider.
var cliReconciledExternalNameConfigs = map[string]config.ExternalName{}

// resourceConfigurator applies all external name configs
// listed in the table terraformPluginSDKExternalNameConfigs and
// cliReconciledExternalNameConfigs and sets the version
// of those resources to v1beta1. For those resource in
// terraformPluginSDKExternalNameConfigs, it also sets
// config.Resource.UseNoForkClient to `true`.
func resourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// if configured both for the no-fork and CLI based architectures,
		// no-fork configuration prevails
		e, configured := terraformPluginSDKExternalNameConfigs[r.Name]
		if !configured {
			e, configured = cliReconciledExternalNameConfigs[r.Name]
		}
		if !configured {
			return
		}
		r.Version = common.VersionV1Beta1
		r.ExternalName = e
	}
}
