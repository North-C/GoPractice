### Variable

`var name type = expression`

#### Pointer
It's same as poniter in C language.

"flag" package is made in the basis of pointer.

#### new() function
`new` is not a keyword, it can be rename in a function.

`p := new(T)` ,new(T) will create a anonymous variable, initialize  its value to be nil/zero, and then return the address of  the variable.

#### Life cycle
for variables in package level, their longevity are same with the  whole program.

But the life cycle of local variables is dynamic, which depends on garbage collector in Go language.

How does the GC know the time that could make garbage collecting?

Using pointer or reference to Accssiblity Analysis.


### Assignment

Assignability:
* Type should be matched, nil can be assigned to each type of pointer or reference variables

### Type

```
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）

```


`type newname typename`


capitalise the first letter of variables----It can be used out of this package.

#### Bool

`true` or `false`

Bool value can't be translated to number 0 or 1 implicitly.

#### String 
String is a unchangeable byte sequence.

It could contain arbitrary character.

`len(string)` return the length of a string

In Go, source files are always encoded with UTF8,and for text string it's same.

So we can write Unicode value into string's  literal value.


##### Unicode
It 's invented to handle diverse text data in all kinds of language.

In GO ,every Unicode code point correspond to a `rune` type value( identical to int32)

An Unicode code point can be stored in 32bit.


##### UTF-8

It encodes Unicode code point to byte sequences.

It is variable-length and compatible with Unicode and ASCII.


There are four important packages: `bytes` `strings` ` strconv` `unicode`

#### Constant

Computation about constants are done in compilation, not running.

##### announcement
`iota` constant generator could initialize constants, or we call it Enumeration

for example:
```go
type WeekDay int

const(
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
/*
Sunday 0
Monday 1
Tuesday 2
...
*/
```
In a `const` announcement, `iota` starts from 0, then every constant announcement will increase by degrees.




### Array & Slice
We can transfer an array parameter in the way of pointer.


As a pointer, we can change the real array using it.

```go
func zero(p *[32]int) {
    *ptr = [32]byte{}
}
```

##### slice
a slice is consist of three parts:
* pointer: point at the bottom array
* length: number of element
* capacity


different slice could share the same array, and refer to the same part of array.

we can't compare slice with slice.

 a slice whose value is nil should be identical to a slice whose value is 0.


 ##### about memory
for an instance:
```go

func nonempty(strings [] string) [] string{
    i := 0
    for _, s := range string{
        string[i] = s
        i++
    }
    return strings[:i]
}
```
In this example, the input slice and the output share the same bottom array.

It can save some memory, but also will override some data.

So, we should use slice as :
```go
runes = append(runes, r)
```
assign  the result to the original slice variable.

According to this feature, we can use slice act as  a stack.



#### Map
In Go,  a map is a reference to a hashtable.

>It's a bad idea to use float as key

format: `map[K]V`

```go
ages := make(map[string]int){
    "alice": 31,
    "charlie": 34,
}
delete(ages, "alice")
```


We can't take the address of map, because with the amount of element increasing, It's possible to allocate bigger memory space for map, and after this phase, the address before may be ineffective.


```go
func equal(x,y map[string]it) bool{
    if len(x) 1= len(y){
        return false
    }
    for k,xv := range x{
        if yv, ok:= y[k]; !ok || yv != xv{
            return false
        }
    }
    return true
}
```


#### Struct/JSON/Text/HTML

To be Continue......


### Function
```go
func name(parameter-list) (result-list){
    //body
}
```

如果两个函数形式**参数列表和返回值列表中的变量类型**一一对应，那么这两个函数被认为有相同的类型或签名。**形参和返回值的变量名不影响函数签名**，也不影响它们是否可以以省略参数类型的形式表示。

实参通过值的方式传递，因此函数的形参是实参的拷贝。对形参进行修改不会影响实参。但是，如果实参包括引用类型，如指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。



