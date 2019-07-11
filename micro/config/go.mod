module foo

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/micro/go-micro v1.6.0
)

replace (
	github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
	github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f
)

replace github.com/micro/go-micro => /home/magodo/github/go-micro
