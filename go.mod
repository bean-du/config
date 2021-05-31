module coco-tool/config

go 1.15

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.1.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/hashicorp/consul/api v1.8.1 // indirect
	github.com/urfave/cli/v2 v2.3.0
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/genproto v0.0.0-20201214200347-8c77b98c765d // indirect
	gorm.io/driver/postgres v1.0.5
	gorm.io/driver/sqlite v1.1.4 // indirect
	gorm.io/gorm v1.20.8
)
