package doris

import "time"

const (
	UsernameConfigPath      = "doris.username"
	HostConfigPath          = "doris.host"
	PortConfigPath          = "doris.port"
	PasswordConfigPath      = "doris.password"
	DataBaseConfigPath      = "doris.database"
	DriverNameConfigPath    = "doris.driver_name"
	CharSetConfigPath       = "doris.charset"
	LocalConfigPath         = "doris.local"
	MaxOpenConnsConfigPath  = "doris.max_open_conns"
	MaxIdleConnsConfigPath  = "doris.max_idle_conns"
	MaxLifetimeConfigPath   = "doris.max_life_time"
	MaxIdleTimeConfigPath   = "doris.max_idle_time"
	DisableMetricConfigPath = "doris.disable_metric"
	DisableTraceConfigPath  = "doris.disable_trace"
)

var (
	Username      string
	Host          string
	Port          string
	Password      string
	DataBase      string
	DriverName    string
	CharSet       string
	Local         string
	MaxOpenConns  int
	MaxIdleConns  int
	MaxLifetime   time.Duration
	MaxIdleTime   time.Duration
	DisableMetric bool
	DisableTrace  bool
)
