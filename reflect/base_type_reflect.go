package reflect

import (
	"fmt"
	"reflect"
)

// 利用包里的TypeOf方法可以获取变量的类型信息
func reflect_typeof(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is:%v\n", t)

	k := t.Kind() //利用Kind() 可以获取t的类型
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64\n")
	case reflect.String:
		fmt.Printf("a is string\n")
	}
}

// 利用包里的ValueOf方法可以获取变量的值信息
func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)

	k := v.Kind() //利用Kind() 可以获取t的类型
	switch k {
	case reflect.Int64:
		fmt.Printf("a is Int64, store value is:%d\n", v.Int())
	case reflect.String:
		fmt.Printf("a is String, store value is:%s\n", v.String())
	}
}
