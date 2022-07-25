package api

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/server"
	"github.com/DevBoxFanBoy/opists/pkg/config"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/urfave/negroni"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

type currentUser struct {
	username string
	password string
}

type apiFeature struct {
	user      *currentUser
	resp      *httptest.ResponseRecorder
	apiRouter *negroni.Negroni
	config    config.Config
}

func (a *apiFeature) resetResponse(*godog.Scenario) {
	a.resp = httptest.NewRecorder()
	a.user = nil
}

func (a *apiFeature) iSendRequestTo(method, endpoint string, body string) (err error) {
	req, err := http.NewRequest(method, endpoint, requestBody(body))
	if a.user != nil {
		req.SetBasicAuth(strings.TrimSpace(a.user.username), a.user.password)
	}
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

func requestBody(body string) io.Reader {
	if len(body) == 0 {
		return nil
	} else {
		return strings.NewReader(body)
	}
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

func (a *apiFeature) theResponseHeaderShouldMatch(headerName string, expectedValue string) (err error) {
	actual := a.resp.Header().Get(headerName)
	if !reflect.DeepEqual(expectedValue, actual) {
		return fmt.Errorf("expected header %v does not match actual, \nexpect: %v vs. \nactual: %v", headerName, expectedValue, actual)
	}
	return nil
}

func (a *apiFeature) asUser(user string) error {
	a.user = &currentUser{strings.TrimSpace(user), a.config.Opists.Security.AdminPassword}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	var cfg config.Config
	if err := cleanenv.ReadConfig("testconfig.yml", &cfg); err != nil {
		fmt.Printf("Server config error: %v\n", err)
		os.Exit(1)
	}
	apiFeature := &apiFeature{config: cfg}
	apiFeature.apiRouter = server.NewIssueTrackingSystem(cfg)

	ctx.BeforeScenario(apiFeature.resetResponse)

	//TODO Given ctx.Step(`^there is ...`
	ctx.Step(`^As "([^"]*)" User$`, apiFeature.asUser)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)" with body '([^']*)'$`, apiFeature.iSendRequestTo)
	ctx.Step(`^the response code should be (\d+)$`, apiFeature.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, apiFeature.theResponseShouldMatchJSON)
	ctx.Step(`^the response header "([^"]*)" match value "([^"]*)"$`, apiFeature.theResponseHeaderShouldMatch)
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
		Name:                "all",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
