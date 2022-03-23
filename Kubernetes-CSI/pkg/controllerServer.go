package pkg

import (
	"context"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/sirupsen/logrus"
	"k8s.io/utils/mount"
)

// Controller Server 实际上对应着 Provisioning and Deleting 阶段(核心的创建/删除卷、快照等都应在此做实现)

type ControllerServer struct {
	Driver *archetype
	// Users add fields as needed.
	//
	// In the NFS CSI implementation, we need to mount the nfs server to the local,
	// so we need a mounter instance.
	//
	// In the CSI implementation of other storage vendors, you may need to add other
	// instances, such as the api client of Alibaba Cloud Storage.
	mounter mount.Interface
}

// ControllerGetCapabilities 返回在创建驱动时设置的 cscap
func (cs *ControllerServer) ControllerGetCapabilities(_ context.Context, _ *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	logrus.Infof("ControllerGetCapabilities: %s", cs.Driver.cscap)
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: cs.Driver.cscap,
	}, nil
}
