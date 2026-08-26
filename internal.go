package main

import "context"

type Dispatcher struct {
	router *Router
}

func NewDispatcher(router *Router) *Dispatcher {
	return &Dispatcher{router: router}
}

func (d *Dispatcher) Dispatch(ctx context.Context, signals []Signal) error {
	for _, signal := range signals {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if err := d.router.Enqueue(signal); err != nil {
			return err
		}
	}
	return nil
}
