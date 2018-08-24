/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type ServiceBrokerConfig struct {
	Insecure        bool                           `json:"Insecure,omitempty"`
	Context         map[string]string              `json:"Context,omitempty"`
	AuthConfig      *ServiceBrokerConfigAuthConfig `json:"AuthConfig,omitempty"`
	CacheExpiration string                         `json:"CacheExpiration,omitempty"`
}