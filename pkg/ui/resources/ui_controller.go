package projects

import (
	"errors"
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"github.com/DevBoxFanBoy/opists/pkg/business/usecase/issue"
	"github.com/DevBoxFanBoy/opists/pkg/business/usecase/project"
	"github.com/DevBoxFanBoy/opists/pkg/ui"
	"github.com/DevBoxFanBoy/opists/pkg/ui/projects"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// A UIController binds http requests to an UI service and writes the service results to the http response
type UIController struct {
	projectUseCase project.UseCase
	issueUseCase   issue.UseCase
}

func NewUIController() router.Router {
	p := project.GetUseCaseControllerInstance()
	i := issue.GetUseCaseControllerInstance()
	return &UIController{projectUseCase: p, issueUseCase: i}
}

// Routes returns all of the api route for the UIController
func (u *UIController) Routes() router.Routes {
	return router.Routes{
		{
			"GetIndex",
			strings.ToUpper("Get"),
			"/index.html",
			u.GetIndex,
		},
		{
			"GetAllProjects",
			strings.ToUpper("Get"),
			"/projects.html",
			u.GetIndex,
		},
		{
			"CreateProject",
			strings.ToUpper("Post"),
			"/projects.html",
			u.CreateProject,
		},
		{
			"GetIssuesFromProject",
			strings.ToUpper("Get"),
			"/{projectKey}/issues.html",
			u.GetIssuesFromProject,
		},
		{
			"GetJSResource",
			strings.ToUpper("Get"),
			"/js/{resource}",
			u.GetJSResource,
		},
		{
			"GetCSSResource",
			strings.ToUpper("Get"),
			"/css/{resource}",
			u.GetCSSResource,
		},
		{
			"GetAssetsResource",
			strings.ToUpper("Get"),
			"/assets/{folder}/{resource}",
			u.GetAssetsResource,
		},
	}
}

func (u *UIController) GetIndex(w http.ResponseWriter, r *http.Request) {
	projects.GetAllProject(w, r, u.projectUseCase)
}

func (u *UIController) CreateProject(w http.ResponseWriter, r *http.Request) {
	projects.CreateProjectForm(w, r, u.projectUseCase)
}

func (u *UIController) GetIssuesFromProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectKey := params["projectKey"]
	if err := validateParameter(projectKey, "ProjectKey"); err != nil {
		ui.HardNotFound(w)
		return
	}
	df, err := u.projectUseCase.GetProject(projectKey)
	if err != nil {
		t, e := template.ParseFiles("ui/management/500.html")
		if e != nil {
			ui.HardInternalServerError(w, e)
			return
		}
		t.Execute(w, struct{ ErrorStr string }{e.Error()})
		return
	}
	issues, err := u.issueUseCase.GetProjectIssues(projectKey)
	if err != nil {
		t, e := template.ParseFiles("ui/management/500.html")
		if e != nil {
			ui.HardInternalServerError(w, e)
			return
		}
		t.Execute(w, struct{ ErrorStr string }{e.Error()})
		return
	}
	is := issues.(model.Issues)
	data := struct {
		Project     model.Project
		Issues      model.Issues
		IssuesCount int
	}{Project: df.(model.Project), Issues: is, IssuesCount: len(is.Issues)}
	t, err := template.ParseFiles("ui/management/issues.html")
	if err != nil {
		ui.HardInternalServerError(w, err)
		return
	}
	t.Execute(w, data)
}

func (u *UIController) GetJSResource(w http.ResponseWriter, r *http.Request) {
	loadResource(w, r, "js", "application/javascript")
}

func (u *UIController) GetCSSResource(w http.ResponseWriter, r *http.Request) {
	loadResource(w, r, "css", "text/css")
}

func (u *UIController) GetAssetsResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	folder := params["folder"]
	if err := validateResource(folder); err != nil {
		ui.HardNotFound(w)
		return
	}
	contentType := ""
	if reflect.DeepEqual("demo", folder) {
		contentType = "application/javascript"
	}
	assets := fmt.Sprintf("assets/%v", folder)
	loadResource(w, r, assets, contentType)
}

func validateParameter(parameter string, what string) error {
	if len(parameter) == 0 {
		err := errors.New(fmt.Sprintf("%v %v is invalid!", what, parameter))
		return err
	}
	return nil
}

func validateResource(resource string) error {
	return validateParameter(resource, "Resource")
}

func loadResource(w http.ResponseWriter, r *http.Request, folder string, contentType string) {
	params := mux.Vars(r)
	resource := params["resource"]
	if err := validateResource(resource); err != nil {
		ui.HardNotFound(w)
		return
	}
	if err := validateResource(folder); err != nil {
		ui.HardFolderNotFound(w, folder)
		return
	}
	filename := fmt.Sprintf("ui/management/%v/%v", folder, resource)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		ui.HardInternalServerError(w, err)
		return
	}
	if len(contentType) > 0 {
		w.Header().Add("Content-Type", contentType)
	}
	fmt.Fprintf(w, "%s", body)
}
