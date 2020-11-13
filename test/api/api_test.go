package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/server"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

type apiFeature struct {
	resp      *httptest.ResponseRecorder
	apiRouter *mux.Router
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
}

func (a *apiFeature) iSendRequestTo(method, endpoint string) (err error) {
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// handle panic
	defer func() {
		switch t := recover().(type) {
		case string:
			err = fmt.Errorf(t)
		case error:
			err = t
			fmt.Println(err.Error())
		}
	}()

	a.apiRouter.ServeHTTP(a.resp, req)
	return
}

func (a *apiFeature) theResponseCodeShouldBe(code int) error {
	if code != a.resp.Code {
		return fmt.Errorf("expected response code to be: %d, but actual is: %d", code, a.resp.Code)
	}
	return nil
}

func (a *apiFeature) theResponseShouldMatchJSON(body *godog.DocString) (err error) {
	var expected, actual interface{}

	// re-encode expected response
	if err = json.Unmarshal([]byte(body.Content), &expected); err != nil {
		return
	}

	// re-encode actual response too
	if err = json.Unmarshal(a.resp.Body.Bytes(), &actual); err != nil {
		return
	}

	// the matching may be adapted per different requirements.
	if !reflect.DeepEqual(expected, actual) {
		return fmt.Errorf("expected JSON does not match actual, \nexpect: %v vs. \nactual: %v", expected, actual)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	apiFeature := &apiFeature{}
	apiFeature.apiRouter = server.NewIssueTrackingSystemRouterV1()

	ctx.BeforeScenario(apiFeature.resetResponse)

	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, apiFeature.iSendRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, apiFeature.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, apiFeature.theResponseShouldMatchJSON)
}

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	status := godog.TestSuite{
		Name:                "projects",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
