//go:build wireinject
// +build wireinject

package main

import (
	"service_api/internal/auth"
	"service_api/internal/cache"
	"service_api/internal/components"
	"service_api/internal/components_admin"
	"service_api/internal/controller"
	"service_api/internal/persist"
	"service_api/internal/server"
	"service_api/internal/stop_list"
	"service_api/internal/validator"

	queue "bitbucket.org/klaraeng/package_queue/ampq"

	conf "bitbucket.org/klaraeng/package_config/conf"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init application.
func wireApp(
	*conf.Server,
	*conf.Data,
	*conf.Params,
	log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		validator.ProviderSet,
		cache.ProviderSet,
		stop_list.ProviderSet,
		persist.ProviderSet,
		components.ProviderSet,
		auth.ProviderSet,
		queue.ProviderSet,
		components_admin.ProviderSet,
		controller.ProviderSet,
		newApp))
}
