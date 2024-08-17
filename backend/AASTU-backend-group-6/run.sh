# Run the application
go get github.com/air-verse/air

export PATH=$PATH:$(go env GOPATH)/bin
fuser -k 8080/tcp
air