package server

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewIssueTrackingSystemRouterV1() *mux.Router {
	IssuesApiService := service.NewIssuesApiService()
	IssuesApiController := rest.NewIssuesApiController(IssuesApiService)

	ProjectsApiService := service.NewProjectsApiService()
	ProjectsApiController := rest.NewProjectsApiController(ProjectsApiService)

	return router.NewRouter(IssuesApiController, ProjectsApiController)
}

func NewIssueTrackingSystemServer() {
	log.Fatal(http.ListenAndServe(":8080", NewIssueTrackingSystemRouterV1()))
}
