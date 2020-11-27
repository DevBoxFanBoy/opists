package project

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCaseController_CreateProject(t *testing.T) {
	underTest := UseCaseController{make(map[string]model.Project)}

	result, err := underTest.CreateProject(model.Project{
		Key:         "A",
		Name:        "B",
		Description: "C",
		Versions:    nil,
		Components:  nil,
		Sprints:     nil,
	})
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	project := result.(model.Project)
	assert.Equal(t, "A", project.Key)
	assert.Equal(t, "B", project.Name)
	assert.Equal(t, "C", project.Description)
	assert.Nil(t, project.Versions)
	assert.Nil(t, project.Components)
	assert.Nil(t, project.Sprints)
}

func TestUseCaseController_CreateProject_ConflictProjectAlreadyExists(t *testing.T) {
	underTest := UseCaseController{map[string]model.Project{
		"B": {
			Key:         "B",
			Name:        "",
			Description: "",
			Versions:    nil,
			Components:  nil,
			Sprints:     nil,
		},
	}}
	result, err := underTest.CreateProject(model.Project{
		Key:         "B",
		Name:        "Other",
		Description: "",
		Versions:    nil,
		Components:  nil,
		Sprints:     nil,
	})
	if err == nil {
		t.Error("Test failed: expect error is not nil.")
		return
	}
	errorResponse := result.(model.ErrorResponse)
	assert.Equal(t, "Project with Key B already exists!", err.Error())
	assert.Equal(t, int32(409), errorResponse.Code)
	assert.Equal(t, "Project with Key B already exists!", errorResponse.Message)
}

func TestUseCaseController_GetProject(t *testing.T) {
	underTest := UseCaseController{map[string]model.Project{
		"B": {
			Key:         "B",
			Name:        "",
			Description: "",
			Versions:    nil,
			Components:  nil,
			Sprints:     nil,
		},
	}}
	result, err := underTest.GetProject("B")
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	project := result.(model.Project)
	assert.Equal(t, "B", project.Key)
	assert.Equal(t, "", project.Name)
	assert.Equal(t, "", project.Description)
	assert.Nil(t, project.Versions)
	assert.Nil(t, project.Components)
	assert.Nil(t, project.Sprints)
}

func TestUseCaseController_GetProject_NotFound(t *testing.T) {
	underTest := UseCaseController{map[string]model.Project{}}
	result, err := underTest.GetProject("B")
	if err == nil {
		t.Error("Test failed: expect error is not nil.")
		return
	}
	errorResponse := result.(model.ErrorResponse)
	assert.Equal(t, "Project with Key B not found!", err.Error())
	assert.Equal(t, int32(404), errorResponse.Code)
	assert.Equal(t, "Project with Key B not found!", errorResponse.Message)
}

func TestUseCaseController_GetAllProject(t *testing.T) {
	underTest := UseCaseController{map[string]model.Project{
		"B": {
			Key:         "B",
			Name:        "A",
			Description: "C",
			Versions:    nil,
			Components:  nil,
			Sprints:     nil,
		},
	}}
	result, err := underTest.GetAllProject()
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	allProjects := result.(model.Projects)
	assert.Equal(t, 1, len(allProjects.Projects))
	assert.NotNil(t, allProjects)
	assert.NotNil(t, allProjects.Projects[0])
	assert.Equal(t, "B", allProjects.Projects[0].Key)
	assert.Equal(t, "A", allProjects.Projects[0].Name)
	assert.Equal(t, "C", allProjects.Projects[0].Description)
	assert.Nil(t, allProjects.Projects[0].Versions)
	assert.Nil(t, allProjects.Projects[0].Components)
	assert.Nil(t, allProjects.Projects[0].Sprints)
}

func TestUseCaseController_GetAllProject_Empty(t *testing.T) {
	underTest := UseCaseController{map[string]model.Project{}}
	result, err := underTest.GetAllProject()
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	allProjects := result.(model.Projects)
	assert.Equal(t, 0, len(allProjects.Projects))
}
