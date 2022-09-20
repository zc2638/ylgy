# ylgy

羊了个羊 通关程序，新版本接口已被修改。

支持两种方式通关：

- 1、[重写第二关的方式](./docs/rewrite.md)
- 2、按照以下文档步骤操作执行通关程序

## 一、获取 UID

- 1、登陆游戏
- 2、点击 `俺的名片`
- 3、查看ID

## 二、安装

### 1、源码

```shell
go install github.com/zc2638/ylgy/cmd/ylgy@latest
```

### 2、安装包

通过 [Release](https://github.com/zc2638/ylgy/releases) 页面下载

## 三、使用

### 1、命令行

- 执行以下命令前，将 `<your-uid>` 替换为实际内容
- `--times` 对应设置通关次数，默认为 `1`

```shell
ylgy --uid <your-uid>
```

### 2、Docker

- 执行以下命令前，将 `<your-uid>` 替换为实际内容
- `--times` 对应设置通关次数，默认为 `1`

```shell
docker run --rm -it zc2638/ylgy sh -c 'ylgy --uid <your-uid>' 
```

## 声明

本项目仅供学习交流，严禁用作商业行为！  
因他人私自不正当使用造成的违法违规行为与本人无关！  
如有任何问题可联系本人删除！

## Buy me a Coffee

如果有帮助到您，并且也使您感到开心的话，可以慷慨的为我买杯coffee吗

<img src="./docs/images/alipay.jpeg" alt="alipay" width="200" />
<br/>
<img src="./docs/images/wechat.jpeg" alt="wechat" width="200" />
