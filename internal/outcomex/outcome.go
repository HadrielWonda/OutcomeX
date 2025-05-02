package outcomex

type Outcome[T any] struct {
	value     T
	errors    []Error
	isSuccess bool
}

func Success[T any](value T) Outcome[T] {
	return Outcome[T]{value: value, isSuccess: true}
}

func Failure[T any](errors ...Error) Outcome[T] {
	if len(errors) == 0 {
		panic("at least one error must be provided")
	}
	return Outcome[T]{errors: errors, isSuccess: false}
}

func FromErrors[T any](errors []Error) Outcome[T] {
	return Failure[T](errors...)
}

func (o Outcome[T]) IsSuccess() bool {
	return o.isSuccess
}

func (o Outcome[T]) IsFailure() bool {
	return !o.isSuccess
}

func (o Outcome[T]) Value() T {
	if o.IsFailure() {
		panic("cannot get value from failed outcome")
	}
	return o.value
}

func (o Outcome[T]) Errors() []Error {
	if o.IsSuccess() {
		return nil
	}
	return o.errors
}

func (o Outcome[T]) FirstError() Error {
	if o.IsSuccess() || len(o.errors) == 0 {
		return Error{}
	}
	return o.errors[0]
}