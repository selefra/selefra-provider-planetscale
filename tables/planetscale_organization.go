package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TablePlanetscaleOrganizationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleOrganizationGenerator{}

func (x *TablePlanetscaleOrganizationGenerator) GetTableName() string {
	return "planetscale_organization"
}

func (x *TablePlanetscaleOrganizationGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleOrganizationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleOrganizationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleOrganizationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			items, err := conn.Organizations.List(ctx)
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

func (x *TablePlanetscaleOrganizationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleOrganizationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the organization was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the organization was updated.").Build(),
	}
}

func (x *TablePlanetscaleOrganizationGenerator) GetSubTables() []*schema.Table {
	return nil
}
