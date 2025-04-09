像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值
当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝
改变i不影响j

当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。 
如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。 

package main
import "fmt"

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
}

const a, b, c = 1, false, "str" //多重赋值
常量还可以用作枚举：

const (
    Unknown = 0
    Female = 1
    Male = 2
)
iota 特殊常量，可以认为是一个可以被编译器修改的常量
第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

const (
    a = iota
    b
    c
)

位运算符
位运算符对整数在内存中的二进制位进行操作
假定 A = 60; B = 13; 其二进制数转换为：

A = 0011 1100

B = 0000 1101
-----------------
A&B = 0000 1100
A|B = 0011 1101
A^B = 0011 0001

  for true  {
        fmt.Printf("这是无限循环。\n");
    }

Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
实例
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g)
}

Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
import "fmt"
func main() {
   var a int = 10  
   fmt.Printf("变量的地址: %x\n", &a  )
}


指针的基本用法
指针类型是一种特殊的变量，用于存储其他变量的内存地址‌。指针类型是通过在基本数据类型前加上星号（*）来定义的，例如*int表示指向整型值的指针。指针的主要作用是提供对内存地址的直接访问，从而实现对原始数据的操作和修改，这在系统编程和性能优化中非常有用‌
    ‌声明指针变量‌：在Go中，可以通过在类型前加上星号来声明一个指针变量。例如，var ptr *int声明了一个指向整型的指针变量ptr。
    ‌初始化指针‌：通过取变量的地址来初始化指针。例如，ptr = &a将指针ptr初始化为变量a的地址。
    ‌通过指针访问值‌：通过在指针前加星号来访问它所指向的值。例如，*ptr表示访问指针ptr所指向的值。

unsafe.Pointer的类型和用途

‌unsafe.Pointer是Go语言中的一个特殊类型，用于在不同类型的指针之间进行转换‌。它不能进行指针运算，但可以将任意类型的指针转换为unsafe.Pointer，然后再转换回其他类型的指针。这种类型主要用于底层操作和性能优化，因为它可以绕过Go的垃圾回收机制，直接操作内存地址‌3。
import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}

以上实例执行输出结果为：

a 变量的地址是: 20818a220
ip 变量储存的指针地址: 20818a220
*ip 变量的值: 20

当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。



var numbers [5]int
还可以使用初始化列表来初始化数组的元素：
var numbers = [5]int{1, 2, 3, 4, 5}
numbers := [5]int{1, 2, 3, 4, 5}
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
或
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}
访问数组元素

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值。例如：

var salary float32 = balance[9]



type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {

    // 创建一个新的结构体
    fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

    // 忽略的字段为 0 或 空
   fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}

动态数组切片不需要说明长度，长度是不固定的，可以追加元素，在追加时可能使切片的容量增大
切片初始化
s :=[] int {1,2,3 } 
 make() 函数来创建切片:
var slice1 []type = make([]type, len)
也可以指定容量，其中 capacity 为可选参数， len 是数组的长度并且也是切片的初始长度
make([]T, length, capacity)
也可以简写为
slice1 := make([]type, len)
切片是可索引的，并且可以由 len() 方法获取长度。
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
func main() {
   var numbers = make([]int,3,5)
   printSlice(numbers)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

import "fmt"

func main() {
   /* 创建切片 */
   numbers := []int{0,1,2,3,4,5,6,7,8}  
   printSlice(numbers)

   /* 打印原始切片 */
   fmt.Println("numbers ==", numbers)

   /* 打印子切片从索引1(包含) 到索引4(不包含)*/
   fmt.Println("numbers[1:4] ==", numbers[1:4])

   /* 默认下限为 0*/
   fmt.Println("numbers[:3] ==", numbers[:3])

   /* 默认上限为 len(s)*/
   fmt.Println("numbers[4:] ==", numbers[4:])

   numbers1 := make([]int,0,5)
   printSlice(numbers1)

   /* 打印子切片从索引  0(包含) 到索引 2(不包含) */
   number2 := numbers[:2]
   printSlice(number2)

   /* 打印子切片从索引 2(包含) 到索引 5(不包含) */
   number3 := numbers[2:5]
   printSlice(number3)

}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
append() 和 copy() 函数

如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
下面的代码描述了从拷贝切片的 copy 方法和向切片追加新元素的 append 方法

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
package main
import "fmt"
// 声明一个包含 2 的幂次方的切片
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
   // 遍历 pow 切片，i 是索引，v 是值
   for i, v := range pow {
      // 打印 2 的 i 次方等于 v
      fmt.Printf("2**%d = %d\n", i, v)
   }
}

 range 迭代字符串时，返回每个字符的索引和 Unicode 代码点（rune）。
实例
package main
import "fmt"
func main() {
    for i, c := range "hello" {
        fmt.Printf("index: %d, char: %c\n", i, c)
    }
}

 for 循环的 range 格式可以省略 key 和 value，如下实例：
实例
package main

import "fmt"

func main() {
    // 创建一个空的 map，key 是 int 类型，value 是 float32 类型
    map1 := make(map[int]float32)
   
    // 向 map1 中添加 key-value 对
    map1[1] = 1.0
    map1[2] = 2.0
    map1[3] = 3.0
    map1[4] = 4.0
   
    // 遍历 map1，读取 key 和 value
    for key, value := range map1 {
        // 打印 key 和 value
        fmt.Printf("key is: %d - value is: %f\n", key, value)
    }

    // 遍历 map1，只读取 key
    for key := range map1 {
        // 打印 key
        fmt.Printf("key is: %d\n", key)
    }

    // 遍历 map1，只读取 value
    for _, value := range map1 {
        // 打印 value
        fmt.Printf("value is: %f\n", value)
    }
}
通道（Channel）

range 遍历从通道接收的值，直到通道关闭。
实例
package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)
   
    for v := range ch {
        fmt.Println(v)
    }
}

忽略值

在遍历时可以使用 _ 来忽略索引或值。
实例
package main

import "fmt"

func main() {
    nums := []int{2, 3, 4}
   
    // 忽略索引
    for _, num := range nums {
        fmt.Println("value:", num)
    }
   
    // 忽略值
    for i := range nums {
        fmt.Println("index:", i)
    }
}

Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变

// 创建一个初始容量为 10 的 Map
m := make(map[string]int, 10)
// 使用字面量创建 Map
m := map[string]int{
    "apple": 1,
    "banana": 2,
    "orange": 3,
}

 /* 创建map */
        countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
        fmt.Println("原始地图")
        /* 打印地图 */
        for country := range countryCapitalMap {
                fmt.Println(country, "首都是", countryCapitalMap [ country ])
        }
        /*删除元素*/ delete(countryCapitalMap, "France")


类型转换 
var a int = 10
var b float64 = float64(a)

 字符串类型转换

将一个字符串转换成另一个类型，可以使用以下语法：
var str string = "10"
var num int
num, _ = strconv.Atoi(str)

以上代码将字符串变量 str 转换为整型变量 num。
注意，strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误，我们可以使用空白标识符 _ 来忽略这个错误

实例将字符串转换为浮点数：
import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.14"
    num, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
    }
}

接口
接口类型转换有两种情况：类型断言和类型转换
func main() {
    var i interface{} = "Hello, World"
    //类型断言
    str, ok := i.(string)
    if ok {
        fmt.Printf("'%s' is a string\n", str)
    } else {
        fmt.Println("conversion failed")
    }
}

以上实例中，我们定义了一个接口类型变量 i，并将它赋值为字符串 "Hello, World"。然后，我们使用类型断言将 i 转换为字符串类型，并将转换后的值赋值给变量 str。最后，我们使用 ok 变量检查类型转换是否成功，如果成功，我们打印转换后的字符串；否则，我们打印转换失败的消息


// 定义一个接口 Writer
type Writer interface {
    Write([]byte) (int, error)
}

// 实现 Writer 接口的结构体 StringWriter
type StringWriter struct {
    str string
}

// 实现 Write 方法
func (sw *StringWriter) Write(data []byte) (int, error) {
    sw.str += string(data)
    return len(data), nil
}

func main() {
    // 创建一个 StringWriter 实例并赋值给 Writer 接口变量
    var w Writer = &StringWriter{}
   
    // 将 Writer 接口类型转换为 StringWriter 类型
    sw := w.(*StringWriter)
   
    // 修改 StringWriter 的字段
    sw.str = "Hello, World"
   
    // 打印 StringWriter 的字段值
    fmt.Println(sw.str)
}

接口（interface）是 Go 语言中的一种类型，用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约。

隐式实现：

    Go 中没有关键字显式声明某个类型实现了某个接口。
    只要一个类型实现了接口要求的所有方法，该类型就自动被认为实现了该接口。

接口类型变量：

    接口变量可以存储实现该接口的任意值。
    接口变量实际上包含了两个部分：
        动态类型：存储实际的值类型。
        动态值：存储具体的值。

零值接口：

    接口的零值是 nil。
    一个未初始化的接口变量其值为 nil，且不包含任何动态类型或值。

空接口：

    定义为 interface{}，可以表示任何类型。
接口的常见用法

    多态：不同类型实现同一接口，实现多态行为。
    解耦：通过接口定义依赖关系，降低模块之间的耦合。
    泛化：使用空接口 interface{} 表示任意类型

import (
        "fmt"
        "math"
)

// 定义接口
type Shape interface {
        Area() float64
        Perimeter() float64
}

// 定义一个结构体
type Circle struct {
        Radius float64
}
// Circle 实现 Shape 接口
func (c Circle) Area() float64 {
        return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
        return 2 * math.Pi * c.Radius
}
func main() {
        c := Circle{Radius: 5}
        var s Shape = c // 接口变量可以存储实现了接口的类型
        fmt.Println("Area:", s.Area())
        fmt.Println("Perimeter:", s.Perimeter())
}
