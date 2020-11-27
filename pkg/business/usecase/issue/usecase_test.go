package issue

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUseCaseController_AddIssue(t *testing.T) {
	underTest := UseCaseController{issues: make(map[string]map[int64]model.Issue)}
	id := int64(0)
	points := int32(0)
	input := model.Issue{
		Id:              &id,
		Name:            "B",
		Description:     "C",
		Status:          "",
		Priority:        "",
		ProjectKey:      "B",
		Components:      nil,
		Sprints:         nil,
		EstimatedPoints: &points,
		EstimatedTime:   "",
		AffectedVersion: "",
		FixedVersion:    "",
	}
	location, err := underTest.AddIssue("A", input)
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
	}
	assert.Equal(t, "A/0", location)
}

func TestUseCaseController_AddIssue_NotFoundProject(t *testing.T) {
	underTest := UseCaseController{issues: make(map[string]map[int64]model.Issue)}
	id := int64(0)
	points := int32(0)
	input := model.Issue{
		Id:              &id,
		Name:            "B",
		Description:     "C",
		Status:          "",
		Priority:        "",
		ProjectKey:      "B",
		Components:      nil,
		Sprints:         nil,
		EstimatedPoints: &points,
		EstimatedTime:   "",
		AffectedVersion: "",
		FixedVersion:    "",
	}
	location, err := underTest.AddIssue("A", input)
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
	}
	assert.Equal(t, "A/0", location)
}

func TestUseCaseController_GetIssueById(t *testing.T) {
	id := int64(0)
	points := int32(0)
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{
		"B": {
			0: model.Issue{
				Id:              &id,
				Name:            "B",
				Description:     "C",
				Status:          "",
				Priority:        "",
				ProjectKey:      "B",
				Components:      nil,
				Sprints:         nil,
				EstimatedPoints: &points,
				EstimatedTime:   "",
				AffectedVersion: "",
				FixedVersion:    "",
			},
		},
	}}
	result, err := underTest.GetIssueById("B", 0)
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	issue := result.(model.Issue)
	assert.Equal(t, int64(0), *issue.Id)
	assert.Equal(t, "B", issue.Name)
	assert.Equal(t, "C", issue.Description)
	assert.Equal(t, "", issue.Status)
	assert.Equal(t, "", issue.Priority)
	assert.Equal(t, "B", issue.ProjectKey)
	assert.Nil(t, issue.Components)
	assert.Nil(t, issue.Sprints)
	assert.Equal(t, int32(0), *issue.EstimatedPoints)
	assert.Equal(t, "", issue.EstimatedTime)
	assert.Equal(t, "", issue.AffectedVersion)
	assert.Equal(t, "", issue.FixedVersion)
}

func TestUseCaseController_GetIssueById_NotFoundProject(t *testing.T) {
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{}}
	result, err := underTest.GetIssueById("B", 0)
	if err == nil {
		t.Error("Test failed: expect error is not nil.")
		return
	}
	errorResponse := result.(model.ErrorResponse)
	assert.Equal(t, "Project with Key B not found!", err.Error())
	assert.Equal(t, int32(404), errorResponse.Code)
	assert.Equal(t, "Project with Key B not found!", errorResponse.Message)
}

func TestUseCaseController_GetIssueById_NotFound(t *testing.T) {
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{"B": {}}}
	result, err := underTest.GetIssueById("B", 0)
	if err == nil {
		t.Error("Test failed: expect error is not nil.")
		return
	}
	errorResponse := result.(model.ErrorResponse)
	assert.Equal(t, "Issue with ID 0 not found!", err.Error())
	assert.Equal(t, int32(404), errorResponse.Code)
	assert.Equal(t, "Issue with ID 0 not found!", errorResponse.Message)
}

func TestUseCaseController_DeleteIssue(t *testing.T) {
	id := int64(0)
	points := int32(0)
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{
		"B": {
			0: model.Issue{
				Id:              &id,
				Name:            "B",
				Description:     "C",
				Status:          "",
				Priority:        "",
				ProjectKey:      "B",
				Components:      nil,
				Sprints:         nil,
				EstimatedPoints: &points,
				EstimatedTime:   "",
				AffectedVersion: "",
				FixedVersion:    "",
			},
		},
	}}
	_, err := underTest.DeleteIssue("B", 0)
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	assert.Equal(t, 0, len(underTest.issues["B"]))
}

func TestUseCaseController_DeleteIssue_NotFoundProject(t *testing.T) {
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{}}
	result, err := underTest.DeleteIssue("B", 0)
	if err == nil {
		t.Error("Test failed: expect error is not nil.")
		return
	}
	errorResponse := result.(model.ErrorResponse)
	assert.Equal(t, "Project with Key B not found!", err.Error())
	assert.Equal(t, int32(404), errorResponse.Code)
	assert.Equal(t, "Project with Key B not found!", errorResponse.Message)
}

func TestUseCaseController_DeleteIssue_NotFound(t *testing.T) {
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{"B": {}}}
	result, err := underTest.DeleteIssue("B", 0)
	if err == nil {
		t.Error("Test failed: expect error is not nil.")
		return
	}
	errorResponse := result.(model.ErrorResponse)
	assert.Equal(t, "Issue with ID 0 not found!", err.Error())
	assert.Equal(t, int32(404), errorResponse.Code)
	assert.Equal(t, "Issue with ID 0 not found!", errorResponse.Message)
}

func TestUseCaseController_GetProjectIssues(t *testing.T) {
	id := int64(0)
	points := int32(0)
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{
		"B": {
			0: model.Issue{
				Id:              &id,
				Name:            "B",
				Description:     "C",
				Status:          "",
				Priority:        "",
				ProjectKey:      "B",
				Components:      nil,
				Sprints:         nil,
				EstimatedPoints: &points,
				EstimatedTime:   "",
				AffectedVersion: "",
				FixedVersion:    "",
			},
		},
	}}
	result, err := underTest.GetProjectIssues("B")
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	allIssues := result.(model.Issues)
	assert.Equal(t, 1, len(allIssues.Issues))
	assert.Equal(t, int64(0), *allIssues.Issues[0].Id)
	assert.Equal(t, "B", allIssues.Issues[0].Name)
	assert.Equal(t, "C", allIssues.Issues[0].Description)
	assert.Equal(t, "", allIssues.Issues[0].Status)
	assert.Equal(t, "", allIssues.Issues[0].Priority)
	assert.Equal(t, "B", allIssues.Issues[0].ProjectKey)
	assert.Nil(t, allIssues.Issues[0].Components)
	assert.Nil(t, allIssues.Issues[0].Sprints)
	assert.Equal(t, int32(0), *allIssues.Issues[0].EstimatedPoints)
	assert.Equal(t, "", allIssues.Issues[0].EstimatedTime)
	assert.Equal(t, "", allIssues.Issues[0].AffectedVersion)
	assert.Equal(t, "", allIssues.Issues[0].FixedVersion)
}

func TestUseCaseController_GetProjectIssues_Empty(t *testing.T) {
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{"B": {}}}
	result, err := underTest.GetProjectIssues("B")
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	allIssues := result.(model.Issues)
	assert.Equal(t, 0, len(allIssues.Issues))
}

func TestUseCaseController_UpdateIssue(t *testing.T) {
	id := int64(0)
	points := int32(0)
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{
		"B": {
			0: model.Issue{
				Id:              &id,
				Name:            "B",
				Description:     "C",
				Status:          "",
				Priority:        "",
				ProjectKey:      "B",
				Components:      nil,
				Sprints:         nil,
				EstimatedPoints: &points,
				EstimatedTime:   "",
				AffectedVersion: "",
				FixedVersion:    "",
			},
		},
	}}
	_, err := underTest.UpdateIssue("B", model.Issue{
		Id:              &id,
		Name:            "A",
		Description:     "C",
		Status:          "",
		Priority:        "",
		ProjectKey:      "B",
		Components:      nil,
		Sprints:         nil,
		EstimatedPoints: &points,
		EstimatedTime:   "",
		AffectedVersion: "",
		FixedVersion:    "",
	})
	if err != nil {
		t.Error("Test failed: expect error is nil.", err)
		return
	}
	issue := underTest.issues["B"][0]
	assert.Equal(t, int64(0), *issue.Id)
	assert.Equal(t, "A", issue.Name)
	assert.Equal(t, "C", issue.Description)
	assert.Equal(t, "", issue.Status)
	assert.Equal(t, "", issue.Priority)
	assert.Equal(t, "B", issue.ProjectKey)
	assert.Nil(t, issue.Components)
	assert.Nil(t, issue.Sprints)
	assert.Equal(t, int32(0), *issue.EstimatedPoints)
	assert.Equal(t, "", issue.EstimatedTime)
	assert.Equal(t, "", issue.AffectedVersion)
	assert.Equal(t, "", issue.FixedVersion)
}

func TestUseCaseController_UpdateIssue_IdRequired(t *testing.T) {
	id := int64(0)
	points := int32(0)
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{
		"B": {
			0: model.Issue{
				Id:              &id,
				Name:            "B",
				Description:     "C",
				Status:          "",
				Priority:        "",
				ProjectKey:      "B",
				Components:      nil,
				Sprints:         nil,
				EstimatedPoints: &points,
				EstimatedTime:   "",
				AffectedVersion: "",
				FixedVersion:    "",
			},
		},
	}}
	result, err := underTest.UpdateIssue("B", model.Issue{
		Name:            "A",
		Description:     "C",
		Status:          "",
		Priority:        "",
		ProjectKey:      "B",
		Components:      nil,
		Sprints:         nil,
		EstimatedPoints: &points,
		EstimatedTime:   "",
		AffectedVersion: "",
		FixedVersion:    "",
	})
	errorResponse := result.(model.ErrorResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "Issue's ID is required!", err.Error())
	assert.Equal(t, int32(400), errorResponse.Code)
	assert.Equal(t, "Issue's ID is required!", errorResponse.Message)
}

func TestUseCaseController_UpdateIssue_BadRequest(t *testing.T) {
	id := int64(0)
	points := int32(0)
	underTest := UseCaseController{issues: map[string]map[int64]model.Issue{
		"B": {
			0: model.Issue{
				Id:              &id,
				Name:            "B",
				Description:     "C",
				Status:          "",
				Priority:        "",
				ProjectKey:      "B",
				Components:      nil,
				Sprints:         nil,
				EstimatedPoints: &points,
				EstimatedTime:   "",
				AffectedVersion: "",
				FixedVersion:    "",
			},
		},
	}}
	result, err := underTest.UpdateIssue("B", model.Issue{
		Name:            "A",
		Description:     "C",
		Status:          "",
		Priority:        "",
		ProjectKey:      "A",
		Components:      nil,
		Sprints:         nil,
		EstimatedPoints: &points,
		EstimatedTime:   "",
		AffectedVersion: "",
		FixedVersion:    "",
	})
	errorResponse := result.(model.ErrorResponse)
	assert.NotNil(t, err)
	assert.Equal(t, "Issue's ProjectKey A is not equal to B!", err.Error())
	assert.Equal(t, int32(400), errorResponse.Code)
	assert.Equal(t, "Issue's ProjectKey A is not equal to B!", errorResponse.Message)
}
