/*
 * Service manager- openapi 3.0.0
 *
 * Test task service manager
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

type Service struct {
	PID int `json:"PID"`

	CommandName string `json:"commandName,omitempty"`

	StartTime string `json:"startTime,omitempty"`
}