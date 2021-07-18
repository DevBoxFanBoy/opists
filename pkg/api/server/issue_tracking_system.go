package server

import (
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/issues"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/projects"
	"github.com/DevBoxFanBoy/opists/pkg/config"
	uiResources "github.com/DevBoxFanBoy/opists/pkg/ui/resources"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewIssueTrackingSystemRouterV1(config config.Config) *mux.Router {
	IssuesApiService := issues.NewApiService()
	IssuesApiController := issues.NewApiController(IssuesApiService)

	ProjectsApiService := projects.NewApiService()
	ProjectsApiController := projects.NewApiController(ProjectsApiService)

	r := router.NewRouter(IssuesApiController, ProjectsApiController)
	router.AddRoutes(r, "/ui/management", uiResources.NewUIController())
	router.AddFaviconRoute(r)
	if config.Opists.PProfOn {
		router.AddPProf(r)
	}
	return r
}

func NewIssueTrackingSystemServer(config config.Config) error {
	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	log.Println(fmt.Sprintf("Server bind to address: %s", address))
	err := http.ListenAndServe(address, NewIssueTrackingSystemRouterV1(config))
	return err
}
