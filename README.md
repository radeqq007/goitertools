# goitertools

## Installation

```sh
go get github.com/radeqq007/goitertools
```

## Examples

### Cycle

```go
count := 0
for v := range goitertools.Cycle([]string{"A", "B"}) {
  fmt.Print(v, " ")
  count++
  if count == 5 {
    break
  }
}

// Output: A B A B A
```

### Count

```go
for v := range goitertools.Count(10, 5) {
    fmt.Print(v, " ")
    if v >= 25 {
        break
    }
}
// Output: 10 15 20 25
```

### Repeat

```go
for v := range goitertools.Repeat("Hi", 3) {
    fmt.Print(v, " ")
}
// Output: Hi Hi Hi
```

### Filter

```go
nums := []int{10, 15, 20, 25}
isEven := func(idx int, val int) bool { return val%2 == 0 }

for v := range goitertools.Filter(nums, isEven) {
    fmt.Print(v, " ")
}
// Output: 10 20
```

### FilterFalse

```go
nums := []int{10, 15, 20, 25}
isEven := func(idx int, val int) bool { return val%2 == 0 }

for v := range goitertools.FilterFalse(nums, isEven) {
    fmt.Print(v, " ")
}
// Output: 15 25
```

### Compress

```go
data := []string{"Go", "Python", "Rust"}
pick := []bool{true, false, true}

for v := range goitertools.Compress(data, pick) {
    fmt.Print(v, " ")
}
// Output: Go Rust
```

### DropWhile

```go
nums := []int{1, 3, 4, 1} // Drops 1 and 3 because they are odd
isOdd := func(idx int, val int) bool { return val%2 != 0 }

for v := range goitertools.DropWhile(nums, isOdd) {
    fmt.Print(v, " ")
}
// Output: 4 1
```

### TakeWhile

```go
nums := []int{1, 3, 4, 1} // Stops at 4 because it is not odd
isOdd := func(idx int, val int) bool { return val%2 != 0 }

for v := range goitertools.TakeWhile(nums, isOdd) {
    fmt.Print(v, " ")
}
// Output: 1 3
```


### Chain

```go
s1 := []int{1, 2}
s2 := []int{3, 4}

for v := range goitertools.Chain(s1, s2) {
    fmt.Print(v, " ")
}
// Output: 1 2 3 4
```
