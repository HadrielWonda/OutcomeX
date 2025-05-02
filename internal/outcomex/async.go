package outcomex

import "context"

type AsyncOutcome[T any] struct {
	ch <-chan Outcome[T]
}

func Async[T any](ctx context.Context, fn func() Outcome[T]) AsyncOutcome[T] {
	resultCh := make(chan Outcome[T], 1)
	go func() {
		defer close(resultCh)
		select {
		case <-ctx.Done():
			resultCh <- Failure[T](NewError("timeout", "operation timed out"))
		default:
			resultCh <- fn()
		}
	}()
	return AsyncOutcome[T]{ch: resultCh}
}

func (a AsyncOutcome[T]) ThenAsync(ctx context.Context, transform func(T) Outcome[T]) AsyncOutcome[T] {
	resultCh := make(chan Outcome[T], 1)
	go func() {
		defer close(resultCh)
		select {
		case res := <-a.ch:
			if res.IsFailure() {
				resultCh <- res
				return
			}
			resultCh <- transform(res.Value())
		case <-ctx.Done():
			resultCh <- Failure[T](NewError("timeout", "operation timed out"))
		}
	}()
	return AsyncOutcome[T]{ch: resultCh}
}

func (a AsyncOutcome[T]) Await() Outcome[T] {
	return <-a.ch
}