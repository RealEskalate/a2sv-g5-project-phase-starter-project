package er

var (
	// Blog Title is shorter than allowed.
	TitleTooShort = NewValidation("blog title is too short.")

	// Blog Title is longer than allowed.
	TitleTooLong = NewValidation("blog title is too long.")

	// Blog Content is shorter than allowed.
	ContentTooShort = NewValidation("blog content is too short.")

	// Blog Content is longer than allowed.
	ContentTooLong = NewValidation("blog content is too long.")
)

// NotFound errors
var (
	// Blog does not exist.
	BlogNotFound = NewNotFound("Blog not found.")
)
