package auth

import "errors"

var ErrNoUserWithId = errors.New("no user with this id")
var ErrNoUserWithUsername = errors.New("no username with this username")
var ErrNoUserWithEmail = errors.New("un registered email")
var ErrFailToDecode = errors.New("fail to decode")
var ErrCursorDuringItr = errors.New("error occur during iteration")
var ErrFailToDelete = errors.New("fail to delert the doucument")
var FailToCreateUser = errors.New("fail to create user")
var ErrUnidentifiedToken = errors.New("unidentified token")
