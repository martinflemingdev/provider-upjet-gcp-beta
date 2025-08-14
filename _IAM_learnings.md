## *IAMMember Ref/Selector fields generation 

You’re super close—your r.References["endpoint"] = … is exactly what you need, but two common gotchas are why you’re not seeing endpointRef / endpointSelector show up:

the reference needs an extractor so Upjet knows which field from the referenced Endpoint to inject, and

your Configure function must be registered in config/provider.go or the generator won’t apply it.

Here’s the minimal, working setup.

```go
// config/vertexai/config.go
package vertexai

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds custom config for Vertex AI resources.
func Configure(p *config.Provider) {
	// IAM Member should point at the Endpoint resource via reference.
	p.AddResourceConfigurator("google_vertex_ai_endpoint_iam_member", func(r *config.Resource) {
		r.References["endpoint"] = config.Reference{
			// Tell Upjet which resource this field points to:
			TerraformName: "google_vertex_ai_endpoint",

			// Extract the value TF expects. For Vertex AI IAM, TF uses the Endpoint's *name*
			// (the full resource name projects/.../locations/.../endpoints/...), which is
			// an observed attribute, so extract it from observation:
			Extractor: "github.com/crossplane/upjet/pkg/resource.ExtractParamPath(\"name\", true)",
			// (true => from observation/status.atProvider)
		}

		// Optional: make location explicit if you want to force it at the MR level.
		// If you want this field required in the MR schema:
		// config.MarkAsRequired(r.TerraformResource, "location")
	})
}
```

Why ExtractParamPath("name", true)? The TF docs and examples use endpoint = google_vertex_ai_endpoint.endpoint.name for IAM resources, and in GCP providers the name is typically the full resource name. Upjet will then generate endpointRef and endpointSelector that resolve to that name. (Upjet will create FooRef and FooSelector automatically for a string field foo when you configure a config.Reference.) 
pkg.go.dev

Alternative extractor: if your Endpoint external name already stores the full resource ID you could use resource.ExtractResourceID(), but name is the value TF explicitly documents for these IAM bindings.

2) Make sure your configurator is wired up

Upjet only runs configurators listed in config/provider.go. Add (or ensure) this line is present:


```go
// config/provider.go
package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	// ... other service imports ...
	vertexaicfg "github.com/your/module/path/config/vertexai"
)

func GetProvider() *config.Provider {
	p := config.NewProvider(
		// ... provider metadata ...
	)

	// Register all service configurators here
	for _, configure := range []func(*config.Provider){
		// ... others ...
		vertexaicfg.Configure,   // <-- ensure this is included
	} {
		configure(p)
	}
	return p
}
```

