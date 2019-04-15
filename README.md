# phonearea
 基于Go写的简单的一个查询手机号码归属地的http接口

 数据来源于 ip138.com
 
 源码内使用到了第三方goquery包
 </br>
 </br>
 bin文件夹内为已编译二进制文件</br>
 编译环境:ubuntu18.04</br>
 测试:centos7可直接运行

 执行时可设置监听端口号:`./phonearea 8888` 默认监听8088端口

 linux执行: `nohup 存放路径/phonearea 监听端口号 &>/dev/null &`
 </br>
 访问: `ip:8888/phone?phone=xxxxxx`
 </br>
 记得开放端口哦
