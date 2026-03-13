package client

import (
	"context"
	"io"
	"net/url"

	"github.com/docker/docker/api/types"
)

// ImageDelta creates a delta between two images.
func (cli *Client) ImageDelta(ctx context.Context, src, dest string, options types.ImageDeltaOptions) (io.ReadCloser, error) {
	query := url.Values{}
	query.Set("src", src)
	query.Set("dest", dest)
	query.Set("t", options.Tag)

	resp, err := cli.postRaw(ctx, "/deltas/create", query, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
