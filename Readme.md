# Person Management System

A simple Go program demonstrating struct usage, pointer manipulation, error handling, and time operations.

## Overview

This program defines a `Person` struct with basic information (first name, last name, age, and creation timestamp) and provides functionality to validate and create Person instances.

## Features

- **Person struct** with fields for first name, last name, age, and creation timestamp
- **Input validation** to ensure required fields are not empty
- **Pointer usage** to pass structs by reference
- **Error handling** with custom error messages
- **Time tracking** for recording when each person record is created

## Code Structure

### Types

```go
type Person struct {
    firstName string
    lastName  string
    age       int
    createdAt time.Time
}
```

### Functions

- **`createPerson(person *Person) (Person, error)`** 
  - Validates that age > 0, first name and last name are not empty
  - Creates and returns a new Person instance with the provided data
  - Returns an error if validation fails

- **`main()`**
  - Initializes a Person instance with sample data
  - Sets the creation timestamp to current time
  - Calls createPerson and handles any errors
  - Prints the resulting Person struct

## Usage Example

```go
// Create a person
var person1 Person
person1.firstName = "benjamin"
person1.lastName = "Carson"
person1.age = 45
person1.createdAt = time.Now()

// Validate and process
person, err := createPerson(&person1)
if err != nil {
    log.Fatal(err)
}
fmt.Println(person)
```

## Output

```
{benjamin Carson 45 2024-01-01 12:00:00 ...}
```

## Error Handling

The program will return an error and exit if:
- Age is 0
- First name is empty
- Last name is empty

## Learning Points

This example demonstrates:
- Struct definition and initialization
- Pointer parameters in functions
- Multiple return values (including errors)
- Time package usage
- Basic validation logic
- Error handling patterns in Go

## Requirements

- Go 1.x or higher
- No external dependencies

## Note

The `createdAt` field is set but not currently validated or used in the creation logic. This could be extended to include time-based validation in future versions.

