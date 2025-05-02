package outcomex

import (
    "context"
    "errors"
    "testing"
    "time"
)

func TestSuccess(t *testing.T) {
    o := Success(42)
    if !o.IsSuccess() {
        t.Error("Should be success")
    }
    if o.Value() != 42 {
        t.Error("Incorrect value")
    }
}

func TestFailure(t *testing.T) {
    o := Failure[int](NewError("test", "error"))
    if !o.IsFailure() {
        t.Error("Should be failure")
    }
    if len(o.Errors()) != 1 {
        t.Error("Should have one error")
    }
}

func TestThenChain(t *testing.T) {
    result := Success(2).
        Then(func(n int) Outcome[int] { return Success(n * 3) }).
        Then(func(n int) Outcome[int] { return Success(n + 1) })
        
    if result.Value() != 7 {
        t.Errorf("Expected 7, got %d", result.Value())
    }
}

func TestAsyncTimeout(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
    defer cancel()

    ao := Async(ctx, func() Outcome[int] {
        time.Sleep(20 * time.Millisecond)
        return Success(42)
    })
    
    result := ao.Await()
    if result.IsSuccess() {
        t.Error("Should have timed out")
    }
}

func TestErrorConversion(t *testing.T) {
    err := errors.New("standard error")
    o := ConvertToOutcome(42, err)
    if o.IsSuccess() {
        t.Error("Should convert standard error")
    }
}

// Benchmarking
func BenchmarkSuccessPath(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Success(42).Then(func(n int) Outcome[int] {
            return Success(n * 2)
        })
    }
}

func BenchmarkFailurePath(b *testing.B) {
    err := NewError("test", "error")
    for i := 0; i < b.N; i++ {
        Failure[int](err).Then(func(n int) Outcome[int] {
            return Success(n * 2)
        })
    }
}