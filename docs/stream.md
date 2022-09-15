# IOS 抓包

### 1. 前往 App Store 下载安装 Stream

<img src="./images/1.png" alt="1" width="200" />

### 2. 打开 Stream

<img src="./images/2.png" alt="2" width="200" />

### 3. 配置 HTTPS 抓包

- 需要安装证书，安装成功如下图所示
- 需要点击 清除缓存

<img src="./images/3.png" alt="3" width="200" />

### 4. 点击 开始抓包

<img src="./images/4.jpeg" alt="4" width="200" />

### 5. 打开微信 --> 访问 羊了个羊 --> 打开排行榜

### 6. 点击 停止抓包

### 7. 点击 抓包历史

### 8. 选择 抓包记录

<img src="./images/5.png" alt="5" width="200" />

### 9. 按域名检索，选择 `cat-match.easygame2021.com` 域名记录

<img src="./images/6.png" alt="6" width="200" />

### 10. 在详情的请求中，找到 t 对应的内容，就是后续要用到的身份信息，拷贝下来

<img src="./images/7.png" alt="7" width="200" />

### 11. 使用本工具执行通关

替换 `<your-custom-token>` 为 上一步获取到的 t 对应的内容

```shell
ylgy --token <your-custom-token>
```