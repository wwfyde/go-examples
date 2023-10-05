package main

import "fmt"

func main() {
	members := []string{"1", "2", "3", "wwfyde"}
	months := [...]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	Q2 := months[4:7]
	summer := months[6:9]

	fmt.Printf("%T %[1]v\n", Q2)
	fmt.Println(summer, len(summer), cap(summer))
	fmt.Println(summer[:], summer[1])
	fmt.Println(summer[4])                // index超过了len(summer) = 3, 编译报错
	newSummer := summer[:7]               // 创建新的slice时, 不能超过cap(summer)
	fmt.Println(summer[:7], newSummer[4]) // [June July August September October November December] October

	for index, value := range members {
		fmt.Println(index+1, value)
	}
	n, q := -5, 5
	fmt.Printf("%b, %b\n", n, n>>1)
	fmt.Printf("%b, %b\n", q, q>>1)
	c := len("1111111111111111111111111111111111111111111111111111111111111111")
	println(c)
}
