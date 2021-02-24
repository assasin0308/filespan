# filespan

golang 实现简单地文件上传,下载,信息更新,删除操作.

### 1. MySQL主从

```JSON
// 1. 主节点
 show master status;
+---------------+----------+--------------+------------------+-------------------+
| File          | Position | Binlog_Do_DB | Binlog_Ignore_DB | Executed_Gtid_Set |
+---------------+----------+--------------+------------------+-------------------+
| binlog.000005 |      374 |              |                  |                   |
+---------------+----------+--------------+------------------+-------------------+
// 2. 从节点:
CHANGE MASTER TO MASTER_HOST='192.168.2.101',MASTER_USER='root',MASTER_PASSWORD='a12345',MASTER_LOG_FILE='binlog.000005',MASTER_LOG_POS=0;
// 3. 启动从节点:
 start slave;
// 同:
start/stop slave sql_thread;
start/stop slave io_thread;
// 4. 查看从节点状态:
show slave status\G


*************************** 1. row ***************************
               Slave_IO_State: Connecting to master
                  Master_Host: 192.168.2.101
                  Master_User: root
                  Master_Port: 3306
                Connect_Retry: 60
              Master_Log_File: binlog.000005
          Read_Master_Log_Pos: 4
               Relay_Log_File: 192-relay-bin.000001
                Relay_Log_Pos: 4
        Relay_Master_Log_File: binlog.000005
// 主要看以下两个线程状态:
***********************************************************************
             Slave_IO_Running: Connecting
            Slave_SQL_Running: Yes
***********************************************************************
              Replicate_Do_DB:
          Replicate_Ignore_DB:
           Replicate_Do_Table:
       Replicate_Ignore_Table:
      Replicate_Wild_Do_Table:
  Replicate_Wild_Ignore_Table:
                   Last_Errno: 0
                   Last_Error:
                 Skip_Counter: 0
          Exec_Master_Log_Pos: 4
              Relay_Log_Space: 154
              Until_Condition: None
               Until_Log_File:
                Until_Log_Pos: 0
           Master_SSL_Allowed: No
           Master_SSL_CA_File:
           Master_SSL_CA_Path:
              Master_SSL_Cert:
            Master_SSL_Cipher:
               Master_SSL_Key:
        Seconds_Behind_Master: NULL
Master_SSL_Verify_Server_Cert: No
                Last_IO_Errno: 1045
                Last_IO_Error: error connecting to master 'root@192.168.2.101:3306' - retry-time: 60  retries: 1
               Last_SQL_Errno: 0
               Last_SQL_Error:
  Replicate_Ignore_Server_Ids:
             Master_Server_Id: 0
                  Master_UUID:
             Master_Info_File: /data/mysql/master.info
                    SQL_Delay: 0
          SQL_Remaining_Delay: NULL
      Slave_SQL_Running_State: Slave has read all relay log; waiting for more updates
           Master_Retry_Count: 86400
                  Master_Bind:
      Last_IO_Error_Timestamp: 210319 23:16:48
     Last_SQL_Error_Timestamp:
               Master_SSL_Crl:
           Master_SSL_Crlpath:
           Retrieved_Gtid_Set:
            Executed_Gtid_Set:
                Auto_Position: 0
         Replicate_Rewrite_DB:
                 Channel_Name:
           Master_TLS_Version:

// 5. 测试主从状态:
主: create database test1 default character set utf8;
```

