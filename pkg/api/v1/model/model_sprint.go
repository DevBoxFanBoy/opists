/*
 * OpenSource Issue Träcking System
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Contact: DevBoxFanBoy@github.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
	"time"
)

// Sprint - Issue in sprint scoped.
type Sprint struct {
	Key string `json:"key"`

	Name string `json:"name"`

	// Startdate of the Sprint.
	Start time.Time `json:"start,omitempty"`

	// Enddate of the Sprint.
	End time.Time `json:"end,omitempty"`
}