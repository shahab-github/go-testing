
`mybindings.hcl`

```hcl
resource "//cloudresourcemanager.googleapis.com/projects/clgcporg8-068" {
  roles = [
    "roles/viewer",
  ]
}
```

### Vault command to create the roleset
```bash
vault write gcp/roleset/my-token-roleset \
  project="clgcporg8-068" \
  secret_type="access_token" \
  token_scopes="https://www.googleapis.com/auth/cloud-platform" \
  bindings=@mybindings.hcl
```

### sample command
```
curl -X 'POST' \
  'http://54.160.189.39:8200/v1/gcp/roleset/testroleset' \
  -H 'accept: */*' \
  -H 'Content-Type: application/json' \
  -H 'X-Vault-Token: myroot' \
  -d '{
  "bindings": "resource \"//cloudresourcemanager.googleapis.com/projects/clgcporg8-068\" { \n\t roles = [\n\"roles/viewer\", \n] \n}",
  "project": "clgcporg8-068",
  "secret_type": "access_token",
  "token_scopes": [
    "https://www.googleapis.com/auth/cloud-platform"
  ]
}'
```

## medium page for gcp vault configuration
https://marco-urrea.medium.com/hashicorp-vault-gcp-secrets-engine-70813983e33a


  
Code of the roleset to generate OAuth2 Access Tokens:
```
vault write gcp/roleset/my-token-roleset \
project="medium-vault-gcp" \
secret_type="access_token"  \
token_scopes="https://www.googleapis.com/auth/cloud-platform" \
bindings=-<<EOF
resource "//cloudresourcemanager.googleapis.com/projects/medium-vault-gcp" 
{
roles = ["roles/viewer"]
}
EOF
```

Code of the roleset to generate GCP Service Account Keys:
```
vault write gcp/roleset/my-key-roleset \
project="medium-vault-gcp" \
secret_type="service_account_key"  \
bindings=-<<EOF
resource "//cloudresourcemanager.googleapis.com/projects/medium-vault-gcp" {
roles = ["roles/viewer"]
}
EOF
```

To configure a static account that generates OAuth2 access tokens (preferred):
```
vault write gcp/static-account/my-token-account \
service_account_email="svc-vault679@vault679.iam.gserviceaccount.com" \
secret_type="access_token"  \
token_scopes="https://www.googleapis.com/auth/cloud-platform" \
bindings=-<<EOF
resource "//cloudresourcemanager.googleapis.com/projects/vault679" {
roles = ["roles/viewer"]
}
EOF
```

To configure a static account that generates GCP Service Account keys:
```
vault write gcp/static-account/my-key-account \
service_account_email="svc-vault679@vault679.iam.gserviceaccount.com" \
secret_type="service_account_key"  \
bindings=-<<EOF
resource "//cloudresourcemanager.googleapis.com/projects/vault679" {
roles = ["roles/viewer"]
}
EOF
```