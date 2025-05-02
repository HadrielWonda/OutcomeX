package outcomex_test

import (
    "testing"
    "github.com/hadrielwonda/outcomex/pkg/outcomex"
)

func TestPublicAPI(t *testing.T) {
    // Test using only the public facade
    result := outcomex.Success(42)
    // Add assertions to verify the result i.e
    if result.IsFailure() {
        t.Errorf("Expected success but got failure")
    }

    if result.Value() != 42 {
        t.Errorf("Expected value 42 but got %v", result.Value())
    }

}