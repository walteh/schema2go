package schemas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr[T any](v T) *T {
	return &v
}

func TestSharedFieldsOfOneOfChildren(t *testing.T) {
	// Test 1: OneOf with constant field "model"
	t.Run("with constant field", func(t *testing.T) {
		// Create a parent type with oneOf children that have a common "model" field with const values
		parent := &Type{
			OneOf: []*Type{
				{
					Title: "RGB",
					Properties: map[string]*Type{
						"model": {
							Const: ptr("rgb"),
							Type:  []string{"string"},
						},
						"r": {
							Type: []string{"number"},
						},
						"g": {
							Type: []string{"number"},
						},
						"b": {
							Type: []string{"number"},
						},
					},
				},
				{
					Title: "HSL",
					Properties: map[string]*Type{
						"model": {
							Const: ptr("hsl"),
							Type:  []string{"string"},
						},
						"h": {
							Type: []string{"number"},
						},
						"s": {
							Type: []string{"number"},
						},
						"l": {
							Type: []string{"number"},
						},
					},
				},
			},
		}

		// Set parent references
		for _, child := range parent.OneOf {
			child.SetOneOfParent(parent)
		}

		// Get shared fields
		shared := SharedFieldsOfOneOfChildren(parent.OneOf)

		// Verify results
		assert.Len(t, shared, 1)
		assert.Equal(t, "model", shared[0].Name)
		assert.True(t, shared[0].IsConstant)
		assert.False(t, shared[0].IsRequired, "Should not be marked as required")
		assert.Equal(t, TypeList{"string"}, shared[0].Type.Type)
		assert.Equal(t, map[string]string{
			"RGB": "rgb",
			"HSL": "hsl",
		}, shared[0].ConstantValuesMap)
	})

	// Test 2: OneOf with non-constant shared fields
	t.Run("with non-constant shared field", func(t *testing.T) {
		parent := &Type{
			OneOf: []*Type{
				{
					Properties: map[string]*Type{
						"id": {
							Type: []string{"string"},
						},
						"value": {
							Type: []string{"number"},
						},
					},
				},
				{
					Properties: map[string]*Type{
						"id": {
							Type: []string{"string"},
						},
						"data": {
							Type: []string{"object"},
						},
					},
				},
			},
		}

		// Set parent references
		for _, child := range parent.OneOf {
			child.SetOneOfParent(parent)
		}

		// Get shared fields
		shared := SharedFieldsOfOneOfChildren(parent.OneOf)

		// Verify results
		assert.Len(t, shared, 1)
		assert.Equal(t, "id", shared[0].Name)
		assert.False(t, shared[0].IsConstant)
		assert.False(t, shared[0].IsRequired, "Should not be marked as required")
		assert.Equal(t, TypeList{"string"}, shared[0].Type.Type)
	})

	// Test 3: No shared fields
	t.Run("no shared fields", func(t *testing.T) {
		parent := &Type{
			OneOf: []*Type{
				{
					Properties: map[string]*Type{
						"foo": {
							Type: []string{"string"},
						},
					},
				},
				{
					Properties: map[string]*Type{
						"bar": {
							Type: []string{"number"},
						},
					},
				},
			},
		}

		// Set parent references
		for _, child := range parent.OneOf {
			child.SetOneOfParent(parent)
		}

		// Get shared fields
		shared := SharedFieldsOfOneOfChildren(parent.OneOf)

		// Verify results
		assert.Empty(t, shared)
	})

	// Test 4: Empty oneOf
	t.Run("empty oneOf", func(t *testing.T) {
		parent := &Type{
			OneOf: []*Type{},
		}

		// Get shared fields
		shared := SharedFieldsOfOneOfChildren(parent.OneOf)

		// Verify results
		assert.Empty(t, shared)
	})

	// Test 5: Inconsistent required status
	t.Run("inconsistent required fields", func(t *testing.T) {
		parent := &Type{
			OneOf: []*Type{
				{
					Properties: map[string]*Type{
						"id": {
							Type: []string{"string"},
						},
					},
					Required: []string{"id"},
				},
				{
					Properties: map[string]*Type{
						"id": {
							Type: []string{"string"},
						},
					},
					// No required field
				},
			},
		}

		// Set parent references
		for _, child := range parent.OneOf {
			child.SetOneOfParent(parent)
		}

		// Get shared fields
		shared := SharedFieldsOfOneOfChildren(parent.OneOf)

		// Verify results - should be empty because required status is inconsistent
		assert.Empty(t, shared, "Fields with inconsistent required status should not be considered shared")
	})

	// Test 6: Consistent required status
	t.Run("consistent required fields", func(t *testing.T) {
		parent := &Type{
			OneOf: []*Type{
				{
					Properties: map[string]*Type{
						"id": {
							Type: []string{"string"},
						},
					},
					Required: []string{"id"},
				},
				{
					Properties: map[string]*Type{
						"id": {
							Type: []string{"string"},
						},
					},
					Required: []string{"id"},
				},
			},
		}

		// Set parent references
		for _, child := range parent.OneOf {
			child.SetOneOfParent(parent)
		}

		// Get shared fields
		shared := SharedFieldsOfOneOfChildren(parent.OneOf)

		// Verify results
		assert.Len(t, shared, 1, "Fields with consistent required status should be considered shared")
		assert.Equal(t, "id", shared[0].Name)
		assert.True(t, shared[0].IsRequired, "Should be marked as required")
	})
}
