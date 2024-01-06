module github.com/ilyatbn/keymv-dbengine

go 1.18

require (
	github.com/lithammer/shortuuid/v4 v4.0.0
	golang.org/x/net v0.0.0-20201021035429-f5854403a974
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/gocql/gocql v1.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/ilyatbn/keymv-proto v0.0.0-20220421093344-551e433bd7af // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/ilyatbn/keymv-proto v0.0.0-20220421093344-551e433bd7af => ../keymv-proto
