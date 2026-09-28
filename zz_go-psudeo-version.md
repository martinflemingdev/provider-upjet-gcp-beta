UPDATE GO.MOD 

How Go Module Pseudo-Versions Work

A pseudo-version follows this pattern:

v<major>.<minor>.<patch>-0.<yyyymmddhhmmss>-<commitHash>

git clone https://github.com/hashicorp/terraform-provider-google-beta.git
cd terraform-provider-google-beta
git fetch --tags
git checkout v6.48.0
git show --no-patch --format='%H %cI' v6.48.0

375b47b6d9c1abba1234ef567890abcdef1234567 2025-08-12T17:46:00Z == v1.20.1-0.20250812174600-375b47b6d9c1

(v1.20.1 is the repo’s long-standing base tag used for these pseudo versions.) 

v6.50.0 == v1.20.1-0.20250919180509-956f510990a7

v7.12.0 == v1.20.1-0.20251118174721-4dd1bc5aadbc

v7.20.0 == v1.20.1-0.20260217183254-340083741c96

v7.28.0 == v1.20.1-0.20260414164434-54c21b6734da