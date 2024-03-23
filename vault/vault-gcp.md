
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
