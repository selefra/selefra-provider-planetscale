package provider

import (
	"github.com/selefra/selefra-provider-planetscale/table_schema_generator"
	"github.com/selefra/selefra-provider-planetscale/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TablePlanetscaleOrganizationGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePlanetscaleDatabaseGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePlanetscaleRegionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePlanetscaleAuditLogGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TablePlanetscaleServiceTokenGenerator{}),
	}
}
