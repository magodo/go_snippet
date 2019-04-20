package bar

import (
	"context"

	"fmt"

	"github.com/magodo/go_snippet/grpc_gateway/internal/api/proto/bar"
)

type BarService struct{}

func (s *BarService) Hello(ctx context.Context, req *bar.StringMessage) (*bar.StringMessage, error) {
	return &bar.StringMessage{
		Value: fmt.Sprintf("Bar get message: %s", req.Value),
	}, nil
}
