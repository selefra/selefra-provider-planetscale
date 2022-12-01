package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePlanetscaleBackupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleBackupGenerator{}

func (x *TablePlanetscaleBackupGenerator) GetTableName() string {
	return "planetscale_backup"
}

func (x *TablePlanetscaleBackupGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleBackupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleBackupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleBackupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			branch := task.ParentRawResult.(databaseBranchRow)

			var orgName, dbName, branchName string
			if task.ParentRawResult != nil {
				branch = task.ParentRawResult.(databaseBranchRow)
				orgName = branch.OrganizationName
				dbName = branch.DatabaseName
				branchName = branch.Branch.Name
			} else {
				orgName = planetscale_client.Organization(ctx, taskClient.(*planetscale_client.Client).Config)
				dbName = task.ParentRawResult.(backupRow).DatabaseName
				branchName = task.ParentRawResult.(backupRow).BranchName
			}

			opts := &planetscale.ListBackupsRequest{Organization: orgName, Database: dbName, Branch: branchName}
			items, err := conn.Backups.List(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range items {
				resultChannel <- backupRow{
					OrganizationName: orgName,
					DatabaseName:     dbName,
					BranchName:       branchName,
					Backup:           i,
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type backupRow struct {
	OrganizationName string
	DatabaseName     string
	BranchName       string
	Backup           *planetscale.Backup
}

func (x *TablePlanetscaleBackupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleBackupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the backup was created.").
			Extractor(column_value_extractor.StructSelector("Backup.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the backup was updated.").
			Extractor(column_value_extractor.StructSelector("Backup.UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("branch_name").ColumnType(schema.ColumnTypeString).Description("Name of the database branch.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the backup.").
			Extractor(column_value_extractor.StructSelector("Backup.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("Size of the backup.").
			Extractor(column_value_extractor.StructSelector("Backup.Size")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the backup.").
			Extractor(column_value_extractor.StructSelector("Backup.PublicID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("State of the backup.").
			Extractor(column_value_extractor.StructSelector("Backup.State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("started_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the backup was started.").
			Extractor(column_value_extractor.StructSelector("Backup.StartedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("completed_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the backup was completed.").
			Extractor(column_value_extractor.StructSelector("Backup.CompletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expires_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the backup expires.").
			Extractor(column_value_extractor.StructSelector("Backup.ExpiresAt")).Build(),
	}
}

func (x *TablePlanetscaleBackupGenerator) GetSubTables() []*schema.Table {
	return nil
}
