package server

import (
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/issues"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/rest/projects"
	"github.com/DevBoxFanBoy/opists/pkg/config"
	"github.com/DevBoxFanBoy/opists/pkg/security"
	uiResources "github.com/DevBoxFanBoy/opists/pkg/ui/resources"
	"github.com/casbin/casbin/v2"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
)

func newIssueTrackingSystemRouterV1(config config.Config) *mux.Router {
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

func NewIssueTrackingSystem(config config.Config) *negroni.Negroni {
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewLogger())
	if config.Opists.Security.Enabled {
		n.Use(security.NewAuthenticationMiddleWare(config))
		if config.Opists.Security.EnableUserLogging {
			n.Use(negroni.HandlerFunc(security.LogCurrentUser))
		}
		e, err := casbin.NewEnforcer(config.Opists.Security.AuthzModelFilePath, config.Opists.Security.AuthzPolicyFilePath)
		if err != nil {
			log.Fatalf("could not create casbin enforcer: %v", err)
		}
		n.Use(security.Authorizer(e))
	}
	n.UseHandler(newIssueTrackingSystemRouterV1(config))
	return n
}

func NewIssueTrackingSystemServer(config config.Config) {
	address := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
	n := NewIssueTrackingSystem(config)
	n.Run(address)
}
