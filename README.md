not in 是不能命中索引的
用 EXISTS 或 NOT EXISTS 代替

    select * from test1
       where EXISTS (select * from test2 where id2 = id1 )
     
     
    select * FROM test1
     where NOT EXISTS (select * from test2 where id2 = id1 )

2、用JOIN 代替

    select id1 from test1
       INNER JOIN test2 ON id2 = id1
     
     
    select id1 from test1
       LEFT JOIN test2 ON id2 = id1
       where id2 IS NULL

       
因为HTTP/2底层是采用TCP协议实现的，虽然解决了HTTP队头阻塞的问题，但是对于TCP队头阻塞的问题却无能为力。

TCP传输过程中会把数据拆分为一个个按照顺序排列的数据包，这些数据包通过网络传输到了接收端，接收端再按照顺序将这些数据包组合成原始数据，这样就完成了数据传输。

但是如果其中的某一个数据包没有按照顺序到达，接收端会一直保持连接等待数据包返回，这时候就会阻塞后续请求。这就发生了TCP队头阻塞。

另外，TCP这种可靠传输是靠三次握手实现的，TCP三次握手的过程客户端和服务器之间需要交互三次，那么也就是说需要消耗1.5 RTT。如果是HTTPS那么消耗的RTT就更多。

而因为很多中间设备比较陈旧，更新换代成本巨大，这就导致TCP协议升级或者采用新的协议基本无法实现。

所以，HTTP/3选择了一种新的技术方案，那就是基于UDP做改造，这种技术叫做QUIC。

那么问题来了，HTTP/3是如何使用的UDP呢？做了哪些改造？如何保证连接的可靠性？UDP协议就没有僵化的问题了吗
————————————————

                            版权声明：本文为博主原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接和本声明。
                        
原文链接：https://blog.csdn.net/hollis_chuang/article/details/111150623

Java 中共有三种变量，分别是类变量、成员变量和局部变量。他们分别存放在JVM的方法区、堆内存和栈内存


JPA 设计为在同一个事务中对同一个 ID 的实体对象进行任何属性修改后再次查询时获得更新后的数据，主要是为了保证事务内数据的一致性视图，提高性能以及简化应用程序的开发复杂度
because：持久化上下文实际上也是一个缓存层，称为第一级缓存。在这个缓存中的实体实例是共享的，即在同一个事务内的多次查询如果涉及到相同的实体实例，则返回的是同一份内存中的对象引用。
这提高了性能，因为它避免了不必要的数据库查询。

没有save，update也能更新数据
entityManager管理数据库实体对象有4个状态：新建，托管，游离，删除。
只要是托管状态的对象，字段发生了变动，entityManager就会在事务结束时将数据的更新，更新到数据库中
游离（Datached）：游离对象，有id值，但没有和持久化上下文（Persistence Context）建立关联。
    托管状态对象提交事务之后，对象状态由托管状态转换为游离状态
    托管状态对象调用em.clear()方法之后，对象状态由托管状态转换为游离状态
    New出来的对象，id赋值之后，也为游离状态
游离态
refresh
    方法可以保证当前的实例与数据库中的实例的内容一致，注意：是反向同步，将数据库中的数据同步到实体中


【强制】生产环境禁止使用System.out或System.err输出或使用e.printStackTrace() 打印异常堆栈。 说明：标准日志输出与标准错误输出文件每次Jboss重启时才滚动，如果大量输出送往这两个文件，容易造成文件大小超过操作系统大小限制。
9.【强制】异常信息应该包括两类信息：案发现场信息和异常堆栈信息。如果不处理，那么通过关键字throws往上抛出。
正例：logger.error("inputParams: {} and errorMessage: {}", 各类参数或者对象toString(), e.getMessage(), e);
10.【强制】日志打印时禁止直接用JSON工具将对象转换成String。
说明：如果对象里某些get方法被覆写，存在抛出异常的情况，则可能会因为打印日志而影响正常业务流程的执行


【强制】超过三个表禁止join。需要join的字段，数据类型保持绝对一致；多表关联查询时，保证被关联的字段需要有索引。
说明：即使双表join也要注意表索引、SQL性能。
3.【强制】在varchar字段上建立索引时，必须指定索引长度，没必要对全字段建立索引，根据实际文本区分度决定索引长度。
说明：索引的长度与区分度是一对矛盾体，一般对字符串类型数据，长度为20的索引，区分度会高达90%以上，可以使用count(distinct left(列名，索引长度)) / count(*) 的区分度来确定。

如果有order by的场景，请注意利用索引的有序性。order by最后的字段是组合索引的一部分，并且放在索引组合顺序的最后，避免出现filesort的情况，影响查询性能。 正例：where a = ? and b = ? order by c；索引：a_b_c 反例：索引如果存在范围查询，那么索引有序性无法利用，如：WHERE a > 10 ORDER BY b；索引a_b无法排序。


调用远程操作必须有超时设置。
说明：类似于HttpClient的超时设置需要自己明确去设置Timeout。根据经验表明，无数次的故障都是因为没有设置 超时时间。

高并发服务器建议调小TCP协议的time_wait超时时间。
说明：操作系统默认240秒后，才会关闭处于time_wait状态的连接，在高并发访问下，服务器端会因为处于time_wait的连接数太多，可能无法建立新的连接，所以需要在服务器上调小此等待值。
正例：在linux服务器上请通过变更/etc/sysctl.conf文件去修改该缺省值（秒）：net.ipv4.tcp_fin_timeout=30

调大服务器所支持的最大文件句柄数（File Descriptor，简写为fd）
说明：主流操作系统的设计是将TCP / UDP连接采用与文件一样的方式去管理，即一个连接对应于一个fd。主流的linux 服务器默认所支持最大fd数量为1024，当并发连接数很大时很容易因为fd不足而出现“open too many files”错误， 导致新的连接无法建立。建议将linux服务器所支持的最大句柄数调高数倍（与服务器的内存数量相关）

给JVM环境参数设置-XX：+HeapDumpOnOutOfMemoryError参数，让JVM碰到OOM场景时输出dump信息。
说明：OOM的发生是有概率的，甚至相隔数月才出现一例，出错时的堆内信息对解决问题非常有帮助

类在设计与实现时要符合单一原则

谨慎使用继承的方式来进行扩展，优先使用聚合/组合的方式来实现。
说明：不得已使用继承的话，必须符合里氏代换原则，此原则说父类能够出现的地方子类一定能够出现，比如，“把钱交出来”，钱的子类美元、欧元、人民币等都可以出现



异步http client
okhttp ，apache httpclient 背后是线程池，而非真异步。  真异步要用jdk的 httpclient，webclient，适用于高并发，节省线程。
JVM中的直接内存（Direct Memory）是什么？

直接内存（Direct Memory）不是JVM堆内存的一部分，它是通过在Java代码中使用NIO库分配的内存，直接在操作系统的物理内存中分配。主要特点和用途：

1、避免内存复制： 直接内存访问避免了JVM堆和本地堆之间的内存复制，提高性能。

2、高效IO操作： 在NIO中，使用直接内存可以提高文件的读写效率。

3、内存管理： 直接内存的分配和回收不受JVM垃圾回收器管理，需要手动释放。


ReentrantLock‌ 则适用于以下场景：

    ‌需要更灵活的锁控制‌：ReentrantLock提供了更多的锁控制功能，如tryLock()方法，可以尝试非阻塞地获取锁‌12。
    ‌尝试非阻塞地获取锁‌：与synchronized不同，ReentrantLock允许你尝试获取锁而不阻塞当前线程，这在某些需要高效并发处理的场景中非常有用‌2。
    ‌响应中断‌：ReentrantLock支持响应中断的锁获取方式，即如果线程在等待锁的过程中被中断，它可以立即响应中断并放弃锁的获取，而synchronized则无法响应中断‌

    
数据库不用红黑树的主要原因包括平衡性不如B树、磁盘I/O效率低、以及复杂度高和空间利用率低‌。

首先，红黑树虽然是一种自平衡的二叉搜索树，能够在插入、删除和查找操作中保持O(log n)的时间复杂度，但在实际应用中，特别是在处理大规模数据的数据库系统中，其平衡性不如B树。B树及其变种（如B+树）通过多路分支结构，能够大幅降低树的高度，使得数据存储和检索的效率显著提升‌12。

其次，红黑树在磁盘I/O效率方面表现不佳。由于红黑树的高度较大，导致在涉及大量数据的数据库操作时，需要更多的磁盘I/O操作。而数据库的设计需要考虑到如何最小化磁盘I/O操作以提高性能。B树及其变种在这方面表现得更好，因为它们可以将大量数据存储在一个节点中，从而减少树的高度和磁盘读取次数‌23。

此外，红黑树的复杂度也相对较高。虽然其插入和删除操作的时间复杂度为O(log n)，但在实际应用中，这些操作的实现相对复杂，且可能需要频繁地进行树的平衡调整。相比之下，B树及其变种在插入、删除和更新操作方面能够保持较高的效率，且减少了重平衡的次数‌


系统性能参数

    ‌文件描述符限制‌：
        fs.file-max：系统级别的最大文件描述符数量。
        fs.nr_open：每个进程可以打开的最大文件描述符数量。
        这些参数可以通过/etc/sysctl.conf文件或sysctl命令进行调整。

    ‌虚拟内存参数‌：
        vm.swappiness：控制内核使用交换空间的倾向。值越高，越倾向于使用交换空间。对于需要高性能的应用服务器，可以设置为较低的值，如10。

    ‌内存管理参数‌：
        vm.overcommit_memory：控制内核是否允许过度分配内存。对于某些需要大量内存分配的应用，可以设置为1以允许过度分配，但需谨慎使用。

    ‌网络参数‌：
        net.ipv4.tcp_max_tw_buckets：控制系统中TIME_WAIT套接字的最大数量。在高并发情况下，可能需要增加此值。

    ‌进程和线程限制‌：
        kernel.pid_max：系统中进程ID的最大值。
        kernel.threads-max：系统中线程的最大数量。
        这些参数同样可以通过/etc/sysctl.conf文件或sysctl命令进行调整。

2. 系统安全参数
    ‌SELinux或AppArmor‌：
        使用SELinux（安全增强型Linux）或AppArmor等强制访问控制工具，进一步增强应用程序的安全性。




现有一批邮件需要发送给订阅顾客，且有一个集群（集群的节点数不定，会动态扩容缩容）来负责具体的邮件发送任务，如何让系统尽快地完成发送？
借助消息中间件，通过发布者订阅者模式来进行任务分配
B. master-slave 部署，由 master 来分配任务
C. 不借助任何中间件，且所有节点均等。通过数据库的 update-returning，从而实现节点之间任务的互斥
分配任务的记录存在数据库。  一个节点用update returning 操作，因为是原子的。另一个就无法执行成功。实现了节点之间任务排斥。


在Linux系统下，我关注过以下内核参数：
‌核心参数包括但不限于‌：
    ‌vm.swappiness‌：设置虚拟内存(swap)使用率，用于控制系统在内存不足时，将页面交换到磁盘的程度‌12。
    ‌net.core.wmem_default, net.core.wmem_max, net.core.rmem_default, net.core.rmem_max‌：这些流控参数用于控制网络连接中的数据传输的缓冲区大小‌1。
    ‌kernel.pid_max‌：设置系统中进程PID的最大值‌1。
    ‌net.ipv4.tcp_syncookies‌：启用或禁用TCP SYN Cookies，可防止SYN flood攻击‌13。
    ‌fs.file-max‌：设置系统中打开文件的最大数量‌14。
    ‌vm.overcommit_memory‌：设置虚拟内存overcommit模式，控制系统是否允许超额分配内存‌12。

此外，还有一些与网络性能相关的关键参数，如：
    ‌net.ipv4.tcp_max_syn_backlog‌：SYN队列长度，用于控制TCP连接请求的最大排队数量‌3。
    ‌net.ipv4.tcp_fin_timeout‌：TCP连接关闭的超时时间，控制主动关闭方FIN-WAIT-2状态的超时时长‌23。
    ‌net.ipv4.tcp_tw_reuse‌ 和 ‌net.ipv4.tcp_tw_recycle‌：这两个参数分别用于开启TIME-WAIT重用和TIME-WAIT快速回收功能，有助于优化网络性能‌23。

以及内存管理方面的参数，例如：

    ‌vm.dirty_ratio‌ 和 ‌vm.dirty_background_ratio‌：这两个参数分别控制系统脏页占内存的比例和后台写入脏页的比例，有助于平衡内存使用和磁盘I/O性能‌2。

这些内核参数对于优化Linux系统的性能、安全性和稳定性至关重要。通过合理调整这些参数，可以显著提升系统的整体表现。需要注意的是，调整内核参数需要谨慎进行，以避免对系统造成不良影响‌24。



TCP是基于字节流的协议‌：TCP不维护消息边界，只保证数据的有序性和可靠性。发送方发送的多个数据包可能会被TCP协议组合成一个大的数据块发送，或者拆分成多个小块发送‌2。

需要根据实际需求来决定。如果需要直观地查看和理解IP地址，那么使用VARCHAR类型字段存储IP地址的方式比较好。如果需要在数据库中存储大量的IP地址，并且需要进行高效的查询，那么使用INT类型字段存储IP地址的方式更为合适。


本项目是本人参加BAT等其他公司电话、现场面试之后总结出来的针对Java面试的知识点或真题，每个点或题目都是在面试中被问过的。

除开知识点，一定要准备好以下套路：  
1. **个人介绍**，需要准备1分钟和5分钟两个版本，包括学习经历、工作经历、项目经历、个人优势、一句话总结。一定要自己背得滚瓜烂熟，张口就来
2. **抽象概念**，当面试官问你是如何理解多线程的时候，你要知道从定义、来源、实现、问题、优化、应用方面系统性地回答
3. **项目强化**，至少与知识点的比例是五五开，所以必须针对简历中的两个以上的项目，形成包括【架构和实现细节】，【正常流程和异常流程的处理】，【难点+坑+复盘优化】三位一体的组合拳
4. **压力练习**，面试的时候难免紧张，可能会严重影响发挥，通过平时多找机会参与交流分享，或找人做压力面试来改善
5. **表达练习**，表达能力非常影响在面试中的表现，能否简练地将答案告诉面试官，可以通过给自己讲解的方式刻意练习
6. **重点针对**，面试官会针对简历提问，所以请针对简历上写的所有技术点进行重点准备

### Java基础
* [JVM原理](https://github.com/xbox1994/Java-Interview/blob/master/MD/Java基础-JVM原理.md)
* [集合](https://github.com/xbox1994/Java-Interview/blob/master/MD/Java基础-集合.md)
* [多线程](https://github.com/xbox1994/Java-Interview/blob/master/MD/Java基础-多线程.md)
* [IO](https://github.com/xbox1994/Java-Interview/blob/master/MD/Java基础-IO.md)
* [问题排查](http://www.wangtianyi.top/blog/2018/07/20/javasheng-chan-huan-jing-xia-wen-ti-pai-cha/?utm_source=github&utm_medium=github)
### Web框架、数据库
* [Spring](https://github.com/xbox1994/Java-Interview/blob/master/MD/Web框架-Spring.md)
* [MySQL](https://github.com/xbox1994/Java-Interview/blob/master/MD/数据库-MySQL.md)
* [Redis](https://github.com/xbox1994/Java-Interview/blob/master/MD/数据库-Redis.md)
### 通用基础
* [操作系统](https://github.com/xbox1994/Java-Interview/blob/master/MD/通用基础-操作系统.md)
* [网络通信协议](https://github.com/xbox1994/Java-Interview/blob/master/MD/通用基础-网络通信协议.md)
* [排序算法](https://github.com/xbox1994/Java-Interview/blob/master/MD/通用基础-排序算法.md)
* [常用设计模式](https://github.com/xbox1994/Java-Interview/blob/master/MD/通用基础-设计模式.md)
* [从URL到看到网页的过程](http://www.wangtianyi.top/blog/2017/10/22/cong-urlkai-shi-,ding-wei-shi-jie/?utm_source=github&utm_medium=github)
### 分布式
* [CAP理论](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-CAP理论.md)
* [锁](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-锁.md)
* [事务](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-事务.md)
* [消息队列](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-消息队列.md)
* [协调器](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-协调器.md)
* [ID生成方式](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-ID生成方式.md)
* [一致性hash](https://github.com/xbox1994/Java-Interview/blob/master/MD/分布式-一致性hash.md)
### 微服务
* [微服务介绍](http://www.wangtianyi.top/blog/2017/04/16/microservies-1-introduction-to-microservies/?utm_source=github&utm_medium=github)
* [服务发现](https://github.com/xbox1994/Java-Interview/blob/master/MD/微服务-服务注册与发现.md)
* [API网关](https://github.com/xbox1994/Java-Interview/blob/master/MD/微服务-网关.md)
* [服务容错保护](https://github.com/xbox1994/Java-Interview/blob/master/MD/微服务-服务容错保护.md)
* [服务配置中心](https://github.com/xbox1994/Java-Interview/blob/master/MD/微服务-服务配置中心.md)
### 算法（头条必问）
* [数组-快速排序-第k大个数](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-数组-快速排序-第k大个数.md)
* [数组-对撞指针-最大蓄水](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-数组-对撞指针-最大蓄水.md)
* [数组-滑动窗口-最小连续子数组](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-数组-滑动窗口-最小连续子数组.md)
* [数组-归并排序-合并有序数组](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-数组-归并排序-合并有序数组.md)
* [链表-链表反转-链表相加](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-链表-反转链表-链表相加.md)
* [链表-双指针-删除倒数第n个](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-链表-双指针-删除倒数第n个.md)
* [二叉树-递归-二叉树反转](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-二叉树-递归-二叉树反转.md)
* [动态规划-连续子数组最大和](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-动态规划-连续子数组最大和.md)
* [数据结构-LRU淘汰算法](https://github.com/xbox1994/Java-Interview/blob/master/MD/算法-数据结构-LRU淘汰算法.md)
### 项目举例
* [秒杀架构](https://github.com/xbox1994/Java-Interview/blob/master/MD/秒杀架构.md)
### 系统设计
* [系统设计-高并发抢红包](https://github.com/xbox1994/Java-Interview/blob/master/MD/系统设计-高并发抢红包.md)
* [系统设计-答题套路](https://github.com/donnemartin/system-design-primer/blob/master/README-zh-Hans.md#%E5%A6%82%E4%BD%95%E5%A4%84%E7%90%86%E4%B8%80%E4%B8%AA%E7%B3%BB%E7%BB%9F%E8%AE%BE%E8%AE%A1%E7%9A%84%E9%9D%A2%E8%AF%95%E9%A2%98)
* [系统设计-在AWS上扩展到数百万用户的系统](https://www.wangtianyi.top/blog/2019/03/06/zai-awsshang-kuo-zhan-dao-shu-bai-mo-yong-hu-de-xi-tong/?utm_source=github&utm_medium=github)
* [系统设计-从面试者角度设计一个系统设计题](http://www.wangtianyi.top/blog/2018/08/31/xi-tong-she-ji-mian-shi-ti-zong-he-kao-cha-mian-shi-zhe-de-da-zhao/?utm_source=github&utm_medium=github)
### 智力题
* [概率p输出1，概率1-p输出0，等概率输出0和1](https://blog.csdn.net/qq_29108585/article/details/60765640)
* [判断点是否在多边形内部](https://www.cnblogs.com/muyefeiwu/p/11260366.html)

欢迎光临[我的博客](http://www.wangtianyi.top/?utm_source=github&utm_medium=github)，发现更多技术资源~
