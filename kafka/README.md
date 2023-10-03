 本示例逐步参考[Getting Started with Apache Kafka and Go](https://developer.confluent.io/get-started/go/)文档并实现
 
思路总结
1. 环境配置, 使用Confluent Cloud快速创建一个试验性的kafka集群; 
2. 本地创建一个go project with module, 并导入 kafka client for go;
3. 按照文档创建`getting-started.properties`配置文件,用于访问kafka集群;
4. 在cloud集群中创建一个topic, kafka使用topic组织和管理事件(event), topic支持分区, 示例中设置为1;
5. 构建Producer代码和util代码用于读取配置文件中的配置, 并编译;
6. 构建Consumer代码, 编译时携带util代码;
7. 编译Producer/Consumer可执行文件(5,6步骤中已执行, 忽略);
8. 运行producer应用和consumer应用
9. 完成, Finished!