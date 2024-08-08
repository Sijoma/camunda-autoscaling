package scalingclient

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/sijoma/camunda-scaling-operator/pkg/scalingclient/zbmgmt"
)

type ZeebeMgmtClient struct {
	api *zbmgmt.APIClient
}

// Todo: Pass in options not a service
func NewZeebeMgmtClient(gwSvc corev1.Service) *ZeebeMgmtClient {
	cfg := zbmgmt.NewConfiguration()
	cfg.Scheme = "http"
	cfg.Host = fmt.Sprintf("%s.%s:%d", gwSvc.Name, gwSvc.Namespace, gwSvc.Spec.Ports[0].Port)
	api := zbmgmt.NewAPIClient(cfg)
	zeebeClient := ZeebeMgmtClient{
		api,
	}
	return &zeebeClient
}

func (z ZeebeMgmtClient) SendScaleRequest(ctx context.Context, brokerIds []int32) error {
	logger := log.FromContext(ctx)
	operation, resp, err := z.api.DefaultAPI.BrokersPost(ctx).RequestBody(brokerIds).Execute()
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("sendScaleRequest: scaling failed with code %d", resp.StatusCode)
	}

	logger.Info("sendScaleRequest: scaling succeeded", "changeId", operation.ChangeId)

	return nil
}

func (z ZeebeMgmtClient) Topology(ctx context.Context) (*zbmgmt.GetTopologyResponse, error) {
	topology, resp, err := z.api.DefaultAPI.RootGet(ctx).Execute()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Topology: failed with code %d", resp.StatusCode)
	}

	return topology, nil
}
