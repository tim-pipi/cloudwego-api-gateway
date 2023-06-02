package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type args interface {
	GetFirstArgument() interface{}
}

type result interface {
	GetResult() interface{}
}

func MiddleWareLogger(s string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			klog.Infof("MiddlewareLogger: %s", s)
			return next(ctx, req, resp)
		}
	}
}

func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Infof("real request: %+v\n", req.(args).GetFirstArgument())
		// get local service information
		klog.Infof("local service name: %v\n", ri.From().ServiceName())
		// get remote service information
		klog.Infof("remote service name: %v, remote method: %v\n", ri.To().ServiceName(), ri.To().Method())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v\n", resp.(result).GetResult())
		return nil
	}
}
