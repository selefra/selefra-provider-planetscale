package tables

import (
	"context"
	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePlanetscaleDatabaseGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleDatabaseGenerator{}

func (x *TablePlanetscaleDatabaseGenerator) GetTableName() string {
	return "planetscale_database"
}

func (x *TablePlanetscaleDatabaseGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleDatabaseGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleDatabaseGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleDatabaseGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)

			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &planetscale.ListDatabasesRequest{Organization: planetscale_client.Organization(ctx, taskClient.(*planetscale_client.Client).Config)}
			items, err := conn.Databases.List(ctx, opts)
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

func (x *TablePlanetscaleDatabaseGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleDatabaseGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("notes").ColumnType(schema.ColumnTypeString).Description("Notes for the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the database was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the database was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_slug").ColumnType(schema.ColumnTypeString).Description("Region where the database is located.").
			Extractor(column_value_extractor.StructSelector("Region.Slug")).Build(),
	}
}

func (x *TablePlanetscaleDatabaseGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TablePlanetscaleDeployRequestGenerator{}),
		table_schema_generator.GenTableSchema(&TablePlanetscaleDatabaseBranchGenerator{}),
	}
}
