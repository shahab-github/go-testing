
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
