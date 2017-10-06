package client

import (
	"io"
	"net/url"

	"golang.org/x/net/context"
)

// ImageImport creates a new image based in the source options.
// It returns the JSON content in the response body.
func (cli *Client) ImageDelta(ctx context.Context, src, dest string) (io.ReadCloser, error) {
	query := url.Values{}
	query.Set("src", src)
	query.Set("dest", dest)

	resp, err := cli.postRaw(ctx, "/images/delta", query, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
