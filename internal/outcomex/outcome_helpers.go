package outcomex

func ConvertToOutcome[T any](val T, err error) Outcome[T] {
	if err != nil {
		return Failure[T](NewError("conversion_error", err.Error()))
	}
	return Success(val)
}

func FromResult[T any](val T, err error) Outcome[T] {
	return ConvertToOutcome(val, err)
}

func WrapOperation[T any](operation func() (T, error)) Outcome[T] {
	val, err := operation()
	return FromResult(val, err)
}