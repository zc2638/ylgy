# ylgy

羊了个羊 通关程序

**提醒：** 需要使用抓包软件获取`身份凭证`

## 一、抓包

抓取 `cat-match.easygame2021.com` 的请求包内 Header 为 `t` 的内容

### 1、电脑

[Charles抓包教程](https://www.jianshu.com/p/ff85b3dac157)

### 2、IOS

[Stream 保姆级教程](./docs/stream.md)

## 二、安装

### 1、源码

```shell
go install github.com/zc2638/ylgy/cmd/ylgy@latest
```

### 2、安装包

通过 [Release](https://github.com/zc2638/ylgy/releases) 页面下载

## 三、使用

### 1、命令行

- 执行以下命令前，将 `<your-custom-token>` 替换为实际内容
- `--times` 对应设置通关次数，默认为 `1`

```shell
ylgy --token <your-custom-token> --times 1
```

### 2、Docker

- 执行以下命令前，将 `<your-custom-token>` 替换为实际内容
- `--times` 对应设置通关次数，默认为 `1`

```shell
docker run --rm -it zc2638/ylgy sh -c 'ylgy --token <your-custom-token> --times 1' 
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
