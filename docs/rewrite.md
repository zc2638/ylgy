# 重写第二关

### 1、通过下面的链接操作, 让`Charles`可以抓取手机请求。

[Charles抓包教程](https://www.jianshu.com/p/ff85b3dac157)

### 2、设置 Charles 重写规则

#### 配置 `Location`

![](./docs/images/rewrite1.png)

#### 配置 `Rewrite Rule`

![](./docs/images/rewrite2.png)

替换的内容：

```
{"err_code":0,"err_msg":"","data":{"map_md5":["046ef1bab26e5b9bfe2473ded237b572","046ef1bab26e5b9bfe2473ded237b572"],"map_seed":[3042065807,873076577,3555997730,3291121329]}}
```

#### 3、在手机上正常闯关

玩得开心！