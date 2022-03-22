package interceptors

import (
	"context"
	"github.com/bhankey/go-utils/pkg/logger"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type LoggerInterceptor struct {
	log logger.Logger
}

func NewLoggerInterceptor(log logger.Logger) *ErrorHandlingInterceptor {
	return &ErrorHandlingInterceptor{
		log: log,
	}
}

func (i *LoggerInterceptor) ServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestID := ""
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if headers, ok := md[RequestID]; ok && len(headers) > 0 {
				requestID = headers[0]
			}
		}

		if requestID == "" {
			i.log.Warnf("failed to get request id")
		}

		log := i.log.WithFields(logrus.Fields{
			"method":     info.FullMethod,
			"request_id": requestID,
		},
		)

		start := time.Now()

		resp, err := handler(ctx, req)

		log = log.WithFields(logrus.Fields{
			"duration": time.Since(start),
			"error":    err,
		})

		log.Infof("response.info")

		return resp, nil
	}
}
