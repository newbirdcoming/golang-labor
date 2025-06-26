# 一、gomock简介与安装

## 1.1 gomock是什么？

gomock是Go语言的一个模拟框架，它允许开发者创建接口的模拟实现，以便在测试中替代真实的依赖项。gomock包含两个主要部分：

1. **gomock包**：提供创建和管理模拟对象的API
2. **mockgen工具**：用于从接口定义生成模拟实现代码

虽然官方在2023年6月停止维护了gomock，但Uber团队维护了一个活跃的分支，社区仍在广泛使用。

## 1.2 安装gomock

```
# 安装gomock包
go get -u github.com/golang/mock/gomock

# 安装mockgen工具（Go 1.16+）
go install github.com/golang/mock/mockgen@v1.6.0

# 验证安装
mockgen -version
```



# 二、 gomock基本使用

## 2.1 定义接口

ps:这里的接口可以使用gomock在使用的时候进行接口的模拟实现

> 这里举例一个person接口 包含eat和sleep
>
> student复用了person 包含name和person

**person.go**

```
type Person interface {
	Eat() string
	Sleep(name string) string
}
```

**student.go**

```
type Student struct {
	Name string
	p    person.Person
}

func (s *Student) Eat() string {
	return s.p.Eat()
}

func (s *Student) Sleep() string {
	re
```





# 三、生成mock代码

ps:这里生成person接口的mock代码 然后student复用person的具体mock实现进行相应的操作

**修改person.go**

```
//go:generate mockgen -destination=../mocks/mock_person.go -package=mocks -source=person.go
type Person interface {
	Eat() string
	Sleep(name string) string
}
```

> //go:generate mockgen -destination=../mocks/mock_person.go -package=mocks -source=person.go
>  这是一种生成mock代码的注释 执行go generate ./... 会自动扫描并将所有带有 //go:generate 注释的代码接口进行mock实现
>
> 1） -destination=../mocks/mock_person.go 表示mock代码的文件名
>
> 2）-package=mocks   表示代码的父文件夹名
>
> 3）-source=person.go 表示包含//go:generate注释代码的文件名
>
> 4）上述都使用的**相对路径** 相对于当前文件



## 3.1 编写测试用例

```
func Test_Eat(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPerson := mocks.NewMockPerson(ctrl)
	defer ctrl.Finish()
	mockStudent := Student{
		Name: "张三",
		p:    mockPerson,
	}
	mockPerson.EXPECT().Eat().Return("张三在吃饭")
	content := mockStudent.Eat()
	fmt.Println("结果为" + content)
	first := mockPerson.EXPECT().Sleep("张三")
	mockStudent.Sleep()
	mockPerson.EXPECT().Sleep("XXX").After(first)
	mockStudent.Name = "XXX"
	mockStudent.Sleep()
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
```



> 1）    
>
> ​    // 创建一个 gomock 控制器
>
> ​    ctrl := gomock.NewController(t)
>
> ​    // 使用 mock 生成器创建一个 Person 接口的 mock 实现
>
> ​    mockPerson := mocks.NewMockPerson(ctrl)
>
> ​    // 测试结束后释放控制器
>
> ​    defer ctrl.Finish()
>
> 2）
>
> ​    // 创建一个测试用的学生实例，注入 mock 的 Person 接口
>
> ​    mockStudent := Student{
>
> ​        Name: "张三",
>
> ​        p:    mockPerson,
>
> ​    }
>
> 3）// 设置对 Eat 方法的调用期望
>
> ​    mockPerson.EXPECT().Eat().Return("张三在吃饭")
>
> 4） // 调用学生的 Eat 方法
>
> ​    content := mockStudent.Eat()
>
> ​    fmt.Println("结果为" + content)

**！！！！上面仅仅是简单的演示gomock**