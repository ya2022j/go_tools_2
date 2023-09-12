要打印结构体的字段名和对应的值，可以使用反射（reflect）包来实现。下面是一个示例代码：

go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	student := Student{
		Name: "Alice",
		Age:  20,
	}

	PrintStruct(student)
}

func PrintStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("Key: %s, Value: %v\n", field.Name, value)
	}
}
在上述示例代码中，我们定义了一个名为Student的结构体，包含Name和Age两个字段。

然后，我们通过反射机制，在PrintStruct()函数中打印结构体的字段名和对应的值。首先，我们使用reflect.ValueOf()函数获取到结构体的reflect.Value类型对象。然后，通过reflect.Type类型的NumField()方法获取结构体中字段的数量，并使用循环依次处理每个字段。在循环中，我们使用reflect.Type的Field()方法获取字段的元数据信息，包括字段名。然后，通过reflect.Value的Field()方法获取字段的值，并使用Interface()方法将其转换为接口类型，即原始的字段值。最后，我们使用fmt.Printf()函数打印字段名和值。

在main()函数中，我们创建了一个Student类型的实例，并将其传递给PrintStruct()函数进行处理。

通过这个示例代码，你可以了解如何使用反射来打印结构体的字段名和对应的值。请注意，在实际使用时，确保导入了reflect包，并根据需要修改结构体的定义和处理逻辑。
