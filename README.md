# Go Examples

Based on [Udemy Course](https://www.udemy.com/course/learn-how-to-code/).

[GO Blog](https://go.dev/blog/)


## Commands

### Go Doc

````
go doc fmt
godoc -http=:8080
cd exampleDoc && go doc
cd exampleDoc && go doc exampleDoc.GoDocExample
````


## Go Test

````
go test
go test -bench . -cover ./...
````


## Go Lint

````
go lint
````


## Go Vet

Code correctness.

````
go vet ./...
````


## Go Coverage

````
go test ./... -coverprofile=c.out
go tool cover -html=c.out 
````