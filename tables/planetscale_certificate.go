package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePlanetscaleCertificateGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleCertificateGenerator{}

func (x *TablePlanetscaleCertificateGenerator) GetTableName() string {
	return "planetscale_certificate"
}

func (x *TablePlanetscaleCertificateGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleCertificateGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleCertificateGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleCertificateGenerator) GetDataSource() *schema.DataSource {
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

			opts := &planetscale.ListDatabaseBranchCertificateRequest{Organization: orgName, Database: dbName, Branch: branchName}
			items, err := conn.Certificates.List(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range items {
				resultChannel <- certificateRow{
					OrganizationName: orgName,
					DatabaseName:     dbName,
					BranchName:       branchName,
					Certificate:      i,
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type certificateRow struct {
	OrganizationName string
	DatabaseName     string
	BranchName       string
	Certificate      *planetscale.DatabaseBranchCertificate
}
type databaseBranchRow struct {
	OrganizationName string
	DatabaseName     string
	Branch           *planetscale.DatabaseBranch
}

func (x *TablePlanetscaleCertificateGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleCertificateGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("certificate").ColumnType(schema.ColumnTypeString).Description("Certificate string.").
			Extractor(column_value_extractor.StructSelector("Certificate.Certificate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the certificate was deleted.").
			Extractor(column_value_extractor.StructSelector("Certificate.DeletedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("branch_name").ColumnType(schema.ColumnTypeString).Description("Name of the database branch.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("Role for the certificate.").
			Extractor(column_value_extractor.StructSelector("Certificate.Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the certificate was created.").
			Extractor(column_value_extractor.StructSelector("Certificate.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the certificate.").
			Extractor(column_value_extractor.StructSelector("Certificate.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("ID of the certificate.").
			Extractor(column_value_extractor.StructSelector("Certificate.PublicID")).Build(),
	}
}

func (x *TablePlanetscaleCertificateGenerator) GetSubTables() []*schema.Table {
	return nil
}
