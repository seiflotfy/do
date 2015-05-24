# GoGo
A collection of essential non-built-in helper Functions

## Map
```
import ("github.com/geekyogre/gogo") gg

...

iterable := []int{9, 8, 7, 6}

values := gg.Map(func(v int) bool { return v%2 == 0 }, iterable)
// values => [false, true, false, true]

```

## MapParallel
```
import ("github.com/geekyogre/gogo") gg

...

iterable := []int{9, 8, 7, 6, 5, 6, 6, 1}

values := gg.MapParallel(func(v int) bool { return v%2 == 0 }, iterable, 4)
// values => [false, true, false, true, false, true, true, false]

```


## Reduce
```
import ("github.com/geekyogre/gogo") gg

...

iterable := []int{9, 8, 7, 6}

values := gg.Reduce(func(v1 int, v2 int) int { return v1 * v2 }, iterable)
// values => 3024

```

## Filter
```
import ("github.com/geekyogre/gogo") gg

...

iterable := []int{9, 8, 7, 6}

values := gg.Filter(func(v int) bool { return v%2 == 0 }, iterable)
// values => [8, 6]

```

## Any
```
import ("github.com/geekyogre/gogo") gg

...

iterable := []bool{false, false, false, true}

values := gg.Any(iterable)
// values => true

```

## All
```
import ("github.com/geekyogre/gogo") gg

...

iterable := []bool{false, false, false, true}

values := gg.All(iterable)
// values => false

```

## Range
```
import ("github.com/geekyogre/gogo") gg

...

values := gg.Range(5)
// values => [0, 1, 2, 3, 4]

values := gg.Range(5, 10)
// values => [5, 6, 7, 8, 9]

values := gg.Range(5, 25, 5)
// values => [5, 10, 15, 20]


```

## Round
```
import ("github.com/geekyogre/gogo") gg

...

value := gg.Round(5.08, 1)
// value => 5.1

value := gg.Round(1.0001, 1)
// value => 1.0

value := gg.Round(8.5645, 2)
// value => 8.56


```