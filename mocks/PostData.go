// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	post "simple-social-media-API/features/post"

	mock "github.com/stretchr/testify/mock"
)

// PostData is an autogenerated mock type for the PostData type
type PostData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, newPost
func (_m *PostData) Add(userID uint, newPost post.Core) (post.Core, error) {
	ret := _m.Called(userID, newPost)

	var r0 post.Core
	if rf, ok := ret.Get(0).(func(uint, post.Core) post.Core); ok {
		r0 = rf(userID, newPost)
	} else {
		r0 = ret.Get(0).(post.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, post.Core) error); ok {
		r1 = rf(userID, newPost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AllPosts provides a mock function with given fields:
func (_m *PostData) AllPosts() ([]post.MyPostsResp, error) {
	ret := _m.Called()

	var r0 []post.MyPostsResp
	if rf, ok := ret.Get(0).(func() []post.MyPostsResp); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]post.MyPostsResp)
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

// Delete provides a mock function with given fields: postID, userID
func (_m *PostData) Delete(postID uint, userID uint) error {
	ret := _m.Called(postID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(postID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPostById provides a mock function with given fields: postID, userID
func (_m *PostData) GetPostById(postID uint, userID uint) (post.MyPostsResp, error) {
	ret := _m.Called(postID, userID)

	var r0 post.MyPostsResp
	if rf, ok := ret.Get(0).(func(uint, uint) post.MyPostsResp); ok {
		r0 = rf(postID, userID)
	} else {
		r0 = ret.Get(0).(post.MyPostsResp)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(postID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyPosts provides a mock function with given fields: userID
func (_m *PostData) MyPosts(userID uint) ([]post.MyPostsResp, error) {
	ret := _m.Called(userID)

	var r0 []post.MyPostsResp
	if rf, ok := ret.Get(0).(func(uint) []post.MyPostsResp); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]post.MyPostsResp)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: postID, userID, updatedPost
func (_m *PostData) Update(postID uint, userID uint, updatedPost post.Core) (post.Core, error) {
	ret := _m.Called(postID, userID, updatedPost)

	var r0 post.Core
	if rf, ok := ret.Get(0).(func(uint, uint, post.Core) post.Core); ok {
		r0 = rf(postID, userID, updatedPost)
	} else {
		r0 = ret.Get(0).(post.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, post.Core) error); ok {
		r1 = rf(postID, userID, updatedPost)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostData interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostData creates a new instance of PostData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostData(t mockConstructorTestingTNewPostData) *PostData {
	mock := &PostData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
