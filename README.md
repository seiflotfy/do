# GoGo
A collection of essential helper Functions

## Map
```
import ("github.com/geekyogre/gogo") gg

...

	iterable := []int{9, 8, 7, 6}
	values := gg.Map(func(v int) bool { return v%2 == 0 }, iterable)
	>>> values => [false, true, false, true]

```

## Reduce
```
import ("github.com/geekyogre/gogo") gg

...

	iterable := []int{9, 8, 7, 6}
	values := gg.Reduce(func(v1 int, v2 int) int { return v1 * v2 }, iterable)
	>>> values => 3024

```