# Error Code Conventions

## Standard Codes

Code | Description
-----|------------
validation | Input validation failure
not_found | Resource missing
conflict | State conflict
unexpected | Unhandled exception
timeout | Operation timed out

## Custom Codes

```go
// Register custom codes at init
func init() {
    outcomex.RegisterCode("custom_code", "Custom description")
}