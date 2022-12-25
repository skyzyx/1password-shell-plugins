package infracost

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func InfracostCLI() schema.Executable {
	return schema.Executable{
		Name:      "Infracost CLI",
		Runs:      []string{"infracost"},
		DocsURL:   sdk.URL("https://www.infracost.io/docs/"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.APIKey,
			},
		},
	}
}
