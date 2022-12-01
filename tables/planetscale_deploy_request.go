package tables

import (
	"context"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TablePlanetscaleDeployRequestGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleDeployRequestGenerator{}

func (x *TablePlanetscaleDeployRequestGenerator) GetTableName() string {
	return "planetscale_deploy_request"
}

func (x *TablePlanetscaleDeployRequestGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleDeployRequestGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleDeployRequestGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleDeployRequestGenerator) GetDataSource() *schema.DataSource {
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

			opts := &planetscale.ListDeployRequestsRequest{Organization: org, Database: dbName}
			items, err := conn.DeployRequests.List(ctx, opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, i := range items {
				resultChannel <- deployRequestRow{
					OrganizationName: org,
					DatabaseName:     dbName,
					DeployRequest:    i,
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type deployRequestRow struct {
	OrganizationName string
	DatabaseName     string
	DeployRequest    *planetscale.DeployRequest
}

func (x *TablePlanetscaleDeployRequestGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleDeployRequestGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("branch").ColumnType(schema.ColumnTypeString).Description("Deploy request branch.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.Branch")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("into_branch").ColumnType(schema.ColumnTypeString).Description("Deploy request into branch.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.IntoBranch")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approved").ColumnType(schema.ColumnTypeBool).Description("True if the deploy request is approved.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.Approved")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the deploy request was created.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.CreatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the deploy request was updated.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.UpdatedAt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_name").ColumnType(schema.ColumnTypeString).Description("Name of the organization.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_name").ColumnType(schema.ColumnTypeString).Description("Name of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number").ColumnType(schema.ColumnTypeInt).Description("Number for this deploy request.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.Number")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique ID for the deplloy request.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("State of the deploy request.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notes").ColumnType(schema.ColumnTypeString).Description("Notes for the deploy request.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.Notes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment").ColumnType(schema.ColumnTypeJSON).Description("Details of the deployment.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.Deployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("closed_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the deploy request was closed.").
			Extractor(column_value_extractor.StructSelector("DeployRequest.ClosedAt")).Build(),
	}
}

func (x *TablePlanetscaleDeployRequestGenerator) GetSubTables() []*schema.Table {
	return nil
}
