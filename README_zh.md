# bk-client-go
Apache Bookkeeper的Go客户端,目前还在初期设计阶段.  

需要实现以下的功能:  
1. 连接池管理  
2. 支持连接到多个bookie后端以及读写
3. 协议封装

后续希望将客户端集成到Pulsar-function-go当中,做到go function支持有状态function.