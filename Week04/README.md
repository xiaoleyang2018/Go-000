学习笔记
# 工程项目结构
## Layout
- 初始化的工具，自动生成项目，用工具约束
- /cmd 里放 项目文件夹（与可执行文件名字相同）main.go的文件
- /cmd 里不放太多代码，不对外 。只服务启动，关闭，配置，初始化，日志，listerhttp，监听信号
- /pkg 对外重用  /internal 不对外
- 建立公司统一的 kit project
- Package Oriented Design (必读) kit 工程不依赖第三方
  “To this end, the Kit project is not allowed to have a vendor folder. If any of packages are dependent on 3rd party packages, they must always build against the latest version of those dependences.”
   kit 特点： 统一 ；标准库方式布局 ；高度抽象 ；支持插件 
- 服务树，应用名称和相关管理人的名字（类似与DNS）
- 微服务中app类型分四类：interface（对外） ， service（对内）， admin（对运营侧）， job（kafka等）， task（定时任务）
- 依赖倒置？依赖注入？
- 三层架构 + 贫血模型
- DO领域对象
- DDD核心有领域层，做业务逻辑
- kratos的目录结构
```
|- api                  # api目录是对外保留的proto文件及生成的pb.go文件
|- cmd		        # 项目主干，main所在
|   |-- myapp
|      |--- main.go
|- configs		 # configs 为配置文件目录
| --db.toml					
|- internal              # 项目内部包
|   |--dao               # dao层，用户数据库、cache、MQ等资源访问
|   |--di	         # 依赖注入层 采用wire静态分析依赖
|      |--- wire.go      # wire 声明
|      |--- wire_gen.go  # go generate 生成的代码
|   |--model		 # model 层，用于声明业务结构体
|   |--server            # server层，用于初始化grpc和http serverå
|   |--service           # service层，用于业务逻辑处理
|- test                  # 测试资源层
```
# API 设计
# 配置管理
# 包管理
# 测试
- 任意顺序，并行