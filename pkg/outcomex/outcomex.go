package outcomex

// Simplified access to most common functionality
import internal "github.com/hadrielwonda/outcomex/internal/outcomex"

type Outcome[T any] = internal.Outcome[T]
type Error = internal.Error
type AsyncOutcome[T any] = internal.AsyncOutcome[T]

func Success[T any](value T) internal.Outcome[T] {
    return internal.Success(value)
}

var (
    Failure         = internal.Failure[any]
    NewError        = internal.NewError
    ValidationError = internal.ValidationError
    WrapOperation   = internal.WrapOperation[string] // Replace 'string' with the desired type
)