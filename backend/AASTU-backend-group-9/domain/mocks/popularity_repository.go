package mocks

import (
    "context"

    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "blog/domain"
)

type PopularityRepository struct {
    mock.Mock
}

func (_m *PopularityRepository) HasUserLiked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, error) {
    ret := _m.Called(ctx, id, userID)

    var r0 bool
    if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID) bool); ok {
        r0 = rf(ctx, id, userID)
    } else {
        r0 = ret.Get(0).(bool)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, primitive.ObjectID) error); ok {
        r1 = rf(ctx, id, userID)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *PopularityRepository) HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) (bool, error) {
    ret := _m.Called(ctx, id, userID)

    var r0 bool
    if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID) bool); ok {
        r0 = rf(ctx, id, userID)
    } else {
        r0 = ret.Get(0).(bool)
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, primitive.ObjectID) error); ok {
        r1 = rf(ctx, id, userID)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *PopularityRepository) UserInteractionsAdder(ctx context.Context, user domain.UserInteraction) error {
    ret := _m.Called(ctx, user)

    var r0 error
    if rf, ok := ret.Get(0).(func(context.Context, domain.UserInteraction) error); ok {
        r0 = rf(ctx, user)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

func (_m *PopularityRepository) UserInteractionsDelete(ctx context.Context, user domain.UserInteraction) error {
    ret := _m.Called(ctx, user)

    var r0 error
    if rf, ok := ret.Get(0).(func(context.Context, domain.UserInteraction) error); ok {
        r0 = rf(ctx, user)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}