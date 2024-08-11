package scalingclient

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/sijoma/camunda-scaling-operator/pkg/scalingclient/zbmgmt"
)

type ZeebeMgmtClient struct {
	api *zbmgmt.APIClient
}

// WithHost allows to specify a host, this is useful for local development / port-forwarding scenarios
func WithHost(host string) func(cfg *zbmgmt.Configuration) {
	return func(cfg *zbmgmt.Configuration) {
		cfg.Host = host
	}
}

func WithTimeout(timeout time.Duration) func(cfg *zbmgmt.Configuration) {
	return func(cfg *zbmgmt.Configuration) {
		cfg.HTTPClient.Timeout = timeout
	}
}

func NewZeebeMgmtClient(opts ...func(*zbmgmt.Configuration)) *ZeebeMgmtClient {
	cfg := zbmgmt.NewConfiguration()
	cfg.Scheme = "http"
	// API calls inside the reconcile loop should not block too long
	// (we still only have 1 worker so no concurrency of the loop)
	cfg.HTTPClient = &http.Client{Timeout: time.Second * 15}

	for _, option := range opts {
		option(cfg)
	}

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
