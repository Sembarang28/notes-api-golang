package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONB []string

func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONB) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Failed to unmarshal JSONB value")
	}
	return json.Unmarshal(source, j)
}
