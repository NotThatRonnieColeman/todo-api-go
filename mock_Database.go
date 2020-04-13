// Code generated by mockery v1.0.0. DO NOT EDIT.

package main

import mock "github.com/stretchr/testify/mock"

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

// allItems provides a mock function with given fields:
func (_m *MockDatabase) allItems() ([]Item, error) {
	ret := _m.Called()

	var r0 []Item
	if rf, ok := ret.Get(0).(func() []Item); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Item)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// close provides a mock function with given fields:
func (_m *MockDatabase) close() {
	_m.Called()
}

// createItem provides a mock function with given fields: item
func (_m *MockDatabase) createItem(item Item) (Item, error) {
	ret := _m.Called(item)

	var r0 Item
	if rf, ok := ret.Get(0).(func(Item) Item); ok {
		r0 = rf(item)
	} else {
		r0 = ret.Get(0).(Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Item) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// deleteItem provides a mock function with given fields: id
func (_m *MockDatabase) deleteItem(id uint) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// getItem provides a mock function with given fields: id
func (_m *MockDatabase) getItem(id uint) (Item, error) {
	ret := _m.Called(id)

	var r0 Item
	if rf, ok := ret.Get(0).(func(uint) Item); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// init provides a mock function with given fields:
func (_m *MockDatabase) init() {
	_m.Called()
}

// updateItem provides a mock function with given fields: id, td
func (_m *MockDatabase) updateItem(id uint, td Item) (Item, error) {
	ret := _m.Called(id, td)

	var r0 Item
	if rf, ok := ret.Get(0).(func(uint, Item) Item); ok {
		r0 = rf(id, td)
	} else {
		r0 = ret.Get(0).(Item)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, Item) error); ok {
		r1 = rf(id, td)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
