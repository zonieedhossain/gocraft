package utils

import (
	"encoding/json"
	"os"
)

type GocraftState struct {
	ModuleName string `json:"module_name"`
	Framework  string `json:"framework"`
	ORM        string `json:"orm,omitempty"`
	Database   string `json:"database,omitempty"`
}

const stateFile = ".gocraft.json"

func SaveState(state GocraftState) error {
	file, err := os.Create(stateFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(state)
}

func LoadState() (*GocraftState, error) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	var state GocraftState
	if err := json.Unmarshal(data, &state); err != nil {
		return nil, err
	}
	return &state, nil
}
