package service

import (
	"context"

	greeter "github.com/magodo/rpc_gateway/internal/api/rpc"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *greeter.Request, resp *greeter.Response) error {
	resp.Msg = "Hello " + req.Name
	return nil
}
