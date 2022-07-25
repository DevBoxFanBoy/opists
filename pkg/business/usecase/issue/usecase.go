package issue

import (
	"errors"
	"fmt"
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
	"github.com/DevBoxFanBoy/opists/pkg/business/usecase/project"
	"sync"
)

type UseCase interface {
	AddIssue(string, model.Issue) (interface{}, error)
	DeleteIssue(string, int64) (interface{}, error)
	GetIssueById(string, int64) (interface{}, error)
	GetProjectIssues(string) (interface{}, error)
	UpdateIssue(string, model.Issue) (interface{}, error)
	//TODO PatchIssue(string, model.Issue) (interface{}, error)
}

var once sync.Once
var instance UseCaseController
var projectReader project.UseCase

func GetUseCaseControllerInstance() UseCase {
	once.Do(func() {
		issue := createIssueModel(0)
		issues := map[string]map[int64]model.Issue{"DF": {0: issue}}
		instance = UseCaseController{issues: issues}
		projectReader = project.GetUseCaseControllerInstance()
	})
	return &instance
}

type UseCaseController struct {
	issues map[string]map[int64]model.Issue
}

func (u *UseCaseController) AddIssue(projectKey string, issue model.Issue) (interface{}, error) {
	if res, err := projectReader.GetProject(projectKey); err != nil {
		return res, err
	}
	if issue.ProjectKey != projectKey {
		err := errors.New(fmt.Sprintf("Issue's ProjectKey %v is not equal to %v!", issue.ProjectKey, projectKey))
		return model.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		}, err
	}
	issueId := *issue.Id
	if _, ok := u.issues[projectKey]; !ok {
		u.issues[projectKey] = make(map[int64]model.Issue)
	}
	if issueId < 0 {
		issueId = int64(len(u.issues[projectKey]))
		issue.Id = &issueId
	}
	if _, ok := u.issues[projectKey][issueId]; !ok {
		u.issues[projectKey][issueId] = issue
	}
	return issue, nil
}

func (u *UseCaseController) DeleteIssue(projectKey string, id int64) (interface{}, error) {
	if res, err := projectReader.GetProject(projectKey); err != nil {
		return res, err
	}
	if projectIssues, ok := u.issues[projectKey]; ok {
		if _, ok := projectIssues[id]; ok {
			delete(projectIssues, id)
			return nil, nil
		}
	}
	err := errors.New(fmt.Sprintf("Issue with ID %v not found!", id))
	return model.ErrorResponse{
		Code:    404,
		Message: err.Error(),
	}, err
}

func (u *UseCaseController) GetIssueById(projectKey string, id int64) (interface{}, error) {
	if res, err := projectReader.GetProject(projectKey); err != nil {
		return res, err
	}
	if projectIssues, ok := u.issues[projectKey]; ok {
		if issue, ok := projectIssues[id]; ok {
			return issue, nil
		}
	}
	err := errors.New(fmt.Sprintf("Issue with ID %v not found!", id))
	return model.ErrorResponse{
		Code:    404,
		Message: err.Error(),
	}, err
}

func (u *UseCaseController) GetProjectIssues(projectKey string) (interface{}, error) {
	if res, err := projectReader.GetProject(projectKey); err != nil {
		return res, err
	}
	var issuesModel = model.Issues{}
	if projectIssues, ok := u.issues[projectKey]; ok {
		for _, element := range projectIssues {
			issuesModel.Issues = append(issuesModel.Issues, element)
		}
	}
	return issuesModel, nil
}

func (u *UseCaseController) UpdateIssue(projectKey string, issue model.Issue) (interface{}, error) {
	if res, err := projectReader.GetProject(projectKey); err != nil {
		return res, err
	}
	if issue.Id == nil {
		err := errors.New(fmt.Sprintf("Issue's ID is required!"))
		return model.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		}, err
	}
	if issue.ProjectKey != projectKey {
		err := errors.New(fmt.Sprintf("Issue's ProjectKey %v is not equal to %v!", issue.ProjectKey, projectKey))
		return model.ErrorResponse{
			Code:    400,
			Message: err.Error(),
		}, err
	}
	if projectIssues, ok := u.issues[projectKey]; ok {
		if _, ok := projectIssues[*issue.Id]; !ok {
			err := errors.New(fmt.Sprintf("Issue with ID %v not found!", *issue.Id))
			return model.ErrorResponse{
				Code:    404,
				Message: err.Error(),
			}, err
		}
		projectIssues[*issue.Id] = issue
	}
	return nil, nil
}

func createIssueModel(id int64) model.Issue {
	points := int32(0)
	return model.Issue{
		Id:              &id,
		Name:            "New Bug",
		Description:     "An error raise when...",
		Status:          "open",
		Priority:        "Highest",
		ProjectKey:      "DF",
		Components:      []string{"DrinkOwnChampagne", "EatMyOwnApplication"},
		Sprints:         []string{"Sprint2"},
		EstimatedPoints: &points,
		EstimatedTime:   "0h",
		AffectedVersion: "1.2.3",
		FixedVersion:    "1.2.4",
	}
}
