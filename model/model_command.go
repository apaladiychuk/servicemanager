/*
 * Service manager- openapi 3.0.0
 *
 * Test task service manager
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package model

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args,omitempty"`
}
