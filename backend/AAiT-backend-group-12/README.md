#

#### Coverage Profile on Controllers
```bash
go test -v -coverprofile="cover.out" -coverpkg="blog_api/controller"  blog_api/tests
```

#### Coverage Profile on Usecases
```bash
go test -v -coverprofile="cover.out" -coverpkg="blog_api/usecase"  blog_api/tests
```

#### Coverage Profile on Repositories
```bash
go test -v -coverprofile="cover.out" -coverpkg="blog_api/repository"  blog_api/tests
```

#### Coverage Profile Rendering
```bash
go tool cover -html="cover.out" -o coverage.html
```