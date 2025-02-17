// Code generated by mockery v2.51.0. DO NOT EDIT.

package mockery

import (
	mock "github.com/stretchr/testify/mock"
	generator "github.com/walteh/schema2go/pkg/generator"
)

// MockField_generator is an autogenerated mock type for the Field type
type MockField_generator struct {
	mock.Mock
}

type MockField_generator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockField_generator) EXPECT() *MockField_generator_Expecter {
	return &MockField_generator_Expecter{mock: &_m.Mock}
}

// DefaultValue provides a mock function with no fields
func (_m *MockField_generator) DefaultValue() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultValue")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockField_generator_DefaultValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultValue'
type MockField_generator_DefaultValue_Call struct {
	*mock.Call
}

// DefaultValue is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) DefaultValue() *MockField_generator_DefaultValue_Call {
	return &MockField_generator_DefaultValue_Call{Call: _e.mock.On("DefaultValue")}
}

func (_c *MockField_generator_DefaultValue_Call) Run(run func()) *MockField_generator_DefaultValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_DefaultValue_Call) Return(_a0 *string) *MockField_generator_DefaultValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_DefaultValue_Call) RunAndReturn(run func() *string) *MockField_generator_DefaultValue_Call {
	_c.Call.Return(run)
	return _c
}

// DefaultValueComment provides a mock function with no fields
func (_m *MockField_generator) DefaultValueComment() *string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DefaultValueComment")
	}

	var r0 *string
	if rf, ok := ret.Get(0).(func() *string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	return r0
}

// MockField_generator_DefaultValueComment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DefaultValueComment'
type MockField_generator_DefaultValueComment_Call struct {
	*mock.Call
}

// DefaultValueComment is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) DefaultValueComment() *MockField_generator_DefaultValueComment_Call {
	return &MockField_generator_DefaultValueComment_Call{Call: _e.mock.On("DefaultValueComment")}
}

func (_c *MockField_generator_DefaultValueComment_Call) Run(run func()) *MockField_generator_DefaultValueComment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_DefaultValueComment_Call) Return(_a0 *string) *MockField_generator_DefaultValueComment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_DefaultValueComment_Call) RunAndReturn(run func() *string) *MockField_generator_DefaultValueComment_Call {
	_c.Call.Return(run)
	return _c
}

// Description provides a mock function with no fields
func (_m *MockField_generator) Description() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Description")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockField_generator_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type MockField_generator_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) Description() *MockField_generator_Description_Call {
	return &MockField_generator_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *MockField_generator_Description_Call) Run(run func()) *MockField_generator_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_Description_Call) Return(_a0 string) *MockField_generator_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_Description_Call) RunAndReturn(run func() string) *MockField_generator_Description_Call {
	_c.Call.Return(run)
	return _c
}

// EnumTypeName provides a mock function with no fields
func (_m *MockField_generator) EnumTypeName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnumTypeName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockField_generator_EnumTypeName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnumTypeName'
type MockField_generator_EnumTypeName_Call struct {
	*mock.Call
}

// EnumTypeName is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) EnumTypeName() *MockField_generator_EnumTypeName_Call {
	return &MockField_generator_EnumTypeName_Call{Call: _e.mock.On("EnumTypeName")}
}

func (_c *MockField_generator_EnumTypeName_Call) Run(run func()) *MockField_generator_EnumTypeName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_EnumTypeName_Call) Return(_a0 string) *MockField_generator_EnumTypeName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_EnumTypeName_Call) RunAndReturn(run func() string) *MockField_generator_EnumTypeName_Call {
	_c.Call.Return(run)
	return _c
}

// EnumValues provides a mock function with no fields
func (_m *MockField_generator) EnumValues() []*generator.EnumValue {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for EnumValues")
	}

	var r0 []*generator.EnumValue
	if rf, ok := ret.Get(0).(func() []*generator.EnumValue); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*generator.EnumValue)
		}
	}

	return r0
}

// MockField_generator_EnumValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnumValues'
type MockField_generator_EnumValues_Call struct {
	*mock.Call
}

// EnumValues is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) EnumValues() *MockField_generator_EnumValues_Call {
	return &MockField_generator_EnumValues_Call{Call: _e.mock.On("EnumValues")}
}

func (_c *MockField_generator_EnumValues_Call) Run(run func()) *MockField_generator_EnumValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_EnumValues_Call) Return(_a0 []*generator.EnumValue) *MockField_generator_EnumValues_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_EnumValues_Call) RunAndReturn(run func() []*generator.EnumValue) *MockField_generator_EnumValues_Call {
	_c.Call.Return(run)
	return _c
}

// IsEnum provides a mock function with no fields
func (_m *MockField_generator) IsEnum() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsEnum")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockField_generator_IsEnum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsEnum'
type MockField_generator_IsEnum_Call struct {
	*mock.Call
}

// IsEnum is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) IsEnum() *MockField_generator_IsEnum_Call {
	return &MockField_generator_IsEnum_Call{Call: _e.mock.On("IsEnum")}
}

func (_c *MockField_generator_IsEnum_Call) Run(run func()) *MockField_generator_IsEnum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_IsEnum_Call) Return(_a0 bool) *MockField_generator_IsEnum_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_IsEnum_Call) RunAndReturn(run func() bool) *MockField_generator_IsEnum_Call {
	_c.Call.Return(run)
	return _c
}

// IsRequired provides a mock function with no fields
func (_m *MockField_generator) IsRequired() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsRequired")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockField_generator_IsRequired_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsRequired'
type MockField_generator_IsRequired_Call struct {
	*mock.Call
}

// IsRequired is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) IsRequired() *MockField_generator_IsRequired_Call {
	return &MockField_generator_IsRequired_Call{Call: _e.mock.On("IsRequired")}
}

func (_c *MockField_generator_IsRequired_Call) Run(run func()) *MockField_generator_IsRequired_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_IsRequired_Call) Return(_a0 bool) *MockField_generator_IsRequired_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_IsRequired_Call) RunAndReturn(run func() bool) *MockField_generator_IsRequired_Call {
	_c.Call.Return(run)
	return _c
}

// JSONName provides a mock function with no fields
func (_m *MockField_generator) JSONName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for JSONName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockField_generator_JSONName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'JSONName'
type MockField_generator_JSONName_Call struct {
	*mock.Call
}

// JSONName is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) JSONName() *MockField_generator_JSONName_Call {
	return &MockField_generator_JSONName_Call{Call: _e.mock.On("JSONName")}
}

func (_c *MockField_generator_JSONName_Call) Run(run func()) *MockField_generator_JSONName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_JSONName_Call) Return(_a0 string) *MockField_generator_JSONName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_JSONName_Call) RunAndReturn(run func() string) *MockField_generator_JSONName_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with no fields
func (_m *MockField_generator) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockField_generator_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockField_generator_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) Name() *MockField_generator_Name_Call {
	return &MockField_generator_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockField_generator_Name_Call) Run(run func()) *MockField_generator_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_Name_Call) Return(_a0 string) *MockField_generator_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_Name_Call) RunAndReturn(run func() string) *MockField_generator_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Type provides a mock function with no fields
func (_m *MockField_generator) Type() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockField_generator_Type_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Type'
type MockField_generator_Type_Call struct {
	*mock.Call
}

// Type is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) Type() *MockField_generator_Type_Call {
	return &MockField_generator_Type_Call{Call: _e.mock.On("Type")}
}

func (_c *MockField_generator_Type_Call) Run(run func()) *MockField_generator_Type_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_Type_Call) Return(_a0 string) *MockField_generator_Type_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_Type_Call) RunAndReturn(run func() string) *MockField_generator_Type_Call {
	_c.Call.Return(run)
	return _c
}

// ValidationRules provides a mock function with no fields
func (_m *MockField_generator) ValidationRules() []*generator.ValidationRule {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ValidationRules")
	}

	var r0 []*generator.ValidationRule
	if rf, ok := ret.Get(0).(func() []*generator.ValidationRule); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*generator.ValidationRule)
		}
	}

	return r0
}

// MockField_generator_ValidationRules_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidationRules'
type MockField_generator_ValidationRules_Call struct {
	*mock.Call
}

// ValidationRules is a helper method to define mock.On call
func (_e *MockField_generator_Expecter) ValidationRules() *MockField_generator_ValidationRules_Call {
	return &MockField_generator_ValidationRules_Call{Call: _e.mock.On("ValidationRules")}
}

func (_c *MockField_generator_ValidationRules_Call) Run(run func()) *MockField_generator_ValidationRules_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockField_generator_ValidationRules_Call) Return(_a0 []*generator.ValidationRule) *MockField_generator_ValidationRules_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockField_generator_ValidationRules_Call) RunAndReturn(run func() []*generator.ValidationRule) *MockField_generator_ValidationRules_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockField_generator creates a new instance of MockField_generator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockField_generator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockField_generator {
	mock := &MockField_generator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
