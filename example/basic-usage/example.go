package example // Example usage of the OutcomeX library, demonstrating the use of the library's features and what it should work like

// import "github.com/HadrielWonda/OutcomeX" // NOTE: Ensure this matches the module path in the OutcomeX repository's go.mod file

// func someOperation() (int, error) {
//     // Example implementation of someOperation
//     return 42, nil // Sample logic, replace with actual operation
// }

// func main() {
//     result := outcomex.WrapOperation(func() (int, error) {
//         return someOperation()
//     }).
//     Then(func(n int) outcomex.Outcome[int] {
//         return outcomex.Success(n * 2)
//     }).
//     MatchResult(
//         func(v int) any { return v },
//         func(errs []outcomex.Error) any { return errs },
//     )
//     _ = result // Use or handle the result as needed
// }