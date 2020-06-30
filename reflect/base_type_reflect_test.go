package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBaseTypeReflect(T *testing.T) {

	T.Log("******************测试通过反射操作基本类型1******************")
	var x int64 = 3
	reflect_typeof(x)

	var y string = "hello"
	reflect_typeof(y)
}

func TestBaseValueReflect(T *testing.T) {

	T.Log("******************测试通过反射操作基本类型2******************")
	var x int64 = 3
	reflect_value(x)

	var y string = "hello"
	reflect_value(y)
}

func TestSetBaseValue(T *testing.T) {

	T.Log("******************测试通过反射操作基本类型3******************")
	// 利用包里的函数改变基础类型的值

	var x float64 = 3.14
	v := reflect.ValueOf(&x)
	fmt.Printf("type of v is %v\n", v.Type())

	// 反射里面取地址需要通过 Elem() 方法进行取地址
	v.Elem().SetFloat(6.28)
	fmt.Printf("After set x is %v", x)
}
