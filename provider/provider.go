package provider

import (
	"context"
	"os"

	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "planetscale",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var planetscaleConfig planetscale_client.Configs

				err := config.Unmarshal(&planetscaleConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(planetscaleConfig.Providers) == 0 {
					planetscaleConfig.Providers = append(planetscaleConfig.Providers, planetscale_client.Config{})
				}

				if planetscaleConfig.Providers[0].Token == "" {
					planetscaleConfig.Providers[0].Token = os.Getenv("PLANETSCALE_TOKEN")
				}

				if planetscaleConfig.Providers[0].Token == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing token in configuration")
				}

				if planetscaleConfig.Providers[0].Organization == "" {
					planetscaleConfig.Providers[0].Organization = os.Getenv("PLANETSCALE_ORGANIZATION")
				}

				if planetscaleConfig.Providers[0].Organization == "" {
					return nil, schema.NewDiagnostics().AddErrorMsg("missing organization in configuration")
				}

				clients, err := planetscale_client.NewClients(planetscaleConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# token: "<YOUR_ACCESS_TOKEN>"
# organization: "<YOUR_ORGANIZATION>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var clientConfig planetscale_client.Configs
				err := config.Unmarshal(&clientConfig.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
