package infracost

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "infracost",
		Platform: schema.PlatformInfo{
			Name:     "Infracost",
			Homepage: sdk.URL("https://infracost.io"),
		},
		Credentials: []schema.CredentialType{
			APIKey(),
		},
		Executables: []schema.Executable{
			InfracostCLI(),
		},
	}
}
