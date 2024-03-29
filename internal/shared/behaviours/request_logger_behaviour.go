package behaviours

import (
	"context"
	"log"

	"github.com/danmaxdanilov/zts.shared/cmd/mediatr"
)

type RequestLoggerBehaviour struct {
}

func (r *RequestLoggerBehaviour) Handle(ctx context.Context, request interface{}, next mediatr.RequestHandlerFunc) (interface{}, error) {
	log.Printf("logging some stuff before handling the request")

	response, err := next()
	if err != nil {
		return nil, err
	}

	log.Println("logging some stuff after handling the request")

	return response, nil
}
