package er

// Validation errors
var (
	// Username is shorter than allowed.
	UsernameTooShort = NewValidation("username is too short.")

	// Username is longer than allowed.
	UsernameTooLong = NewValidation("username is too long.")

	// Password is susceptible to attack.
	WeakPassword = NewValidation("password is too weak.")

	// Username is not UUID.
	UsernameInvalidFormat = NewValidation("username has an invalid format.")

	// Invalid Email
	EmailInvalidFormat = NewValidation("email has an invalid format.")

	// Username is shorter than allowed.
	FirstNameTooShort = NewValidation("first is too short.")

	// Username is longer than allowed.
	FirstNameTooLong = NewValidation("first is too long.")
)

// Conflict errors
var (
	// User with a similar username exists.
	UsernameConflict = NewConflict("username already taken.")
)

// NotFound errors
var (
	// User is does not exist.
	UserNotFound = NewNotFound("User not found.")
)

// Unexpected errors
var (
	// Unexpected error occurred while hashing.
	Hash = NewUnexpected("error hashing password.")
)
