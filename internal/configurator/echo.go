// Package configurator provides dependency injection definitions.
package configurator

import (
	viperBundle "github.com/gozix/viper/v2"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/sarulabs/di/v2"
)

// DefEchoConfiguratorName is a definition name.
const DefEchoConfiguratorName = "echo.configurator.echo"

// DefEchoConfigurator is echo custom configurator definition getter.
func DefEchoConfigurator() di.Def {
	return di.Def{
		Name: DefEchoConfiguratorName,
		Tags: []di.Tag{{
			Name: TagConfigurator,
		}},
		Build: func(ctn di.Container) (interface{}, error) {
			return func(e *echo.Echo) (err error) {
				var cfg *viperBundle.Viper
				if err = ctn.Fill(viperBundle.BundleName, &cfg); err != nil {
					return err
				}

				if cfg.IsSet("echo.debug") {
					e.Debug = cfg.GetBool("echo.debug")
				}

				switch cfg.GetString("echo.level") {
				case "debug":
					e.Logger.SetLevel(log.DEBUG)
				case "info":
					e.Logger.SetLevel(log.INFO)
				case "warn":
					e.Logger.SetLevel(log.WARN)
				case "error":
					e.Logger.SetLevel(log.ERROR)
				case "off":
					e.Logger.SetLevel(log.OFF)
				}

				if cfg.IsSet("echo.static") {
					e.Static(
						cfg.GetString("echo.static.prefix"),
						cfg.GetString("echo.static.root"),
					)
				}

				if cfg.IsSet("echo.hide_banner") {
					e.HideBanner = cfg.GetBool("echo.hide_banner")
				}

				if cfg.IsSet("echo.hide_port") {
					e.HidePort = cfg.GetBool("echo.hide_port")
				}

				return nil
			}, nil
		},
	}
}
