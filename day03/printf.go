package main

import "fmt"

type Student struct {
	x, y int
}

func main() {
	// %v   值的默认格式表示
	// %+v  类似%v，但输出结构体时会添加字段名
	// %#v  值的GO语法表示
	// %T   值的类型的GO语法表示
	// %t   布尔类型true或false
	// %c   字符
	// %b   二进制
	// %o   八进制
	// %d   十进制
	// %8d  十进制八位，前边补空格
	// %08d 十进制八位，前边补0
	// %x   十六进制 a-f小写
	// %X   十六进制 a-f大写
	// %f   浮点数
	// %s   string []byte
	// %q   转义输出
	str := "alex"
	fmt.Printf("%T, %v\n", str, str)

	var a rune = '一'
	fmt.Printf("%T, %v\n", a, a)

	p := Student{1, 2}
	fmt.Printf("%T, %v\n", p, p)

	fmt.Printf("%T, %t\n", true, true)

	fmt.Printf("%T, %d\n", 233, 233)
	fmt.Printf("%T, %5d\n", 233, 233)
	fmt.Printf("%T, %05d\n", 233, 233)
	fmt.Printf("%T, %05d\n", 233, 233)

	b := fmt.Sprintf("%b", 123)
	fmt.Printf("%T, %s\n", b, b)

	fmt.Printf("%s\n", "你好，世界\n")
	fmt.Printf("%q\n", "你好，世界\n")

}
