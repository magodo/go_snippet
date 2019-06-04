package internal

import (
	"context"
	"fmt"

	pb "foo/proto/greeter"
)

type Handler struct {
	Name string
}

func (s *Handler) Greet(ctx context.Context, req *pb.Request, resp *pb.Response) error {
	resp.Msg = fmt.Sprintf("%s: hello", s.Name)
	return nil
}
