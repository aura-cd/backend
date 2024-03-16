package router

import (
	"net/http"

	usage "github.com/aura-cd/backend/internal/adapters/controllers/application/usage"
	"github.com/aura-cd/backend/internal/router/middleware"
	usecase "github.com/aura-cd/backend/internal/usecases/application/usage"
)

func (r *router) Usage() {
	uc := usage.New(
		usecase.New(r.controllerConn),
	)

	r.mux.Handle("GET /organizations/{organization}/deployments/{deploymentName}/usage", middleware.BuildChain(
		r.middleware.Integrate(
			http.HandlerFunc(uc.Usage),
		),
	))
}
