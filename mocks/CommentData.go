// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	comment "simple-social-media-API/features/comment"

	mock "github.com/stretchr/testify/mock"
)

// CommentData is an autogenerated mock type for the CommentData type
type CommentData struct {
	mock.Mock
}

// Add provides a mock function with given fields: newComment
func (_m *CommentData) Add(newComment comment.Core) (comment.Core, error) {
	ret := _m.Called(newComment)

	var r0 comment.Core
	if rf, ok := ret.Get(0).(func(comment.Core) comment.Core); ok {
		r0 = rf(newComment)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comment.Core) error); ok {
		r1 = rf(newComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userId, postId, commentId
func (_m *CommentData) Delete(userId uint, postId uint, commentId uint) error {
	ret := _m.Called(userId, postId, commentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, uint) error); ok {
		r0 = rf(userId, postId, commentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetComments provides a mock function with given fields:
func (_m *CommentData) GetComments() ([]comment.Core, error) {
	ret := _m.Called()

	var r0 []comment.Core
	if rf, ok := ret.Get(0).(func() []comment.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comment.Core)
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

// Update provides a mock function with given fields: updComment
func (_m *CommentData) Update(updComment comment.Core) (comment.Core, error) {
	ret := _m.Called(updComment)

	var r0 comment.Core
	if rf, ok := ret.Get(0).(func(comment.Core) comment.Core); ok {
		r0 = rf(updComment)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comment.Core) error); ok {
		r1 = rf(updComment)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentData interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentData creates a new instance of CommentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentData(t mockConstructorTestingTNewCommentData) *CommentData {
	mock := &CommentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}