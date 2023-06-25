package kubeconfig

import (
	"encoding/json"
	"os"

	"sigs.k8s.io/yaml"
)

type KubeConfig struct {
	ApiVersion     string                 `json:"apiVersion"`
	Clusters       []ClusterWithName      `json:"clusters"`
	Contexts       []ContextWithName      `json:"contexts"`
	CurrentContext string                 `json:"current-context,omitempty"`
	Kind           string                 `json:"kind"`
	Preferences    map[string]interface{} `json:"preferences,omitempty"`
	Users          []UserWithName         `json:"users"`
}

func Load(configPath string) (KubeConfig, error) {
	configContents, err := os.ReadFile(configPath)
	if err != nil {
		return KubeConfig{}, err
	}

	configJSON, err := yaml.YAMLToJSON(configContents)
	if err != nil {
		return KubeConfig{}, err
	}

	var config KubeConfig
	if err := json.Unmarshal(configJSON, &config); err != nil {
		return KubeConfig{}, err
	}

	return config, nil
}

func Write(config KubeConfig, configPath string) error {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return err
	}

	configYAML, err := yaml.JSONToYAML(configJSON)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, configYAML, 0644)
}
