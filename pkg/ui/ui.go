package ui

import (
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"net/http"
)

// ProjectsUIServicer defines the ui actions for the ProjectsApi service
type ProjectsUIServicer interface {
	GetAllProject() (interface{}, error)
	CreateProject() (interface{}, error)
	CreateProjectForm(project model.Project) (interface{}, error)
}

func HardInternalServerError(w http.ResponseWriter, err error) {
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "500 Internal Server Error", err.Error())
}

func HardNotFound(w http.ResponseWriter) {
	fmt.Fprintf(w, "<h1>%s</h1>", "404 Not Found")
}

func HardFolderNotFound(w http.ResponseWriter, folder string) {
	fmt.Fprintf(w, "<h1>%s</h1><div>Folder: %s</div>", "404 Not Found", folder)
}
