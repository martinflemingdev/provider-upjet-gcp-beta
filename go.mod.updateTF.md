# How Go Module Pseudo-Versions Work

A pseudo-version follows this pattern:

```
v{major}.{minor}.{patch}-0.{yyyymmddhhmmss}-{commitHash}
```

To demonstrate how a pseudo-version is derived for a specific commit, consider the following example using the `terraform-provider-google-beta` repository and tag `v6.48.0` tag:

```shell
git clone https://github.com/hashicorp/terraform-provider-google-beta.git
cd terraform-provider-google-beta
git fetch --tags
git checkout v6.48.0
git show --no-patch --format='%H %cI' v6.48.0
```

**Output:**

```
375b47b6d9c1abba1234ef567890abcdef1234567 2025-08-12T17:46:00Z 
```

This output can be translated into the corresponding Go module pseudo-version as follows:
```
v1.20.1-0.20250812174600-375b47b6d9c1
```

*(Note: `v1.20.1` is the repository’s long-standing base tag used for these pseudo versions.)*

The Go mod file entry will look like this:

```
github.com/hashicorp/terraform-provider-google-beta v1.20.1-0.20250812173325-375b47bde290
```
