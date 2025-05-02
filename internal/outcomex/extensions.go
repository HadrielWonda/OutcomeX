package outcomex

func (o Outcome[T]) Where(predicate func(T) bool, error Error) Outcome[T] {
	if o.IsFailure() {
		return o
	}
	if predicate(o.Value()) {
		return o
	}
	return Failure[T](error)
}

func (o Outcome[T]) Recover(recoverFn func([]Error) T) T {
	if o.IsSuccess() {
		return o.Value()
	}
	return recoverFn(o.Errors())
}

func Combine(outcomes ...Outcome[any]) Outcome[[]any] {
	var results []any
	var errors []Error

	for _, outcome := range outcomes {
		if outcome.IsFailure() {
			errors = append(errors, outcome.Errors()...)
		} else {
			results = append(results, outcome.Value())
		}
	}

	if len(errors) > 0 {
		return Failure[[]any](errors...)
	}
	return Success[[]any](results)
}