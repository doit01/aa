有索引的时候就是行锁，没有索引的时候就是表索

间隙锁
记录锁：即锁住记录本身

2、间隙锁：锁住一段没有记录的间隙，可以是两条记录的中间部分，也可以是第一条记录的前置部分或最后一条记录的后续部分

    2.1 需要注意的是，间隙锁仅仅是阻塞对该间隙进行插入操作，而不阻塞对该间隙的查询操作；如有表tab，有索引idx，对(5,10)加入间隙锁，执行下面两条命令：

    insert into tab (idx) values(7); 阻塞

    select * from tab where idx = 7 for update; 不阻塞但无数据

3、next-key锁：是记录锁与间隙锁的结合，特点为左开右闭，如(4,10]，由间隙锁(4,10)和记录锁idx = 10组成
———————————————— 
普通索引的加锁规则
若查询记录不存在，则next-key锁退化为间隙锁，锁住记录所在空隙，此处空隙是指与所查询记录（如果存在的话）相邻的两条记录之间的空隙，这两条记录除了可以是数据库表中实际存在的，也可以是其他记录拟插入的
范围查询：与唯一索引的区别在于：
1、>=会将另一侧的间隙也一起锁住；
2、<会将锁住第一条不满足条件的记录
  如下表：表名为foo，uid为主键索引，也即唯一索引，idx为普通索引
  +-----+-----+------+
  | uid | age | idx  |
  +-----+-----+------+
  |   1 |   1 |    1 |
  |   3 |   3 |    3 |
  |   4 |   4 |    4 |
  |  10 |  10 |   10 |
  |  16 |  16 |   16 |

如事务A执行：select * from foo where idx>= 10 and idx < 12 for update;
其中idx为普通索引，按照上表，(4,10)之间会加上间隙锁，且16会被锁住，也即[10,16]锁住
————————————————
行锁的加锁技巧

事务加锁时，其他事务进行写操作时会受到阻塞，这个阻塞时间当然是越短越好，那么对于一个事务当中不同需要加锁的语句，可以采用以下方式控制：

    对语句所需要锁住的记录条数进行预估，在不影响业务的情况下，将锁住记录多的语句排在锁住记录少的后面，锁粒度大的容易发生冲突，这样安排可以减少与其他事务冲突时间
    将热点记录的加锁操作排在事务后面执行
    批量操作分几次进行，减少锁冲突的概率和时间
————————————————

                            版权声明：本文为博主原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接和本声明。
                        
原文链接：https://blog.csdn.net/m0_54864585/article/details/126076199

                            版权声明：本文为博主原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接和本声明。
                        
原文链接：https://blog.csdn.net/m0_54864585/article/details/126076199


springboot 启动加载过程：
@EnableAutoConfiguaration 通过AutoConfigurationImportSelector，SpringFactoriesLoader搜集配置文件中的配置工厂类（定义在jar包的spring.factories下），通过classloader加载组件工厂类，组件工厂类实例化bean，如jdbcFatorybean，工厂类创建出datasource，redis，filter之类的bean。

SpringFactoriesLoader为Spring工厂加载器，该对象提供了loadFactoryNames方法，入参为factoryClass和classLoader即需要传入工厂类名称和对应的类加载器，方法会根据指定的classLoader，加载该类加器搜索路径下的指定文件，即spring.factories文件，传入的工厂类为接口，而文件中对应的类则是接口的实现类，或最终作为实现类。



@ConditionalOnClass({ SqlSessionFactory.class, SqlSessionFactoryBean.class})这个注解的意思是：当存在SqlSessionFactory.class, SqlSessionFactoryBean.class这两个类时才解析MybatisAutoConfiguration配置类,否则不解析这一个配置类。我们需要mybatis为我们返回会话对象，就必须有会话工厂相关类


@CondtionalOnBean(DataSource.class):只有处理已经被声明为bean的dataSource


@ConditionalOnMissingBean(MapperFactoryBean.class)这个注解的意思是如果容器中不存在name指定的bean则创建bean注入，否则不执行

作者：祖大帅
链接：https://juejin.cn/post/6844903652201594887
来源：稀土掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

devops for linux
https://github.com/1Panel-dev/1Panel

以下是Tomcat中max-connections与server.tomcat.threads.max（即max-threads）的核心区别与职责解析：

1. max-connections（最大连接数）‌
定义‌：控制Tomcat‌同时建立的TCP连接总数上限‌，包括活跃请求和空闲连接。
触发场景‌：
当连接数达到此阈值，新连接将被‌立即拒绝‌（而非排队）。
默认值：
NIO模式‌：10,000（Tomcat 8+）
BIO模式：与max-threads值相同。
类比‌：类似电话总机的‌最大接入线路数‌，超出后直接占线。
2. server.tomcat.threads.max（最大工作线程数）‌
定义‌：控制‌同时处理请求的线程数量‌，即线程池大小。
触发场景‌：
所有线程忙碌时，新请求进入accept-count队列等待（若队列未满）。
默认值：200（根据Tomcat版本可能动态调整）。
性能影响‌：
增加线程数可提升并发处理能力，但过度增加会导致‌CPU频繁切换开销‌。
调优建议‌：通常设为 CPU核心数 × 200 左右（I/O密集型场景）。
关键区别总结‌
参数‌	作用层级	拒绝请求的触发条件	默认值
max-connections	‌TCP连接层‌	连接总数超限	NIO：10,000
server.tomcat.threads.max	‌请求处理层‌	线程全忙‌且‌等待队列满	200

📌 ‌调优注意事项‌：

高并发场景‌：优先调整max-threads（受限于CPU和内存）；
长连接服务‌：需关注max-connections防止连接耗尽；
队列缓冲作用依赖accept-count参数，需与max-threads协同配置。



多线程环境下实现事务回滚
手动事务管理（推荐方案） 

@Autowired
private DataSource dataSource;

public void multiThreadTransaction() throws Exception {
    Connection conn = dataSource.getConnection();
    conn.setAutoCommit(false); // 关闭自动提交
    ExecutorService executor = Executors.newFixedThreadPool(4);
    List<Future<Boolean>> futures = new ArrayList<>();
    
    try {
        // 提交子线程任务
        for (int i = 0; i < 10; i++) {
            futures.add(executor.submit(() -> {
                try (Connection innerConn = dataSource.getConnection()) {
                    innerConn.setAutoCommit(false);
                    // 执行数据库操作
                    return true; // 操作成功
                } catch (Exception e) {
                    return false; // 标记失败
                }
            }));
        }
        
        // 检查所有子线程结果
        for (Future<Boolean> future : futures) {
            if (!future.get()) { // 任一子线程失败
                conn.rollback(); // 主线程回滚
                return;
            }
        }
        conn.commit(); // 全部成功则提交
    } catch (Exception e) {
        conn.rollback();
    } finally {
        conn.close();
        executor.shutdown();
    }
}




ROW_NUMBER

    ‌功能‌：ROW_NUMBER函数为查询结果集中的每一行分配一个唯一的连续序号。即使排序字段值相同，ROW_NUMBER也会赋予不同的序号‌12。
    ‌适用场景‌：适用于需要严格唯一序号的场景，如分页查询、需要绝对唯一标识的情况‌12。

RANK

    ‌功能‌：RANK函数为结果集中的每一行分配一个排名。当遇到相同的排序字段值时，RANK会赋予相同的排名，并且后续排名会跳跃，留出空缺位置‌12。
    ‌适用场景‌：适用于允许并列排名的场景，如体育赛事排名等，允许排名不连续的情
 WITH DepartmentSales AS (
    SELECT 
        e.department_id,
        s.employee_id,
        s.amount,
        RANK() OVER (PARTITION BY e.department_id ORDER BY s.amount DESC) as rank
    FROM 
        sales s
    JOIN 
        employees e ON s.employee_id = e.employee_id
)
SELECT 
    e.department_id,
    e.name,
    ds.amount
FROM 
    DepartmentSales ds
JOIN 
    employees e ON ds.employee_id = e.employee_id
WHERE 
    ds.rank = 1;



    

    
tail -n 10 filename.txt | head -n 5
awk 'NR>=5 && NR<=10' xx.log
sed -n '5,10p' xx.log

查找字符串数组中的最长公共前缀
横向扫描的基本思路是，将数组的第一个字符串作为基准字符串，然后与数组中的其他字符串逐个字符进行比较，直到找到不匹配的字符或达到最短字符串的末尾
public class Solution {
    public String longestCommonPrefix(String[] strs) {
        if (strs == null || strs.length == 0) return "";
        
        // 以第一个字符串作为基准
        String prefix = strs[0];
        
        for (int i = 1; i < strs.length; i++) {
            while (strs[i].indexOf(prefix) != 0) {
                // 如果当前字符串不以prefix开始，则缩短prefix
                prefix = prefix.substring(0, prefix.length() - 1);
                // 如果prefix为空，则直接返回""
                if (prefix.isEmpty()) return "";
            }
        }
        return prefix;
    }
}



用synchronized修饰方法可以把整个方法变为同步代码块，synchronized方法加锁对象是this；
通过合理的设计和数据封装可以让一个类变为“线程安全”；一个类没有特殊说明，默认不是thread-safe；


thread local
ThreadLocal表示线程的“局部变量”，它确保每个线程的ThreadLocal变量都是各自独立的；
给每个方法增加一个context参数非常麻烦，而且有些时候，如果调用链有无法修改源码的第三方库，User对象就传不进去了。

Java标准库提供了一个特殊的ThreadLocal，它可以在一个线程中传递同一个对象

ThreadLocal适合在一个线程的处理流程中保持上下文（避免了同一参数在所有方法中传递）；

使用ThreadLocal要用try ... finally结构，并在finally中清除
https://liaoxuefeng.com/books/java/threading/thread-local/index.html
CAP理论‌
定义‌：分布式系统无法同时满足‌一致性（C）‌、‌可用性（A）‌、‌分区容错性（P）‌，需在CA/CP/AP中取舍，通常优先保证P‌。
应用场景‌：
CP系统‌：如ZooKeeper，强一致但可能牺牲可用性‌。
AP系统‌：如Eureka，高可用但可能数据短暂不一致‌。

BASE理论‌
核心思想‌：通过‌基本可用（BA）‌、‌软状态（S）‌、‌最终一致性（E）‌实现高可用性，适用于弱一致性场景（如电商库存扣减）‌。 
BASE应用‌：
    ‌基本可用‌：即使支付网关短暂故障，用户仍可提交订单，后续异步处理支付状态‌57。
    ‌软状态‌：允许订单在支付完成前保持“支付中”状态，避免因强一致性阻塞流程‌57。
    ‌最终一致性‌：支付结果通过消息队列异步通知，确保数据最终一致（如支付宝/微信支付回调）‌
场景‌：广告服务故障时，需保证核心功能可用。
‌BASE应用‌：
    ‌基本可用‌：广告模块故障时，页面仍展示商品列表，仅隐藏广告位（如淘宝首页降级策略）‌35。
    ‌最终一致性‌：广告数据异步恢复后，通过定时任务重新加载到缓存‌    
秒杀活动预扣库存与最终一致性实现方案‌

在高并发秒杀场景中，库存扣减需兼顾性能与数据准确性。以下是基于预扣库存模式的典型实现流程及技术要点：

一、预扣库存核心流程‌
请求拦截与预扣库存‌
Redis原子预扣‌：用户请求到达时，通过Lua脚本在Redis中执行原子操作扣减库存，并记录预扣流水（如用户ID、商品ID、预扣时间）‌。
设置预扣有效期‌：预扣库存标记为“已锁定”状态，并设置TTL（如15分钟），超时未支付则自动释放库存‌。

异步处理订单‌
消息队列异步下单‌：预扣成功后，通过RocketMQ发送异步消息创建订单，避免直接操作数据库导致性能瓶颈‌。
延迟消息回滚‌：订单创建失败时，触发RocketMQ延迟消息（如1分钟重试），若多次重试仍失败则回滚Redis库存‌。

库存最终确认‌
数据库事务确认‌：订单支付成功后，通过数据库事务扣减实际库存，并清理Redis中的预扣记录‌。
异常补偿机制‌：若数据库扣减失败，通过定时任务或消息队列补偿Redis库存，确保数据最终一致‌。
二、关键一致性保障技术‌
Redis与数据库数据同步‌
对账表设计‌：引入对账表记录所有预扣流水，定时扫描Redis预扣记录与数据库实际库存差异，触发自动补偿‌。
最终一致性表‌：异步任务定期比对Redis与数据库库存，修复不一致数据（如Redis释放但数据库未更新）‌。

异常场景处理‌
服务宕机恢复‌：通过Redis预扣流水日志恢复中断操作，结合消息队列重试机制保证流程连续性‌。
网络分区容错‌：采用TCC模式（Try-Confirm-Cancel）实现柔性事务，预扣失败时通过Cancel阶段回滚库存‌。
三、技术选型与优化‌
组件‌	‌作用‌	‌关键配置‌
Redis‌	实现高并发预扣库存，支持Lua脚本保证原子性	使用Hash结构存储商品库存，ZSet记录预扣流水‌
RocketMQ‌	异步解耦订单创建与库存确认，支持事务消息与延迟重试	配置事务监听器处理本地事务与消息回查‌
Seata‌	分布式事务协调（可选），适用于复杂补偿场景	采用Saga模式编排库存扣减与订单创建流程‌
四、典型异常处理策略‌
预扣成功但订单创建失败‌
自动回滚‌：通过消息消费者监听异常事件，调用Redis API释放预扣库存‌。
人工介入‌：对账表记录异常数据，触发告警并支持人工修正‌。

支付超时释放库存‌
延迟队列监听‌：ZSet按预扣时间排序，定时任务扫描超时记录并释放库存‌。
总结‌
秒杀场景通过‌预扣库存+异步确认‌实现高并发与最终一致性：
性能优先‌：Redis预扣避免数据库直接承压‌。
兜底设计‌：对账表+补偿任务覆盖所有异常场景‌。
平衡选择‌：放弃强一致性换取吞吐量提升，通过技术组合实现业务可接受的数据准确性‌。    
    
二、分布式事务处理‌

2PC（两阶段提交）‌：
阶段1（投票）‌：协调者询问参与者是否可提交；‌阶段2（提交/回滚）‌：根据投票结果执行提交或回滚，存在同步阻塞和单点故障问题‌。

分布式事务解决方案‌
TCC（补偿事务）‌：
Try‌：预留资源；‌Confirm‌：确认提交；‌Cancel‌：失败回滚，需业务代码实现补偿逻辑（如订单状态回滚）‌。
消息队列（MQ）‌：
本地消息表‌：事务与消息发送绑定，通过异步重试保证最终一致性（如订单创建后发送消息扣减库存）‌。
Seata框架‌：支持AT（自动补偿）、TCC、XA等模式，通过全局事务ID管理跨服务事务‌。
三、分布式系统解决方案‌
分布式锁实现‌
Redis实现‌：
使用SET key value NX EX命令加锁，Lua脚本保证原子性解锁，需解决锁续期问题（如Redisson看门狗机制）‌。
ZooKeeper实现‌：通过临时顺序节点监听机制，避免锁失效导致的并发问题‌。
分布式缓存与一致性‌
缓存穿透‌：空值缓存+布隆过滤器拦截非法请求‌。
缓存雪崩‌：随机过期时间+多级缓存（如本地缓存+Redis）‌。
数据一致性‌：采用双写策略+消息队列异步校验（如先更新DB再删除缓存）‌。
四、场景与实战问题‌
如何设计高并发秒杀系统？‌
分层削峰‌：CDN静态资源缓存+网关限流（如令牌桶算法）‌。
异步处理‌：请求队列化（如RabbitMQ）+库存预扣减（Redis原子操作）‌。
最终一致性‌：订单状态异步通知+库存回补机制（防止超卖）‌。

分布式ID生成方案‌

雪花算法（Snowflake）‌：时间戳+机器ID+序列号，保证全局唯一且趋势递增‌。
Redis原子操作‌：INCR命令生成分段ID，需解决持久化问题‌。
五、高频代码示例‌
java
Copy Code
// TCC模式示 例（Try阶段预留资源）
public boolean tryReserveInventory(Long productId, Integer count) {
    // 冻结库存，非实际扣减
    String key = "inventory:freeze:" + productId;
    return redisTemplate.opsForValue().increment(key, count) >= 0;
}




jdk17 default gc is G1 
java -XX:+PrintCommandLineFlags -version
-XX:InitialHeapSize=16777216 -XX:MaxHeapSize=268435456 -XX:MinHeapSize=6815736 -XX:+PrintCommandLineFlags -XX:ReservedCodeCacheSize=251658240 -XX:+SegmentedCodeCache -XX:+UseCompressedClassPointers -XX:+UseCompressedOops -XX:+UseSerialGC 
分区回收‌：G1将堆内存划分为多个Region，根据优先级进行回收，减少全局停顿时间。
‌并发和并行处理‌：G1支持并发标记和并发预处理，减少垃圾回收对应用的影响。

G1 的基本原理与核心设计‌

    ‌Region 分区机制‌
    G1 将堆划分为多个大小相等的 Region（默认 2048 个），逻辑上分为 Eden、Survivor、Old 和 Humongous（大对象区）。通过动态调整回收优先级（标记垃圾最多的 Region 优先回收），实现低停顿目标‌38。
    ‌并发标记与混合回收‌
        ‌并发标记阶段‌：与应用程序并行执行，标记存活对象，避免全堆停顿‌48。
        ‌混合回收（Mixed GC）‌：回收年轻代 Region 和部分老年代 Region，避免老年代完全回收（非 Full GC）‌38。

‌2. 如何实现可预测的停顿时间？‌

    ‌停顿时间模型‌：通过 -XX:MaxGCPauseMillis 设定目标停顿时间（默认 200ms），G1 根据历史回收数据动态调整 Region 回收数量和顺序，优先处理垃圾比例高的 Region‌38。
    ‌增量回收‌：将回收任务拆分为多个小阶段（如初始标记、并发标记、重新标记、清理），分批次完成，避免单次长时间停顿‌48。

‌3. Region 设计的优势‌

    ‌内存利用率高‌：支持动态分配 Region 类型（如 Eden→Survivor→Old），避免传统分代模型的内存浪费‌38。
    ‌并行与并发优化‌：多个回收线程可同时处理不同 Region，降低线程竞争，提升吞吐量‌38。

‌4. Mixed GC 的触发条件与流程‌

    ‌触发阈值‌：当老年代占用堆比例达到 -XX:InitiatingHeapOccupancyPercent（默认 45%）时触发‌48。
    ‌执行流程‌：
        ‌初始标记（STW）‌：标记根对象，伴随一次年轻代 GC‌8。
        ‌并发标记‌：并行标记存活对象‌8。
        ‌最终标记（STW）‌：修正并发标记期间变动的对象引用‌8。
        ‌筛选回收（Evacuation）‌：复制存活对象到空闲 Region，清理垃圾 Region‌48。

‌5. G1 与 CMS 的核心区别‌
‌特性‌ 	‌G1‌ 	‌CMS‌
‌算法‌ 	标记-整理（整体） + 复制（局部） 	标记-清除
‌内存模型‌ 	Region 逻辑分区 	物理分代（年轻代 + 老年代）
‌停顿时间‌ 	可控（可预测模型） 	不可控（依赖堆碎片情况）
‌适用场景‌ 	大堆、低延迟需求 	中小堆、高吞吐需求
‌内存碎片处理‌ 	通过 Region 复制整理减少碎片 	需 Full GC 整理碎片
‌引用‌ 	‌34 	‌34
‌6. 大对象（Humongous）处理机制‌

    ‌定义‌：对象大小超过 Region 50% 则判定为大对象，分配在连续的 Humongous Region‌38。
    ‌回收策略‌：在年轻代 GC 或 Mixed GC 中回收无引用的大对象，避免占用过多连续空间‌38。

‌7. 调优关键参数‌

    ‌目标停顿时间‌：-XX:MaxGCPauseMillis=200（单位：ms）‌38。
    ‌混合回收阈值‌：-XX:InitiatingHeapOccupancyPercent=45（老年代占用堆比例）‌48。
    ‌Region 大小‌：-XX:G1HeapRegionSize=2M（建议为 1MB~32MB，需为 2 的幂）‌48。

‌8. Full GC 触发条件及应对‌

    ‌触发场景‌：
        Mixed GC 回收速度跟不上对象分配速度，导致老年代占满‌8。
        并发标记失败（如堆内存不足）‌8。
    ‌优化方案‌：
        增大堆内存或降低 -XX:InitiatingHeapOccupancyPercent‌8。
        避免频繁大对象分配，减少 Humongous Region 碎片‌38。

‌9. 适用场景与局限性‌

    ‌适用场景‌：
        堆内存 ≥ 6GB，需低延迟（如实时交易系统）‌38。
        对停顿时间敏感的应用（如金融、游戏服务）‌48。
    ‌局限性‌：
        小堆场景性能可能不如 CMS/Parallel GC‌34。
        内存占用较高（需维护 Region 元数据）‌

String.intern() 方法的作用‌
将字符串对象添加到常量池‌：
调用 intern() 方法时，若字符串常量池中已存在内容相同的字符串，则直接返回池中的引用；若不存在，则将该字符串对象添加到池中，并返回池中的引用。 

String s1 = new String("abc");  // 堆中创建新对象
String s2 = s1.intern();        // 返回常量池中的 "abc" 引用
System.out.println(s1 == s2);   // false（s1在堆，s2在池）

String s3 = "abc";              // 直接使用常量池
System.out.println(s2 == s3);   // true（两者均指向池中同一对象）
 内存优化原理‌
    ‌减少重复字符串的内存占用‌：
通过 intern() 可强制将字符串放入常量池，避免重复字符串在堆中创建多个对象
/ 未使用 intern()
String a = new String("hello");  // 堆中对象
String b = new String("hello");  // 堆中另一个对象
System.out.println(a == b);      // false

// 使用 intern()
String c = new String("world").intern();  // 池中对象
String d = "world";                       // 池中同一对象
System.out.println(c == d);               // true
字符串常量池移至 ‌堆内存‌，允许通过垃圾回收管理未引用的常量字符串，减少内存泄漏风险。
适用场景‌

    ‌大量重复字符串处理‌：如解析 CSV/JSON 数据时，对重复字段值调用 intern() 可显著减少内存占用。
    ‌高频字符串比较‌：若需频繁使用 equals() 比较字符串内容，可先调用 intern() 后用 == 比较引用，提升性能（需权衡池化开销）。

注意事项‌

    ‌性能开销‌：intern() 的底层实现依赖哈希表（Java 7+ 使用 ConcurrentHashMap），高并发场景下可能成为瓶颈。
    ‌内存风险‌：过度池化唯一字符串（如 UUID）会导致常量池膨胀，反而增加内存压力。
    ‌版本差异‌：Java 6 的常量池容量固定（默认 1009），易触发 OutOfMemoryError；Java 7+ 支持动态扩展。
核心目的‌：通过字符串常量池复用相同内容的字符串，节省内存。
‌适用场景‌：处理大量重复字符串时（如日志分析、数据解析）。
‌避坑指南‌：避免池化唯一或动态生成的字符串，优先在 Java 7+ 中使用。
‌性能权衡‌：在内存节省与池化开销之间找到平衡点。


类加载的三个阶段‌

    ‌加载‌：通过全限定名获取二进制字节流（Class 文件），将静态结构转换为方法区的运行时数据结构，并在堆中生成 java.lang.Class 对象作为访问入口‌36。
    ‌验证‌：检查字节码是否符合 JVM 规范（如魔数、版本号、常量池合法性），确保无安全漏洞‌36。
    ‌准备‌：为类变量（static 变量）分配内存并设置初始值（如 int 初始化为 0，final 直接赋常量池值）‌35。

‌连接与初始化‌

    ‌解析‌：将符号引用（如类名、方法名）转换为直接引用（内存地址）‌78。
    ‌初始化‌：执行 <clinit> 方法（静态变量赋值、静态代码块），是类加载的最后一步‌

类加载器与双亲委派模型‌

    ‌类加载器分类‌
        ‌启动类加载器（Bootstrap ClassLoader）‌：加载 JRE/lib 下的核心类库（如 rt.jar）‌27。
        ‌扩展类加载器（Extension ClassLoader）‌：加载 JRE/lib/ext 目录的扩展类‌26。
        ‌应用类加载器（Application ClassLoader）‌：加载用户类路径（ClassPath）的类‌26。
        ‌自定义类加载器‌：继承 ClassLoader，重写 findClass() 实现动态加载（如热部署、加密解密）‌27。

    ‌双亲委派机制‌
        ‌流程‌：子类加载器收到请求后，优先委派父类加载器处理，父类无法完成时才由子类加载‌26。
        ‌优点‌：避免类重复加载，保护核心类库安全（如防止自定义 java.lang.String 覆盖 JVM 实现
必须立即初始化的 5 种情况‌

    使用 new 实例化对象、读取/设置类的静态字段、调用类的静态方法‌67。
    反射调用类（如 Class.forName()）且类未初始化‌67。
    初始化子类时发现父类未初始化（先触发父类初始化）‌67。
    JVM 启动时指定的主类（包含 main() 方法的类）‌78。
    JDK 动态语言支持（如 Lambda 表达式涉及的类）‌

如何打破双亲委派模型？‌
    ‌场景‌：Tomcat 为隔离不同 Web 应用的类，每个应用使用独立类加载器‌27。
    ‌方法‌：重写 loadClass() 方法，直接加载特定类（如 SPI 服务发现）‌
类卸载条件‌

    类的所有实例已被回收。
    类的 Class 对象未被引用。
    加载该类的类加载器已被回收‌
JVM 如何加载动态生成的类？‌
    通过 ByteArrayOutputStream 生成字节码，调用 defineClass() 方法动态加载‌
热部署实现原理‌
    自定义类加载器加载修改后的类，旧类无引用后由垃圾回收器回收，新类生效‌



‌一、JMM 核心概念‌

    ‌内存划分与交互规则‌
        ‌主内存‌：存储所有共享变量（如类静态变量、实例对象），所有线程均可访问‌15。
        ‌工作内存‌：线程私有，存储主内存变量的副本，线程操作变量需先拷贝至工作内存，修改后同步回主内存‌15。
        ‌交互操作‌：read（从主内存读取）、load（加载到工作内存）、use（使用）、assign（赋值）、store（存储回主内存）、write（主内存更新）‌58。

    ‌三大特性‌
        ‌原子性‌：基本类型（int、boolean 等）的读写操作不可分割（long/double 在 32 位系统中非原子）‌58。
        ‌可见性‌：volatile 保证变量修改后立即刷新到主内存，其他线程可见；synchronized 和 final 也能实现可见性‌15。
        ‌有序性‌：volatile 禁止指令重排序，synchronized 通过锁保证代码块串行执行‌15。

    ‌先行发生原则（Happens-Before）‌
        程序顺序规则、锁规则（解锁先于加锁）、volatile 规则（写先于读）、线程启动规则（start() 先于线程代码）、传递性规则等‌

状态转换‌：新建（NEW）→ 就绪（RUNNABLE）→ 运行（RUNNING）→ 阻塞（BLOCKED/WAITING）→ 终止（TERMINATED）‌36。
‌阻塞场景‌：等待锁（synchronized）、Object.wait()、Thread.sleep()、IO 操作等‌37。

同步机制‌

    ‌synchronized‌：基于对象监视器锁（Monitor），修饰代码块或方法，保证原子性和可见性‌57。
    ‌ReentrantLock‌：可中断、支持公平锁、绑定多个条件变量（Condition），需手动释放锁‌78。
    ‌volatile‌：仅保证可见性和有序性，不保证复合操作原子性（
原子类与 CAS‌
    ‌AtomicInteger 等‌：基于 CAS（Compare-And-Swap）实现无锁原子操作，避免线程阻塞‌57。
    ‌CAS 问题‌：ABA 问题（可通过版本号解决）、自旋开销‌
避免死锁策略‌
    按固定顺序获取锁、设置锁超时（tryLock）、死锁检测（如 JStack 分析

线程不安全场景‌

    ‌复合操作‌：非原子操作（如 HashMap 并发扩容）导致数据丢失或覆盖‌47。
    ‌对象逃逸‌：未正确同步的共享对象被多线程修改（如未加锁的 ArrayList 并发添加元素
ReentrantLock 对比 synchronized‌
‌特性‌ 	synchronized 	ReentrantLock
‌锁获取方式‌ 	JVM 隐式管理 	需手动 lock() 和 unlock()‌78
‌公平性‌ 	非公平锁（默认） 	支持公平锁与非公平锁‌27
‌条件变量‌ 	不支持 	       支持多个 Condition‌78
‌可中断性‌ 	不可中断 	支持 lockInterruptibly()

    减少锁竞争‌
        ‌缩小锁粒度‌：使用分段锁（如 ConcurrentHashMap）或细粒度锁（如只锁共享变量）‌27。
        ‌无锁编程‌：基于 CAS 的原子类（如 AtomicInteger）实现线程安全‌78。

    ‌锁消除与锁粗化‌
        ‌锁消除‌：JIT 编译器对不可能存在竞争的锁进行消除（如局部对象锁）‌27。
        ‌锁粗化‌：合并多个相邻锁操作，减少频繁加锁/解锁的开销‌27。

    ‌避免死锁‌
        ‌固定顺序加锁‌：按全局统一顺序获取多把锁（如按 hash 值排序）‌78。
        ‌超时机制‌：使用 tryLock(timeout) 避免无限等待（如 ReentrantLock 支持）‌78。

‌四、高频进阶问题‌

    ‌如何实现线程安全的单例模式？‌
        ‌双重检查锁（DCL）‌：结合 volatile 禁止指令重排序，保证单例唯一性‌27。

        javaCopy Code
        public class Singleton {
            private volatile static Singleton instance;
            public static Singleton getInstance() {
                if (instance == null) {
                    synchronized (Singleton.class) {
                        if (instance == null) {
                            instance = new Singleton();
                        }
                    }
                }
                return instance;
            }
        }

    ‌CAS 的 ABA 问题如何解决？‌
        ‌版本号机制‌：使用 AtomicStampedReference 记录变量修改版本号，避免值被其他线程修改后恢复原值‌78。

    ‌线程池如何避免资源耗尽？‌
        ‌参数配置‌：合理设置核心线程数、最大线程数、队列容量及拒绝策略（如 ThreadPoolExecutor.AbortPolicy）‌57。

‌五、实战调优案例‌

    ‌高并发计数器优化‌
        ‌场景‌：多线程频繁累加计数器导致性能瓶颈。
        ‌优化‌：使用 LongAdder 替代 AtomicLong，通过分段累加减少 CAS 竞争‌78。

    ‌死锁检测与排查‌
        ‌工具‌：通过 jstack 导出线程堆栈，分析线程阻塞链；或使用 Arthas 在线诊断工具定位死锁‌78。

‌总结‌

线程安全与锁优化的核心在于 ‌平衡性能与安全‌，需重点掌握：

    synchronized 锁升级机制及适用场景‌27。
    ReentrantLock 的灵活特性（如公平锁、条件变量）‌78。
    锁优化策略（如无锁编程、锁消除/粗化）‌27。
    面试中可结合源码（如 AbstractQueuedSynchronizer 实现）和实际案例（如高并发计数器、死锁排查）深入阐述。

要解决单例 Bean 依赖原型 Bean 时原型 Bean 生命周期被破坏的问题（即单例 Bean 始终使用同一个原型实例）
单例 Bean 初始化时仅注入一次原型 Bean，导致后续操作始终复用同一实例，违背原型 Bean ‌每次获取新对象‌ 的设计初衷
Jakarta Provider 接口‌（推荐）
import jakarta.inject.Provider;

public class SingletonBean {
    @Resource
    private Provider<PrototypeBean> prototypeBeanProvider; // 注入Provider

    public void execute() {
        PrototypeBean prototype = prototypeBeanProvider.get(); // 每次get()创建新实例
        prototype.doSomething();
    }
}


Spring 事务传播机制中 REQUIRES_NEW 的实现原理‌
    创建新事务时，暂停当前事务（通过 TransactionSynchronizationManager 解绑资源），新事务提交后恢复原事务‌
    
Metaspace 内存泄漏的常见原因及排查方法‌
    ‌原因‌：动态生成类未卸载（如大量使用反射或 CGLIB 代理）‌35。
    ‌排查‌：使用 jcmd <pid> GC.class_stats 统计类加载信息，或通过 MAT 分析堆转储‌58

CAP 理论中 BASE 理论的应用场景‌
    ‌核心‌：基本可用（Basically Available）、软状态（Soft State）、最终一致性（Eventual Consistency）。
    ‌场景‌：电商库存允许超卖后补偿，优先保证系统可用性‌78。
    如：通知用户：通过邮件、短信等方式通知用户订单无法处理，并提供解决方案，如重新下单或退款。


    
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

只有被托管的对象才可以被refresh
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
