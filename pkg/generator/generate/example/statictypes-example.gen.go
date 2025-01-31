package generator

var _ Schema = &StaticSchema{}

type StaticSchema struct {
	Enums_   []*EnumModel
	Imports_ []string
	Package_ string
	Structs_ []Struct
}

// Enums implements Schema.
func (b *StaticSchema) Enums() []*EnumModel {
	return b.Enums_
}

// Imports implements Schema.
func (b *StaticSchema) Imports() []string {
	return b.Imports_
}

// Package implements Schema.
func (b *StaticSchema) Package() string {
	return b.Package_
}

// Structs implements Schema.
func (b *StaticSchema) Structs() []Struct {
	return b.Structs_
}

var _ Struct = &StaticStruct{}

type StaticStruct struct {
	Name_                string
	Description_         string
	Fields_              []Field
	HasAllOf_            bool
	HasCustomMarshaling_ bool
	HasDefaults_         bool
	HasValidation_       bool
}

// Description implements Struct.
func (b *StaticStruct) Description() string {
	return b.Description_
}

// Fields implements Struct.
func (b *StaticStruct) Fields() []Field {
	return b.Fields_
}

// HasAllOf implements Struct.
func (b *StaticStruct) HasAllOf() bool {
	return b.HasAllOf_
}

// HasCustomMarshaling implements Struct.
func (b *StaticStruct) HasCustomMarshaling() bool {
	return b.HasCustomMarshaling_
}

// HasDefaults implements Struct.
func (b *StaticStruct) HasDefaults() bool {
	return b.HasDefaults_
}

// HasValidation implements Struct.
func (b *StaticStruct) HasValidation() bool {
	return b.HasValidation_
}

// Name implements Struct.
func (b *StaticStruct) Name() string {
	return b.Name_
}

var _ Field = &StaticField{}

type StaticField struct {
	Name_                string
	JSONName_            string
	Description_         string
	Type_                string
	IsRequired_          bool
	IsEnum_              bool
	EnumTypeName_        string
	EnumValues_          []*EnumValue
	DefaultValue_        *string
	DefaultValueComment_ *string
	ValidationRules_     []*ValidationRule
}

// DefaultValue implements Field.
func (f *StaticField) DefaultValue() *string {
	return f.DefaultValue_
}

// DefaultValueComment implements Field.
func (f *StaticField) DefaultValueComment() *string {
	return f.DefaultValueComment_
}

// EnumTypeName implements Field.
func (f *StaticField) EnumTypeName() string {
	return f.EnumTypeName_
}

// EnumValues implements Field.
func (f *StaticField) EnumValues() []*EnumValue {
	return f.EnumValues_
}

// IsEnum implements Field.
func (f *StaticField) IsEnum() bool {
	return f.IsEnum_
}

// IsRequired implements Field.
func (f *StaticField) IsRequired() bool {
	return f.IsRequired_
}

// JSONName implements Field.
func (f *StaticField) JSONName() string {
	return f.JSONName_
}

// Name implements Field.
func (f *StaticField) Name() string {
	return f.Name_
}

// Type implements Field.
func (f *StaticField) Type() string {
	return f.Type_
}

// ValidationRules implements Field.
func (f *StaticField) ValidationRules() []*ValidationRule {
	return f.ValidationRules_
}

func (f *StaticField) Description() string {
	return f.Description_
}

func NewStaticSchema(schema Schema) *StaticSchema {
	stat := &StaticSchema{
		Package_: schema.Package(),
		Enums_:   schema.Enums(),
		Imports_: schema.Imports(),
	}

	for _, strt := range schema.Structs() {
		stat.Structs_ = append(stat.Structs_, NewStaticStruct(strt))
	}

	return stat
}

func NewStaticStruct(strt Struct) *StaticStruct {
	stat := &StaticStruct{
		Name_:                strt.Name(),
		Description_:         strt.Description(),
		Fields_:              strt.Fields(),
		HasAllOf_:            strt.HasAllOf(),
		HasCustomMarshaling_: strt.HasCustomMarshaling(),
		HasDefaults_:         strt.HasDefaults(),
		HasValidation_:       strt.HasValidation(),
	}

	for _, field := range strt.Fields() {
		stat.Fields_ = append(stat.Fields_, NewStaticField(field))
	}

	return stat
}

func NewStaticField(field Field) *StaticField {
	stat := &StaticField{
		Name_:                field.Name(),
		JSONName_:            field.JSONName(),
		Description_:         field.Description(),
		Type_:                field.Type(),
		IsRequired_:          field.IsRequired(),
		IsEnum_:              field.IsEnum(),
		EnumTypeName_:        field.EnumTypeName(),
		EnumValues_:          field.EnumValues(),
		DefaultValue_:        field.DefaultValue(),
		DefaultValueComment_: field.DefaultValueComment(),
		ValidationRules_:     field.ValidationRules(),
	}

	return stat
}
