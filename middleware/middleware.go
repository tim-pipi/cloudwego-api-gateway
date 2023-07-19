package middleware

import (
	"context"
	"fmt"

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

func ValidatorMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) (err error) {
		if gfa, ok := args.(interface{ GetFirstArgument() interface{} }); ok {
			req := gfa.GetFirstArgument()
			if rv, ok := req.(interface{ IsValid() error }); ok {
				if err := rv.IsValid(); err != nil {
					return fmt.Errorf("request data is not valid:%w", err)
				}
			}
		}
		err = next(ctx, args, result)
		if err != nil {
			return err
		}
		if gr, ok := result.(interface{ GetResult() interface{} }); ok {
			resp := gr.GetResult()
			if rv, ok := resp.(interface{ IsValid() error }); ok {
				if err := rv.IsValid(); err != nil {
					return fmt.Errorf("response data is not valid:%w", err)
				}
			}
		}
		return nil
	}
}
