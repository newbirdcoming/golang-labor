package student

import (
	"fmt"
	"gomock-test/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

// Test_Eat 测试学生的吃饭和睡觉行为
func Test_Eat(t *testing.T) {
	// 创建一个 gomock 控制器
	ctrl := gomock.NewController(t)
	// 使用 mock 生成器创建一个 Person 接口的 mock 实现
	mockPerson := mocks.NewMockPerson(ctrl)
	// 测试结束后释放控制器
	defer ctrl.Finish()

	// 创建一个测试用的学生实例，注入 mock 的 Person 接口
	mockStudent := Student{
		Name: "张三",
		p:    mockPerson,
	}

	// 设置对 Eat 方法的调用期望
	mockPerson.EXPECT().Eat().Return("张三在吃饭")
	// 调用学生的 Eat 方法
	content := mockStudent.Eat()
	fmt.Println("结果为" + content)

	// 设置第一次对 Sleep 方法的调用期望
	first := mockPerson.EXPECT().Sleep("张三")
	// 调用学生的 Sleep 方法
	mockStudent.Sleep()

	// 设置第二次对 Sleep 方法的调用期望，并指定在 first 之后执行
	mockPerson.EXPECT().Sleep("XXX").After(first)
	// 修改学生姓名并再次调用 Sleep
	mockStudent.Name = "XXX"
	mockStudent.Sleep()

	// 以下是被注释的代码，展示了更多 mock 功能的使用方式
	//mockPerson.
	//	EXPECT().
	//	Sleep("张三").
	//	Return("张三在睡觉").
	//	Do(func(name string) {
	//		fmt.Println("模拟调用了Sleep方法，参数为：" + name)
	//	})
	//content = mockStudent.Sleep()
	//fmt.Println("结果为" + content)
}
