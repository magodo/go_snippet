package foo

import (
	"context"

	"fmt"

	"github.com/magodo/go_snippet/grpc_gateway/pkg/api/proto/foo"
)

type FooService struct{}

func (s *FooService) Hello(ctx context.Context, req *foo.StringMessage) (*foo.StringMessage, error) {
	return &foo.StringMessage{
		Value: fmt.Sprintf("Foo get message: %s", req.Value),
	}, nil
}
