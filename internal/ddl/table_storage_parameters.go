package ddl

import "strconv"

// TableStorageParameters influence how table data is stored and maintained
type TableStorageParameters struct {
	Fillfactor                        *int     `hcl:"fillfactor,optional"`
	ParallelWorkers                   *int     `hcl:"parallel_workers,optional"`
	AutovacuumEnabled                 *bool    `hcl:"autovacuum_enabled,optional"`
	AutovacuumVacuumThreshold         *int     `hcl:"autovacuum_vacuum_threshold,optional"`
	AutovacuumVacuumScaleFactor       *float64 `hcl:"autovacuum_vacuum_scale_factor,optional"`
	AutovacuumAnalyzeThreshold        *int     `hcl:"autovacuum_analyze_threshold,optional"`
	AutovacuumAnalyzeScaleFactor      *float64 `hcl:"autovacuum_analyze_scale_factor,optional"`
	AutovacuumVacuumCostDelay         *int     `hcl:"autovacuum_vacuum_cost_limit,optional"`
	AutovacuumFreezeMinAge            *int     `hcl:"autovacuum_freeze_min_age,optional"`
	AutovacuumFreezeMaxAge            *int     `hcl:"autovacuum_freeze_max_age,optional"`
	AutovacuumFreezeTableAge          *int     `hcl:"autovacuum_freeze_table_age,optional"`
	AutovacuumMultixactFreezeMinAge   *int     `hcl:"autovacuum_multixact_freeze_min_age,optional"`
	AutovacuumMultixactFreezeMaxAge   *int     `hcl:"autovacuum_multixact_freeze_max_age,optional"`
	AutovacuumMultixactFreezeTableAge *int     `hcl:"autovacuum_multixact_freeze_table_age,optional"`
	LogAutovacuumMinDuration          *int     `hcl:"log_autovacuum_min_duration,optional"`
	UserCatalogTable                  *bool    `hcl:"user_catalog_table,optional"`
}

//nolint:gocyclo
func (p *TableStorageParameters) Set(key, value string) error {
	switch key {
	case "fillfactor":
		tmp, err := strconv.Atoi(value)
		p.Fillfactor = &tmp
		return err
	case "parallel_workers":
		tmp, err := strconv.Atoi(value)
		p.ParallelWorkers = &tmp
		return err
	case "autovacuum_enabled":
		tmp, err := strconv.ParseBool(value)
		p.AutovacuumEnabled = &tmp
		return err
	case "autovacuum_vacuum_threshold":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumVacuumThreshold = &tmp
		return err
	case "autovacuum_vacuum_scale_factor":
		tmp, err := strconv.ParseFloat(value, 64)
		p.AutovacuumVacuumScaleFactor = &tmp
		return err
	case "autovacuum_analyze_threshold":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumAnalyzeThreshold = &tmp
		return err
	case "autovacuum_analyze_scale_factor":
		tmp, err := strconv.ParseFloat(value, 64)
		p.AutovacuumAnalyzeScaleFactor = &tmp
		return err
	case "autovacuum_vacuum_cost_limit":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumVacuumCostDelay = &tmp
		return err
	case "autovacuum_freeze_min_age":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumFreezeMinAge = &tmp
		return err
	case "autovacuum_freeze_max_age":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumFreezeMaxAge = &tmp
		return err
	case "autovacuum_freeze_table_age":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumFreezeTableAge = &tmp
		return err
	case "autovacuum_multixact_freeze_min_age":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumMultixactFreezeMinAge = &tmp
		return err
	case "autovacuum_multixact_freeze_max_age":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumMultixactFreezeMaxAge = &tmp
		return err
	case "autovacuum_multixact_freeze_table_age":
		tmp, err := strconv.Atoi(value)
		p.AutovacuumMultixactFreezeTableAge = &tmp
		return err
	case "log_autovacuum_min_duration":
		tmp, err := strconv.Atoi(value)
		p.LogAutovacuumMinDuration = &tmp
		return err
	case "user_catalog_table":
		tmp, err := strconv.ParseBool(value)
		p.UserCatalogTable = &tmp
		return err
	}
	return nil
}
