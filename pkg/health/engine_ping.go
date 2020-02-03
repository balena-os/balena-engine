package health

import (
	"context"
)

// enginePing talks to the /_ping endpoint, evaluating that the engine API is up.
type enginePing struct {}

func (c enginePing) Description() string {
	return "Is engine online"
}

func (c enginePing) Check(ctx context.Context) error {
	cli, err := NewEngineClient()
	if err != nil {
		return err
	}
	_, err = cli.Ping(ctx)
	if err != nil {
		return err
	}
	return nil
}
