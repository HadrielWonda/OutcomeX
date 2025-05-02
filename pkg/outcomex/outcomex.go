package outcomex

// Simplified access to most common functionality
import internal "github.com/hadrielwonda/outcomex/internal/outcomex"

type Outcome[T any] = internal.Outcome[T]
type Error = internal.Error
type AsyncOutcome[T any] = internal.AsyncOutcome[T]

var (
    Success         = internal.Success
    Failure         = internal.Failure
    NewError        = internal.NewError
    ValidationError = internal.ValidationError
    WrapOperation   = internal.WrapOperation
)