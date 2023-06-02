package client

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/longhorn/longhorn-instance-manager/pkg/api"
	rpc "github.com/longhorn/longhorn-instance-manager/pkg/imrpc"
	"github.com/longhorn/longhorn-instance-manager/pkg/meta"
	"github.com/longhorn/longhorn-instance-manager/pkg/types"
	"github.com/longhorn/longhorn-instance-manager/pkg/util"
)

type DiskServiceContext struct {
	cc      *grpc.ClientConn
	service rpc.DiskServiceClient
}

func (c *DiskServiceClient) Close() error {
	if c.cc == nil {
		return nil
	}
	return c.cc.Close()
}

func (c *DiskServiceClient) getControllerServiceClient() rpc.DiskServiceClient {
	return c.service
}

type DiskServiceClient struct {
	serviceURL string
	tlsConfig  *tls.Config
	DiskServiceContext
}

func NewDiskServiceClient(serviceURL string, tlsConfig *tls.Config) (*DiskServiceClient, error) {
	getDiskServiceContext := func(serviceUrl string, tlsConfig *tls.Config) (DiskServiceContext, error) {
		connection, err := util.Connect(serviceUrl, tlsConfig)
		if err != nil {
			return DiskServiceContext{}, errors.Wrapf(err, "cannot connect to Disk Service %v", serviceUrl)
		}

		return DiskServiceContext{
			cc:      connection,
			service: rpc.NewDiskServiceClient(connection),
		}, nil
	}

	serviceContext, err := getDiskServiceContext(serviceURL, tlsConfig)
	if err != nil {
		return nil, err
	}

	return &DiskServiceClient{
		serviceURL:         serviceURL,
		tlsConfig:          tlsConfig,
		DiskServiceContext: serviceContext,
	}, nil
}

func NewDiskServiceClientWithTLS(serviceURL, caFile, certFile, keyFile, peerName string) (*DiskServiceClient, error) {
	tlsConfig, err := util.LoadClientTLS(caFile, certFile, keyFile, peerName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load tls key pair from file")
	}

	return NewDiskServiceClient(serviceURL, tlsConfig)
}

func (c *DiskServiceClient) DiskCreate(diskType, diskName, diskPath string, blockSize int64) (*api.DiskInfo, error) {
	if diskName == "" || diskPath == "" {
		return nil, fmt.Errorf("failed to create disk: missing required parameters")
	}

	t, ok := rpc.DiskType_value[diskType]
	if !ok {
		return nil, fmt.Errorf("failed to get disk info: invalid disk type %v", diskType)
	}

	client := c.getControllerServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	resp, err := client.DiskCreate(ctx, &rpc.DiskCreateRequest{
		DiskType:  rpc.DiskType(t),
		DiskName:  diskName,
		DiskPath:  diskPath,
		BlockSize: blockSize,
	})
	if err != nil {
		return nil, err
	}

	return &api.DiskInfo{
		ID:          resp.GetId(),
		UUID:        resp.GetUuid(),
		Path:        resp.GetPath(),
		Type:        resp.GetType(),
		TotalSize:   resp.GetTotalSize(),
		FreeSize:    resp.GetFreeSize(),
		TotalBlocks: resp.GetTotalBlocks(),
		FreeBlocks:  resp.GetFreeBlocks(),
		BlockSize:   resp.GetBlockSize(),
		ClusterSize: resp.GetClusterSize(),
	}, nil
}

func (c *DiskServiceClient) DiskGet(diskType, diskName, diskPath string) (*api.DiskInfo, error) {
	if diskName == "" || diskPath == "" {
		return nil, fmt.Errorf("failed to get disk info: missing required parameters")
	}

	t, ok := rpc.DiskType_value[diskType]
	if !ok {
		return nil, fmt.Errorf("failed to get disk info: invalid disk type %v", diskType)
	}

	client := c.getControllerServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	resp, err := client.DiskGet(ctx, &rpc.DiskGetRequest{
		DiskType: rpc.DiskType(t),
		DiskName: diskName,
		DiskPath: diskPath,
	})
	if err != nil {
		return nil, err
	}

	return &api.DiskInfo{
		ID:          resp.GetId(),
		UUID:        resp.GetUuid(),
		Path:        resp.GetPath(),
		Type:        resp.GetType(),
		TotalSize:   resp.GetTotalSize(),
		FreeSize:    resp.GetFreeSize(),
		TotalBlocks: resp.GetTotalBlocks(),
		FreeBlocks:  resp.GetFreeBlocks(),
		BlockSize:   resp.GetBlockSize(),
		ClusterSize: resp.GetClusterSize(),
	}, nil
}

func (c *DiskServiceClient) DiskDelete(diskType, diskName, diskUUID string) error {
	if diskName == "" || diskUUID == "" {
		return fmt.Errorf("failed to delete disk: missing required parameters")
	}

	client := c.getControllerServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	_, err := client.DiskDelete(ctx, &rpc.DiskDeleteRequest{
		DiskType: rpc.DiskType(rpc.DiskType_value[diskType]),
		DiskName: diskName,
		DiskUuid: diskUUID,
	})
	return err
}

func (c *DiskServiceClient) DiskReplicaInstanceList(diskType, diskName string) (map[string]*api.ReplicaStorageInstance, error) {
	if diskName == "" {
		return nil, fmt.Errorf("failed to list replica instances on disk: missing required parameter")
	}

	client := c.getControllerServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	resp, err := client.DiskReplicaInstanceList(ctx, &rpc.DiskReplicaInstanceListRequest{
		DiskType: rpc.DiskType(rpc.DiskType_value[diskType]),
		DiskName: diskName,
	})
	if err != nil {
		return nil, err
	}

	instances := map[string]*api.ReplicaStorageInstance{}
	for name, instance := range resp.ReplicaInstances {
		instances[name] = &api.ReplicaStorageInstance{
			Name:       instance.Name,
			UUID:       instance.Uuid,
			DiskName:   instance.DiskName,
			DiskUUID:   instance.DiskUuid,
			SpecSize:   instance.SpecSize,
			ActualSize: instance.ActualSize,
		}
	}

	return instances, nil
}

func (c *DiskServiceClient) DiskReplicaInstanceDelete(diskType, diskName, diskUUID, replciaInstanceName string) error {
	if diskName == "" || diskUUID == "" || replciaInstanceName == "" {
		return fmt.Errorf("failed to delete replica instance on disk: missing required parameters")
	}

	client := c.getControllerServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	_, err := client.DiskReplicaInstanceDelete(ctx, &rpc.DiskReplicaInstanceDeleteRequest{
		DiskType:            rpc.DiskType(rpc.DiskType_value[diskType]),
		DiskName:            diskName,
		DiskUuid:            diskUUID,
		ReplciaInstanceName: replciaInstanceName,
	})

	return err
}

func (c *DiskServiceClient) VersionGet() (*meta.DiskServiceVersionOutput, error) {
	client := c.getControllerServiceClient()
	ctx, cancel := context.WithTimeout(context.Background(), types.GRPCServiceTimeout)
	defer cancel()

	resp, err := client.VersionGet(ctx, &empty.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get disk service version")
	}

	return &meta.DiskServiceVersionOutput{
		Version:   resp.Version,
		GitCommit: resp.GitCommit,
		BuildDate: resp.BuildDate,

		InstanceManagerDiskServiceAPIVersion:    int(resp.InstanceManagerDiskServiceAPIVersion),
		InstanceManagerDiskServiceAPIMinVersion: int(resp.InstanceManagerDiskServiceAPIMinVersion),
	}, nil
}
