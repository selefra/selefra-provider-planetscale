package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePlanetscalePasswordGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscalePasswordGenerator{}

func (x *TablePlanetscalePasswordGenerator) GetTableName() string {
	return "planetscale_password"
}

func (x *TablePlanetscalePasswordGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscalePasswordGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscalePasswordGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscalePasswordGenerator) GetDataSource() *schema.DataSource {
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

			opts := &planetscale.ListDatabaseBranchPasswordRequest{Organization: orgName, Database: dbName, Branch: branchName}
			items, err := conn.Passwords.List(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range items {
				resultChannel <- passwordRow{
					OrganizationName: orgName,
					DatabaseName:     dbName,
					BranchName:       branchName,
					Password:         i,
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type passwordRow struct {
	OrganizationName string
	DatabaseName     string
	BranchName       string
	Password         *planetscale.DatabaseBranchPassword
}

func (x *TablePlanetscalePasswordGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscalePasswordGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the password was deleted.").
			Extractor(column_value_extractor.StructSelector("Password.DeletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_strings").ColumnType(schema.ColumnTypeJSON).Description("Connection strings for the branch.").
			Extractor(column_value_extractor.StructSelector("Password.ConnectionStrings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the password.").
			Extractor(column_value_extractor.StructSelector("Password.PublicID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("Role for the password.").
			Extractor(column_value_extractor.StructSelector("Password.Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the password was created.").
			Extractor(column_value_extractor.StructSelector("Password.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("branch_name").ColumnType(schema.ColumnTypeString).Description("Name of the database branch.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the password.").
			Extractor(column_value_extractor.StructSelector("Password.Name")).Build(),
	}
}

func (x *TablePlanetscalePasswordGenerator) GetSubTables() []*schema.Table {
	return nil
}
