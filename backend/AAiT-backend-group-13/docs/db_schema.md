# Schema Documentation

## Blog

Represents a blog post with metadata such as title, content, tags, and counts for likes, dislikes, and comments.

- **id**: `UUID` - The unique identifier for the blog post.
- **title**: `string` - The title of the blog post. Length must be between `3` and `100` characters.
- **content**: `string` - The content of the blog post. Length must be between `50` and `10000` characters.
- **tags**: `[]string` - A list of tags associated with the blog post.
- **createdDate**: `time.Time` - The date and time when the blog post was created.
- **updatedDate**: `time.Time` - The date and time when the blog post was last updated.
- **userID**: `UUID` - The unique identifier of the user who created the blog post.
- **likeCount**: `int` - The number of likes the blog post has received.
- **dislikeCount**: `int` - The number of dislikes the blog post has received.
- **commentCount**: `int` - The number of comments on the blog post.

## Comment

Represents a comment on a blog post.

- **id**: `UUID` - The unique identifier for the comment.
- **content**: `string` - The content of the comment. Length must be between `5` and `500` characters.
- **userID**: `UUID` - The unique identifier of the user who made the comment.
- **blogID**: `UUID` - The unique identifier of the blog post being commented on.

## Reaction

Represents a user's reaction (like or dislike) to a blog post.

- **id**: `UUID` - The unique identifier for the reaction.
- **isLike**: `bool` - Indicates whether the reaction is a like (`true`) or a dislike (`false`).
- **userID**: `UUID` - The unique identifier of the user who made the reaction.
- **blogID**: `UUID` - The unique identifier of the blog post being reacted to.

## ResetCode

Represents a code used for password resets.

- **CodeHash**: `string` - The hashed reset code.
- **Expr**: `time.Time` - The expiration time of the reset code.

## User

Represents a system user.

- **id**: `UUID` - The unique identifier for the user.
- **firstName**: `string` - The user's first name. Length must be between `3` and `250` characters.
- **lastName**: `string` - The user's last name.
- **username**: `string` - The user's username. Must match the pattern `^[a-zA-Z0-9_]+$` and be between `3` and `20` characters in length.
- **email**: `string` - The user's email address. Must match the pattern `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`.
- **passwordHash**: `string` - The hashed password of the user.
- **isAdmin**: `bool` - Indicates whether the user has administrative privileges.
- **resetCode**: `*ResetCode` - An optional reset code used for password recovery.
- **createdAt**: `time.Time` - The date and time when the user account was created.
- **updatedAt**: `time.Time` - The date and time when the user account was last updated.
- **isActive**: `bool` - Indicates whether the user email is verified or not.
