package common

import (
	"encoding/json"
	"io/ioutil"
)

// ParseConfig parses a configuration file and loads into a config
func ParseConfig(confPath string) (JSONConfig, error) {
	file, err := ioutil.ReadFile(confPath)
	if err != nil {
		return Config, err
	}
	err = json.Unmarshal(file, &Config)
	if err != nil {
		return Config, err
	}
	return Config, nil
}

// CreateConfig generates a default JSON config file and writes to the given path
func CreateConfig(confPath string) (JSONConfig, error) {
	c := JSONConfig{
		Hostname: "streamer.xplex.online",
		RigURL:   "https://rig-dev.xplex.online",
		RTMPPort: 1935,
		HTTPPort: 8086,
		AgentKey: "alongsharedkeyhereusedforagents",
	}
	j, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return c, err
	}
	err = ioutil.WriteFile(confPath, j, 0644)
	if err != nil {
		return c, err
	}
	return c, nil
}
