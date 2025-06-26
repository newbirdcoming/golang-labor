# golang-labor

本仓库是一个用于记录 CSDN 博客中相关源码的 GitHub 仓库，方便读者查阅和学习。每个目录对应一篇或多篇 CSDN 博客的源码实现，并附有简要说明和对应的博客链接。

## 目录说明

### zap-test

- **简介**：该目录包含了基于 [zap](https://github.com/uber-go/zap) 日志库的 Go 项目示例，演示了如何通过配置文件灵活管理日志输出、日志分割、日志级别等功能。
- **相关博客**：
  - [Go高性能日志库zap的使用与配置详解（含源码）](https://blog.csdn.net/LanJieZhiFu/article/details/148922806?sharetype=blogdetail&sharerId=148922806&sharerefer=PC&sharesource=LanJieZhiFu&spm=1011.2480.3001.8118)  
- **主要内容**：
  - 日志配置文件（支持热更新）
  - 控制台与文件多路输出
  - 日志切割与归档
  - 日志级别与格式自定义

### gomock-test

- **简介**：该目录演示了如何使用 [gomock](https://blog.csdn.net/LanJieZhiFu/article/details/148769776?sharetype=blogdetail&sharerId=148769776&sharerefer=PC&sharesource=LanJieZhiFu&spm=1011.2480.3001.8118) 进行 Go 接口的单元测试与 mock 代码生成。
- **主要内容**：
  - `person/`：定义需要 mock 的接口
  - `mocks/`：通过 mockgen 工具生成的 mock 实现
  - `student/`：依赖接口的业务结构体及其单元测试
- **用法简述**：
  1. 安装依赖：  
     `go get github.com/golang/mock/gomock`  
     `go install github.com/golang/mock/mockgen@latest`
  2. 生成 mock 代码：  
     `go generate ./...`
  3. 运行测试：  
     `go test ./student`
- **相关博客**：如有对应博客可补充链接

---

后续会持续补充更多 CSDN 博客源码及其目录说明，欢迎关注与交流！
