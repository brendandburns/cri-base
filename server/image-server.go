package server

import (
	"github.com/brendandburns/cri-base/generated"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

type ImageServer struct {
}

// ListImages lists existing images.
func (i *ImageServer) ListImages(ctx context.Context, req *generated.ListImagesRequest) (*generated.ListImagesResponse, error) {
     glog.Infof("ListImages %v\n", *req)
     return &generated.ListImagesResponse{}, nil
}

// ImageStatus returns the status of the image. If the image is not
// present, returns a response with ImageStatusResponse.Image set to
// nil.
func (i *ImageServer) ImageStatus(ctx context.Context, req *generated.ImageStatusRequest) (*generated.ImageStatusResponse, error) {
     glog.Infof("ImageStatus %v\n", *req)
     return &generated.ImageStatusResponse{}, nil
}

// PullImage pulls an image with authentication config.
func (i *ImageServer) PullImage(ctx context.Context, req *generated.PullImageRequest) (*generated.PullImageResponse, error) {
     	glog.Infof("PullImage %v\n", *req)
	return &generated.PullImageResponse{}, nil
}

// RemoveImage removes the image.
// This call is idempotent, and must not return an error if the image has
// already been removed.
func (i *ImageServer) RemoveImage(ctx context.Context, req *generated.RemoveImageRequest) (*generated.RemoveImageResponse, error) {
     glog.Infof("RemoveImage %v\n", *req)
	return &generated.RemoveImageResponse{}, nil
}

// ImageFSInfo returns information of the filesystem that is used to store images.
func (i *ImageServer) ImageFsInfo(ctx context.Context, req *generated.ImageFsInfoRequest) (*generated.ImageFsInfoResponse, error) {
     glog.Infof("ImageFsInfo %v\n", *req)
	return &generated.ImageFsInfoResponse{}, nil
}
