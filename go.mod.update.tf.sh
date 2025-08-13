# How Go Module Pseudo-Versions Work

# A pseudo-version follows this pattern:

# v<major>.<minor>.<patch>-0.<yyyymmddhhmmss>-<commitHash>

git clone https://github.com/hashicorp/terraform-provider-google-beta.git
cd terraform-provider-google-beta
git fetch --tags
git checkout v6.48.0
git show --no-patch --format='%H %cI' v6.48.0

# Output:
# 375b47b6d9c1abba1234ef567890abcdef1234567 2025-08-12T17:46:00Z 
# == 
# v1.20.1-0.20250812174600-375b47b6d9c1

# (v1.20.1 is the repo’s long-standing base tag used for these pseudo versions.) 

github.com/hashicorp/terraform-provider-google-beta v1.20.1-0.20250812173325-375b47bde290