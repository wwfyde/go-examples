// 输出其命令行参数
package main

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator")

func main() {
	fmt.Println(reflect.TypeOf(n), reflect.TypeOf(sep))
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
