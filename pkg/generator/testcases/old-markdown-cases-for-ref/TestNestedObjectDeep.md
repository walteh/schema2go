# json-schema

```json
{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"title": "Location",
	"type": "object",
	"required": ["address"],
	"properties": {
		"address": {
			"type": "object",
			"required": ["coordinates"],
			"properties": {
				"coordinates": {
					"type": "object",
					"required": ["latitude", "longitude"],
					"properties": {
						"latitude": {
							"type": "number"
						},
						"longitude": {
							"type": "number"
						}
					}
				}
			}
		}
	}
}
```

---

# go-code

```go
// Code generated by schema2go. DO NOT EDIT.
// 🏗️ Generated from JSON Schema

package models

import (
    "encoding/json"
    "gitlab.com/tozd/go/errors"
)

type LocationAddressCoordinates struct {
    Latitude  float64 `json:"latitude"`  // Required
    Longitude float64 `json:"longitude"` // Required
}

// Validate ensures all required fields are present
func (x *LocationAddressCoordinates) Validate() error {
    return nil
}

type LocationAddress struct {
    Coordinates LocationAddressCoordinates `json:"coordinates"` // Required
}

// Validate ensures all required fields are present
func (x *LocationAddress) Validate() error {
    if err := x.Coordinates.Validate(); err != nil {
        return errors.Errorf("validating coordinates: %w", err)
    }
    return nil
}

type Location struct {
    Address LocationAddress `json:"address"` // Required
}

// Validate ensures all required fields are present
func (x *Location) Validate() error {
    if err := x.Address.Validate(); err != nil {
        return errors.Errorf("validating address: %w", err)
    }
    return nil
}
```
