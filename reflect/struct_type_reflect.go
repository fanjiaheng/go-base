package reflect

import "fmt"

type Student struct {
	Name  string `json:"jsName" db:"dbName"`
	Sex   int
	Age   int
	Score float32
}

func (s *Student) SetName(name string) {
	fmt.Printf("有参数方法 通过反射进行调用:%v\n", s)
	s.Name = name
}

func (s *Student) PrintStudent() {
	fmt.Printf("无参数方法 通过反射进行调用:%v\n", s)
}
