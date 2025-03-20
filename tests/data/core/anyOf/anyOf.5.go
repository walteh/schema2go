// Code generated by github.com/walteh/schema2go, DO NOT EDIT.

package test

import "encoding/json"
import "errors"
import "fmt"
import yaml "gopkg.in/yaml.v3"

// Text provided to or from an LLM.
type TextContent struct {
	// The text content of the message.
	Text string `json:"text" yaml:"text" mapstructure:"text"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TextContent) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if _, ok := raw["text"]; raw != nil && !ok {
		return fmt.Errorf("field text in TextContent: required")
	}
	type Plain TextContent
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = TextContent(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *TextContent) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if _, ok := raw["text"]; raw != nil && !ok {
		return fmt.Errorf("field text in TextContent: required")
	}
	type Plain TextContent
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = TextContent(plain)
	return nil
}

type CallToolResultContentElem_0 = TextContent

// The server's response to a tool call
type CallToolResult struct {
	// Content corresponds to the JSON schema field "content".
	Content []CallToolResultContentElem `json:"content,omitempty" yaml:"content,omitempty" mapstructure:"content,omitempty"`
}

// Text provided to or from an LLM.
type CallToolResultContentElem struct {
	// The text content of the message.
	Text string `json:"text" yaml:"text" mapstructure:"text"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CallToolResultContentElem) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	var callToolResultContentElem_0 CallToolResultContentElem_0
	var errs []error
	if err := callToolResultContentElem_0.UnmarshalJSON(value); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 1 {
		return fmt.Errorf("all validators failed: %s", errors.Join(errs...))
	}
	type Plain CallToolResultContentElem
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = CallToolResultContentElem(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *CallToolResultContentElem) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	var callToolResultContentElem_0 CallToolResultContentElem_0
	var errs []error
	if err := callToolResultContentElem_0.UnmarshalYAML(value); err != nil {
		errs = append(errs, err)
	}
	if len(errs) == 1 {
		return fmt.Errorf("all validators failed: %s", errors.Join(errs...))
	}
	type Plain CallToolResultContentElem
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = CallToolResultContentElem(plain)
	return nil
}
