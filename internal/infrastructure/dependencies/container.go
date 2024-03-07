package dependencies

import (
	"github.com/abaykerimov/test_kmf/internal/app/services"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/clients/nat_api"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/providers/db"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/repositories"
	"runtime"
)

type Container struct {
	DB         *db.DB
	Client     *nat_api.Client
	Repository *repositories.Repository
	Service    *services.Service
}

func NewDI() *Container {
	container := new(Container).Init()

	runtime.SetFinalizer(container, func(cf *Container) {
		cf.DB.Close()
	})

	return container
}

func (d *Container) Init() *Container {
	d.DB = new(db.DB).Initialize()
	d.Client = nat_api.NewClient()
	d.Repository = repositories.NewRepository(d.DB)
	d.Service = services.NewService(d.Repository, d.Client)

	return d
}
