package project

import (
	"errors"
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"sync"
	"time"
)

type UseCase interface {
	GetAllProject() (interface{}, error)
	GetProject(string) (interface{}, error)
	CreateProject(project model.Project) (interface{}, error)
	UpdateProject(projectKey string, project model.Project) (interface{}, error)
	DeleteProject(projectKey string) (interface{}, error)
}

var once sync.Once
var instance UseCaseController

type UseCaseController struct {
	projects map[string]model.Project
}

func GetUseCaseControllerInstance() UseCase {
	once.Do(func() {
		project, _ := createProjectModel("DF")
		projects := map[string]model.Project{
			"DF": project,
		}
		instance = UseCaseController{projects: projects}
	})
	return &instance
}

func (u *UseCaseController) GetAllProject() (interface{}, error) {
	var projectsModel = model.Projects{}
	for _, element := range u.projects {
		projectsModel.Projects = append(projectsModel.Projects, element)
	}
	return projectsModel, nil
}

func (u *UseCaseController) GetProject(projectKey string) (interface{}, error) {
	project, ok := u.projects[projectKey]
	if !ok {
		err := errors.New(fmt.Sprintf("Project with Key %v not found!", projectKey))
		return model.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		}, err
	}
	return project, nil
}

func (u *UseCaseController) CreateProject(project model.Project) (interface{}, error) {
	if _, ok := u.projects[project.Key]; !ok {
		u.projects[project.Key] = project
	} else {
		err := errors.New(fmt.Sprintf("Project with Key %v already exists!", project.Key))
		return model.ErrorResponse{
			Code:    409,
			Message: err.Error(),
		}, err
	}
	return project, nil
}

func (u *UseCaseController) UpdateProject(projectKey string, project model.Project) (interface{}, error) {
	if projectKey != project.Key {
		err := errors.New(fmt.Sprintf("Project's Key %v is not equal to %v!", project.Key, projectKey))
		return model.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		}, err
	}
	if _, ok := u.projects[project.Key]; !ok {
		err := errors.New(fmt.Sprintf("Project with Key %v not found!", project.Key))
		return model.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		}, err
	}
	u.projects[project.Key] = project
	return nil, nil
}

func (u *UseCaseController) DeleteProject(projectKey string) (interface{}, error) {
	_, ok := u.projects[projectKey]
	if !ok {
		err := errors.New(fmt.Sprintf("Project with Key %v not found!", projectKey))
		return model.ErrorResponse{
			Code:    404,
			Message: err.Error(),
		}, err
	}
	delete(u.projects, projectKey)
	//TODO delete all issues of this project, too
	return nil, nil
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
