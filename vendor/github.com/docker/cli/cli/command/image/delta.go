package image

import (
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

type deltaOptions struct {
	src  string
	dest string
}

// NewDeltaCommand creates a new `docker delta` command
func NewDeltaCommand(dockerCli command.Cli) *cobra.Command {
	var options deltaOptions

	cmd := &cobra.Command{
		Use:   "delta [OPTIONS] SRC_IMAGE DEST_IMAGE",
		Short: "Create a binary delta between SRC_IMAGE and DEST_IMAGE",
		Args:  cli.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			options.src = args[0]
			options.dest = args[1]
			return runDelta(dockerCli, options)
		},
	}

	return cmd
}

func runDelta(dockerCli command.Cli, options deltaOptions) error {
	clnt := dockerCli.Client()

	responseBody, err := clnt.ImageDelta(context.Background(), options.src, options.dest)
	if err != nil {
		return err
	}
	defer responseBody.Close()

	return jsonmessage.DisplayJSONMessagesToStream(responseBody, dockerCli.Out(), nil)
}
