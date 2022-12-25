package infracost

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestAPIKeyProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, APIKey().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.APIKey: "KJfyPvXDzYvxyTBGeUUKaKbOSEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"INFRACOST_API_KEY": "KJfyPvXDzYvxyTBGeUUKaKbOSEXAMPLE",
				},
			},
		},
	})
}

func TestAPIKeyImporter(t *testing.T) {
	plugintest.TestImporter(t, APIKey().Importer, map[string]plugintest.ImportCase{
		"INFRACOST_API_KEY": {
			Environment: map[string]string{
				"INFRACOST_API_KEY": "KJfyPvXDzYvxyTBGeUUKaKbOSEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.APIKey: "KJfyPvXDzYvxyTBGeUUKaKbOSEXAMPLE",
					},
				},
			},
		},
	})
}
