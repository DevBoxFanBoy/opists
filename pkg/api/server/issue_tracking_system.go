package server

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/issues"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/projects"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewIssueTrackingSystemRouterV1() *mux.Router {
	IssuesApiService := issues.NewApiService()
	IssuesApiController := issues.NewApiController(IssuesApiService)

	ProjectsApiService := projects.NewApiService()
	ProjectsApiController := projects.NewApiController(ProjectsApiService)

	return router.NewRouter(IssuesApiController, ProjectsApiController)
}

func NewIssueTrackingSystemServer() {
	log.Fatal(http.ListenAndServe(":8080", NewIssueTrackingSystemRouterV1()))
}
