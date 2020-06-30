package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 获取结构体的字段
func TestGetStructField(T *testing.T) {

	T.Log("******************测试通过反射操作结构体类型1******************")
	//创建一个结构体变量
	var s Student = Student{
		Name:  "BigOrange",
		Sex:   1,
		Age:   10,
		Score: 80.1,
	}

	v := reflect.ValueOf(s)
	t := v.Type()
	kind := t.Kind()

	//分析s变量的类型，如果是结构体类型，那么遍历所有的字段
	switch kind {
	case reflect.Int64:
		fmt.Printf("s is int64\n")
	case reflect.Float32:
		fmt.Printf("s is int64\n")
	case reflect.Struct:
		fmt.Printf("s is struct\n")
		fmt.Printf("field num of s is %d\n", v.NumField())
		//NumFiled()获取字段数，v.Field(i)可以取得下标位置的字段信息，返回的是一个Value类型的值
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			//打印字段的名称、类型以及值
			fmt.Printf("name:%s type:%v value:%v\n",
				t.Field(i).Name, field.Type().Kind(), field.Interface())
		}
	default:
		fmt.Printf("default\n")
	}
}

// 对结构体内的字段进行赋值操作
func TestSetStruct(T *testing.T) {

	T.Log("******************测试通过反射操作结构体类型2******************")
	s := Student{
		Name:  "BigOrange",
		Sex:   1,
		Age:   10,
		Score: 80.1,
	}

	fmt.Printf("Name:%v, Sex:%v,Age:%v,Score:%v \n", s.Name, s.Sex, s.Age, s.Score)
	v := reflect.ValueOf(&s) //这里传的是地址！！！

	v.Elem().Field(0).SetString("ChangeName")
	v.Elem().FieldByName("Score").SetFloat(99.9)

	fmt.Printf("Name:%v, Sex:%v,Age:%v,Score:%v \n", s.Name, s.Sex, s.Age, s.Score)
}

// 获取结构体里面的方法
func TestGetStructMethod(T *testing.T) {

	T.Log("******************测试通过反射操作结构体类型3******************")
	s := Student{
		Name:  "BigOrange",
		Sex:   1,
		Age:   10,
		Score: 80.1,
	}

	v := reflect.ValueOf(&s)
	//取得Type信息
	t := v.Type()

	fmt.Printf("struct student have %d methods\n", t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("struct %d method, name:%s type:%v\n", i, method.Name, method.Type)
	}
}
