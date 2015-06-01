# Do
[![GoDoc](https://godoc.org/github.com/geekyogre/do?status.svg)](https://godoc.org/github.com/geekyogre/do)
Do is collection of essential non-built-in helper Functions for Go, mostly inspired by the Python built-in functions


## All
```
import ("github.com/geekyogre/do")
...

iterable := []bool{false, false, false, true}

values : do.All(iterable)
// values => false

```

## Any
```
import ("github.com/geekyogre/do")
...

iterable := []bool{false, false, false, true}

values : do.Any(iterable)
// values => true

```

## Filter
```
import ("github.com/geekyogre/do")
...

iterable := []int{9, 8, 7, 6}

values : do.Filter(func(v int) bool { return v%2 == 0 }, iterable)
// values => [8, 6]

```

## Map
```
import ("github.com/geekyogre/do")
...

iterable := []int{9, 8, 7, 6}

values : do.Map(func(v int) bool { return v%2 == 0 }, iterable)
// values => [false, true, false, true]

```

## MapParallel
```
import ("github.com/geekyogre/do")
...

iterable := []int{9, 8, 7, 6, 5, 6, 6, 1}

values : do.MapParallel(func(v int) bool { return v%2 == 0 }, iterable, 4)
// values => [false, true, false, true, false, true, true, false]

```

## Range
```
import ("github.com/geekyogre/do")

...
values : do.Range(5)
// values => [0, 1, 2, 3, 4]

values : do.Range(5, 10)
// values => [5, 6, 7, 8, 9]

values : do.Range(5, 25, 5)
// values => [5, 10, 15, 20]

```

## Reduce
```
import ("github.com/geekyogre/do")
...

iterable := []int{9, 8, 7, 6}

values : do.Reduce(func(v1 int, v2 int) int { return v1 * v2 }, iterable)
// values => 3024

```

## Reversed
```
import ("github.com/geekyogre/do")
...

value := do.Reversed([]int {1, 2, 4, 8, 16})
// value => []int {16, 8, 4, 2, 1}

```

## Round
```
import ("github.com/geekyogre/do")
...

value : do.Round(5.08, 1)
// value => 5.1

value : do.Round(1.0001, 1)
// value => 1.0

value : do.Round(8.5645, 2)
// value => 8.56

```

## Slice
```
import ("github.com/geekyogre/do")
...

value := do.Slice([]int {1, 2, 4, 8, 16, 32, 64, 128}, 1, 5, 3)
// value => [4 32]

```

## Sum
```
import ("github.com/geekyogre/do")
...

value := do.Sum([]float32{1.1, 2, 4, 8, 16, 32, 64, 128})
// float32(Sum(value).(float64)) => 255.1

// *Note: Sum will return an interface that can be assigned only to int64 or float64*

```

## Unique
```
import ("github.com/geekyogre/do")
...

iterable := []int{8, 9, 8, 5, 7, 6, 5}

values : do.Unique(func(v int) int { return v }, iterable)
// values => [8, 9, 5, 7, 6]

```
