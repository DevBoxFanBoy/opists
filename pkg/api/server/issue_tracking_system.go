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

func NewIssueTrackingSystemRouterV1() *mux.Router {
	IssuesApiService := issues.NewApiService()
	IssuesApiController := issues.NewApiController(IssuesApiService)

	ProjectsApiService := projects.NewApiService()
	ProjectsApiController := projects.NewApiController(ProjectsApiService)

	r := router.NewRouter(IssuesApiController, ProjectsApiController)
	router.AddRoutes(r, "/ui/management", uiResources.NewUIController())
	router.AddFaviconRoute(r)
	return r
}

func NewIssueTrackingSystemServer(config config.Config) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port),
		NewIssueTrackingSystemRouterV1()))
}
