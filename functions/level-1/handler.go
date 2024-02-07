package main

import (
	"context"
	"time"
)

type Handler struct {
}

func NewHandler() (*Handler, error) {

	return &Handler{}, nil
}

func (h *Handler) Handle(ctx context.Context) error {

	if deadline, ok := ctx.Deadline(); ok {
		var cancel context.CancelFunc
		// Reduce the deadline to give time for metrics and other cleanup work to happen before the Lambda function times out
		ctx, cancel = context.WithDeadline(ctx, deadline.Add(-200*time.Millisecond))
		defer cancel()
	}

	return nil
}
