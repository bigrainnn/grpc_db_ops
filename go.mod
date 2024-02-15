module github.com/bigrainnn/grpc_db_ops

go 1.20

require (
	github.com/bigrainnn/db_ops_common v1.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/pelletier/go-toml/v2 v2.0.6
	go.uber.org/zap v1.24.0
	google.golang.org/grpc v1.61.0
	google.golang.org/protobuf v1.32.0
	gorm.io/driver/mysql v1.5.0
	gorm.io/gorm v1.24.7-0.20230306060331-85eaf9eeda11
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240125205218-1f4bbc51befe // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240125205218-1f4bbc51befe // indirect
)

replace github.com/bigrainnn/db_ops_common => /Users/dygao/GolandProjects/db_ops
