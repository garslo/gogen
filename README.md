# gogen

A simplification of Go's `go/ast` package that allows for some
interesting code generation. Currently very rough.

# Examples

See the
[examples](https://github.com/garslo/gogen/tree/master/examples)
directory for examples and a build/run script.

```sh
$ ./run-example.sh for_loop.go
CODE:
package main

import "os"
import "fmt"

func main() {
	var i int
	for i = 0; i <= 10; i++ {
		fmt.Println(i)
	}
	os.Exit(i)
}
RUN RESULT:
0
1
2
3
4
5
6
7
8
9
10
exit status 11
```
