### 现有问题

----

### 服务端待开发列表
+ 增加addlist.txt文件，在执行的时候将AddList压入到Map中，然后进行比对，有的从map中删除，所有的list行读取完之后，再执行map中的行
+ 无法解析的地址，做第二次查询dns信息，如果仍然未查询到，从list.txt中删除
+ 新建一个空的addlist文件，覆盖云盘，将有效的list文件，覆盖云盘的list文件

### 客户端待开发列表
+ list.txt从云盘上下载
+ Client端增加－A 功能，维护本地的addList.txt文件，从云盘上下载之后，追加条目，然后覆盖
+ -C 命令运行最后，重启dns服务
+ 增加-A命令，通过代理端口访问 http://域名.ipaddress.com 然后通过goquery去获取dns地址信息