module github.com/rajabhishekmaurya/ecommerce-microservices/order-service

go 1.25.0

require github.com/rajabhishekmaurya/ecommerce-microservices/common v0.0.0

replace github.com/rajabhishekmaurya/ecommerce-microservices/common => ../common

require (
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/labstack/echo/v4 v4.15.4
	google.golang.org/grpc v1.82.0
)

require (
	github.com/labstack/gommon v0.5.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/rajabhishekmaurya/ecommerce-microservices/api-gateway v0.0.0-20260714170004-1f5ea64edf54
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260414002931-afd174a4e478 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
