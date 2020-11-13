/*
 * OpenSource Issue Träcking System
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Contact: DevBoxFanBoy@github.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package opists

import (
	"errors"
)

// ProjectsApiService is a service that implents the logic for the ProjectsApiServicer
// This service should implement the business logic for every endpoint for the ProjectsApi API.
// Include any external packages or services that will be required by this service.
type ProjectsApiService struct {
}

// NewProjectsApiService creates a default api service
func NewProjectsApiService() ProjectsApiServicer {
	return &ProjectsApiService{}
}

// GetAllProject - Returns all projects
func (s *ProjectsApiService) GetAllProject() (interface{}, error) {
	// TODO - update GetAllProject with the required logic for this service method.
	// Add api_projects_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetAllProject' not implemented")
}

// GetProject - Returns the Project by key
func (s *ProjectsApiService) GetProject(projectKey string) (interface{}, error) {
	// TODO - update GetProject with the required logic for this service method.
	// Add api_projects_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetProject' not implemented")
}
