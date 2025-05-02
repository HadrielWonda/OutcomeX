package outcomex

// Simplified access to most common functionality
import (
    "fmt"
    internal "github.com/hadrielwonda/outcomex/internal/outcomex"
)

type Outcome[T any] = internal.Outcome[T]
type Error = internal.Error
type AsyncOutcome[T any] = internal.AsyncOutcome[T]

func Success[T any](value T) internal.Outcome[T] {
    return internal.Success(value)
}


func ExampleSuccess() {
    result := Success(42)
    fmt.Println(result.IsSuccess())
    // Output: true
}

func ExampleFailure() {
    err := NewError("test", "example")
    result := internal.Failure[int](err)
    fmt.Println(result.IsFailure())
    // Output: true
}

var (
    Failure         = internal.Failure[any]
    NewError        = internal.NewError
    ValidationError = internal.ValidationError
    WrapOperation   = internal.WrapOperation[string] // Replace 'string' with the desired type
)