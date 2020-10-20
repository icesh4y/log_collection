# log_collection
日志收集demo

jdk : 11.08\
kafka: 2.12-2.20

###1. config/zookeeper.properites 配置文件
  - 16 ： dataDir=D:/fafuka/tmp/zookeeper 修改存放地址
  
###2. config/server.properites 配置文件
  - 60 ： log.dirs=D:/kafuka/tmp/kafka-logs  修改存放地址

###3. 切换至安装路径
  - bin\windows\zookeeper-server-start.bat config\zookeeper.properties  启动zookeeper
  - bin\windows\kafka-server-start.bat config\server.properties 启动卡夫卡
  
###4. 卡夫卡终端消费
  - bin\windows\kafka-console-consumer.bat --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
