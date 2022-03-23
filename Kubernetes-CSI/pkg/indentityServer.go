package pkg

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type archetype struct {
	name     string
	nodeID   string
	version  string
	endpoint string

	// Add CSI plugin parameters here
	parameter1 string
	parameter2 int
	parameter3 time.Duration

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

type IdentityServer struct {
	Driver *archetype
}

// GetPluginIno 获取插件名称和版本
// 版本号只能是 1.1.1 不能是 v1.1.1 格式
func (ids *IdentityServer) GetPluginIno(_ context.Context, _ *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginInfoResponse, error) {
	if ids.Driver.name == "" {
		return nil, status.Error(codes.Unavailable, "Driver name not configured")
	}

	if ids.Driver.version == "" {
		return nil, status.Error(codes.Unavailable, "Driver is missing version")
	}

	return &csi.GetPluginInfoResponse{
		Name:          ids.Driver.name,
		VendorVersion: ids.Driver.version,
	}, nil
}

// Probe 接口用来做健康检测
// TODO 应该需要做一部分的逻辑判活
func (ids *IdentityServer) Probe(_ context.Context, _ *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	return &csi.ProbeResponse{}, nil
}

// GetPluginCapabilities
func (ids *IdentityServer) GetPluginCapabilities(_ context.Context, _ *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	logrus.Infof("Using default capabilities")
	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
		},
	}, nil
}
