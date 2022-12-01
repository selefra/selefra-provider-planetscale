package tables

import (
	"context"

	"github.com/planetscale/planetscale-go/planetscale"
	"github.com/selefra/selefra-provider-planetscale/planetscale_client"
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TablePlanetscaleAuditLogGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TablePlanetscaleAuditLogGenerator{}

func (x *TablePlanetscaleAuditLogGenerator) GetTableName() string {
	return "planetscale_audit_log"
}

func (x *TablePlanetscaleAuditLogGenerator) GetTableDescription() string {
	return ""
}

func (x *TablePlanetscaleAuditLogGenerator) GetVersion() uint64 {
	return 0
}

func (x *TablePlanetscaleAuditLogGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TablePlanetscaleAuditLogGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := planetscale_client.Connect(ctx, taskClient.(*planetscale_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			opts := &planetscale.ListAuditLogsRequest{Organization: planetscale_client.Organization(ctx, taskClient.(*planetscale_client.Client).Config)}

			startingAfter := ""
			for {

				resp, err := conn.AuditLogs.List(ctx, opts, planetscale.WithStartingAfter(startingAfter), planetscale.WithLimit(1000))
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, i := range resp.Data {
					resultChannel <- i
				}
				if resp.HasNext {
					startingAfter = *resp.CursorEnd
				} else {
					break
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TablePlanetscaleAuditLogGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TablePlanetscaleAuditLogGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("action").ColumnType(schema.ColumnTypeString).Description("Short action for this audit record, e.g. created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remote_ip").ColumnType(schema.ColumnTypeString).Description("IP address the action was requested from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_id").ColumnType(schema.ColumnTypeString).Description("ID of the resource type for this audit record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_display_name").ColumnType(schema.ColumnTypeString).Description("Display name for the target resoruce, e.g. test_db.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the audit record was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actor_type").ColumnType(schema.ColumnTypeString).Description("Type of the actor, e.g. User.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auditable_type").ColumnType(schema.ColumnTypeString).Description("Resource type the audit entry is for, e.g. Branch.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auditable_display_name").ColumnType(schema.ColumnTypeString).Description("Display name of the resource for this audit entry, e.g. test_branch.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actor_display_name").ColumnType(schema.ColumnTypeString).Description("Display name of the actor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_type").ColumnType(schema.ColumnTypeString).Description("Resource type for this audit record, e.g. Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Metadata for the audit record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actor_id").ColumnType(schema.ColumnTypeString).Description("Unique ID of the actor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("audit_action").ColumnType(schema.ColumnTypeString).Description("Full action for this audit record, e.g. deploy_request.created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Description("Geographic location the action was requested from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the audit record was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique ID of the log entry.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Type of log entry, e.g. AuditLogEvent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auditable_id").ColumnType(schema.ColumnTypeString).Description("Unique ID for the resource type of the audit entry.").Build(),
	}
}

func (x *TablePlanetscaleAuditLogGenerator) GetSubTables() []*schema.Table {
	return nil
}
