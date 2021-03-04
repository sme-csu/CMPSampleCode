### CMP平台基于Go SDK的Sample Code

#### 相关资料如下
1. disk/main.go

    Azure SDK for Go 快速入门，创建磁盘的一个例子。

2. query/main.go
    
    批量查询的示例，本示例将获取订阅下所有虚机实例。

3. deployment/main.go

    使用ARM模版部署Azure资源的示例。
    
4. ARM Template/vms.json

    批量部署虚拟机的ARM模版

#### 如何运行

1. 安装Go语言。[官方地址](https://golang.org/doc/install)
2. 安装Azure SDK for Go。[官方地址](https://docs.microsoft.com/zh-cn/azure/developer/go/azure-sdk-install)
3. 修改代码中的订阅相关的常量参数。
4. 使用VS Code打开go代码文件，直接运行。

#### 代码中的常量说明

- tenantID    租户ID
- subscriptionID 订阅ID
- clientID    应用程序ID
- clientSecret   应用程序密码
- templateURL    调用远程ARM模版的地址
    - 例如：https://raw.githubusercontent.com/sme-csu/CMPSampleCode/main/ARM Template/vms.json