/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

// Newly created app information
type AppCreateResponse struct {

	Status string `json:"status,omitempty"`

	RepositoryUrl string `json:"repository_url,omitempty"`

	Ip string `json:"ip,omitempty"`
}
