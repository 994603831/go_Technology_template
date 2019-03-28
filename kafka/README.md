# windows10下Kafka环境搭建



  包含JDK+Zookeeper+Kafka三部分。

1) 安装包：Java SE Development Kit 9.0.1
      下载地址：http://www.oracle.com/technetwork/java/javase/downloads/jdk9-downloads-3848520.html



2)    配置环境：（与之前版本设置有差异）

![img](https://img-blog.csdn.net/20171231205549337?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

​       ![img](https://img-blog.csdn.net/20171231204810251?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

​        ![img](https://img-blog.csdn.net/20171231204814006?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



3)    打开Dos界面，运行java：

​       ![img](https://img-blog.csdn.net/20171231204818416?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



4)    运行javac：             

​           ![img](https://img-blog.csdn.net/20171231204822633?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



5)    配置完成。


Zookeeper：
1)    建议下载稳定版。

       下载地址：http://mirrors.hust.edu.cn/apache/zookeeper/


​       ![img](https://img-blog.csdn.net/20171231204826214?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

2)    下载后解压到一个目录：eg: D:\Java\Tool\zookeeper-3.4.10



3)    在zookeeper-3.4.10目录下，新建文件夹，并命名(eg: data).(路径为：D:\Java\Tool\zookeeper-3.4.10\conf\data)



4)    进入Zookeeper设置目录，eg: D:\Java\Tool\zookeeper-3.4.10\conf

       复制“zoo_sample.cfg”副本à并将副本重命名为“zoo.cfg”
    
       在任意文本编辑器（eg：记事本）中打开zoo.cfg
    
       找到并编辑dataDir=D:\\Java\\Tool\\zookeeper-3.4.10\\data



5)    添加系统环境变量：

       在系统变量中添加ZOOKEEPER_HOME = D:\Java\Tool\zookeeper-3.4.10
    
       编辑path系统变量，添加为路径%ZOOKEEPER_HOME%\bin



6)    在zoo.cfg文件中修改默认的Zookeeper端口（默认端口2181）



7)    Dos下运行：zkserver

​        ![img](https://img-blog.csdn.net/20171231204830449?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

8)    搭建成功。




Kafka：
1)    安装包：kafka_2.12-1.0.0.tgz

       下载地址：http://kafka.apache.org/downloads.html
    
       推荐版本：kafka_2.12-1.0.0.tgz


​          ![img](https://img-blog.csdn.net/20180101001411237?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

2)    下载后解压缩。eg: D:\Tools\kafka_2.11-1.0.0\



3)    建立一个空文件夹 logs. eg: D:\Tools\kafka_2.11-1.0.0\logs



4)    进入config目录，编辑 server.properties文件(eg: 用“写字板”打开)。

       找到并编辑log.dirs= D:\\Tools\\kafka_2.11-1.0.0\\logs
    
       找到并编辑zookeeper.connect=localhost:2181。表示本地运行。
    
       (Kafka会按照默认，在9092端口上运行，并连接zookeeper的默认端口：2181)



运行：请确保在启动Kafka服务器前，Zookeeper实例已经准备好并开始运行。（就是开着Zookeeper窗口不要关）

1)    在 D:\WorkSoftware\kafka_2.11-1.0.0下，按住shift+鼠标右键。

       选择“在此处打开Powershell窗口（S）”（如果没有此选项，在此处打开命令窗口）。



2)    运行：.\bin\windows\kafka-server-start.bat .\config\server.properties



3)    可能会报错：“找不到或无法加载主类 Files\java\jdk-9.0.1\lib;C:\Program”

​        ![img](https://img-blog.csdn.net/20171231204837402?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)



4)    解决（3）的办法：

       在kafka安装目录中找到bin\windows目录中的kafka-run-class.bat为%CLASSPATH%加上双引号(可用Matlab打开，并进行搜索)
    
       修改前：setCOMMAND=%JAVA%%KAFKA_HEAP_OPTS% %KAFKA_JVM_PERFORMANCE_OPTS% %KAFKA_JMX_OPTS%%KAFKA_LOG4J_OPTS% -cp%CLASSPATH% %KAFKA_OPTS% %*   
    
       修改后：SetCOMMAND=%JAVA%%KAFKA_HEAP_OPTS% %KAFKA_JVM_PERFORMANCE_OPTS% %KAFKA_JMX_OPTS%%KAFKA_LOG4J_OPTS% -cp"%CLASSPATH%"%KAFKA_OPTS% %*



5)    再次运行：.\bin\windows\kafka-server-start.bat.\config\server.properties

​        ![img](https://img-blog.csdn.net/20171231204840551?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvdGlhbm1hbmNobg==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

6）搭建成功。


原文：https://blog.csdn.net/tianmanchn/article/details/78943147 
