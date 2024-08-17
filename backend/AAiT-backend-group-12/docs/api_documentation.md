# Getting Started



`localhost:8080/api/v1/blog`
`localhost:8080/api/v1/auth`

### Auth

- `localhost:8080/api/v1/auth/signup` 3.1.1
- `localhost:8080/api/v1/auth/login` 3.1.2 
- `localhost:8080/api/v1/auth/forget-password` 3.1.3
- `localhost:8080/api/v1/auth/logout` 3.1.4
- `localhost:8080/api/v1/auth/promotiom` 3.1.5
- `localhost:8080/api/v1/auth/demotion` 3.1.6
- `localhost:8080/api/v1/auth/update` 3.1.7

### Blog API

- `localhost:8080/api/v1/blog` POST for get

Pagination:
- Page number
- Number of documents per page

Search by:
- title
- author name

Filter by:
- tag
- date
- dislike count
- like count
- comments count
- view count

Sort by:
- date
- like count
- dislike count

### AI:
- suggest improvement: takes content and title -> performs improvement ands back the data
- create content: takes content title -> writes the blog content
- generate content ideas: query for a good blog topic