# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.18.3] - 2024-01-22
### Changed
- 修复设置偏移量死锁的bug.

## [1.18.2] - 2024-01-22
### Changed
- RollingOnce 新增按照小时和分钟设置偏移量.
- Update changelog.

## [1.18.1] - 2024-01-22
### Changed
- 这两都提示推荐使用RollingOnce.
- 优化RollingOnce, 增加可重置偏移量的方法.

## [1.18.0] - 2024-01-20
### Changed
- 删除部分废弃的代码.
- 注释掉部分告警日志.
- 补充Pool接口注释.
- 优化连接池Pool部分代码.

## [1.17.9] - 2024-01-18
### Changed
- 增加Touch函数.

## [1.17.8] - 2024-01-17
### Changed
- Bar公开方法增加注释.
- Bar结构增加部分注释.

## [1.17.7] - 2024-01-17
### Changed
- 修复未初始化的bug.
- 修复未初始化的bug.

## [1.17.6] - 2024-01-17
### Changed
- 进度条增加结束等待方法.

## [1.17.5] - 2024-01-14
### Changed
- 优化可以忽略panic的异常捕获函数.

## [1.17.4] - 2024-01-14
### Changed
- 调整解析参数的时机, 只在异常发生时解析.

## [1.17.3] - 2024-01-14
### Changed
- 允许CatchPanic传入可变参数.

## [1.17.2] - 2024-01-13
### Changed
- 更新依赖库版本.

## [1.17.1] - 2024-01-11
### Changed
- 删除废弃的代码.

## [1.17.0] - 2024-01-11
### Changed
- 修订Changelog.
- 删除废弃的代码.
- Update changelog.

## [1.16.9] - 2024-01-09
### Changed
- Update changelog.
- 优化部分代码, 删除非必要的import.

## [1.16.8] - 2024-01-08
### Changed
- Update changelog.
- 优化部分代码.
- 修订注释.
- 新增RollingOnce, 相对于PeriodicOnce优化了时间戳的处理方式.
- 优化timestamp时间戳.
- 调整部分代码.
- 新增时间戳(毫秒数)的功能函数.
- 更新依赖库golang.org/x/exp版本.
- 修订ParseTime的函数注释.

## [1.16.7] - 2024-01-02
### Changed
- Update changelog.
- 调整util包中的uuid能到pkg.
- 调整util包中的时间格式功能到api包.
- 增加功能性函数, 捕获panic, 忽略异常, 继续执行.
- 优化应用程序名的获取方式, 改成懒加载.

## [1.16.6] - 2023-12-31
### Changed
- Update changelog.
- 更新依赖库pkg版本.
- 更新依赖库pkg版本.
- 更新依赖库pkg版本.
- 调整因删除errors包对其它功能的影响.
- 删除errors包.
- 迁移: github.com/smartystreets-prototypes/go-disruptor@v0.0.0-20231024205940-61200af675a1 到 pkg.
- 迁移: github.com/valyala/fastjson@v1.6.4 到 pkg.

## [1.16.5] - 2023-12-30
### Changed
- Update changelog.
- 调整缓存csv逻辑, 如果切片为空, 直接返回.

## [1.16.4] - 2023-12-30
### Changed
- Update changelog.
- 调整文件名, 字面意义和功能保持一致.

## [1.16.3] - 2023-12-30
### Changed
- Update changelog.
- 修复num.Decimals函数负浮点四舍五入错误的bug.

## [1.16.2] - 2023-12-23
### Changed
- Update changelog.
- 更新依赖库版本.
- 更新依赖库版本.
- 更新依赖库版本.

## [1.16.1] - 2023-12-23
### Changed
- Update changelog.
- 修复windows变量引用的错误.
- 优化filestat中时间的处理方法.
- 从exp包中复制maps的Keys和Values函数, 意在移除对exp的依赖.

## [1.16.0] - 2023-12-23
### Changed
- Update changelog.
- 优化部分代码.

## [1.15.9] - 2023-12-22
### Changed
- Update changelog.
- 修复windows缺少error返回值的bug.
- 新增获取文件状态(时间)的函数.
- 增加windows操作系统的文件时间戳的获取方法.

## [1.15.8] - 2023-12-16
### Changed
- Update changelog.
- PeriodOnce默认在非runtime的debug模式下不输出日志.

## [1.15.7] - 2023-12-15
### Changed
- Update changelog.
- 定时任务默认在非runtime的debug模式下不输出日志.

## [1.15.6] - 2023-12-14
### Changed
- Update changelog.
- 优化去重处理方式.
- 给获取应用程序文件名的函数增加注释.

## [1.15.5] - 2023-12-12
### Changed
- Update changelog.
- 增加文件路径中的日期格式, 数据中的日期格式.
- 删除旧版本的滑动Once功能.

## [1.15.4] - 2023-12-12
### Changed
- Update changelog.
- 更新pkg版本同步go版本.

## [1.15.3] - 2023-12-12
### Changed
- Update changelog.
- 更新go1.21.5.

## [1.15.2] - 2023-12-07
### Changed
- Update changelog.
- 更新依赖库pkg版本.

## [1.15.1] - 2023-12-05
### Changed
- Update changelog.
- 更新依赖库版本.

## [1.15.0] - 2023-12-04
### Changed
- Update changelog.
- Post方法去掉返回值中的lastModified.

## [1.14.9] - 2023-12-04
### Changed
- Update changelog.
- 修复json判断失败的bug.

## [1.14.8] - 2023-12-04
### Changed
- Update changelog.
- 增加一个容错机制的HttpPost函数.

## [1.14.7] - 2023-12-04
### Changed
- Update changelog.
- 优化部分功能函数.

## [1.14.6] - 2023-12-04
### Changed
- Update changelog.
- HttpRequest函数增加可以传入header.
- Http新增一个独立的Get方法, 允许传入header.

## [1.14.5] - 2023-12-04
### Changed
- Update changelog.
- Http工具包增加post方法.

## [1.14.4] - 2023-12-04
### Changed
- Update changelog.
- 运行时允许重置debug状态.

## [1.14.3] - 2023-12-03
### Changed
- Update changelog.
- 增加自旋锁.

## [1.14.2] - 2023-12-03
### Changed
- Update changelog.
- 增加具有滑动窗口功能的WaitGroup.

## [1.14.1] - 2023-12-03
### Changed
- Update changelog.
- 迁移gocsv从github.com/gocarina/gocsv到gitee.com/quant1x/pkg/gocsv.

## [1.14.0] - 2023-11-27
### Changed
- Update changelog.
- 移除终端二维码工具库到pkg.

## [1.13.9] - 2023-11-26
### Changed
- Update changelog.
- 更新依赖库版本.

## [1.13.8] - 2023-11-23
### Changed
- 增加其它运行时需要的工具函数.
- 增加编译选项中调试开关.
- Update changelog.
- 增加编译选项中调试开关.
- 实验多时段定时任务.

## [1.13.7] - 2023-10-28
### Changed
- Update changelog.
- Logger缓存map改为sync.map.

## [1.13.6] - 2023-10-28
### Changed
- Update changelog.
- Treemap的clear方法增加互斥锁.

## [1.13.5] - 2023-10-27
### Changed
- Update changelog.
- 调整csv文件的关闭方式.

## [1.13.4] - 2023-10-26
### Changed
- Update changelog.
- 定时任务增加重置日志.

## [1.13.3] - 2023-10-22
### Changed
- Update changelog.
- 调整调度任务, 增加计时.

## [1.13.2] - 2023-10-22
### Changed
- Update changelog.
- 调整skip调度任务策略.
- 增加获取func信息的函数.
- 新增runtime包.
- 增加获取当前代码的函数名, 文件名以及行号的函数.

## [1.13.1] - 2023-10-21
### Changed
- Update changelog.
- 增加应用退出等待机制.

## [1.13.0] - 2023-10-20
### Changed
- Update changelog.
- 调整任务未执行完成跳过的函数封装.

## [1.12.9] - 2023-10-20
### Changed
- Update changelog.
- 调整任务未执行完成跳过的函数封装.

## [1.12.8] - 2023-10-19
### Changed
- Update changelog.
- 修订可延迟执行的定时调度组件.

## [1.12.7] - 2023-10-19
### Changed
- Update changelog.
- 修订日期重置错乱的bug.

## [1.12.6] - 2023-10-16
### Changed
- Update changelog.
- 修正周期初始化时间为9点整.

## [1.12.5] - 2023-10-16
### Changed
- Update changelog.
- 调整周期初始化锁.

## [1.12.4] - 2023-10-15
### Changed
- Update changelog.
- 调整函数名.

## [1.12.3] - 2023-10-15
### Changed
- Update changelog.
- 更新依赖版本.
- 调整hashmap的代码, 引入github.com/orcaman/concurrent-map/v2.
- 删除废弃的代码.

## [1.12.2] - 2023-10-15
### Changed
- Update changelog.
- 增加协程安全的hashmap.

## [1.12.1] - 2023-10-10
### Changed
- Update changelog.
- 收录go-runewidth组件.

## [1.12.0] - 2023-10-08
### Changed
- Update changelog.
- 增加not found判断.

## [1.11.9] - 2023-10-07
### Changed
- Update changelog.
- 增加线程安全的TreeMap.
- 优化http client参数.
- 优化http client.

## [1.11.8] - 2023-10-05
### Changed
- Update changelog.
- 新增embed封装函数.

## [1.11.7] - 2023-10-01
### Changed
- Update changelog.
- 删除废弃的代码.

## [1.11.6] - 2023-10-01
### Changed
- Update changelog.
- 优化fastqueue的push为异步方式.

## [1.11.5] - 2023-09-29
### Changed
- Update changelog.
- 优化滑动窗口锁, 窗口期内只初始化一次.
- 增加context.Context的封装.
- 增加注释, 从1.12版本开始将移除MultiOnce.

## [1.11.4] - 2023-09-29
### Changed
- Update changelog.
- 增加滑动窗口式的加载锁.

## [1.11.3] - 2023-09-15
### Changed
- Update changelog.
- 调整homedir, 所有操作系统设置了GOX_HOME都会优先返回.

## [1.11.2] - 2023-09-15
### Changed
- Windows服务安装时创建系统环境变量GOX_HOME.

## [1.11.1] - 2023-09-15
### Changed
-  启用环境变量GOX_HOME是为了Windows服务以系统账户运行时无法获取登录用户的宿主目录而预备的.

## [1.11.0] - 2023-09-15
### Changed
- 去掉多余的import.

## [1.10.9] - 2023-09-15
### Changed
- 屏蔽检测连接池已打开数量的日志.
- Windows服务屏蔽使用本地用户登录, 本地用户登录有一个问题, 密码更换后会造成服务运行不正常.

## [1.10.8] - 2023-09-15
### Changed
- Windows 服务属性增加本地用户名.

## [1.10.7] - 2023-09-13
### Changed
- Update changelog.
- 获取连接增加告警日志.

## [1.10.6] - 2023-09-12
### Changed
- Update changelog.
- 更换golang.org/x/exp/slices为系统标准库.

## [1.10.5] - 2023-09-10
### Changed
- Update changelog.
- 升级依赖库版本.
- 升级string和bytes转换函数.

## [1.10.4] - 2023-09-10
### Changed
- Update changelog.
- 升级string和bytes转换函数.

## [1.10.3] - 2023-08-24
### Changed
- Update changelog.
- 增加linux cpu 序列号获取方式, 用第一块网卡的mac地址代替.

## [1.10.2] - 2023-08-16
### Changed
- Update changelog.
- 修订进度条结束逻辑, 先复写进度条, 再结束.

## [1.10.1] - 2023-08-13
### Changed
- Update changelog.
- 升级go版本到1.21.0.

## [1.10.0] - 2023-08-02
### Changed
- Update changelog.
- 调整服务的运行顺序.

## [1.9.9] - 2023-08-01
### Changed
- Update changelog.
- 服务组件去掉日志.

## [1.9.8] - 2023-08-01
### Changed
- Update changelog.
- 增加日志初始化.

## [1.9.7] - 2023-08-01
### Changed
- Update changelog.
- 增加日志.

## [1.9.6] - 2023-08-01
### Changed
- Update changelog.
- 调整windows服务的运行方式.

## [1.9.5] - 2023-07-21
### Changed
- Update changelog.
- 恢复daemon的demo.
- 删除daemon的demo.

## [1.9.4] - 2023-07-20
### Changed
- Update changelog.
- 新增daemon工具库.

## [1.9.3] - 2023-07-08
### Changed
- Update changelog.
- 更新依赖库版本.

## [1.9.2] - 2023-07-08
### Changed
- Update changelog.
- 修复季度编码的bug.

## [1.9.1] - 2023-07-07
### Changed
- Update changelog.
- 判断float是否NaN.

## [1.9.0] - 2023-07-06
### Changed
- Update changelog.
- 优化slice去重函数.
- 调整获取周、月开始和结束时间的函数.

## [1.8.9] - 2023-07-02
### Changed
- Update changelog.
- 修复文件句柄未关闭的bug.
- 修复文件句柄未关闭的bug.

## [1.8.8] - 2023-06-30
### Changed
- Update changelog.
- 修复死锁的bug.

## [1.8.7] - 2023-06-30
### Changed
- Update changelog.
- 增加默认初始化日期函数.

## [1.8.6] - 2023-06-30
### Changed
- Update changelog.
- 增加日期切换功能, 默认不开启.

## [1.8.5] - 2023-06-29
### Changed
- Update changelog.
- 没必要加锁, 加锁是个多余的操作.

## [1.8.4] - 2023-06-29
### Changed
- Update changelog.
- 重置计数器加锁.

## [1.8.3] - 2023-06-27
### Changed
- Update changelog.
- 修复chanel阻塞的bug.

## [1.8.2] - 2023-06-27
### Changed
- Update changelog.
- 修复死锁的bug.

## [1.8.1] - 2023-06-27
### Changed
- Update changelog.
- 连接池增加关闭所有链接的方法.

## [1.8.0] - 2023-06-27
### Changed
- Update changelog.
- 增加base64算法.

## [1.7.9] - 2023-06-27
### Changed
- Update changelog.
- 调整package.

## [1.7.8] - 2023-06-27
### Changed
- Update changelog.
- 增加CPU序列号的获取函数.

## [1.7.7] - 2023-06-26
### Changed
- Update changelog.
- 新增项目内的消息队列.

## [1.7.6] - 2023-06-24
### Changed
- Update changelog.
- 季度函数增加财报季的返回值.

## [1.7.5] - 2023-06-21
### Changed
- Update changelog.
- 新增一个定时调度任务, 回调函数不会并发执行.

## [1.7.4] - 2023-06-21
### Changed
- Update changelog.
- 增加计算指定日期的季度开始和结束时间函数.

## [1.7.3] - 2023-06-17
### Changed
- Update changelog.
- 修复字符串数据反射结构体存在数组长度和字段数量不匹配导致数据缺失的bug.

## [1.7.2] - 2023-06-16
### Changed
- Update changelog.
- 新增结构体tag的反射缓存.

## [1.7.1] - 2023-06-16
### Changed
- Update changelog.
- 更新依赖库.

## [1.7.0] - 2023-06-14
### Changed
- Update changelog.
- 新增slice range函数.

## [1.6.9] - 2023-06-14
### Changed
- Update changelog.
- SliceUnique函数第一个参数约束为指针.

## [1.6.8] - 2023-06-14
### Changed
- Update changelog.
- 新增文件系统相关的检测函数.

## [1.6.7] - 2023-06-14
### Changed
- Update changelog.
- 新增slice和csv文件互转的函数.

## [1.6.6] - 2023-06-14
### Changed
- Update changelog.
- 新增MultiOnce组件, 可以重置的Sync.Once.

## [1.6.5] - 2023-06-13
### Changed
- Update changelog.
- 新增排序和去重两个新函数.

## [1.6.4] - 2023-06-13
### Changed
- Update changelog.
- 微调字符串数组转结构体的函数.
- 修订部分警告信息.

## [1.6.3] - 2023-06-11
### Changed
- Update changelog.
- 修订Copy的两个入参, 限制必须是指针.

## [1.6.2] - 2023-06-07
### Changed
- Update changelog.
- 修订部分util工具库.

## [1.6.1] - 2023-06-06
### Changed
- Update changelog.
- 新增四舍五入, 日期时间等函数.

## [1.6.0] - 2023-06-03
### Changed
- Update changelog.
- 调整vek目录为num.

## [1.5.1] - 2023-05-13
### Changed
- Update changelog.
- 收录vek汇编工具库.

## [1.5.0] - 2023-05-13
### Changed
- Update changelog.
- 迁移代码仓库到gitee.

## [1.3.33] - 2023-05-12
### Changed
- Update changelog.
- 更新依赖库版本号.

## [1.3.32] - 2023-05-12
### Changed
- Update changelog.
- 修订http客户端header中accept字段错误的bug.

## [1.3.31] - 2023-05-11
### Changed
- Update changelog.
- 调整进度条检测机制.

## [1.3.30] - 2023-05-10
### Changed
- Update changelog.
- 更新依赖库版本号.

## [1.3.29] - 2023-05-07
### Changed
- Update changelog.
- 调整CheckFilepath参数名.

## [1.3.28] - 2023-05-07
### Changed
- Update changelog.
- 调整array tag反射机制代码.
- 删除早期的测试代码.

## [1.3.27] - 2023-05-07
### Changed
- Update changelog.
- 调整git仓库同步脚本.

## [1.3.26] - 2023-05-06
### Changed
- Update changelog.
- 增加切片唯一性排序.

## [1.3.25] - 2023-04-26
### Changed
- Update changelog.
- 调整源文件名.
- Update changelog.

## [1.3.24] - 2023-04-26
### Changed
- 增加slice 过滤函数.

## [1.3.23] - 2023-04-24
### Changed
- Update changelog.
- Treemap加锁.

## [1.3.22] - 2023-04-23
### Changed
- Update changelog.
- 调整bar代码.
- 调整切片反转函数测试代码.

## [1.3.21] - 2023-04-23
### Changed
- Update changelog.
- 增加切片反转函数.

## [1.3.20] - 2023-04-23
### Changed
- Update changelog.
- 修复计时的bug.
- 提升bar更新速度.

## [1.3.19] - 2023-04-23
### Changed
- Update changelog.
- Update changelog.
- 提升bar更新速度.

## [1.3.18] - 2023-04-03
### Changed
- 收敛信号统一处理机制.

## [1.3.17] - 2023-03-20
### Changed
- 修复时区的bug.

## [1.3.16] - 2023-03-20
### Changed
- 优化http client.

## [1.3.15] - 2023-03-18
### Changed
- 警告不推荐使用SetLogPath.

## [1.3.14] - 2023-03-18
### Changed
- Logger增加初始化函数.
- 增加字符串判断空的函数.

## [1.3.13] - 2023-03-18
### Changed
- 判断closer是否为nil.

## [1.3.12] - 2023-03-18
### Changed
- 增加判断是否debug模式.

## [1.3.11] - 2023-03-18
### Changed
- 增加日志组件退出时可能出现panic.

## [1.3.10] - 2023-03-09
### Changed
- 增加检测文件路径的函数.

## [1.3.9] - 2023-02-28
### Changed
- 修订cron版本.

## [1.3.8] - 2023-02-23
### Changed
- 剔除redis的工具库.
- 收录github.com/modern-go/reflect2, 原作者已不维护更新.
- 收录github.com/modern-go/reflect2, 原作者已不维护更新.

## [1.3.7] - 2023-02-21
### Changed
- 屏蔽主进程关闭, 进度条等着结束信号是产生的panic.

## [1.3.6] - 2023-02-19
### Changed
- 增加信号对操作系统的识别.
- 增加信号对操作系统的识别.

## [1.3.5] - 2023-02-19
### Changed
- 增加progressbar工具, github.com/qianlnk/pgbar, 原库有个问题, 就是不能定位在同一行进行循环展示进度条, 原因是maxline一直在增加.

## [1.3.4] - 2023-02-19
### Changed
- 去除部分废弃的代码.
- 增加两个判断前后缀的函数.

## [1.3.3] - 2023-02-18
### Changed
- 增加exception.
- 增加判断chan是否关闭的函数.

## [1.3.2] - 2023-02-18
### Changed
- 修订ioutil.readall函数, 调整到io.readall, 优化部分closer接口.

## [1.3.1] - 2023-01-29
### Changed
- 修订README.
- 修订主机地址错误的问题.
- 增加推送所有tag的命令.

## [1.3.0] - 2023-01-29
### Changed
- 推送后显示当前的远程仓库地址.
- 在"set -e"之后出现的代码，一旦出现了返回值非零，整个脚本就会立即退出.
- 增加gitee和github两个仓库的代码同步脚本.
- 去掉多余的转换.
- 整理部分package.
- 调整包路径.
- 改变package, 将c-struct和struc并入encoding/binary中.
- 不常用的package的归档在labs.

## [1.2.7] - 2023-01-29
### Changed
- 修订reflect2.

## [1.2.6] - 2023-01-27
### Changed
- 增加一对字节数组和字符串互转的函数.

## [1.2.5] - 2023-01-24
### Changed
- Import github.com/mitchellh/go-homedir.

## [1.2.4] - 2023-01-15
### Changed
- 增加lambda工具包.

## [1.2.3] - 2023-01-15
### Changed
- C struct增加测试用例.

## [1.2.2] - 2023-01-15
### Changed
- Fork github.com/fananchong/cstruct-go.

## [1.2.1] - 2023-01-14
### Changed
- 新增struc工具包.
- 更新pool版本.
- 去掉rune内建关键字的警告.
- 去掉rune内建关键字的警告.

## [1.2.0] - 2023-01-01
### Changed
- 修订支持1.20.

## [1.1.21] - 2021-08-23
### Changed
- 增加string书写风格转换.

## [1.1.20] - 2021-08-21
### Changed
- 增加http组件.

## [1.1.19] - 2021-07-29
### Changed
- 调整package.

## [1.1.18] - 2021-07-29
### Changed
- 改变常量的目录.

## [1.1.17] - 2021-07-28
### Changed
- 增加数组相关的函数.
- 调整目录结构.

## [1.1.16] - 2021-07-26
### Changed
- 增加单测.
- 增加测试代码.
- 增加并发hashmap.

## [1.1.15] - 2021-07-16
### Changed
- 恢复gls的单测.

## [1.1.14] - 2021-07-16
### Fixed
- README.

### Changed
- 修订GO111MODULE为auto.
- 调整travis 从org转到com.
- Merge branch 'master' of https://github.com/mymmsc/gox.
- V1.1.x (#1).

## [1.1.13] - 2021-07-15
### Fixed
- README.

### Changed
- 增加go:nocheckptr.
- 修订README.

## [1.1.12] - 2021-07-15
### Changed
- 修订代码检测流水.

## [1.1.11] - 2021-07-15
### Changed
- 调整travis.

## [1.1.10] - 2021-07-15
### Changed
- 修订组件依赖.
- 暂时通过改单测golang原文件名_test.go结尾的方式屏蔽有问题的测试.
- 修订依赖组件.
- 修订orm框架的引用.

## [1.1.9] - 2021-07-10
### Changed
- 增加struct拷贝功能.

## [1.1.8] - 2021-07-07
### Changed
- 整型忽略非数字部分.

## [1.1.7] - 2021-07-07
### Changed
- 整型忽略非数字部分.

## [1.1.6] - 2021-07-07
### Changed
- 整理目录.

## [1.1.5] - 2021-07-07
### Changed
- 修订方法名.

## [1.1.3] - 2021-07-07
### Changed
- 增加数值转字符串的方法.
- 增加安全的字符串转数值的方法.
- 调整基础工具的package.
- 增加忽略异常关闭I/O.
- 去掉没有引用的package.

## [1.1.2] - 2021-06-29
### Changed
- 调整package路径.
- Bou.ke/monkey的目录结构.
- 修改golang版本.
- Add AOP.

## [1.1.1] - 2020-08-01
### Changed
- Add encoding.

## [1.1.0] - 2020-08-01
### Changed
- Fix golang version.

## [1.0.28] - 2020-03-30
### Changed
- Fix version.

## [1.0.27] - 2019-08-17
### Changed
- Fix init.

## [1.0.26] - 2019-08-17
### Changed
- Fix init.

## [1.0.25] - 2019-08-17
### Changed
- Fix MDC.

## [1.0.24] - 2019-08-14
### Changed
- Fix logs.

## [1.0.23] - 2019-08-10
### Changed
- Fix package.

## [1.0.22] - 2019-08-10
### Changed
- Fix package.

## [1.0.21] - 2019-08-10
### Changed
- Add 增加滑动窗口式的waitgroup.

## [1.0.20] - 2019-05-09
### Changed
- Fix fastjson data.

## [1.0.19] - 2019-05-09
### Changed
- 调整slf4g包名.

## [1.0.18] - 2019-05-09
### Changed
- 调整slf4g包名.
- 格式化代码.

## [1.0.17] - 2019-04-23
### Changed
- Fix 备份历史文件的日期.
- Merge branch 'v1.0.x' of https://github.com/mymmsc/gox into v1.0.x.
- Fix sign.
- 增加测试中文.
- Fix package.
- Fix.

## [1.0.16] - 2019-04-14
### Changed
- Fix go test.

## [1.0.15] - 2019-04-14
### Changed
- Add github.com/ddliu/go-httpclient.

## [1.0.14] - 2019-04-14
### Changed
- Add github.com/robfig/cron.

## [1.0.13] - 2019-04-13
### Changed
- Fix package.

## [1.0.12] - 2019-04-13
### Changed
- Add GetString.

## [1.0.11] - 2019-04-13
### Changed
- Add fastjson.

## [1.0.10] - 2019-04-12
### Changed
- Fix timestamp.
- Delete filetype.
- Add filetype.

## [1.0.9] - 2019-04-05
### Changed
- Fix util error.

## [1.0.8] - 2019-04-05
### Changed
- Fix package.
- Add github.com/emirpasic/gods README.
- Fix README.
- Fix package.
- Fix.
- Fix package.
- Fix package.
- Fix util.
- Fix travis.
- Fix travis.
- 修改仓库地址.

## [1.0.7] - 2019-03-31
### Changed
- Fix module list.
- 调整测试脚本.
- 调整测试脚本.
- Change to execute.
- 调整go module依赖.
- 忽略go test测试结果.

## [1.0.6] - 2019-03-31
### Changed
- 删除废弃的代码.
- Add thread local.

## [1.0.5] - 2019-03-30
### Changed
- 删除golang.org/x/sys的引用.

## [1.0.4] - 2019-03-30
### Changed
- Add logger.

## [1.0.3] - 2019-03-30
### Changed
- Add redis pool.

## [1.0.2] - 2019-03-24
### Changed
- Add testing.

## [1.0.1] - 2019-03-24
### Changed
- Fix version.
- Fix version.

## [1.0.0] - 2019-03-24
### Changed
- Fix codecov.
- Fix license.
- Fix codecov.
- Fix codecov.
- Add test code.
- Add test code.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix codecov.
- Fix license.
- Fix travis.
- Fix travis.
- Fix travis.
- Fix golang version.
- Fix.
- Fix.
- Fix.
- Fix.
- Fix travis.
- Add travis-ci.
- Add go-xorm 反转数据库结构 用法.
- Add go-xorm mysql template.
- 修改二维码路径.
- Add QR terminal.
- Add support go module.
- Add gitingore.
- Initial commit.

[Unreleased]: https://gitee.com/quant1x/gox/compare/v1.18.3...HEAD
[1.18.3]: https://gitee.com/quant1x/gox/compare/v1.18.2...v1.18.3
[1.18.2]: https://gitee.com/quant1x/gox/compare/v1.18.1...v1.18.2
[1.18.1]: https://gitee.com/quant1x/gox/compare/v1.18.0...v1.18.1
[1.18.0]: https://gitee.com/quant1x/gox/compare/v1.17.9...v1.18.0
[1.17.9]: https://gitee.com/quant1x/gox/compare/v1.17.8...v1.17.9
[1.17.8]: https://gitee.com/quant1x/gox/compare/v1.17.7...v1.17.8
[1.17.7]: https://gitee.com/quant1x/gox/compare/v1.17.6...v1.17.7
[1.17.6]: https://gitee.com/quant1x/gox/compare/v1.17.5...v1.17.6
[1.17.5]: https://gitee.com/quant1x/gox/compare/v1.17.4...v1.17.5
[1.17.4]: https://gitee.com/quant1x/gox/compare/v1.17.3...v1.17.4
[1.17.3]: https://gitee.com/quant1x/gox/compare/v1.17.2...v1.17.3
[1.17.2]: https://gitee.com/quant1x/gox/compare/v1.17.1...v1.17.2
[1.17.1]: https://gitee.com/quant1x/gox/compare/v1.17.0...v1.17.1
[1.17.0]: https://gitee.com/quant1x/gox/compare/v1.16.9...v1.17.0
[1.16.9]: https://gitee.com/quant1x/gox/compare/v1.16.8...v1.16.9
[1.16.8]: https://gitee.com/quant1x/gox/compare/v1.16.7...v1.16.8
[1.16.7]: https://gitee.com/quant1x/gox/compare/v1.16.6...v1.16.7
[1.16.6]: https://gitee.com/quant1x/gox/compare/v1.16.5...v1.16.6
[1.16.5]: https://gitee.com/quant1x/gox/compare/v1.16.4...v1.16.5
[1.16.4]: https://gitee.com/quant1x/gox/compare/v1.16.3...v1.16.4
[1.16.3]: https://gitee.com/quant1x/gox/compare/v1.16.2...v1.16.3
[1.16.2]: https://gitee.com/quant1x/gox/compare/v1.16.1...v1.16.2
[1.16.1]: https://gitee.com/quant1x/gox/compare/v1.16.0...v1.16.1
[1.16.0]: https://gitee.com/quant1x/gox/compare/v1.15.9...v1.16.0
[1.15.9]: https://gitee.com/quant1x/gox/compare/v1.15.8...v1.15.9
[1.15.8]: https://gitee.com/quant1x/gox/compare/v1.15.7...v1.15.8
[1.15.7]: https://gitee.com/quant1x/gox/compare/v1.15.6...v1.15.7
[1.15.6]: https://gitee.com/quant1x/gox/compare/v1.15.5...v1.15.6
[1.15.5]: https://gitee.com/quant1x/gox/compare/v1.15.4...v1.15.5
[1.15.4]: https://gitee.com/quant1x/gox/compare/v1.15.3...v1.15.4
[1.15.3]: https://gitee.com/quant1x/gox/compare/v1.15.2...v1.15.3
[1.15.2]: https://gitee.com/quant1x/gox/compare/v1.15.1...v1.15.2
[1.15.1]: https://gitee.com/quant1x/gox/compare/v1.15.0...v1.15.1
[1.15.0]: https://gitee.com/quant1x/gox/compare/v1.14.9...v1.15.0
[1.14.9]: https://gitee.com/quant1x/gox/compare/v1.14.8...v1.14.9
[1.14.8]: https://gitee.com/quant1x/gox/compare/v1.14.7...v1.14.8
[1.14.7]: https://gitee.com/quant1x/gox/compare/v1.14.6...v1.14.7
[1.14.6]: https://gitee.com/quant1x/gox/compare/v1.14.5...v1.14.6
[1.14.5]: https://gitee.com/quant1x/gox/compare/v1.14.4...v1.14.5
[1.14.4]: https://gitee.com/quant1x/gox/compare/v1.14.3...v1.14.4
[1.14.3]: https://gitee.com/quant1x/gox/compare/v1.14.2...v1.14.3
[1.14.2]: https://gitee.com/quant1x/gox/compare/v1.14.1...v1.14.2
[1.14.1]: https://gitee.com/quant1x/gox/compare/v1.14.0...v1.14.1
[1.14.0]: https://gitee.com/quant1x/gox/compare/v1.13.9...v1.14.0
[1.13.9]: https://gitee.com/quant1x/gox/compare/v1.13.8...v1.13.9
[1.13.8]: https://gitee.com/quant1x/gox/compare/v1.13.7...v1.13.8
[1.13.7]: https://gitee.com/quant1x/gox/compare/v1.13.6...v1.13.7
[1.13.6]: https://gitee.com/quant1x/gox/compare/v1.13.5...v1.13.6
[1.13.5]: https://gitee.com/quant1x/gox/compare/v1.13.4...v1.13.5
[1.13.4]: https://gitee.com/quant1x/gox/compare/v1.13.3...v1.13.4
[1.13.3]: https://gitee.com/quant1x/gox/compare/v1.13.2...v1.13.3
[1.13.2]: https://gitee.com/quant1x/gox/compare/v1.13.1...v1.13.2
[1.13.1]: https://gitee.com/quant1x/gox/compare/v1.13.0...v1.13.1
[1.13.0]: https://gitee.com/quant1x/gox/compare/v1.12.9...v1.13.0
[1.12.9]: https://gitee.com/quant1x/gox/compare/v1.12.8...v1.12.9
[1.12.8]: https://gitee.com/quant1x/gox/compare/v1.12.7...v1.12.8
[1.12.7]: https://gitee.com/quant1x/gox/compare/v1.12.6...v1.12.7
[1.12.6]: https://gitee.com/quant1x/gox/compare/v1.12.5...v1.12.6
[1.12.5]: https://gitee.com/quant1x/gox/compare/v1.12.4...v1.12.5
[1.12.4]: https://gitee.com/quant1x/gox/compare/v1.12.3...v1.12.4
[1.12.3]: https://gitee.com/quant1x/gox/compare/v1.12.2...v1.12.3
[1.12.2]: https://gitee.com/quant1x/gox/compare/v1.12.1...v1.12.2
[1.12.1]: https://gitee.com/quant1x/gox/compare/v1.12.0...v1.12.1
[1.12.0]: https://gitee.com/quant1x/gox/compare/v1.11.9...v1.12.0
[1.11.9]: https://gitee.com/quant1x/gox/compare/v1.11.8...v1.11.9
[1.11.8]: https://gitee.com/quant1x/gox/compare/v1.11.7...v1.11.8
[1.11.7]: https://gitee.com/quant1x/gox/compare/v1.11.6...v1.11.7
[1.11.6]: https://gitee.com/quant1x/gox/compare/v1.11.5...v1.11.6
[1.11.5]: https://gitee.com/quant1x/gox/compare/v1.11.4...v1.11.5
[1.11.4]: https://gitee.com/quant1x/gox/compare/v1.11.3...v1.11.4
[1.11.3]: https://gitee.com/quant1x/gox/compare/v1.11.2...v1.11.3
[1.11.2]: https://gitee.com/quant1x/gox/compare/v1.11.1...v1.11.2
[1.11.1]: https://gitee.com/quant1x/gox/compare/v1.11.0...v1.11.1
[1.11.0]: https://gitee.com/quant1x/gox/compare/v1.10.9...v1.11.0
[1.10.9]: https://gitee.com/quant1x/gox/compare/v1.10.8...v1.10.9
[1.10.8]: https://gitee.com/quant1x/gox/compare/v1.10.7...v1.10.8
[1.10.7]: https://gitee.com/quant1x/gox/compare/v1.10.6...v1.10.7
[1.10.6]: https://gitee.com/quant1x/gox/compare/v1.10.5...v1.10.6
[1.10.5]: https://gitee.com/quant1x/gox/compare/v1.10.4...v1.10.5
[1.10.4]: https://gitee.com/quant1x/gox/compare/v1.10.3...v1.10.4
[1.10.3]: https://gitee.com/quant1x/gox/compare/v1.10.2...v1.10.3
[1.10.2]: https://gitee.com/quant1x/gox/compare/v1.10.1...v1.10.2
[1.10.1]: https://gitee.com/quant1x/gox/compare/v1.10.0...v1.10.1
[1.10.0]: https://gitee.com/quant1x/gox/compare/v1.9.9...v1.10.0
[1.9.9]: https://gitee.com/quant1x/gox/compare/v1.9.8...v1.9.9
[1.9.8]: https://gitee.com/quant1x/gox/compare/v1.9.7...v1.9.8
[1.9.7]: https://gitee.com/quant1x/gox/compare/v1.9.6...v1.9.7
[1.9.6]: https://gitee.com/quant1x/gox/compare/v1.9.5...v1.9.6
[1.9.5]: https://gitee.com/quant1x/gox/compare/v1.9.4...v1.9.5
[1.9.4]: https://gitee.com/quant1x/gox/compare/v1.9.3...v1.9.4
[1.9.3]: https://gitee.com/quant1x/gox/compare/v1.9.2...v1.9.3
[1.9.2]: https://gitee.com/quant1x/gox/compare/v1.9.1...v1.9.2
[1.9.1]: https://gitee.com/quant1x/gox/compare/v1.9.0...v1.9.1
[1.9.0]: https://gitee.com/quant1x/gox/compare/v1.8.9...v1.9.0
[1.8.9]: https://gitee.com/quant1x/gox/compare/v1.8.8...v1.8.9
[1.8.8]: https://gitee.com/quant1x/gox/compare/v1.8.7...v1.8.8
[1.8.7]: https://gitee.com/quant1x/gox/compare/v1.8.6...v1.8.7
[1.8.6]: https://gitee.com/quant1x/gox/compare/v1.8.5...v1.8.6
[1.8.5]: https://gitee.com/quant1x/gox/compare/v1.8.4...v1.8.5
[1.8.4]: https://gitee.com/quant1x/gox/compare/v1.8.3...v1.8.4
[1.8.3]: https://gitee.com/quant1x/gox/compare/v1.8.2...v1.8.3
[1.8.2]: https://gitee.com/quant1x/gox/compare/v1.8.1...v1.8.2
[1.8.1]: https://gitee.com/quant1x/gox/compare/v1.8.0...v1.8.1
[1.8.0]: https://gitee.com/quant1x/gox/compare/v1.7.9...v1.8.0
[1.7.9]: https://gitee.com/quant1x/gox/compare/v1.7.8...v1.7.9
[1.7.8]: https://gitee.com/quant1x/gox/compare/v1.7.7...v1.7.8
[1.7.7]: https://gitee.com/quant1x/gox/compare/v1.7.6...v1.7.7
[1.7.6]: https://gitee.com/quant1x/gox/compare/v1.7.5...v1.7.6
[1.7.5]: https://gitee.com/quant1x/gox/compare/v1.7.4...v1.7.5
[1.7.4]: https://gitee.com/quant1x/gox/compare/v1.7.3...v1.7.4
[1.7.3]: https://gitee.com/quant1x/gox/compare/v1.7.2...v1.7.3
[1.7.2]: https://gitee.com/quant1x/gox/compare/v1.7.1...v1.7.2
[1.7.1]: https://gitee.com/quant1x/gox/compare/v1.7.0...v1.7.1
[1.7.0]: https://gitee.com/quant1x/gox/compare/v1.6.9...v1.7.0
[1.6.9]: https://gitee.com/quant1x/gox/compare/v1.6.8...v1.6.9
[1.6.8]: https://gitee.com/quant1x/gox/compare/v1.6.7...v1.6.8
[1.6.7]: https://gitee.com/quant1x/gox/compare/v1.6.6...v1.6.7
[1.6.6]: https://gitee.com/quant1x/gox/compare/v1.6.5...v1.6.6
[1.6.5]: https://gitee.com/quant1x/gox/compare/v1.6.4...v1.6.5
[1.6.4]: https://gitee.com/quant1x/gox/compare/v1.6.3...v1.6.4
[1.6.3]: https://gitee.com/quant1x/gox/compare/v1.6.2...v1.6.3
[1.6.2]: https://gitee.com/quant1x/gox/compare/v1.6.1...v1.6.2
[1.6.1]: https://gitee.com/quant1x/gox/compare/v1.6.0...v1.6.1
[1.6.0]: https://gitee.com/quant1x/gox/compare/v1.5.1...v1.6.0
[1.5.1]: https://gitee.com/quant1x/gox/compare/v1.5.0...v1.5.1
[1.5.0]: https://gitee.com/quant1x/gox/compare/v1.3.33...v1.5.0
[1.3.33]: https://gitee.com/quant1x/gox/compare/v1.3.32...v1.3.33
[1.3.32]: https://gitee.com/quant1x/gox/compare/v1.3.31...v1.3.32
[1.3.31]: https://gitee.com/quant1x/gox/compare/v1.3.30...v1.3.31
[1.3.30]: https://gitee.com/quant1x/gox/compare/v1.3.29...v1.3.30
[1.3.29]: https://gitee.com/quant1x/gox/compare/v1.3.28...v1.3.29
[1.3.28]: https://gitee.com/quant1x/gox/compare/v1.3.27...v1.3.28
[1.3.27]: https://gitee.com/quant1x/gox/compare/v1.3.26...v1.3.27
[1.3.26]: https://gitee.com/quant1x/gox/compare/v1.3.25...v1.3.26
[1.3.25]: https://gitee.com/quant1x/gox/compare/v1.3.24...v1.3.25
[1.3.24]: https://gitee.com/quant1x/gox/compare/v1.3.23...v1.3.24
[1.3.23]: https://gitee.com/quant1x/gox/compare/v1.3.22...v1.3.23
[1.3.22]: https://gitee.com/quant1x/gox/compare/v1.3.21...v1.3.22
[1.3.21]: https://gitee.com/quant1x/gox/compare/v1.3.20...v1.3.21
[1.3.20]: https://gitee.com/quant1x/gox/compare/v1.3.19...v1.3.20
[1.3.19]: https://gitee.com/quant1x/gox/compare/v1.3.18...v1.3.19
[1.3.18]: https://gitee.com/quant1x/gox/compare/v1.3.17...v1.3.18
[1.3.17]: https://gitee.com/quant1x/gox/compare/v1.3.16...v1.3.17
[1.3.16]: https://gitee.com/quant1x/gox/compare/v1.3.15...v1.3.16
[1.3.15]: https://gitee.com/quant1x/gox/compare/v1.3.14...v1.3.15
[1.3.14]: https://gitee.com/quant1x/gox/compare/v1.3.13...v1.3.14
[1.3.13]: https://gitee.com/quant1x/gox/compare/v1.3.12...v1.3.13
[1.3.12]: https://gitee.com/quant1x/gox/compare/v1.3.11...v1.3.12
[1.3.11]: https://gitee.com/quant1x/gox/compare/v1.3.10...v1.3.11
[1.3.10]: https://gitee.com/quant1x/gox/compare/v1.3.9...v1.3.10
[1.3.9]: https://gitee.com/quant1x/gox/compare/v1.3.8...v1.3.9
[1.3.8]: https://gitee.com/quant1x/gox/compare/v1.3.7...v1.3.8
[1.3.7]: https://gitee.com/quant1x/gox/compare/v1.3.6...v1.3.7
[1.3.6]: https://gitee.com/quant1x/gox/compare/v1.3.5...v1.3.6
[1.3.5]: https://gitee.com/quant1x/gox/compare/v1.3.4...v1.3.5
[1.3.4]: https://gitee.com/quant1x/gox/compare/v1.3.3...v1.3.4
[1.3.3]: https://gitee.com/quant1x/gox/compare/v1.3.2...v1.3.3
[1.3.2]: https://gitee.com/quant1x/gox/compare/v1.3.1...v1.3.2
[1.3.1]: https://gitee.com/quant1x/gox/compare/v1.3.0...v1.3.1
[1.3.0]: https://gitee.com/quant1x/gox/compare/v1.2.7...v1.3.0
[1.2.7]: https://gitee.com/quant1x/gox/compare/v1.2.6...v1.2.7
[1.2.6]: https://gitee.com/quant1x/gox/compare/v1.2.5...v1.2.6
[1.2.5]: https://gitee.com/quant1x/gox/compare/v1.2.4...v1.2.5
[1.2.4]: https://gitee.com/quant1x/gox/compare/v1.2.3...v1.2.4
[1.2.3]: https://gitee.com/quant1x/gox/compare/v1.2.2...v1.2.3
[1.2.2]: https://gitee.com/quant1x/gox/compare/v1.2.1...v1.2.2
[1.2.1]: https://gitee.com/quant1x/gox/compare/v1.2.0...v1.2.1
[1.2.0]: https://gitee.com/quant1x/gox/compare/v1.1.21...v1.2.0
[1.1.21]: https://gitee.com/quant1x/gox/compare/v1.1.20...v1.1.21
[1.1.20]: https://gitee.com/quant1x/gox/compare/v1.1.19...v1.1.20
[1.1.19]: https://gitee.com/quant1x/gox/compare/v1.1.18...v1.1.19
[1.1.18]: https://gitee.com/quant1x/gox/compare/v1.1.17...v1.1.18
[1.1.17]: https://gitee.com/quant1x/gox/compare/v1.1.16...v1.1.17
[1.1.16]: https://gitee.com/quant1x/gox/compare/v1.1.15...v1.1.16
[1.1.15]: https://gitee.com/quant1x/gox/compare/v1.1.14...v1.1.15
[1.1.14]: https://gitee.com/quant1x/gox/compare/v1.1.13...v1.1.14
[1.1.13]: https://gitee.com/quant1x/gox/compare/v1.1.12...v1.1.13
[1.1.12]: https://gitee.com/quant1x/gox/compare/v1.1.11...v1.1.12
[1.1.11]: https://gitee.com/quant1x/gox/compare/v1.1.10...v1.1.11
[1.1.10]: https://gitee.com/quant1x/gox/compare/v1.1.9...v1.1.10
[1.1.9]: https://gitee.com/quant1x/gox/compare/v1.1.8...v1.1.9
[1.1.8]: https://gitee.com/quant1x/gox/compare/v1.1.7...v1.1.8
[1.1.7]: https://gitee.com/quant1x/gox/compare/v1.1.6...v1.1.7
[1.1.6]: https://gitee.com/quant1x/gox/compare/v1.1.5...v1.1.6
[1.1.5]: https://gitee.com/quant1x/gox/compare/v1.1.3...v1.1.5
[1.1.3]: https://gitee.com/quant1x/gox/compare/v1.1.2...v1.1.3
[1.1.2]: https://gitee.com/quant1x/gox/compare/v1.1.1...v1.1.2
[1.1.1]: https://gitee.com/quant1x/gox/compare/v1.1.0...v1.1.1
[1.1.0]: https://gitee.com/quant1x/gox/compare/v1.0.28...v1.1.0
[1.0.28]: https://gitee.com/quant1x/gox/compare/v1.0.27...v1.0.28
[1.0.27]: https://gitee.com/quant1x/gox/compare/v1.0.26...v1.0.27
[1.0.26]: https://gitee.com/quant1x/gox/compare/v1.0.25...v1.0.26
[1.0.25]: https://gitee.com/quant1x/gox/compare/v1.0.24...v1.0.25
[1.0.24]: https://gitee.com/quant1x/gox/compare/v1.0.23...v1.0.24
[1.0.23]: https://gitee.com/quant1x/gox/compare/v1.0.22...v1.0.23
[1.0.22]: https://gitee.com/quant1x/gox/compare/v1.0.21...v1.0.22
[1.0.21]: https://gitee.com/quant1x/gox/compare/v1.0.20...v1.0.21
[1.0.20]: https://gitee.com/quant1x/gox/compare/v1.0.19...v1.0.20
[1.0.19]: https://gitee.com/quant1x/gox/compare/v1.0.18...v1.0.19
[1.0.18]: https://gitee.com/quant1x/gox/compare/v1.0.17...v1.0.18
[1.0.17]: https://gitee.com/quant1x/gox/compare/v1.0.16...v1.0.17
[1.0.16]: https://gitee.com/quant1x/gox/compare/v1.0.15...v1.0.16
[1.0.15]: https://gitee.com/quant1x/gox/compare/v1.0.14...v1.0.15
[1.0.14]: https://gitee.com/quant1x/gox/compare/v1.0.13...v1.0.14
[1.0.13]: https://gitee.com/quant1x/gox/compare/v1.0.12...v1.0.13
[1.0.12]: https://gitee.com/quant1x/gox/compare/v1.0.11...v1.0.12
[1.0.11]: https://gitee.com/quant1x/gox/compare/v1.0.10...v1.0.11
[1.0.10]: https://gitee.com/quant1x/gox/compare/v1.0.9...v1.0.10
[1.0.9]: https://gitee.com/quant1x/gox/compare/v1.0.8...v1.0.9
[1.0.8]: https://gitee.com/quant1x/gox/compare/v1.0.7...v1.0.8
[1.0.7]: https://gitee.com/quant1x/gox/compare/v1.0.6...v1.0.7
[1.0.6]: https://gitee.com/quant1x/gox/compare/v1.0.5...v1.0.6
[1.0.5]: https://gitee.com/quant1x/gox/compare/v1.0.4...v1.0.5
[1.0.4]: https://gitee.com/quant1x/gox/compare/v1.0.3...v1.0.4
[1.0.3]: https://gitee.com/quant1x/gox/compare/v1.0.2...v1.0.3
[1.0.2]: https://gitee.com/quant1x/gox/compare/v1.0.1...v1.0.2
[1.0.1]: https://gitee.com/quant1x/gox/compare/v1.0.0...v1.0.1
[1.0.0]: https://gitee.com/quant1x/gox/releases/tag/v1.0.0
