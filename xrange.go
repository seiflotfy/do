// XRange...
package gogo

import "fmt"

func XRange(args... int) []int {
  start, end := 0, 0
  steps := 1;
  if len(args) >= 1 {
		end = args[0]
	}
	if len(args) >= 2 {
		start = args[0]
		end = args[1]
		if end <= start {
			panic(fmt.Errorf("start %d is larger than end %d", start, end))
		}
	}
	if len(args) == 3 {
		steps = args[2]
	}

  if (steps == 0) {
    steps = 1
  }
	if (end == 0) {
    end = start
    start = 0
  }
	var result []int
  i := start
  idx := 0
	for ; i < end; idx++ {
    result[idx] = i
    i += steps
	}
	return result
}
