/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package health

import (
	"github.com/containerd/containerd/cmd/ctr/commands"
	"github.com/urfave/cli"
	health "google.golang.org/grpc/health/grpc_health_v1"
)

// Command runs a container
var Command = cli.Command{
	Name:  "health",
	Usage: "checks containerd health",
	Action: func(context *cli.Context) error {
		client, ctx, cancel, err := commands.NewClient(context)
		if err != nil {
			return err
		}
		defer cancel()

		hs := client.HealthService()
		response, err := hs.Check(ctx, &health.HealthCheckRequest{})
		if err != nil {
			return err
		}

		status := response.Status
		if status != health.HealthCheckResponse_SERVING {
			return cli.NewExitError(status.String(), 1)
		}

		return cli.NewExitError(status.String(), 0)
	},
}
