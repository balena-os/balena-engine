package health

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"github.com/containerd/containerd"
	"github.com/containerd/containerd/pkg/dialer"
	"github.com/pkg/errors"
)

// containerdHealth contacts the containerd health service and fails if the
// reported status is not "serving".
type containerdHealth struct{}

func (c containerdHealth) Description() string {
	return "Is containerd healthy"
}


const (
	containerdAddr string = "/var/run/balena-engine/containerd/balena-engine-containerd.sock"
	containerdNs   string = "moby"
)

func (c containerdHealth) Check(ctx context.Context) error {
	cli, err := containerd.New(
		containerdAddr,
		containerd.WithDefaultNamespace(containerdNs),
		containerd.WithDialOpts([]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBackoffMaxDelay(3*time.Second),
			grpc.WithContextDialer(func(_ context.Context, addr string) (net.Conn, error) {
				if deadline, ok := ctx.Deadline(); ok {
					return dialer.Dialer(addr, time.Until(deadline))
				}
				return dialer.Dialer(addr, 0)
			}),
		}),
	)
	if err != nil {
		return err
	}

	hs := cli.HealthService()
	response, err := hs.Check(ctx, &health.HealthCheckRequest{})
	if err != nil {
		return err
	}
	status := response.GetStatus()
	if status != health.HealthCheckResponse_SERVING {
		return errors.New(status.String())
	}
	return nil
}
