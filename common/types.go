package common

// Config holds instance of parsed config
var Config JSONConfig

// JSONConfig holds the structure for configuration
type JSONConfig struct {
	Hostname string `json:"hostname"`
	RigURL   string `json:"rigUrl"`
	RTMPPort int    `json:"rtmpPort"`
	HTTPPort int    `json:"httpPort"`
	AgentKey string `json:"agentKey"`
}
