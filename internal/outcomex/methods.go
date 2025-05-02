package outcomex

func (o Outcome[T]) Then(transform func(T) Outcome[T]) Outcome[T] {
	if o.IsFailure() {
		return o
	}
	return transform(o.Value())
}

func (o Outcome[T]) ThenConvert(transform func(T) (any, error)) Outcome[any] {
	if o.IsFailure() {
		return Failure[any](o.Errors()...)
	}
	
	val, err := transform(o.Value())
	if err != nil {
		return Failure[any](NewError("conversion_error", err.Error()))
	}
	return Success[any](val)
}

func (o Outcome[T]) Match(onSuccess func(T), onFailure func([]Error)) {
	if o.IsSuccess() {
		onSuccess(o.Value())
	} else {
		onFailure(o.Errors())
	}
}

func (o Outcome[T]) MatchResult(onSuccess func(T) any, onFailure func([]Error) any) any {
	if o.IsSuccess() {
		return onSuccess(o.Value())
	}
	return onFailure(o.Errors())
}