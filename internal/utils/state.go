package utils

import (
	"encoding/json"
	"os"
)

type GoCraftState struct {
	ModuleName string `json:"module_name"`
	Framework  string `json:"framework"`
	GoVersion  string `json:"go_version,omitempty"`
	ORM        string `json:"orm,omitempty"`
	Database   string `json:"database,omitempty"`
}

const stateFile = ".gocraft.json"

func SaveState(state GoCraftState) error {
	file, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(state)
}

func LoadState() (*GoCraftState, error) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	var state GoCraftState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	return &state, nil
}
