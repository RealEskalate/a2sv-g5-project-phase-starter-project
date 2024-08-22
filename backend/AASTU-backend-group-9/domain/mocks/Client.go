package mocks

import (
    "context"
    "blog/database"

    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/mongo"
)

// Client is a manually created mock type for the Client type
type Client struct {
    mock.Mock
}
type mockConstructorTestingTNewClient interface {
    mock.TestingT
    Cleanup(func())
}

// Connect provides a mock function with given fields: _a0
func (_m *Client) Connect(_a0 context.Context) error {
    ret := _m.Called(_a0)

    var r0 error
    if rf, ok := ret.Get(0).(func(context.Context) error); ok {
        r0 = rf(_a0)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

// Database provides a mock function with given fields: _a0
func (_m *Client) Database(_a0 string) database.Database {
    ret := _m.Called(_a0)

    var r0 database.Database
    if rf, ok := ret.Get(0).(func(string) database.Database); ok {
        r0 = rf(_a0)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(database.Database)
        }
    }

    return r0
}

// Disconnect provides a mock function with given fields: _a0
func (_m *Client) Disconnect(_a0 context.Context) error {
    ret := _m.Called(_a0)

    var r0 error
    if rf, ok := ret.Get(0).(func(context.Context) error); ok {
        r0 = rf(_a0)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

// Ping provides a mock function with given fields: _a0
func (_m *Client) Ping(_a0 context.Context) error {
    ret := _m.Called(_a0)

    var r0 error
    if rf, ok := ret.Get(0).(func(context.Context) error); ok {
        r0 = rf(_a0)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

// StartSession provides a mock function with given fields:
func (_m *Client) StartSession() (mongo.Session, error) {
    ret := _m.Called()

    var r0 mongo.Session
    if rf, ok := ret.Get(0).(func() mongo.Session); ok {
        r0 = rf()
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(mongo.Session)
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

// UseSession provides a mock function with given fields: ctx, fn
func (_m *Client) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
    ret := _m.Called(ctx, fn)

    var r0 error
    if rf, ok := ret.Get(0).(func(context.Context, func(mongo.SessionContext) error) error); ok {
        r0 = rf(ctx, fn)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
    mock := &Client{}
    mock.Mock.Test(t)

    t.Cleanup(func() { mock.AssertExpectations(t) })

    return mock
}