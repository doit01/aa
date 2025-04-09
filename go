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


panic 和 recover

Go 的 panic 用于处理不可恢复的错误，recover 用于从 panic 中恢复。

panic:

    导致程序崩溃并输出堆栈信息。
    常用于程序无法继续运行的情况。

recover:

    捕获 panic，避免程序崩溃。

实例
package main

import "fmt"

func safeFunction() {
        defer func() {
                if r := recover(); r != nil {
                        fmt.Println("Recovered from panic:", r)
                }
        }()
        panic("something went wrong")
}

func main() {
        fmt.Println("Starting program...")
        safeFunction()
        fmt.Println("Program continued after panic")
}

 使用 errors.Is 和 errors.As

从 Go 1.13 开始，errors 包引入了 errors.Is 和 errors.As 用于处理错误链：
errors.Is

检查某个错误是否是特定错误或由该错误包装而成。

fmt 包与错误格式化

fmt 包提供了对错误的格式化输出支持：

    %v：默认格式。
    %+v：如果支持，显示详细的错误信息。
    %s：作为字符串输出



Go 语言支持并发，通过 goroutines 和 channels 提供了一种简洁且高效的方式来实现并发。

Goroutines：

    Go 中的并发执行单位，类似于轻量级的线程。
    Goroutine 的调度由 Go 运行时管理，用户无需手动分配线程。
    使用 go 关键字启动 Goroutine。
    Goroutine 是非阻塞的，可以高效地运行成千上万个 Goroutine。

Channel：

    Go 中用于在 Goroutine 之间通信的机制。
    支持同步和数据共享，避免了显式的锁机制。
    使用 chan 关键字创建，通过 <- 操作符发送和接收数据。

Scheduler（调度器）：

Go 的调度器基于 GMP 模型，调度器会将 Goroutine 分配到系统线程中执行，并通过 M 和 P 的配合高效管理并发。

    G：Goroutine。
    M：系统线程（Machine）。
    P：逻辑处理器（Processor）。

Goroutine

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式：

go 函数名( 参数列表 )

例如：

go f(x, y, z)

开启一个新的 goroutine:

f(x, y, z)

Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
实例
package main

import (
        "fmt"
        "time"
)

func sayHello() {
        for i := 0; i < 5; i++ {
                fmt.Println("Hello")
                time.Sleep(100 * time.Millisecond)
        }
}

func main() {
        go sayHello() // 启动 Goroutine
        for i := 0; i < 5; i++ {
                fmt.Println("Main")
                time.Sleep(100 * time.Millisecond)
        }
}

执行以上代码，你会看到输出的 Main 和 Hello。输出是没有固定先后顺序，因为它们是两个 goroutine 在执行：

Main
Hello
Main
Hello
...

通道（Channel）

通道（Channel）是用于 Goroutine 之间的数据传递。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。

使用 make 函数创建一个 channel，使用 <- 操作符发送和接收数据。如果未指定方向，则为双向通道。

ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v

声明一个通道很简单，我们使用chan关键字即可，通道在使用前必须先创建：

ch := make(chan int)

注意：默认情况下，通道是不带缓冲区的。发送端发送数据，同时必须有接收端相应的接收数据。

以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：
实例
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // 把 sum 发送到通道 c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // 从通道 c 中接收

    fmt.Println(x, y, x+y)
}

输出结果为：

-5 17 12

通道缓冲区

通道可以设置缓冲区，通过 make 的第二个参数指定缓冲区大小：

ch := make(chan int, 100)

带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

注意：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。
实例
package main

import "fmt"

func main() {
    // 这里我们定义了一个可以存储整数类型的带缓冲通道
    // 缓冲区大小为2
    ch := make(chan int, 2)

    // 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
    // 而不用立刻需要去同步读取数据
    ch <- 1
    ch <- 2

    // 获取这两个数据
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

执行输出结果为：

1
2

Go 遍历通道与关闭通道

Go 通过 range 关键字来实现遍历读取到的数据，类似于与数组或切片。格式如下：

v, ok := <-ch

如果通道接收不到数据后 ok 就为 false，这时通道就可以使用 close() 函数来关闭。
实例
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
    // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
    // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
    // 会结束，从而在接收第 11 个数据的时候就阻塞了。
    for i := range c {
        fmt.Println(i)
    }
}

执行输出结果为：

0
1
1
2
3
5
8
13
21
34

Select 语句

select 语句使得一个 goroutine 可以等待多个通信操作。select 会阻塞，直到其中的某个 case 可以继续执行：
实例
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}

以上代码中，fibonacci goroutine 在 channel c 上发送斐波那契数列，当接收到 quit channel 的信号时退出。

执行输出结果为：

0
1
1
2
3
5
8
13
21
34
quit

使用 WaitGroup

sync.WaitGroup 用于等待多个 Goroutine 完成。

同步多个 Goroutine：
实例
package main

import (
        "fmt"
        "sync"
)

func worker(id int, wg *sync.WaitGroup) {
        defer wg.Done() // Goroutine 完成时调用 Done()
        fmt.Printf("Worker %d started\n", id)
        fmt.Printf("Worker %d finished\n", id)
}

func main() {
        var wg sync.WaitGroup

        for i := 1; i <= 3; i++ {
                wg.Add(1) // 增加计数器
                go worker(i, &wg)
        }

        wg.Wait() // 等待所有 Goroutine 完成
        fmt.Println("All workers done")
}

以上代码，执行输出结果如下：

Worker 1 started
Worker 1 finished
Worker 2 started
Worker 2 finished
Worker 3 started
Worker 3 finished
All workers done

高级特性

Buffered Channel：

创建有缓冲的 Channel。

ch := make(chan int, 2)

Context：

用于控制 Goroutine 的生命周期。

context.WithCancel、context.WithTimeout。

Mutex 和 RWMutex：

sync.Mutex 提供互斥锁，用于保护共享资源。

var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()

并发编程小结

Go 语言通过 Goroutine 和 Channel 提供了强大的并发支持，简化了传统线程模型的复杂性。配合调度器和同步工具，可以轻松实现高性能并发程序。

    Goroutines 是轻量级线程，使用 go 关键字启动。
    Channels 用于 goroutines 之间的通信。
    Select 语句 用于等待多个 channel 操作。

常见问题

死锁 (Deadlock)：

    示例：所有 Goroutine 都在等待，但没有任何数据可用。
    解决：避免无限等待、正确关闭通道。

数据竞争 (Data Race)：

    示例：多个 Goroutine 同时访问同一变量。
    解决：使用 Mutex 或 Channel 同步访问。
