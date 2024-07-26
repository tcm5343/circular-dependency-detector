module github.com/tcm5343/circular-dependency-detector

go 1.22

require (
	gonum.org/v1/gonum v0.15.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
)

// replace github.com/tcm5343/circular-dependency-detector/pkg/dot => ./pkg/dot
