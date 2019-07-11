module foo

go 1.12

require (
	git.ucloudadmin.com/udb/micro/common v0.0.0-20190610021846-1c86d1bad26b
	github.com/opentracing/opentracing-go v1.1.0
	go.uber.org/zap v1.10.0
)

replace git.ucloudadmin.com/udb/micro/common => /media/storage/UCloud/Project/udb-micro/common

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.0.0-20181115231424-8e868ca12c0f

replace github.com/golang/lint => github.com/golang/lint v0.0.0-20190227174305-8f45f776aaf1
