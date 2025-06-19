🔧 自定义 Spring Boot Starter 步骤与原理详解
🛠️ 一、自定义 Starter 实现步骤（遵循官方规范）

命名规范‌
采用 xxx-spring-boot-starter 格式（如 my-spring-boot-starter），区别于官方 spring-boot-starter-xxx。

创建 Maven 项目‌
添加核心依赖：

xml
Copy Code
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter</artifactId>
</dependency>
<!-- 配置元数据支持 -->
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-configuration-processor</artifactId>
    <optional>true</optional>
</dependency>
```:ml-citation{ref="1,3" data="citationList"}



定义配置属性类‌
通过 @ConfigurationProperties 绑定配置文件：

java
Copy Code
@ConfigurationProperties(prefix = "demo")
public class DemoProperties {
    private String sayWhat;
    private String toWho;
    // getters/setters
}
```:ml-citation{ref="3" data="citationList"}



编写业务服务‌
实现核心功能逻辑：

java
Copy Code
public class DemoService {
    private final String sayWhat;
    private final String toWho;
    
    public String say() {
        return sayWhat + "! " + toWho;
    }
}
```:ml-citation{ref="3" data="citationList"}



创建自动配置类‌
集成配置与业务逻辑：

java
Copy Code
@Configuration
@EnableConfigurationProperties(DemoProperties.class)
@ConditionalOnClass(DemoService.class)
public class DemoAutoConfiguration {
    
    @Bean
    @ConditionalOnMissingBean
    public DemoService demoService(DemoProperties properties) {
        return new DemoService(properties.getSayWhat(), properties.getToWho());
    }
}
```:ml-citation{ref="2,4" data="citationList"}



注册自动配置‌
在 resources/META-INF/spring.factories 添加：

properties
Copy Code
org.springframework.boot.autoconfigure.EnableAutoConfiguration=\
  com.example.DemoAutoConfiguration


新版本支持 META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports 文件（一行一个配置类）。

⚙️ 二、Starter 被依赖使用的工作原理

启动触发‌
Spring Boot 应用启动时，@SpringBootApplication 组合了 @EnableAutoConfiguration，激活自动配置机制。

自动配置加载‌
AutoConfigurationImportSelector 扫描所有依赖包中的：

META-INF/spring.factories
META-INF/spring/org.springframework.boot.autoconfigure.AutoConfiguration.imports
加载声明的配置类。

条件装配‌
配置类中的 @ConditionalOnXXX 注解（如 @ConditionalOnClass）动态判断：

类路径是否存在所需类
容器是否缺少目标 Bean
配置属性是否启用
满足条件才注册 Bean。

属性绑定‌
@EnableConfigurationProperties + @ConfigurationProperties 将 application.yml 中的配置（如 demo.sayWhat=Hello）注入属性对象。

服务注入‌
最终将业务 Bean（如 DemoService）自动装配到 Spring 容器，用户可直接 @Autowired 使用。

📦 三、最佳实践建议

模块化拆分‌

核心逻辑与自动配置分离，避免强耦合
Starter 项目仅包含配置类和少量必要代码

条件注解精准控制‌
合理使用组合条件注解，避免意外装配：

java
Copy Code
@ConditionalOnProperty(prefix = "demo", name = "enabled", havingValue = "true")
@ConditionalOnWebApplication


配置元数据提示‌
在 resources/META-INF 添加 additional-spring-configuration-metadata.json，提供 IDE 配置提示：

json
Copy Code
{
  "properties": [{
    "name": "demo.say-what",
    "type": "java.lang.String",
    "description": "设置问候语"
  }]
}
```:ml-citation{ref="1" data="citationList"}



💎 ‌总结‌：自定义 Starter 通过约定优于配置+条件装配，实现“开箱即用”。其核心是将配置、服务、装配逻辑封装为独立模块，使用者只需添加依赖即可自动注入服务。
