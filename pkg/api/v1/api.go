/*
 * OpenSource Issue Träcking System
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Contact: DevBoxFanBoy@github.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1

import (
	"github.com/DevBoxFanBoy/opists/pkg/api/v1/model"
)

// IssuesApiServicer defines the api actions for the IssuesApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type IssuesApiServicer interface {
	AddIssue(string, model.Issue) (interface{}, error)
	DeleteIssue(string, int64) (interface{}, error)
	GetIssueById(string, int64) (interface{}, error)
	GetProjectIssues(string) (interface{}, error)
	UpdateIssue(string, model.Issue) (interface{}, error)
}

// ProjectsApiServicer defines the api actions for the ProjectsApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ProjectsApiServicer interface {
	GetAllProject() (interface{}, error)
	GetProject(string) (interface{}, error)
	UpdateProject(string, model.Project) (interface{}, error)
	CreateProject(project model.Project) (interface{}, error)
	DeleteProject(projectKey string) (interface{}, error)
}
