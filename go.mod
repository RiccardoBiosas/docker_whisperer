module docker_whisperer

go 1.16

require (
	github.com/containerd/containerd v1.5.1 // indirect
	github.com/docker/docker v20.10.6+incompatible
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/go-kit/kit v0.9.0
	github.com/sirupsen/logrus v1.8.1 // indirect
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200910180754-dd1b699fc489
	google.golang.org/grpc v1.37.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0 // indirect
	google.golang.org/protobuf v1.28.0
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kubernetes v1.13.0
)
