

# Advanced Usage

## Chained Validation

```go
type User struct {
    Name string
    Age  int
}

func validateUser(u User) outcomex.Outcome[User] {
    return outcomex.Success(u).
        Where(func(u User) bool { return u.Age >= 18 }, 
            outcomex.ValidationError("underage")).
        Where(func(u User) bool { return u.Name != "" },
            outcomex.ValidationError("empty_name"))
}