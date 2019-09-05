package sliceutils_test

import (
	"fmt"
	"github.com/theplant/sliceutils"
)


/*
# `sliceutils`

Wrap, Unwrap convert `[]interface{}` to `[]*MyStruct` and back
*/
func ExampleWrap_01demo() {
	type Val struct {
		val int
		str string
	}

	var s []string
	sliceutils.Unwrap([]interface{}{"1", "2"}, &s)

	fmt.Println("s =", s)

	istrs := sliceutils.Wrap(s)
	fmt.Println("istrs =", istrs)


	var i []int
	sliceutils.Unwrap([]interface{}{1, 2}, &i)
	fmt.Println("i =", i)


	iis := sliceutils.Wrap(i)
	fmt.Println("iis =", iis)


	var vals []*Val
	sliceutils.Unwrap([]interface{}{&Val{1, "1"}, &Val{2, "2"}}, &vals)
	fmt.Printf("vals = %+v, %+v\n", vals[0], vals[1])


	ivals := sliceutils.Wrap(vals)
	fmt.Printf("ivals = %+v, %+v\n", ivals[0], ivals[1])

	var vals2 []Val
	sliceutils.Unwrap([]interface{}{Val{1, "1"}, Val{2, "2"}}, &vals2)
	fmt.Printf("vals2 = %+v\n", vals2)

	ivals2 := sliceutils.Wrap(vals2)
	fmt.Printf("ivals2 = %+v\n", ivals2)

	//Output:
	//s = [1 2]
	//istrs = [1 2]
	//i = [1 2]
	//iis = [1 2]
	//vals = &{val:1 str:1}, &{val:2 str:2}
	//ivals = &{val:1 str:1}, &{val:2 str:2}
	//vals2 = [{val:1 str:1} {val:2 str:2}]
	//ivals2 = [{val:1 str:1} {val:2 str:2}]
}
