module github.com/igorok-follow/analytics-service

go 1.16

require (
	github.com/daviddengcn/go-colortext v1.0.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/gorilla/securecookie v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.9
	github.com/rs/cors v1.8.0
	gitlab.tarology.me/tarodev/tarology-service v0.0.0-20230531110232-05661e09da01
	go.opentelemetry.io/otel v1.15.0
	go.opentelemetry.io/otel/exporters/jaeger v1.15.0
	go.opentelemetry.io/otel/sdk v1.15.0
	go.opentelemetry.io/otel/trace v1.15.0
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v3 v3.0.1
)
