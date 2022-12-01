package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePlanetscaleDatabaseBranchGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleDatabaseBranchGenerator{}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetTableName() string {
	return "planetscale_database_branch"
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)

			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			org := planetscale_client.Organization(ctx, taskClient.(*planetscale_client.Client).Config)

			var dbName string
			if task.ParentRawResult != nil {
				dbName = task.ParentRawResult.(*planetscale.Database).Name
			} else if task.ParentRawResult.(backupRow).DatabaseName != "" {
				dbName = task.ParentRawResult.(backupRow).DatabaseName
			}

			opts := &planetscale.ListDatabaseBranchesRequest{Organization: org, Database: dbName}
			items, err := conn.DatabaseBranches.List(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range items {
				resultChannel <- databaseBranchRow{
					OrganizationName: org,
					DatabaseName:     dbName,
					Branch:           i,
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_branch").ColumnType(schema.ColumnTypeString).Description("Parent of this branch.").
			Extractor(column_value_extractor.StructSelector("Branch.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_slug").ColumnType(schema.ColumnTypeString).Description("Region where the database is located.").
			Extractor(column_value_extractor.StructSelector("Branch.Region.Slug")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("production").ColumnType(schema.ColumnTypeBool).Description("True if this branch is in production.").
			Extractor(column_value_extractor.StructSelector("Branch.Production")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the branch was created.").
			Extractor(column_value_extractor.StructSelector("Branch.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the branch was updated.").
			Extractor(column_value_extractor.StructSelector("Branch.UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the branch.").
			Extractor(column_value_extractor.StructSelector("Branch.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ready").ColumnType(schema.ColumnTypeBool).Description("True if the branch is ready.").
			Extractor(column_value_extractor.StructSelector("Branch.Ready")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_host_url").ColumnType(schema.ColumnTypeString).Description("Host name to access the database.").
			Extractor(column_value_extractor.StructSelector("Branch.AccessHostURL")).Build(),
	}
}

func (x *TablePlanetscaleDatabaseBranchGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TablePlanetscaleCertificateGenerator{}),
		table_schema_generator.GenTableSchema(&TablePlanetscalePasswordGenerator{}),
		table_schema_generator.GenTableSchema(&TablePlanetscaleBackupGenerator{}),
	}
}
