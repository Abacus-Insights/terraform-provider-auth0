module github.com/auth0/terraform-provider-auth0

go 1.16

require (
	github.com/auth0/go-auth0 v0.5.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/terraform-plugin-sdk v1.16.1
)

replace github.com/auth0/go-auth0 => /Users/justinusmenzel/dev/Abacus/cms-interop/terraform-provider-auth0/../auth0
