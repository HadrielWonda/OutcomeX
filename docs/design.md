# Architecture Decisions

## Core Principles

1. **Immutability**: All Outcome instances are immutable after creation
2. **Composability**: Operations can be chained without side effects
3. **Context Awareness**: Async operations respect context cancellation
4. **Type Safety**: Generic implementation prevents type mismatches

## Error Aggregation

```go
type Outcome[T any] struct {
    errors []Error  // Contains all accumulated errors
}