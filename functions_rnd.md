# How Many Functions Exists

## Find all func with explanations and one-liner

### 1. **Normal Function**

**One-liner**: A reusable block of code with a name and parameters.

```go
func add(a int, b int) int {
 return a + b
}
```

---

### 2. **Anonymous Function**

**One-liner**: A function without a name, often assigned to a variable or used immediately.

```go
sum := func(a, b int) int { return a + b }
fmt.Println(sum(3, 4))
```

---

### 3. **Immediately Invoked Function Expression (IIFE)**

**One-liner**: An anonymous function called right after it's defined.

```go
result := func(a, b int) int { return a + b }(5, 6)
fmt.Println(result)
```

---

### 4. **Method (Function with Receiver)**

**One-liner**: A function associated with a type.

```go
type Person struct {
 name string
}

func (p Person) greet() string {
 return "Hello, " + p.name
}
```

---

### 5. **Function as a Parameter**

**One-liner**: Passing functions to other functions.

```go
func operate(a, b int, fn func(int, int) int) int {
 return fn(a, b)
}
```

---

### 6. **Function Returning a Function (Closure)**

**One-liner**: A function that returns another function with access to its outer scope.

```go
func multiplier(factor int) func(int) int {
 return func(val int) int {
  return factor * val
 }
}
```

---

### 7. **Variadic Function**

**One-liner**: Accepts any number of arguments.

```go
func sumAll(nums ...int) int {
 total := 0
 for _, n := range nums {
  total += n
 }
 return total
}
```

---

### 8. **Defer Function**

**One-liner**: Scheduled to run after the surrounding function completes.

```go
func logExit() {
 defer fmt.Println("Function exited")
 fmt.Println("Inside function")
}
```

---

### 9. **Recursive Function**

**One-liner**: A function that calls itself.

```go
func factorial(n int) int {
 if n == 0 {
  return 1
 }
 return n * factorial(n-1)
}
```

---

### 10. **Method with Pointer Receiver**

**One-liner**: Allows modifying the original value.

```go
func (p *Person) setName(newName string) {
 p.name = newName
}
```

---

### 11. **Init Function**

**One-liner**: Automatically runs before main; used for setup.

```go
func init() {
 fmt.Println("Init called")
}
```

---

### 12. **Main Function**

**One-liner**: The entry point of a Go program.

```go
func main() {
 fmt.Println("Hello, Go!")
}
```
