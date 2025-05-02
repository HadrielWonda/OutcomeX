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
    t.Parallel()
    
    // Use shorter timings with buffer
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
    defer cancel()

    slowFunc := func() Outcome[int] {
        time.Sleep(10 * time.Millisecond) // Longer than context timeout
        return Success(42)
    }

    ao := Async(ctx, slowFunc)
    result := ao.Await()

    if result.IsSuccess() {
        t.Error("Expected timeout failure but got success")
    }
    
    // Verify error type
    firstErr := result.FirstError()
    if firstErr.Code != "timeout" {
        t.Errorf("Expected timeout error, got %s", firstErr.Code)
    }
}

func TestAsyncSuccess(t *testing.T) {
    t.Parallel()
    
    ctx := context.Background()
    ao := Async(ctx, func() Outcome[int] {
        return Success(42)
    })
    
    result := ao.Await()
    if result.Value() != 42 {
        t.Error("Async success value mismatch")
    }
}


func TestThenAsync(t *testing.T) {
    t.Parallel()
    
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel()

    ao := Async(ctx, func() Outcome[int] {
        return Success(2)
    }).
    ThenAsync(ctx, func(n int) Outcome[int] {
        return Success(n * 3)
    })

    result := ao.Await()
    if result.Value() != 6 {
        t.Error("ThenAsync chain failed")
    }
}

func TestCombineSuccess(t *testing.T) {
    outcomes := []Outcome[any]{
        Success[any]("test"),
        Success[any](42),
    }
    
    combined := Combine(outcomes...)
    if len(combined.Value()) != 2 {
        t.Error("Combine success count mismatch")
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