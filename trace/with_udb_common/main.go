package main

import (
	"context"

	"git.ucloudadmin.com/udb/micro/common/logger"
	"git.ucloudadmin.com/udb/micro/common/trace"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {

	name := "foo"
	// init logger
	logger.Init(name)
	// init tracer
	trace.Init(name)
	defer trace.Close()

	ctx := context.Background()
	//logger.For(ctx).Info("test")
	span, ctx := opentracing.StartSpanFromContext(ctx, "op")
	defer span.Finish()
	span.LogEvent("evt")

	span1, _ := opentracing.StartSpanFromContext(ctx, "sub-op1")
	defer span1.Finish()
	ext.SpanKindRPCClient.Set(span1)
	ext.PeerService.Set(span1, "mysql")
	span1.LogEvent("sub-evt1")
}
