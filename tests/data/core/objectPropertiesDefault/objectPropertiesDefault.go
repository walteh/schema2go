// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

import "encoding/json"
import "errors"
import "fmt"
import yaml "gopkg.in/yaml.v3"
import "reflect"

type DecoratedPlanner struct {
	// Decorator corresponds to the JSON schema field "decorator".
	Decorator DecoratedPlannerDecorator `json:"decorator,omitempty" yaml:"decorator,omitempty" mapstructure:"decorator,omitempty"`

	// Event corresponds to the JSON schema field "event".
	Event *Event `json:"event,omitempty" yaml:"event,omitempty" mapstructure:"event,omitempty"`
}

type DecoratedPlannerDecorator struct {
	// Color corresponds to the JSON schema field "color".
	Color string `json:"color,omitempty" yaml:"color,omitempty" mapstructure:"color,omitempty"`

	// Theme corresponds to the JSON schema field "theme".
	Theme *string `json:"theme,omitempty" yaml:"theme,omitempty" mapstructure:"theme,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DecoratedPlannerDecorator) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain DecoratedPlannerDecorator
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if v, ok := raw["color"]; !ok || v == nil {
		plain.Color = "#ffffff"
	}
	*j = DecoratedPlannerDecorator(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *DecoratedPlannerDecorator) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain DecoratedPlannerDecorator
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if v, ok := raw["color"]; !ok || v == nil {
		plain.Color = "#ffffff"
	}
	*j = DecoratedPlannerDecorator(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *DecoratedPlanner) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain DecoratedPlanner
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if v, ok := raw["decorator"]; !ok || v == nil {
		plain.Decorator = DecoratedPlannerDecorator{
			Color: "#ffffff",
			Theme: nil,
		}
	}
	*j = DecoratedPlanner(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DecoratedPlanner) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain DecoratedPlanner
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if v, ok := raw["decorator"]; !ok || v == nil {
		plain.Decorator = DecoratedPlannerDecorator{
			Color: "#ffffff",
			Theme: nil,
		}
	}
	*j = DecoratedPlanner(plain)
	return nil
}

type DefaultPlanner struct {
	// Event corresponds to the JSON schema field "event".
	Event *Event `json:"event,omitempty" yaml:"event,omitempty" mapstructure:"event,omitempty"`
}

type Event struct {
	// Name corresponds to the JSON schema field "name".
	Name *EventName `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// Tags corresponds to the JSON schema field "tags".
	Tags []EventTagsElem `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`
}

type EventName string

const EventNameBIRTHDAY EventName = "BIRTHDAY"
const EventNameGAME EventName = "GAME"
const EventNameHOLIDAY EventName = "HOLIDAY"

var enumValues_EventName = []interface{}{
	"BIRTHDAY",
	"GAME",
	"HOLIDAY",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventName) UnmarshalJSON(value []byte) error {
	var v string
	if err := json.Unmarshal(value, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventName {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventName, v)
	}
	*j = EventName(v)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EventName) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventName {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventName, v)
	}
	*j = EventName(v)
	return nil
}

type EventTagsElem string

const EventTagsElemCITY EventTagsElem = "CITY"
const EventTagsElemCOUNTRY EventTagsElem = "COUNTRY"
const EventTagsElemPERSON EventTagsElem = "PERSON"
const EventTagsElemREGION EventTagsElem = "REGION"

var enumValues_EventTagsElem = []interface{}{
	"COUNTRY",
	"REGION",
	"CITY",
	"PERSON",
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *EventTagsElem) UnmarshalYAML(value *yaml.Node) error {
	var v string
	if err := value.Decode(&v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventTagsElem {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventTagsElem, v)
	}
	*j = EventTagsElem(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventTagsElem) UnmarshalJSON(value []byte) error {
	var v string
	if err := json.Unmarshal(value, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventTagsElem {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventTagsElem, v)
	}
	*j = EventTagsElem(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Event) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	type Plain Event
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	if v, ok := raw["tags"]; !ok || v == nil {
		plain.Tags = []EventTagsElem{}
	}
	*j = Event(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *Event) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	type Plain Event
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	if v, ok := raw["tags"]; !ok || v == nil {
		plain.Tags = []EventTagsElem{}
	}
	*j = Event(plain)
	return nil
}

type ObjectPropertiesDefault struct {
	// Active corresponds to the JSON schema field "active".
	Active interface{} `json:"active,omitempty" yaml:"active,omitempty" mapstructure:"active,omitempty"`

	// Planners corresponds to the JSON schema field "planners".
	Planners []ObjectPropertiesDefaultPlannersElem `json:"planners,omitempty" yaml:"planners,omitempty" mapstructure:"planners,omitempty"`
}

type ObjectPropertiesDefaultPlannersElem struct {
	// Decorated corresponds to the JSON schema field "decorated".
	Decorated *DecoratedPlanner `json:"decorated,omitempty" yaml:"decorated,omitempty" mapstructure:"decorated,omitempty"`

	// Plain corresponds to the JSON schema field "plain".
	Plain *DefaultPlanner `json:"plain,omitempty" yaml:"plain,omitempty" mapstructure:"plain,omitempty"`
}

type ObjectPropertiesDefaultPlannersElem_0 struct {
	// Plain corresponds to the JSON schema field "plain".
	Plain *DefaultPlanner `json:"plain,omitempty" yaml:"plain,omitempty" mapstructure:"plain,omitempty"`
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ObjectPropertiesDefaultPlannersElem_0) UnmarshalYAML(value *yaml.Node) error {
	type Plain ObjectPropertiesDefaultPlannersElem_0
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = ObjectPropertiesDefaultPlannersElem_0(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectPropertiesDefaultPlannersElem_0) UnmarshalJSON(value []byte) error {
	type Plain ObjectPropertiesDefaultPlannersElem_0
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = ObjectPropertiesDefaultPlannersElem_0(plain)
	return nil
}

type ObjectPropertiesDefaultPlannersElem_1 struct {
	// Decorated corresponds to the JSON schema field "decorated".
	Decorated *DecoratedPlanner `json:"decorated,omitempty" yaml:"decorated,omitempty" mapstructure:"decorated,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectPropertiesDefaultPlannersElem_1) UnmarshalJSON(value []byte) error {
	type Plain ObjectPropertiesDefaultPlannersElem_1
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = ObjectPropertiesDefaultPlannersElem_1(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ObjectPropertiesDefaultPlannersElem_1) UnmarshalYAML(value *yaml.Node) error {
	type Plain ObjectPropertiesDefaultPlannersElem_1
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = ObjectPropertiesDefaultPlannersElem_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectPropertiesDefaultPlannersElem) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	var objectPropertiesDefaultPlannersElem_0 ObjectPropertiesDefaultPlannersElem_0
	var objectPropertiesDefaultPlannersElem_1 ObjectPropertiesDefaultPlannersElem_1
	var errs []error
	if err := objectPropertiesDefaultPlannersElem_0.UnmarshalJSON(value); err != nil {
		errs = append(errs, err)
	}
	if err := objectPropertiesDefaultPlannersElem_1.UnmarshalJSON(value); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 2 {
		return fmt.Errorf("all validators failed: %s", errors.Join(errs...))
	}
	type Plain ObjectPropertiesDefaultPlannersElem
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = ObjectPropertiesDefaultPlannersElem(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ObjectPropertiesDefaultPlannersElem) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	var objectPropertiesDefaultPlannersElem_0 ObjectPropertiesDefaultPlannersElem_0
	var objectPropertiesDefaultPlannersElem_1 ObjectPropertiesDefaultPlannersElem_1
	var errs []error
	if err := objectPropertiesDefaultPlannersElem_0.UnmarshalYAML(value); err != nil {
		errs = append(errs, err)
	}
	if err := objectPropertiesDefaultPlannersElem_1.UnmarshalYAML(value); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 2 {
		return fmt.Errorf("all validators failed: %s", errors.Join(errs...))
	}
	type Plain ObjectPropertiesDefaultPlannersElem
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = ObjectPropertiesDefaultPlannersElem(plain)
	return nil
}
