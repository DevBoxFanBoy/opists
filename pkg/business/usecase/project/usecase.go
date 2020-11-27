package project

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"time"
)

type UseCase struct {
}

func (p *UseCase) GetAllProject() (interface{}, error) {
	prj, err := createProjectModel("DF")
	if err != nil {
		return nil, err
	}
	return model.Projects{Projects: []model.Project{prj}}, nil
}

func (p *UseCase) GetProject(projectKey string) (interface{}, error) {
	return createProjectModel(projectKey)
}

func createProjectModel(projectKey string) (model.Project, error) {
	var start, err = time.Parse(time.RFC3339, "2020-11-12T07:00:34.700Z")
	if err != nil {
		return model.Project{}, err
	}
	var end, endErr = time.Parse(time.RFC3339, "2020-11-26T15:18:36.330Z")
	if endErr != nil {
		return model.Project{}, endErr
	}
	return model.Project{
		Key:         projectKey,
		Name:        "DogFooding",
		Description: "The Project used intern for Development.",
		Versions:    []string{"1.2.3", "1.2.4"},
		Components: []model.Component{{
			Name:        "DrinkOwnChampagne",
			Description: "Used intern for Development.",
			Versions:    []string{"DOC 1.0.0", "DOC 1.0.1"},
		}},
		Sprints: []model.Sprint{{
			Key:   "Sprint2",
			Name:  "Sprint 2 - Consume DogFooding",
			Start: start,
			End:   end,
		}},
	}, nil
}
