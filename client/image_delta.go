package client

import (
	"io"
	"net/url"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
)

// ImageImport creates a new image based in the source options.
// It returns the JSON content in the response body.
func (cli *Client) ImageDelta(ctx context.Context, src, dest string, options types.ImageDeltaOptions) (io.ReadCloser, error) {
	query, err := cli.imageDeltaOptionsToQuery(options)
	if err != nil {
		return nil, err
	}
	query.Set("src", src)
	query.Set("dest", dest)

	resp, err := cli.postRaw(ctx, "/images/delta", query, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

func (cli *Client) imageDeltaOptionsToQuery(options types.ImageDeltaOptions) (url.Values, error) {
	query := url.Values{}
	query.Set("t", options.Tag)
	return query, nil
}
