module github.com/zackzhangkai/grpc-pubsub-example

go 1.20

//replace github.com/zackzhangkai/grpc-pubsub-example => ../

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	k8s.io/klog/v2 v2.120.1
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
)
