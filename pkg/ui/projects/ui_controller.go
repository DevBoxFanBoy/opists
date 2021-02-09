package projects

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"github.com/DevBoxFanBoy/opists/pkg/business/usecase/project"
	"github.com/DevBoxFanBoy/opists/pkg/ui"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

//// A UIController binds http requests to an UI service and writes the service results to the http response
//type UIController struct {
//	projectUseCase project.UseCase
//}

//func NewUIController() UIController {
//	p := project.GetUseCaseControllerInstance()
//	return &UIController{projectUseCase: p}
//}

// Routes returns all of the api route for the UIController
//func (u *UIController) Routes() router.Routes {
//	return router.Routes{
//		{
//			"GetAllProject",
//			strings.ToUpper("Get"),
//			"/projects.xhtml",
//			u.GetAllProject,
//		},
//		{
//			"CreateProject",
//			strings.ToUpper("Get"),
//			"/create/projects.xhtml",
//			u.CreateProject,
//		},
//		{
//			"CreateProjectForm",
//			strings.ToUpper("Post"),
//			"/create/projects.xhtml",
//			u.CreateProjectForm,
//		},
//	}
//}

func GetAllProject(w http.ResponseWriter, r *http.Request, u project.UseCase) {
	allPrjs, err := u.GetAllProject()
	if err != nil {
		t, e := template.ParseFiles("ui/management/500.html")
		if e != nil {
			ui.HardInternalServerError(w, e)
			return
		}
		t.Execute(w, struct{ ErrorStr string }{e.Error()})
		return
	}
	prjs := allPrjs.(model.Projects)
	data := struct {
		Projects     model.Projects
		ProjectCount int
	}{Projects: prjs, ProjectCount: len(prjs.Projects)}
	t, err := template.ParseFiles("ui/management/index.html")
	if err != nil {
		ui.HardInternalServerError(w, err)
		return
	}
	t.Execute(w, data)
}

func CreateProject(w http.ResponseWriter, r *http.Request) {
	//TODO maybe ajax call?
	params := mux.Vars(r)
	projectKey := params["projectKey"]

	router.EncodeJSONResponse(projectKey, nil, w)
}

func CreateProjectForm(w http.ResponseWriter, r *http.Request, u project.UseCase) {
	r.ParseForm()
	projectKey := r.FormValue("projectKey")
	if len(r.FormValue("name")) == 0 {
		t, err := template.ParseFiles("ui/management/500.html")
		if err != nil {
			ui.HardInternalServerError(w, err)
			return
		}
		t.Execute(w, struct{ ErrorStr string }{err.Error()})
		return
	}
	p := model.Project{
		Key:         projectKey,
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Versions:    make([]string, 0),
		Components:  make([]model.Component, 0),
		Sprints:     make([]model.Sprint, 0),
	}
	if _, err := u.CreateProject(p); err != nil {
		t, err := template.ParseFiles("ui/management/500.html")
		if err != nil {
			ui.HardInternalServerError(w, err)
			return
		}
		t.Execute(w, struct{ ErrorStr string }{err.Error()})
		return
	}
	GetAllProject(w, r, u)
}
