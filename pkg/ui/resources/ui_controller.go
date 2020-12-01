package projects

import (
	"errors"
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/router"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// A UIController binds http requests to an UI service and writes the service results to the http response
type UIController struct {
}

func NewUIController() router.Router {
	return &UIController{}
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
	body, err := ioutil.ReadFile("ui/management/index.html")
	if err != nil {
		HardInternalServerError(w, err)
		return
	}
	fmt.Fprintf(w, "%s", body)
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
		HardNotFound(w)
		return
	}
	contentType := ""
	if reflect.DeepEqual("demo", folder) {
		contentType = "application/javascript"
	}
	assets := fmt.Sprintf("assets/%v", folder)
	loadResource(w, r, assets, contentType)
}

func validateResource(resource string) error {
	if len(resource) == 0 {
		err := errors.New(fmt.Sprintf("Resource %v is invalid!", resource))
		return err
	}
	return nil
}

func loadResource(w http.ResponseWriter, r *http.Request, folder string, contentType string) {
	params := mux.Vars(r)
	resource := params["resource"]
	if err := validateResource(resource); err != nil {
		HardNotFound(w)
		return
	}
	if err := validateResource(folder); err != nil {
		HardFolderNotFound(w, folder)
		return
	}
	filename := fmt.Sprintf("ui/management/%v/%v", folder, resource)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		HardInternalServerError(w, err)
		return
	}
	if len(contentType) > 0 {
		w.Header().Add("Content-Type", contentType)
	}
	fmt.Fprintf(w, "%s", body)
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
