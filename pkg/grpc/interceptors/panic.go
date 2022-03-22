package interceptors

import (
	"context"
	"github.com/bhankey/go-utils/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PanicInterceptor struct {
	log logger.Logger
}

func NewPanicInterceptor(log logger.Logger) *PanicInterceptor {
	return &PanicInterceptor{
		log: log,
	}
}

func (i *PanicInterceptor) ServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				i.log.Errorf("panic happend: %v", r)

				err = status.New(codes.Internal, "panic").Err()
			}
		}()

		return handler(ctx, req)
	}
}
