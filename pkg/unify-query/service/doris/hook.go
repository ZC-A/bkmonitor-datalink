package doris

import (
	"fmt"
	"time"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/eventbus"
	"github.com/spf13/viper"
)

func setDefaultConfig() {
	viper.SetDefault(CharSetConfigPath, "utf8mb4")
	viper.SetDefault(LocalConfigPath, "Local")
	viper.SetDefault(MaxIdleConnsConfigPath, 5)
	viper.SetDefault(MaxOpenConnsConfigPath, 10)
	viper.SetDefault(DriverNameConfigPath, "mysql")
	viper.SetDefault(MaxIdleTimeConfigPath, 30*60)
	viper.SetDefault(MaxLifetimeConfigPath, 60*60)
	viper.SetDefault(DisableMetricConfigPath, false)
	viper.SetDefault(DisableTraceConfigPath, false)
}

func LoadConfig() {
	Username = viper.GetString(UsernameConfigPath)
	Host = viper.GetString(HostConfigPath)
	Port = viper.GetString(PortConfigPath)
	Password = viper.GetString(PasswordConfigPath)
	DataBase = viper.GetString(DataBaseConfigPath)
	DriverName = viper.GetString(DriverNameConfigPath)
	Local = viper.GetString(LocalConfigPath)
	CharSet = viper.GetString(CharSetConfigPath)
	MaxOpenConns = viper.GetInt(MaxIdleConnsConfigPath)
	MaxIdleTime = time.Duration(viper.GetInt(MaxIdleTimeConfigPath))
	MaxIdleConns = viper.GetInt(MaxIdleConnsConfigPath)
	MaxLifetime = time.Duration(viper.GetInt(MaxLifetimeConfigPath))
	DisableMetric = viper.GetBool(DisableMetricConfigPath)
	DisableTrace = viper.GetBool(DisableTraceConfigPath)
}

func init() {
	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPreParse, setDefaultConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for http module for default config, maybe http module won't working.",
			eventbus.EventSignalConfigPreParse,
		)
	}

	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPostParse, LoadConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for http module for new config, maybe http module won't working.",
			eventbus.EventSignalConfigPostParse,
		)
	}
}
