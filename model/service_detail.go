/*
 * Service manager- openapi 3.0.0
 *
 * Test task service manager
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

// depricated
type ServiceDetail struct {
	PID     int      `json:"PID,omitempty"`
	Command string   `json:"command,omitempty"`
	Stdout  []string `json:"stdout,omitempty"`
	Stderr  []string `json:"stderr,omitempty"`
}