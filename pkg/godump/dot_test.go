package godump_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/walteh/schema2go/pkg/godump"
)

func TestDotNotation(t *testing.T) {
	type Child struct {
		Name string
		Age  int
	}

	type Parent struct {
		Name     string
		Children []Child
		Info     map[string]string
	}

	p := Parent{
		Name: "John",
		Children: []Child{
			{Name: "Alice", Age: 10},
			{Name: "Bob", Age: 12},
		},
		Info: map[string]string{
			"city":    "New York",
			"country": "USA",
		},
	}

	d := godump.Dumper{
		DotNotation: true,
	}

	result := d.Sprint(p)
	t.Log("Output:\n", result)

	// Verify dot notation format
	assert.Contains(t, result, "Parent.Name: \"John\"")
	assert.Contains(t, result, "Parent.Children[0].Name: \"Alice\"")
	assert.Contains(t, result, "Parent.Children[0].Age: 10")
	assert.Contains(t, result, "Parent.Children[1].Name: \"Bob\"")
	assert.Contains(t, result, "Parent.Children[1].Age: 12")
	assert.Contains(t, result, "Parent.Info.city: \"New York\"")
	assert.Contains(t, result, "Parent.Info.country: \"USA\"")
}

// Test structures
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	ID       int
	Name     string
	Title    string
	Address  Address
	Contacts []Contact
	Manager  *Employee
	Tags     map[string]string
	Skills   []string
	Active   bool
	Salary   float64
}

func TestDotNotationComprehensive(t *testing.T) {
	// Create a complex nested structure
	manager := &Employee{
		ID:    1,
		Name:  "John Boss",
		Title: "CEO",
		Address: Address{
			Street:  "123 Main St",
			City:    "San Francisco",
			Country: "USA",
			ZipCode: "94105",
		},
		Contacts: []Contact{
			{Email: "john@example.com", Phone: "555-0001"},
		},
		Tags: map[string]string{
			"department": "Executive",
			"level":      "C-Suite",
		},
		Skills: []string{"Leadership", "Strategy"},
		Active: true,
		Salary: 250000.00,
	}

	employee := Employee{
		ID:    2,
		Name:  "Alice Worker",
		Title: "Engineer",
		Address: Address{
			Street:  "456 Tech Ave",
			City:    "San Jose",
			Country: "USA",
			ZipCode: "95113",
		},
		Contacts: []Contact{
			{Email: "alice@example.com", Phone: "555-0002"},
			{Email: "alice.personal@example.com", Phone: "555-0003"},
		},
		Manager: manager,
		Tags: map[string]string{
			"department": "Engineering",
			"team":       "Backend",
			"level":      "Senior",
		},
		Skills: []string{"Go", "Python", "Kubernetes"},
		Active: true,
		Salary: 150000.00,
	}

	d := godump.Dumper{
		DotNotation: true,
	}

	result := d.Sprint(employee)
	t.Log("\nDot Notation Output:\n", result)

	// Test basic fields
	assert.Contains(t, result, "Employee.ID: 2", "should contain employee ID")
	assert.Contains(t, result, "Employee.Name: \"Alice Worker\"", "should contain employee name")
	assert.Contains(t, result, "Employee.Title: \"Engineer\"", "should contain employee title")
	assert.Contains(t, result, "Employee.Active: true", "should contain employee active status")
	assert.Contains(t, result, "Employee.Salary: 150000", "should contain employee salary")

	// Test nested struct (Address)
	assert.Contains(t, result, "Employee.Address.Street: \"456 Tech Ave\"", "should contain address street")
	assert.Contains(t, result, "Employee.Address.City: \"San Jose\"", "should contain address city")
	assert.Contains(t, result, "Employee.Address.Country: \"USA\"", "should contain address country")
	assert.Contains(t, result, "Employee.Address.ZipCode: \"95113\"", "should contain address zip code")

	// Test slice of structs (Contacts)
	assert.Contains(t, result, "Employee.Contacts[0].Email: \"alice@example.com\"", "should contain first contact email")
	assert.Contains(t, result, "Employee.Contacts[0].Phone: \"555-0002\"", "should contain first contact phone")
	assert.Contains(t, result, "Employee.Contacts[1].Email: \"alice.personal@example.com\"", "should contain second contact email")
	assert.Contains(t, result, "Employee.Contacts[1].Phone: \"555-0003\"", "should contain second contact phone")

	// Test map fields
	assert.Contains(t, result, "Employee.Tags.department: \"Engineering\"", "should contain department tag")
	assert.Contains(t, result, "Employee.Tags.team: \"Backend\"", "should contain team tag")
	assert.Contains(t, result, "Employee.Tags.level: \"Senior\"", "should contain level tag")

	// Test string slice
	assert.Contains(t, result, "Employee.Skills[0]: \"Go\"", "should contain first skill")
	assert.Contains(t, result, "Employee.Skills[1]: \"Python\"", "should contain second skill")
	assert.Contains(t, result, "Employee.Skills[2]: \"Kubernetes\"", "should contain third skill")

	// Test nested pointer to struct (Manager)
	assert.Contains(t, result, "Employee.Manager.ID: 1", "should contain manager ID")
	assert.Contains(t, result, "Employee.Manager.Name: \"John Boss\"", "should contain manager name")
	assert.Contains(t, result, "Employee.Manager.Title: \"CEO\"", "should contain manager title")
	assert.Contains(t, result, "Employee.Manager.Address.Street: \"123 Main St\"", "should contain manager address street")
	assert.Contains(t, result, "Employee.Manager.Tags.department: \"Executive\"", "should contain manager department tag")
	assert.Contains(t, result, "Employee.Manager.Skills[0]: \"Leadership\"", "should contain manager first skill")
}

func TestDotNotationNilPointers(t *testing.T) {
	employee := Employee{
		ID:      3,
		Name:    "Bob Newbie",
		Title:   "Intern",
		Manager: nil,
		Tags:    nil,
		Skills:  nil,
	}

	d := godump.Dumper{
		DotNotation: true,
	}

	result := d.Sprint(employee)
	t.Log("\nNil Pointer Output:\n", result)

	// Test nil pointer and slices
	assert.Contains(t, result, "Employee.Manager: (nil)", "should handle nil manager pointer")
	assert.Contains(t, result, "Employee.Tags: (nil)", "should handle nil map")
	assert.Contains(t, result, "Employee.Skills: (nil)", "should handle nil slice")
}

func TestDotNotationCyclicReferences(t *testing.T) {
	employee1 := &Employee{
		ID:    4,
		Name:  "Eve Cyclic",
		Title: "Manager",
	}

	employee2 := &Employee{
		ID:      5,
		Name:    "Charlie Cyclic",
		Title:   "Lead",
		Manager: employee1,
	}

	// Create a cycle
	employee1.Manager = employee2

	d := godump.Dumper{
		DotNotation: true,
	}

	result := d.Sprint(employee1)
	t.Log("\nCyclic Reference Output:\n", result)

	// Test cyclic references
	assert.Contains(t, result, "Employee.Name: \"Eve Cyclic\"", "should contain first employee name")
	assert.Contains(t, result, "Employee.Manager.Name: \"Charlie Cyclic\"", "should contain second employee name")
	assert.Contains(t, result, "Employee.Manager.Manager: &@1", "should handle cyclic reference with pointer tag")
}

// Worker interface for testing interface handling
type Worker interface {
	Work() string
	GetSalary() float64
}

// unexported interface for testing private interface handling
type worker interface {
	work() string
}

// FullTimeEmployee implements Worker with both exported and unexported fields
type FullTimeEmployee struct {
	ID             int
	Name           string
	department     string // unexported
	salary         float64
	secretBonus    float64 // unexported
	privateNotes   []string
	internalRank   int
	publicRating   float64
	manager        *FullTimeEmployee // unexported pointer
	Assistant      *FullTimeEmployee // exported pointer
	implementation worker            // unexported interface
}

func (e *FullTimeEmployee) Work() string {
	return "Working from " + e.department
}

func (e *FullTimeEmployee) GetSalary() float64 {
	return e.salary + e.secretBonus
}

func (e *FullTimeEmployee) work() string {
	return "private work"
}

func TestDotNotationUnexportedAndInterfaces(t *testing.T) {
	assistant := &FullTimeEmployee{
		ID:           101,
		Name:         "Bob Assistant",
		department:   "Support",
		salary:       50000,
		secretBonus:  5000,
		privateNotes: []string{"Good team player"},
		internalRank: 2,
		publicRating: 4.5,
	}

	employee := &FullTimeEmployee{
		ID:           100,
		Name:         "Alice Manager",
		department:   "Engineering",
		salary:       100000,
		secretBonus:  10000,
		privateNotes: []string{"High performer", "Leadership potential"},
		internalRank: 5,
		publicRating: 4.8,
		Assistant:    assistant,
		manager:      nil,
	}
	// Make employee implement its own worker interface
	employee.implementation = employee

	// Test with private fields hidden
	d1 := godump.Dumper{
		DotNotation:       true,
		HidePrivateFields: true,
	}
	resultHidden := d1.Sprint(employee)
	t.Log("\nDot Notation Output (private fields hidden):\n", resultHidden)

	// Verify only public fields are shown
	assert.Contains(t, resultHidden, "FullTimeEmployee.ID: 100", "should contain ID")
	assert.Contains(t, resultHidden, "FullTimeEmployee.Name: \"Alice Manager\"", "should contain Name")
	assert.Contains(t, resultHidden, "FullTimeEmployee.Assistant", "should contain Assistant")
	assert.NotContains(t, resultHidden, "department", "should not contain private department")
	assert.NotContains(t, resultHidden, "salary", "should not contain private salary")
	assert.NotContains(t, resultHidden, "secretBonus", "should not contain private secretBonus")
	assert.NotContains(t, resultHidden, "privateNotes", "should not contain privateNotes")
	assert.NotContains(t, resultHidden, "internalRank", "should not contain internalRank")

	// Test with private fields shown
	d2 := godump.Dumper{
		DotNotation:       true,
		HidePrivateFields: false,
	}
	resultShown := d2.Sprint(employee)
	t.Log("\nDot Notation Output (private fields shown):\n", resultShown)

	// Verify all fields are shown
	assert.Contains(t, resultShown, "FullTimeEmployee.ID: 100", "should contain ID")
	assert.Contains(t, resultShown, "FullTimeEmployee.Name: \"Alice Manager\"", "should contain Name")
	assert.Contains(t, resultShown, "FullTimeEmployee.Assistant.ID: 101", "should contain Assistant's ID")
	assert.Contains(t, resultShown, "FullTimeEmployee.department: \"Engineering\"", "should contain department")
	assert.Contains(t, resultShown, "FullTimeEmployee.salary: 100000", "should contain salary")
	assert.Contains(t, resultShown, "FullTimeEmployee.secretBonus: 10000", "should contain secretBonus")
	assert.Contains(t, resultShown, "FullTimeEmployee.privateNotes[0]: \"High performer\"", "should contain first private note")
	assert.Contains(t, resultShown, "FullTimeEmployee.privateNotes[1]: \"Leadership potential\"", "should contain second private note")
	assert.Contains(t, resultShown, "FullTimeEmployee.internalRank: 5", "should contain internalRank")
	assert.Contains(t, resultShown, "FullTimeEmployee.publicRating: 4.8", "should contain publicRating")
	assert.Contains(t, resultShown, "FullTimeEmployee.manager: (nil)", "should contain manager")
	assert.Contains(t, resultShown, "FullTimeEmployee.implementation: &@1", "should contain implementation with cyclic reference")

	// Test interface handling
	var worker Worker = employee
	d3 := godump.Dumper{
		DotNotation: true,
	}
	resultInterface := d3.Sprint(worker)
	t.Log("\nDot Notation Output (interface):\n", resultInterface)

	// Verify interface shows underlying type's fields
	assert.Contains(t, resultInterface, "FullTimeEmployee.ID: 100", "should contain ID through interface")
	assert.Contains(t, resultInterface, "FullTimeEmployee.Name: \"Alice Manager\"", "should contain Name through interface")
	assert.Contains(t, resultInterface, "FullTimeEmployee.Assistant.ID: 101", "should contain Assistant's ID through interface")
}

// NumberWrapper is a struct that can hold either an integer or float pointer
type NumberWrapper struct {
	Integer *int
	Float   *float64
}

// SchemaProperty represents a property with min/max constraints
type SchemaProperty struct {
	Maximum *NumberWrapper
	Minimum *NumberWrapper
}

func TestDotNotationNestedPointers(t *testing.T) {
	float100 := float64(100)
	float0 := float64(0)

	prop := &SchemaProperty{
		Maximum: &NumberWrapper{
			Integer: nil,
			Float:   &float100,
		},
		Minimum: &NumberWrapper{
			Integer: nil,
			Float:   &float0,
		},
	}

	d := godump.Dumper{
		DotNotation: true,
	}

	result := d.Sprint(prop)
	t.Log("\nNested Pointer Output:\n", result)

	// Test the exact format we want
	assert.Contains(t, result, "SchemaProperty.Maximum.Integer: (nil)", "should show nil integer")
	assert.Contains(t, result, "SchemaProperty.Maximum.Float: &100", "should show float pointer value")
	assert.Contains(t, result, "SchemaProperty.Minimum.Integer: (nil)", "should show nil integer")
	assert.Contains(t, result, "SchemaProperty.Minimum.Float: &0", "should show float pointer value")

	// Verify what we don't want to see
	assert.NotContains(t, result, "SchemaProperty.Maximum: &SchemaProperty.Maximum", "should not repeat path")
	assert.NotContains(t, result, "SchemaProperty.Minimum: &SchemaProperty.Minimum", "should not repeat path")
}
