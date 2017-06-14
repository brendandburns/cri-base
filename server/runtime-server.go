package server

import (
	"github.com/brendandburns/cri-base/generated"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

type RuntimeServer struct {
}

// Version returns the runtime name, runtime version, and runtime API version.
func (s *RuntimeServer) Version(ctx context.Context, req *generated.VersionRequest) (*generated.VersionResponse, error) {
	glog.Infof("Version request: %v\n", *req)
	return &generated.VersionResponse{
		Version:           "v1alpha1",
		RuntimeName:       "cri-base",
		RuntimeVersion:    "0.1.0",
		RuntimeApiVersion: "0.1.0",
	}, nil
}

// RunPodSandbox creates and starts a pod-level sandbox. Runtimes must ensure
// the sandbox is in the ready state on success.
func (s *RuntimeServer) RunPodSandbox(ctx context.Context, req *generated.RunPodSandboxRequest) (*generated.RunPodSandboxResponse, error) {
	glog.Infof("RunPodSandbox request: %v\n", *req)
	return &generated.RunPodSandboxResponse{}, nil
}

// StopPodSandbox stops any running process that is part of the sandbox and
// reclaims network resources (e.g., IP addresses) allocated to the sandbox.
// If there are any running containers in the sandbox, they must be forcibly
// terminated.
// This call is idempotent, and must not return an error if all relevant
// resources have already been reclaimed. kubelet will call StopPodSandbox
// at least once before calling RemovePodSandbox. It will also attempt to
// reclaim resources eagerly, as soon as a sandbox is not needed. Hence,
// multiple StopPodSandbox calls are expected.
func (s *RuntimeServer) StopPodSandbox(ctx context.Context, req *generated.StopPodSandboxRequest) (*generated.StopPodSandboxResponse, error) {
	glog.Infof("StopPodSandbox request: %v\n", *req)
	return &generated.StopPodSandboxResponse{}, nil
}

// RemovePodSandbox removes the sandbox. If there are any running containers
// in the sandbox, they must be forcibly terminated and removed.
// This call is idempotent, and must not return an error if the sandbox has
// already been removed.
func (s *RuntimeServer) RemovePodSandbox(ctx context.Context, req *generated.RemovePodSandboxRequest) (*generated.RemovePodSandboxResponse, error) {
	glog.Infof("RemovePodSandbox request: %v\n", *req)
	return nil, nil
}

// PodSandboxStatus returns the status of the PodSandbox. If the PodSandbox is not
// present, returns an error.
func (s *RuntimeServer) PodSandboxStatus(ctx context.Context, req *generated.PodSandboxStatusRequest) (*generated.PodSandboxStatusResponse, error) {
	glog.Infof("PodSandboxStatus request: %v\n", *req)
	return &generated.PodSandboxStatusResponse{}, nil
}

// ListPodSandbox returns a list of PodSandboxes.
func (s *RuntimeServer) ListPodSandbox(ctx context.Context, req *generated.ListPodSandboxRequest) (*generated.ListPodSandboxResponse, error) {
	glog.Infof("ListPodSandbox request: %v\n", *req)
	return &generated.ListPodSandboxResponse{}, nil
}

// CreateContainer creates a new container in specified PodSandbox
func (s *RuntimeServer) CreateContainer(ctx context.Context, req *generated.CreateContainerRequest) (*generated.CreateContainerResponse, error) {
	glog.Infof("CreateContainer request: %v\n", *req)
	return &generated.CreateContainerResponse{}, nil
}

// StartContainer starts the container.
func (s *RuntimeServer) StartContainer(ctx context.Context, req *generated.StartContainerRequest) (*generated.StartContainerResponse, error) {
	glog.Infof("StartContainer request: %v\n", *req)
	return &generated.StartContainerResponse{}, nil
}

// StopContainer stops a running container with a grace period (i.e., timeout).
// This call is idempotent, and must not return an error if the container has
// already been stopped.
// TODO: what must the runtime do after the grace period is reached?
func (s *RuntimeServer) StopContainer(ctx context.Context, req *generated.StopContainerRequest) (*generated.StopContainerResponse, error) {
	glog.Infof("StopContainer request: %v\n", *req)
	return &generated.StopContainerResponse{}, nil
}

// RemoveContainer removes the container. If the container is running, the
// container must be forcibly removed.
// This call is idempotent, and must not return an error if the container has
// already been removed.
func (s *RuntimeServer) RemoveContainer(ctx context.Context, req *generated.RemoveContainerRequest) (*generated.RemoveContainerResponse, error) {
	glog.Infof("RemoveContainer request: %v\n", *req)
	return &generated.RemoveContainerResponse{}, nil
}

// ListContainers lists all containers by filters.
func (s *RuntimeServer) ListContainers(ctx context.Context, req *generated.ListContainersRequest) (*generated.ListContainersResponse, error) {
	glog.Infof("ListContainers request: %v\n", *req)
	return &generated.ListContainersResponse{}, nil
}

// ContainerStatus returns status of the container. If the container is not
// present, returns an error.
func (s *RuntimeServer) ContainerStatus(ctx context.Context, req *generated.ContainerStatusRequest) (*generated.ContainerStatusResponse, error) {
	glog.Infof("ContainerStatus request: %v\n", *req)
	return &generated.ContainerStatusResponse{}, nil
}

// ExecSync runs a command in a container synchronously.
func (s *RuntimeServer) ExecSync(ctx context.Context, req *generated.ExecSyncRequest) (*generated.ExecSyncResponse, error) {
	glog.Infof("ExecSync request: %v\n", *req)
	return &generated.ExecSyncResponse{}, nil
}

// Exec prepares a streaming endpoint to execute a command in the container.
func (s *RuntimeServer) Exec(ctx context.Context, req *generated.ExecRequest) (*generated.ExecResponse, error) {
	glog.Infof("Exec request: %v\n", *req)
	return &generated.ExecResponse{}, nil
}

// Attach prepares a streaming endpoint to attach to a running container.
func (s *RuntimeServer) Attach(ctx context.Context, req *generated.AttachRequest) (*generated.AttachResponse, error) {
	glog.Infof("Attach request: %v\n", *req)
	return &generated.AttachResponse{}, nil
}

// PortForward prepares a streaming endpoint to forward ports from a PodSandbox.
func (s *RuntimeServer) PortForward(ctx context.Context, req *generated.PortForwardRequest) (*generated.PortForwardResponse, error) {
	glog.Infof("PortForward request: %v\n", *req)
	return &generated.PortForwardResponse{}, nil
}

// UpdateRuntimeConfig updates the runtime configuration based on the given request.
func (s *RuntimeServer) UpdateRuntimeConfig(ctx context.Context, req *generated.UpdateRuntimeConfigRequest) (*generated.UpdateRuntimeConfigResponse, error) {
	glog.Infof("UpdateRuntime request: %v\n", *req)
	return &generated.UpdateRuntimeConfigResponse{}, nil
}

// Status returns the status of the runtime.
func (s *RuntimeServer) Status(ctx context.Context, req *generated.StatusRequest) (*generated.StatusResponse, error) {
	glog.Infof("Status request: %v\n", *req)
	return &generated.StatusResponse{}, nil
}
