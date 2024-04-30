package server

import (
	config "app/src/common/config"
	response "app/src/common/responses"
	types "app/src/common/types"
	db "app/src/infraestructure/db/adapter"
	"go.uber.org/fx"
)

type ProvidersStore struct {
	Providers []fx.Option
}

func (ps *ProvidersStore) Init() {
	ps.Providers = []fx.Option{
		fx.Provide(types.NewHandlersStore),
		fx.Provide(response.NewResult),
		fx.Provide(config.NewConfig),
		fx.Provide(db.NewDBConnection),
	}
}
func (ps *ProvidersStore) AddModule(p []fx.Option) {
	ps.Providers = append(ps.Providers, p...)
}

func (ps *ProvidersStore) Up(lp ...[]fx.Option) {
	ps.Providers = append(ps.Providers, fx.Invoke(NewHttpFiberServer))
	fx.New(ps.Providers...).Run()
}
