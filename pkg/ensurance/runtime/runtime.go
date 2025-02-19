package runtime

import (
	"fmt"

	"google.golang.org/grpc"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
	"k8s.io/klog/v2"

	gprcconnection "github.com/gocrane/crane/pkg/ensurance/grpc"
)

// runtimeEndpoint is CRI server runtime endpoint
func getRuntimeClientConnection(runtimeEndpoint string) (*grpc.ClientConn, error) {
	var runtimeEndpoints []string
	if runtimeEndpoint != "" {
		runtimeEndpoints = append(runtimeEndpoints, runtimeEndpoint)
	}
	runtimeEndpoints = append(runtimeEndpoints, defaultRuntimeEndpoints...)
	klog.V(2).Infof("Runtime connect using endpoints: %v. You can set the endpoint instead.", defaultRuntimeEndpoints)
	return gprcconnection.InitGrpcConnection(runtimeEndpoints)
}

// imageEndpoint is CRI server image endpoint, default same as runtime endpoint
func getImageClientConnection(imageEndpoint string) (*grpc.ClientConn, error) {
	var imageEndpoints []string
	if imageEndpoint != "" {
		imageEndpoints = append(imageEndpoints, imageEndpoint)
	}
	imageEndpoints = append(imageEndpoints, defaultRuntimeEndpoints...)
	klog.V(2).Infof(fmt.Sprintf("Image connect using endpoints: %v. You should set the endpoint instead.", imageEndpoints))
	return gprcconnection.InitGrpcConnection(imageEndpoints)
}

// GetRuntimeClient get the runtime client
func GetRuntimeClient(runtimeEndpoint string) (pb.RuntimeServiceClient, *grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := getRuntimeClientConnection(runtimeEndpoint)
	if err != nil {
		return nil, nil, err
	}
	runtimeClient := pb.NewRuntimeServiceClient(conn)
	return runtimeClient, conn, nil
}

// GetImageClient get the runtime client
func GetImageClient(imageEndpoint string) (pb.ImageServiceClient, *grpc.ClientConn, error) {
	// Set up a connection to the server.
	conn, err := getImageClientConnection(imageEndpoint)
	if err != nil {
		return nil, nil, err
	}
	imageClient := pb.NewImageServiceClient(conn)
	return imageClient, conn, nil
}
