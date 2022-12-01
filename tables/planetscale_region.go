package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TablePlanetscaleRegionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleRegionGenerator{}

func (x *TablePlanetscaleRegionGenerator) GetTableName() string {
	return "planetscale_region"
}

func (x *TablePlanetscaleRegionGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleRegionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleRegionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleRegionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := &planetscale.ListRegionsRequest{}
			items, err := conn.Regions.List(ctx, opts)
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

func (x *TablePlanetscaleRegionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleRegionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Description("Location for the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).Description("True if the region is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("slug").ColumnType(schema.ColumnTypeString).Description("Slug of the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Display name of the region.").Build(),
	}
}

func (x *TablePlanetscaleRegionGenerator) GetSubTables() []*schema.Table {
	return nil
}
