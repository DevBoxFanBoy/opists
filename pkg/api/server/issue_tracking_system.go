package server

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/issues"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/projects"
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
	return r
}

func NewIssueTrackingSystemServer() {
	log.Fatal(http.ListenAndServe(":8080", NewIssueTrackingSystemRouterV1()))
}
