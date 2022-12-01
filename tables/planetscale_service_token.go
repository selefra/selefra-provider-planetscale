package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TablePlanetscaleServiceTokenGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleServiceTokenGenerator{}

func (x *TablePlanetscaleServiceTokenGenerator) GetTableName() string {
	return "planetscale_service_token"
}

func (x *TablePlanetscaleServiceTokenGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleServiceTokenGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleServiceTokenGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleServiceTokenGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &planetscale.ListServiceTokensRequest{Organization: planetscale_client.Organization(ctx, taskClient.(*planetscale_client.Client).Config)}
			items, err := conn.ServiceTokens.List(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range items {
				resultChannel <- i
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TablePlanetscaleServiceTokenGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleServiceTokenGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the service token.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Type of the token.").Build(),
	}
}

func (x *TablePlanetscaleServiceTokenGenerator) GetSubTables() []*schema.Table {
	return nil
}
