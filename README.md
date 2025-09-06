# GOX — an extension of the standard Go library

Number of extensions to the standard Go library that would make development easier.

## Container

Extensions of the [container](https://pkg.go.dev/container) package.

### Heap
A generic heap (aka priority queue) implementation. Not safe for concurrent use.

More information about the heap can be found [here](https://pkg.go.dev/container/heap).

#### Usage
```go
package main

import (
	"fmt"

	"github.com/riabininkf/gox/container"
)

type Task struct {
	priority int
	name     string
}

func main() {
	// Initialize heap
	heap := container.NewHeap[*Task](func(a, b *Task) bool {
		return a.priority < b.priority
	})

	// Initialize values
	heap.Push(&Task{4, "d"})
	heap.Push(&Task{2, "b"})
	heap.Push(&Task{1, "a"})
	heap.Push(&Task{3, "c"})

	// Extract and print values
	for heap.Len() > 0 {
		el, ok := heap.Pop()
		fmt.Println(el, ok)
	}
}
```

Output:
```shell
&{1 a} true
&{2 b} true
&{3 c} true
&{4 d} true
```

### Ring
A generic circular list (ring) wrapper. Not safe for concurrent use.

More information about the ring can be found [here](https://pkg.go.dev/container/ring).

#### Usage
```go
package main

import (
	"fmt"

	"github.com/riabininkf/gox/container"
)

type Element struct {
	Value int
}

func main() {
	// Initialize ring
	r := container.NewRing[*Element](4)

	// Initialize values
	for i := 1; i <= r.Len(); i++ {
		r.Set(&Element{Value: i})
		r.Next()
	}

	// Iterate 2 times around the ring
	for i := 0; i < r.Len()*2; i++ {
		val, ok := r.Value()
		fmt.Println(val, ok)
		r.Next()
	}
}
```
Output:
```shell
&{1} true
&{2} true
&{3} true
&{4} true
&{1} true
&{2} true
&{3} true
&{4} true
```

## Sync
Extensions of the [sync](https://pkg.go.dev/sync) package.

### Map
Generic wrapper for sync.Map.

More information about the sync.Map can be found [here](https://pkg.go.dev/sync#Map).
#### Usage

Without sync.Map
```go
package main

import (
	"sync"
)

func main() {
	m := map[int]struct{}{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		idx := i
		wg.Go(func() {
			m[idx] = struct{}{}
		})
	}

	wg.Wait()
}
```

```shell
fatal error: concurrent map writes

goroutine 13 [running]:

Process finished with the exit code 2
```

With sync.Map
```go
package main

import (
	"sync"

	syncx "github.com/riabininkf/gox/sync"
)

func main() {
	m := syncx.NewMap[int, struct{}]()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		idx := i
		wg.Go(func() {
			m.Store(idx, struct{}{})
		})
	}

	wg.Wait()
}
```

```shell
Process finished with the exit code 0
```