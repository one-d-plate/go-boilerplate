package pkg

import (
	"encoding/json"
	"os"
)

func LoadJSONFile[T any](path string, out *T) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, out)
}
