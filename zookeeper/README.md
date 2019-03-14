测试Golang的zk库
===

1. 启动zk集群：

        💤  zookeeper [master] ⚡  cd DockerCompose
        💤  DockerCompose [master] ⚡  docker-compose up -d
        Creating network "dockercompose_zk_net" with driver "bridge"
        Creating dockercompose_zoo2_1 ... done
        Creating dockercompose_zoo1_1 ... done
        Creating dockercompose_zoo3_1 ... done

2. 编译*watch*下的程序并且拷贝到另一个名为`test`的容器中，将`test`容器加入zk的network，然后运行:

        💤  zookeeper [master] ⚡  cd watch
        💤  watch [master] ⚡  go build
        💤  watch [master] ⚡  docker run --rm -Pdit --name test pg
        5be393c2437988e45b68f6f2ef7499c9b77c08dbbaf8740402162ab701f3ba6d
        💤  watch [master] ⚡  docker cp $PWD/foo test:/root
        💤  watch [master] ⚡  docker network connect dockercompose_zk_net test                                                                                                                                                                                                        
        💤  watch [master] ⚡  docker exec test /root/foo

        2019/03/14 10:26:33 [Session Watcher] {EventSession StateConnecting  <nil> 172.18.0.4:2181}
        2019/03/14 10:26:33 Connected to 172.18.0.4:2181
        2019/03/14 10:26:33 [Session Watcher] {EventSession StateConnected  <nil> 172.18.0.4:2181}
        2019/03/14 10:26:33 Authenticated: id=216175115617828867, timeout=10000
        2019/03/14 10:26:33 Re-submitting `0` credentials after reconnect
        2019/03/14 10:26:33 [Session Watcher] {EventSession StateHasSession  <nil> 172.18.0.4:2181}

    (注意上面的超时时间设置了10s)

3. 将`test`容器从zk网络中移除，模拟网络分区的情况：

        💤  zk_admin [master] docker exec test date && docker network disconnect dockercompose_zk_net test
        Thu Mar 14 10:26:45 UTC 2019

    此时，程序输出：

        2019/03/14 10:26:50 Recv loop terminated: err=read tcp 172.18.0.5:51864->172.18.0.4:2181: i/o timeout
        2019/03/14 10:26:50 Send loop terminated: err=<nil>
        2019/03/14 10:26:50 [Session Watcher] {EventSession StateDisconnected  <nil> 172.18.0.4:2181}
        2019/03/14 10:26:50 [Session Watcher] {EventSession StateConnecting  <nil> 172.18.0.2:2181}
        2019/03/14 10:26:51 Failed to connect to 172.18.0.2:2181: dial tcp 172.18.0.2:2181: i/o timeout
        2019/03/14 10:26:51 [Session Watcher] {EventSession StateConnecting  <nil> 172.18.0.3:2181}
        2019/03/14 10:26:52 Failed to connect to 172.18.0.3:2181: dial tcp 172.18.0.3:2181: i/o timeout
        2019/03/14 10:26:52 [Session Watcher] {EventSession StateConnecting  <nil> 172.18.0.4:2181}
        ...

    可见，程序在5秒以后发现自己与zk server断开，然后自动发起重连。（这里的5s应该和10s的timeout值有关，因为如果把timeout设置的很大，这里的重连会在网络断开很久以后才开始）。

    在10s以后再将网络恢复：

        💤  zk_admin [master] docker network connect dockercompose_zk_net test 

    此时，程序的输出如下：

        2019/03/14 10:27:01 Connected to 172.18.0.4:2181
        2019/03/14 10:27:01 [Session Watcher] {EventSession StateConnected  <nil> 172.18.0.4:2181}
        2019/03/14 10:27:01 Authentication failed: zk: session has been expired by the server
        2019/03/14 10:27:01 [Session Watcher] {EventSession StateExpired  <nil> 172.18.0.4:2181}
        2019/03/14 10:27:01 [Session Watcher] {EventSession StateDisconnected  <nil> 172.18.0.4:2181}
        2019/03/14 10:27:01 [Session Watcher] {EventSession StateConnecting  <nil> 172.18.0.2:2181}
        2019/03/14 10:27:01 [Existence Watcher] (zk.Event) {
         Type: (zk.EventType) EventNotWatching,
         State: (zk.State) StateDisconnected,
         Path: (string) (len=4) "/foo",
         Err: (*errors.errorString)(0xc000086690)(zk: session has been expired by the server),
         Server: (string) ""
        }
        2019/03/14 10:27:01 press Ctrl-C to quit...
        2019/03/14 10:27:01 Connected to 172.18.0.2:2181
        2019/03/14 10:27:01 [Session Watcher] {EventSession StateConnected  <nil> 172.18.0.2:2181}
        2019/03/14 10:27:01 Authenticated: id=72059927541317635, timeout=10000
        2019/03/14 10:27:01 Re-submitting `0` credentials after reconnect
        2019/03/14 10:27:01 [Session Watcher] {EventSession StateHasSession  <nil> 172.18.0.2:2181}

    注意：

        - ExistsWatcher接收到的event为`EventNotWatching`
        - Session Channel会在重连之后接收到expiration event

    而如果，网络在10s内恢复，那么:
        
        - ExistsWatcher不会接收任何event，而是继续watch
        - Session Channel在重连之后不会接收到expiration event，只会接收到disconnect,connect等event

