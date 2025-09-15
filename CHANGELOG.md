# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.24.7] - 2025-09-15
### Changed
- 修改字符串源文件名

## [1.24.6] - 2025-09-15
### Changed
- 修复关闭全部连接的bug
- update changelog

## [1.24.5] - 2025-09-14
### Changed
- 调整进度条demo代码
- update changelog

## [1.24.4] - 2025-09-14
### Changed
- 更新依赖库版本
- update changelog

## [1.24.3] - 2025-09-14
### Changed
- 删除废弃的gitee和github同步脚本
- update changelog

## [1.24.2] - 2025-09-14
### Changed
- 删除原生的gocsv, 对浮点和整数处理的有问题
- update changelog

## [1.24.1] - 2025-09-14
### Changed
- 修订go最低支持的版本
- update changelog

## [1.24.0] - 2025-09-14
### Changed
- 更新依赖库版本,仓库改为github
- update changelog

## [1.23.2] - 2025-09-14
### Changed
- 更新依赖库版本
- update changelog

## [1.23.1] - 2025-08-14
### Changed
- 更新依赖库版本
- update changelog

## [1.23.0] - 2025-08-14
### Changed
- gox最低支持go1.25
- update changelog

## [1.22.13] - 2025-08-11
### Changed
- sort imports
- update changelog

## [1.22.12] - 2025-07-08
### Changed
- 增加内存锁定的注释
- 新增内存锁定功能
- mmap功能迁移至sys/mem
- 调整部分代码的布局
- 修正mmap功能函数
- 调整测试mmap产生的临时文件名, 避免git加入仓库
- 调整内存映射封装的结构体名
- 预备std标准功能
- 早前的cache功能改名到v1版本
- 新增第2个版本的cache功能
- 新增第2个版本的cache功能
- 重构跨平台的mmap功能
- 调整.gitignore配置,.dat文件都要需要git仓库忽略
- 标准库array增加实验性扩容函数
- 调整第3版内存映射
- 新增私有常量, 文件系统权限
- 更新依赖库版本
- 删除废弃的cache代码
- update changelog

## [1.22.11] - 2025-03-10
### Changed
- 优化泛型Channel的数据功能
- update changelog

## [1.22.10] - 2025-03-10
### Changed
- 优化部分代码
- update changelog

## [1.22.9] - 2025-03-10
### Changed
- 修复在NewRollingOnceWithHour方法中存在致命错误
- 检测hour和minute参数的有效性
- 传入参数前对齐数据类型
- update changelog

## [1.22.8] - 2025-03-10
### Changed
- 更新pkg版本到0.5.1
- update changelog

## [1.22.7] - 2025-03-09
### Changed
- 去掉终端输出的调试信息
- 新增系统信号的hook功能, 允许注册和主动发送关闭信号
- 修复信号不准确的bug
- update changelog

## [1.22.6] - 2025-03-09
### Changed
- 更新依赖库版本
- update changelog

## [1.22.5] - 2025-03-09
### Changed
- 拆分原logger
- 单独处理日期更新
- 统一退出控制
- 优化压缩旧文件的处理过程, 修复原文件没有释放的bug
- 修复defer调用文件关闭方法的警告信息
- update changelog

## [1.22.4] - 2025-03-08
### Changed
- 修复Possible misuse of 'unsafe.Pointer'警告
- 修复定时器不能退出的bug
- 新建第2个版本的RollingOnce
- 剔除冗余的判断
- 整理日志记录器的部分代码, 去掉冗余的应用程序名称获取的sync/once用法
- update changelog

## [1.22.3] - 2025-03-02
### Changed
- 删除部分废弃的代码
- 删除部分废弃的代码
- update changelog

## [1.22.2] - 2025-02-26
### Changed
- 依据最大行数预先输出足够的行
- update changelog

## [1.22.1] - 2025-02-26
### Changed
- 修订周期0的注释错误,应该是Sunday=0
- 新增一个获取指定日期的周开始时间和结束时间
- 优化进度条
- update changelog

## [1.22.0] - 2025-02-15
### Changed
- 更新以支持go1.24
- update changelog

## [1.21.9] - 2024-08-06
### Changed
- 更新依赖库版本
- update changelog

## [1.21.8] - 2024-08-06
### Changed
- 更新依赖库版本

## [1.21.7] - 2024-08-06
### Changed
- 更新依赖库版本
- update changelog

## [1.21.6] - 2024-08-06
### Changed
- 新增解析时间类字符串的基准测试
- 新增解析时间类字符串的基准测试代码
- 调整hashmap部分代码
- 调整部分测试代码
- 更新依赖库版本
- update changelog

## [1.21.5] - 2024-07-05
### Changed
- 优化字节数组转泛型切片
- 更新pkg版本到0.2.10
- update changelog

## [1.21.4] - 2024-06-20
### Changed
- 调整mmap组件代码
- update changelog
- update changelog

## [1.21.3] - 2024-06-14
### Changed
- 更新依赖库版本

## [1.21.2] - 2024-05-16
### Changed
- 切片保存csv文件新增强制参数, 默认不强制刷新
- update changelog
- update changelog

## [1.21.1] - 2024-05-11
### Changed
- 修订同步github代码的repo
- 更新依赖库版本
- 调整mmap映射切片的处理方法
- update changelog

## [1.21.0] - 2024-03-30
### Changed
- 修订REAMDE中的徽章
- 修订go版本
- update changelog

## [1.20.9] - 2024-03-21
### Changed
- 切片为空就是清空文件
- update changelog

## [1.20.8] - 2024-03-17
### Changed
- 更新依赖库版本
- update changelog

## [1.20.7] - 2024-03-12
### Changed
- 更新go版本
- update changelog

## [1.20.6] - 2024-03-12
### Changed
- 保存csv文件强制刷新磁盘
- update changelog

## [1.20.5] - 2024-03-12
### Changed
- 非mac版本不优化Now函数
- update changelog

## [1.20.4] - 2024-03-11
### Changed
- 更新依赖库版本
- update changelog

## [1.20.3] - 2024-03-11
### Changed
- 屏蔽v2版本的Now实现
- update changelog

## [1.20.2] - 2024-02-27
### Changed
- 优化Timestamp.Now函数, 去掉nanotime函数的调用
- 优化Catch处理方法, 应对vet对(v ...any)中出现%的错误检查
- update changelog

## [1.20.1] - 2024-02-14
### Changed
- 取消gox对num的依赖
- update changelog

## [1.20.0] - 2024-02-14
### Changed
- 微调部分函数错误定义
- update changelog

## [1.19.9] - 2024-02-12
### Changed
- 独立拆分num工具包
- tags包适配num
- update changelog

## [1.19.8] - 2024-02-10
### Changed
- 更新依赖库版本
- 微调部分代码
- 调整部分测试代码数据
- 调整vek的license和readme源文件名
- 调整vek代码机构, 去掉arm64冗余代码
- 调整部分代码
- 补充arm框架缺失的函数
- num默认开启加速
- update changelog

## [1.19.7] - 2024-02-06
### Changed
- 改interface{}为any
- 新增反射功能函数
- 新增mmap转slice功能
- 修复unmap windows调用错误的bug
- 修订slice长度
- 删除废弃字段定义
- update changelog

## [1.19.6] - 2024-01-27
### Changed
- 进度条收尾阶段的循环处理, 增加sleep防止对于advance造成假死锁的现象
- 调整进度条部分代码
- update changelog

## [1.19.5] - 2024-01-27
### Changed
- 时间戳新增计算年月日时分秒毫秒的方法
- update changelog

## [1.19.4] - 2024-01-25
### Changed
- 优化精度条
- update changelog

## [1.19.3] - 2024-01-25
### Changed
- 优化精度条存在race的问题
- 屏蔽关闭状态的测试代码
- 去掉废弃的代码
- update changelog

## [1.19.2] - 2024-01-25
### Changed
- 调整RollingOnce部分函数和私有成员变量名
- update changelog

## [1.19.1] - 2024-01-25
### Changed
- RollingOnce增加返回当前周期的观察点毫秒数
- update changelog

## [1.19.0] - 2024-01-24
### Changed
- 修复RollingOnce丢失窗口期的计算
- update changelog

## [1.18.9] - 2024-01-24
### Changed
- 修复RollingOnce竞态数据问题
- 修复Logger竞态数据问题
- 改interface{}为any
- 优化进度条
- 删除废弃的代码
- 修订测试代码
- 删除废弃的代码
- 优化捕获panic代码, panic应该是第一时间捕获, 增加隔函数调用recover返回nil
- 去掉记录日志的panic捕获
- 去掉logger初始化捕获异常
- 调整调度任务函数可继续信号的发送位置
- 修复进度条计数器没走完整的bug
- 调整调度函数退出机制
- 修复通道关闭的情况下会继续发送数据
- update changelog

## [1.18.8] - 2024-01-23
### Changed
- 修订日志使用应用程序名时去掉扩展名
- update changelog

## [1.18.7] - 2024-01-23
### Changed
- 屏蔽FastCache代码, syscall.Mmap不跨平台
- update changelog

## [1.18.6] - 2024-01-23
### Changed
- 致命的日志写入文件且输出到控制台
- update changelog

## [1.18.5] - 2024-01-23
### Changed
- cache增加泛型的map和pool
- 优化日志记录器: ①按照应用程序名进一步拆分日志目录, ②日志对象logValue的使用cache.Pool提升写入速度
- 调整变量名为驼峰格式
- update changelog

## [1.18.4] - 2024-01-22
### Changed
- 增加部分关闭和初始化方面的告警类日志
- update changelog

## [1.18.3] - 2024-01-22
### Changed
- 修复设置偏移量死锁的bug
- update changelog

## [1.18.2] - 2024-01-22
### Changed
- update changelog
- RollingOnce 新增按照小时和分钟设置偏移量
- update changelog

## [1.18.1] - 2024-01-22
### Changed
- 优化RollingOnce, 增加可重置偏移量的方法
- 这两都提示推荐使用RollingOnce

## [1.18.0] - 2024-01-20
### Changed
- 优化连接池Pool部分代码
- 补充Pool接口注释
- 注释掉部分告警日志
- 删除部分废弃的代码
- update changelog

## [1.17.9] - 2024-01-18
### Changed
- 增加Touch函数
- update changelog

## [1.17.8] - 2024-01-17
### Changed
- Bar结构增加部分注释
- bar公开方法增加注释
- update changelog

## [1.17.7] - 2024-01-17
### Changed
- 修复未初始化的bug
- 修复未初始化的bug
- update changelog

## [1.17.6] - 2024-01-17
### Changed
- 进度条增加结束等待方法
- update changelog

## [1.17.5] - 2024-01-14
### Changed
- 优化可以忽略panic的异常捕获函数
- update changelog

## [1.17.4] - 2024-01-14
### Changed
- 调整解析参数的时机, 只在异常发生时解析
- update changelog

## [1.17.3] - 2024-01-14
### Changed
- 允许CatchPanic传入可变参数
- update changelog

## [1.17.2] - 2024-01-13
### Changed
- 更新依赖库版本
- update changelog

## [1.17.1] - 2024-01-11
### Changed
- 删除废弃的代码
- update changelog
- update changelog

## [1.17.0] - 2024-01-11
### Changed
- update changelog
- 删除废弃的代码
- 修订Changelog

## [1.16.9] - 2024-01-09
### Changed
- 优化部分代码, 删除非必要的import
- update changelog

## [1.16.8] - 2024-01-08
### Changed
- 修订ParseTime的函数注释
- 更新依赖库golang.org/x/exp版本
- 新增时间戳(毫秒数)的功能函数
- 调整部分代码
- 优化timestamp时间戳
- 新增RollingOnce, 相对于PeriodicOnce优化了时间戳的处理方式
- 修订注释
- 优化部分代码
- update changelog

## [1.16.7] - 2024-01-02
### Changed
- 优化应用程序名的获取方式, 改成懒加载
- 增加功能性函数, 捕获panic, 忽略异常, 继续执行
- 调整util包中的时间格式功能到api包
- 调整util包中的uuid能到pkg
- update changelog

## [1.16.6] - 2023-12-31
### Changed
- 迁移: github.com/valyala/fastjson@v1.6.4 到 pkg
- 迁移: github.com/smartystreets-prototypes/go-disruptor@v0.0.0-20231024205940-61200af675a1 到 pkg
- 删除errors包
- 调整因删除errors包对其它功能的影响
- 更新依赖库pkg版本
- 更新依赖库pkg版本
- 更新依赖库pkg版本
- update changelog

## [1.16.5] - 2023-12-30
### Changed
- 调整缓存csv逻辑, 如果切片为空, 直接返回
- update changelog

## [1.16.4] - 2023-12-30
### Changed
- 调整文件名, 字面意义和功能保持一致
- update changelog

## [1.16.3] - 2023-12-30
### Changed
- 修复num.Decimals函数负浮点四舍五入错误的bug
- update changelog

## [1.16.2] - 2023-12-23
### Changed
- 更新依赖库版本
- 更新依赖库版本
- 更新依赖库版本
- update changelog

## [1.16.1] - 2023-12-23
### Changed
- 从exp包中复制maps的Keys和Values函数, 意在移除对exp的依赖
- 优化filestat中时间的处理方法
- 修复windows变量引用的错误
- update changelog

## [1.16.0] - 2023-12-23
### Changed
- 优化部分代码
- update changelog

## [1.15.9] - 2023-12-22
### Changed
- 增加windows操作系统的文件时间戳的获取方法
- 新增获取文件状态(时间)的函数
- 修复windows缺少error返回值的bug
- update changelog

## [1.15.8] - 2023-12-16
### Changed
- PeriodOnce默认在非runtime的debug模式下不输出日志
- update changelog

## [1.15.7] - 2023-12-15
### Changed
- 定时任务默认在非runtime的debug模式下不输出日志
- update changelog

## [1.15.6] - 2023-12-14
### Changed
- 给获取应用程序文件名的函数增加注释
- 优化去重处理方式
- update changelog

## [1.15.5] - 2023-12-12
### Changed
- 删除旧版本的滑动Once功能
- 增加文件路径中的日期格式, 数据中的日期格式
- update changelog

## [1.15.4] - 2023-12-12
### Changed
- 更新pkg版本同步go版本
- update changelog

## [1.15.3] - 2023-12-12
### Changed
- 更新go1.21.5
- update changelog

## [1.15.2] - 2023-12-07
### Changed
- 更新依赖库pkg版本
- update changelog

## [1.15.1] - 2023-12-05
### Changed
- 更新依赖库版本
- update changelog

## [1.15.0] - 2023-12-04
### Changed
- post方法去掉返回值中的lastModified
- update changelog

## [1.14.9] - 2023-12-04
### Changed
- 修复json判断失败的bug
- update changelog

## [1.14.8] - 2023-12-04
### Changed
- 增加一个容错机制的HttpPost函数
- update changelog

## [1.14.7] - 2023-12-04
### Changed
- 优化部分功能函数
- update changelog

## [1.14.6] - 2023-12-04
### Changed
- http新增一个独立的Get方法, 允许传入header
- HttpRequest函数增加可以传入header
- update changelog

## [1.14.5] - 2023-12-04
### Changed
- http工具包增加post方法
- update changelog

## [1.14.4] - 2023-12-04
### Changed
- 运行时允许重置debug状态
- update changelog

## [1.14.3] - 2023-12-03
### Changed
- 增加自旋锁
- update changelog

## [1.14.2] - 2023-12-03
### Changed
- 增加具有滑动窗口功能的WaitGroup
- update changelog

## [1.14.1] - 2023-12-03
### Changed
- 迁移gocsv从github.com/gocarina/gocsv到gitee.com/quant1x/pkg/gocsv
- update changelog

## [1.14.0] - 2023-11-27
### Changed
- 移除终端二维码工具库到pkg
- update changelog

## [1.13.9] - 2023-11-26
### Changed
- 更新依赖库版本
- update changelog

## [1.13.8] - 2023-11-23
### Changed
- 实验多时段定时任务
- 增加编译选项中调试开关
- update changelog
- 增加编译选项中调试开关
- 增加其它运行时需要的工具函数

## [1.13.7] - 2023-10-28
### Changed
- logger缓存map改为sync.map
- update changelog

## [1.13.6] - 2023-10-28
### Changed
- treemap的clear方法增加互斥锁
- update changelog

## [1.13.5] - 2023-10-27
### Changed
- 调整csv文件的关闭方式
- update changelog

## [1.13.4] - 2023-10-26
### Changed
- 定时任务增加重置日志
- update changelog

## [1.13.3] - 2023-10-22
### Changed
- 调整调度任务, 增加计时
- update changelog

## [1.13.2] - 2023-10-22
### Changed
- 增加获取当前代码的函数名, 文件名以及行号的函数
- 新增runtime包
- 增加获取func信息的函数
- 调整skip调度任务策略
- update changelog

## [1.13.1] - 2023-10-21
### Changed
- 增加应用退出等待机制
- update changelog

## [1.13.0] - 2023-10-20
### Changed
- 调整任务未执行完成跳过的函数封装
- update changelog

## [1.12.9] - 2023-10-20
### Changed
- 调整任务未执行完成跳过的函数封装
- update changelog

## [1.12.8] - 2023-10-19
### Changed
- 修订可延迟执行的定时调度组件
- update changelog

## [1.12.7] - 2023-10-19
### Changed
- 修订日期重置错乱的bug
- update changelog

## [1.12.6] - 2023-10-16
### Changed
- 修正周期初始化时间为9点整
- update changelog

## [1.12.5] - 2023-10-16
### Changed
- 调整周期初始化锁
- update changelog

## [1.12.4] - 2023-10-15
### Changed
- 调整函数名
- update changelog

## [1.12.3] - 2023-10-15
### Changed
- 删除废弃的代码
- 调整hashmap的代码, 引入github.com/orcaman/concurrent-map/v2
- 更新依赖版本
- update changelog

## [1.12.2] - 2023-10-15
### Changed
- 增加协程安全的hashmap
- update changelog

## [1.12.1] - 2023-10-10
### Changed
- 收录go-runewidth组件
- update changelog

## [1.12.0] - 2023-10-08
### Changed
- 增加not found判断
- update changelog

## [1.11.9] - 2023-10-07
### Changed
- 优化http client
- 优化http client参数
- 增加线程安全的TreeMap
- update changelog

## [1.11.8] - 2023-10-05
### Changed
- 新增embed封装函数
- update changelog

## [1.11.7] - 2023-10-01
### Changed
- 删除废弃的代码
- update changelog

## [1.11.6] - 2023-10-01
### Changed
- 优化fastqueue的push为异步方式
- update changelog

## [1.11.5] - 2023-09-29
### Changed
- 增加注释, 从1.12版本开始将移除MultiOnce
- 增加context.Context的封装
- 优化滑动窗口锁, 窗口期内只初始化一次
- update changelog

## [1.11.4] - 2023-09-29
### Changed
- 增加滑动窗口式的加载锁
- update changelog

## [1.11.3] - 2023-09-15
### Changed
- 调整homedir, 所有操作系统设置了GOX_HOME都会优先返回
- update changelog

## [1.11.2] - 2023-09-15
### Changed
- windows服务安装时创建系统环境变量GOX_HOME

## [1.11.1] - 2023-09-15
### Changed
- 启用环境变量GOX_HOME是为了Windows服务以系统账户运行时无法获取登录用户的宿主目录而预备的

## [1.11.0] - 2023-09-15
### Changed
- 去掉多余的import

## [1.10.9] - 2023-09-15
### Changed
- windows服务屏蔽使用本地用户登录, 本地用户登录有一个问题, 密码更换后会造成服务运行不正常
- 屏蔽检测连接池已打开数量的日志

## [1.10.8] - 2023-09-15
### Changed
- windows 服务属性增加本地用户名

## [1.10.7] - 2023-09-13
### Changed
- 获取连接增加告警日志
- update changelog

## [1.10.6] - 2023-09-12
### Changed
- 更换golang.org/x/exp/slices为系统标准库
- update changelog

## [1.10.5] - 2023-09-10
### Changed
- 升级string和bytes转换函数
- 升级依赖库版本
- update changelog

## [1.10.4] - 2023-09-10
### Changed
- 升级string和bytes转换函数
- update changelog

## [1.10.3] - 2023-08-24
### Changed
- 增加linux cpu 序列号获取方式, 用第一块网卡的mac地址代替
- update changelog

## [1.10.2] - 2023-08-16
### Changed
- 修订进度条结束逻辑, 先复写进度条, 再结束
- update changelog

## [1.10.1] - 2023-08-13
### Changed
- 升级go版本到1.21.0
- update changelog

## [1.10.0] - 2023-08-02
### Changed
- 调整服务的运行顺序
- update changelog

## [1.9.9] - 2023-08-01
### Changed
- 服务组件去掉日志
- update changelog

## [1.9.8] - 2023-08-01
### Changed
- 增加日志初始化
- update changelog

## [1.9.7] - 2023-08-01
### Changed
- 增加日志
- update changelog

## [1.9.6] - 2023-08-01
### Changed
- 调整windows服务的运行方式
- update changelog

## [1.9.5] - 2023-07-21
### Changed
- 删除daemon的demo
- 恢复daemon的demo
- update changelog

## [1.9.4] - 2023-07-20
### Changed
- 新增daemon工具库
- update changelog

## [1.9.3] - 2023-07-08
### Changed
- 更新依赖库版本
- update changelog

## [1.9.2] - 2023-07-08
### Changed
- 修复季度编码的bug
- update changelog

## [1.9.1] - 2023-07-07
### Changed
- 判断float是否NaN
- update changelog

## [1.9.0] - 2023-07-06
### Changed
- 调整获取周、月开始和结束时间的函数
- 优化slice去重函数
- update changelog

## [1.8.9] - 2023-07-02
### Changed
- 修复文件句柄未关闭的bug
- 修复文件句柄未关闭的bug
- update changelog

## [1.8.8] - 2023-06-30
### Changed
- 修复死锁的bug
- update changelog

## [1.8.7] - 2023-06-30
### Changed
- 增加默认初始化日期函数
- update changelog

## [1.8.6] - 2023-06-30
### Changed
- 增加日期切换功能, 默认不开启
- update changelog

## [1.8.5] - 2023-06-29
### Changed
- 没必要加锁, 加锁是个多余的操作
- update changelog

## [1.8.4] - 2023-06-29
### Changed
- 重置计数器加锁
- update changelog

## [1.8.3] - 2023-06-27
### Changed
- 修复chanel阻塞的bug
- update changelog

## [1.8.2] - 2023-06-27
### Changed
- 修复死锁的bug
- update changelog

## [1.8.1] - 2023-06-27
### Changed
- 连接池增加关闭所有链接的方法
- update changelog

## [1.8.0] - 2023-06-27
### Changed
- 增加base64算法
- update changelog

## [1.7.9] - 2023-06-27
### Changed
- 调整package
- update changelog

## [1.7.8] - 2023-06-27
### Changed
- 增加CPU序列号的获取函数
- update changelog

## [1.7.7] - 2023-06-26
### Changed
- 新增项目内的消息队列
- update changelog

## [1.7.6] - 2023-06-24
### Changed
- 季度函数增加财报季的返回值
- update changelog

## [1.7.5] - 2023-06-21
### Changed
- 新增一个定时调度任务, 回调函数不会并发执行
- update changelog

## [1.7.4] - 2023-06-21
### Changed
- 增加计算指定日期的季度开始和结束时间函数
- update changelog

## [1.7.3] - 2023-06-17
### Changed
- 修复字符串数据反射结构体存在数组长度和字段数量不匹配导致数据缺失的bug
- update changelog

## [1.7.2] - 2023-06-16
### Changed
- 新增结构体tag的反射缓存
- update changelog

## [1.7.1] - 2023-06-16
### Changed
- 更新依赖库
- update changelog

## [1.7.0] - 2023-06-14
### Changed
- 新增slice range函数
- update changelog

## [1.6.9] - 2023-06-14
### Changed
- SliceUnique函数第一个参数约束为指针
- update changelog

## [1.6.8] - 2023-06-14
### Changed
- 新增文件系统相关的检测函数
- update changelog

## [1.6.7] - 2023-06-14
### Changed
- 新增slice和csv文件互转的函数
- update changelog

## [1.6.6] - 2023-06-14
### Changed
- 新增MultiOnce组件, 可以重置的Sync.Once
- update changelog

## [1.6.5] - 2023-06-13
### Changed
- 新增排序和去重两个新函数
- update changelog

## [1.6.4] - 2023-06-13
### Changed
- 修订部分警告信息
- 微调字符串数组转结构体的函数
- update changelog

## [1.6.3] - 2023-06-11
### Changed
- 修订Copy的两个入参, 限制必须是指针
- update changelog

## [1.6.2] - 2023-06-07
### Changed
- 修订部分util工具库
- update changelog

## [1.6.1] - 2023-06-06
### Changed
- 新增四舍五入, 日期时间等函数
- update changelog

## [1.6.0] - 2023-06-03
### Changed
- 调整vek目录为num
- update changelog

## [1.5.1] - 2023-05-13
### Changed
- 收录vek汇编工具库
- update changelog

## [1.5.0] - 2023-05-13
### Changed
- 迁移代码仓库到gitee
- update changelog

## [1.3.33] - 2023-05-12
### Changed
- 更新依赖库版本号
- update changelog

## [1.3.32] - 2023-05-12
### Changed
- 修订http客户端header中accept字段错误的bug
- update changelog

## [1.3.31] - 2023-05-11
### Changed
- 调整进度条检测机制
- update changelog

## [1.3.30] - 2023-05-10
### Changed
- 更新依赖库版本号
- update changelog

## [1.3.29] - 2023-05-07
### Changed
- 调整CheckFilepath参数名
- update changelog

## [1.3.28] - 2023-05-07
### Changed
- 删除早期的测试代码
- 调整array tag反射机制代码
- update changelog

## [1.3.27] - 2023-05-07
### Changed
- 调整git仓库同步脚本
- update changelog

## [1.3.26] - 2023-05-06
### Changed
- 增加切片唯一性排序
- update changelog

## [1.3.25] - 2023-04-26
### Changed
- update changelog
- 调整源文件名
- update changelog

## [1.3.24] - 2023-04-26
### Changed
- 增加slice 过滤函数

## [1.3.23] - 2023-04-24
### Changed
- treemap加锁
- update changelog

## [1.3.22] - 2023-04-23
### Changed
- 调整切片反转函数测试代码
- 调整bar代码
- update changelog

## [1.3.21] - 2023-04-23
### Changed
- 增加切片反转函数
- update changelog

## [1.3.20] - 2023-04-23
### Changed
- 提升bar更新速度
- 修复计时的bug
- update changelog

## [1.3.19] - 2023-04-23
### Changed
- 提升bar更新速度
- update changelog
- update changelog

## [1.3.18] - 2023-04-03
### Changed
- 收敛信号统一处理机制

## [1.3.17] - 2023-03-20
### Changed
- 修复时区的bug

## [1.3.16] - 2023-03-20
### Changed
- 优化http client

## [1.3.15] - 2023-03-18
### Changed
- 警告不推荐使用SetLogPath

## [1.3.14] - 2023-03-18
### Changed
- 增加字符串判断空的函数
- logger增加初始化函数

## [1.3.13] - 2023-03-18
### Changed
- 判断closer是否为nil

## [1.3.12] - 2023-03-18
### Changed
- 增加判断是否debug模式

## [1.3.11] - 2023-03-18
### Changed
- 增加日志组件退出时可能出现panic

## [1.3.10] - 2023-03-09
### Changed
- 增加检测文件路径的函数

## [1.3.9] - 2023-02-28
### Changed
- 修订cron版本

## [1.3.8] - 2023-02-23
### Changed
- 收录github.com/modern-go/reflect2, 原作者已不维护更新
- 收录github.com/modern-go/reflect2, 原作者已不维护更新
- 剔除redis的工具库

## [1.3.7] - 2023-02-21
### Changed
- 屏蔽主进程关闭, 进度条等着结束信号是产生的panic

## [1.3.6] - 2023-02-19
### Changed
- 增加信号对操作系统的识别
- 增加信号对操作系统的识别

## [1.3.5] - 2023-02-19
### Changed
- 增加progressbar工具, github.com/qianlnk/pgbar, 原库有个问题, 就是不能定位在同一行进行循环展示进度条, 原因是maxline一直在增加

## [1.3.4] - 2023-02-19
### Changed
- 增加两个判断前后缀的函数
- 去除部分废弃的代码

## [1.3.3] - 2023-02-18
### Changed
- 增加判断chan是否关闭的函数
- 增加exception

## [1.3.2] - 2023-02-18
### Changed
- 修订ioutil.readall函数, 调整到io.readall, 优化部分closer接口

## [1.3.1] - 2023-01-29
### Changed
- 增加推送所有tag的命令
- 修订主机地址错误的问题
- 修订README

## [1.3.0] - 2023-01-29
### Changed
- 增加gitee和github两个仓库的代码同步脚本
- 在"set -e"之后出现的代码，一旦出现了返回值非零，整个脚本就会立即退出
- 推送后显示当前的远程仓库地址

## [1.2.7] - 2023-01-29
### Changed
- 不常用的package的归档在labs
- 改变package, 将c-struct和struc并入encoding/binary中
- 调整包路径
- 整理部分package
- 去掉多余的转换

## [1.2.6] - 2023-01-27
### Changed
- 增加一对字节数组和字符串互转的函数

## [1.2.5] - 2023-01-24
### Changed
- import github.com/mitchellh/go-homedir

## [1.2.4] - 2023-01-15
### Changed
- 增加lambda工具包

## [1.2.3] - 2023-01-15
### Changed
- c struct增加测试用例

## [1.2.2] - 2023-01-15
### Changed
- fork github.com/fananchong/cstruct-go

## [1.2.1] - 2023-01-14
### Changed
- 去掉rune内建关键字的警告
- 去掉rune内建关键字的警告
- 更新pool版本
- 新增struc工具包

## [1.2.0] - 2023-01-05
### Changed
- 修订支持1.20

## [1.1.21] - 2021-08-23
### Changed
- 增加string书写风格转换

## [1.1.20] - 2021-08-21
### Changed
- 增加http组件

## [1.1.19] - 2021-07-29
### Changed
- 调整package

## [1.1.18] - 2021-07-29
### Changed
- 改变常量的目录

## [1.1.17] - 2021-07-28
### Changed
- 调整目录结构
- 增加数组相关的函数

## [1.1.16] - 2021-07-26
### Changed
- 增加并发hashmap
- 增加测试代码
- 增加单测

## [1.1.15] - 2021-07-16
### Changed
- 恢复gls的单测

## [1.1.14] - 2021-07-16
### Changed
- Merge branch 'master' of https://github.com/mymmsc/gox
- 调整travis 从org转到com
- 修订GO111MODULE为auto

## [1.1.13] - 2021-07-15
### Changed
- V1.1.x (#1)

* 修订README

* 增加go:nocheckptr

* fix: README
- 修订README
- 增加go:nocheckptr
- fix: README

## [1.1.12] - 2021-07-15
### Changed
- 修订代码检测流水

## [1.1.11] - 2021-07-15
### Changed
- 调整travis

## [1.1.10] - 2021-07-15
### Changed
- 修订orm框架的引用
- 修订依赖组件
- 暂时通过改单测golang原文件名_test.go结尾的方式屏蔽有问题的测试
- 修订组件依赖

## [1.1.9] - 2021-07-10
### Changed
- 增加struct拷贝功能

## [1.1.8] - 2021-07-07
### Changed
- 整型忽略非数字部分

## [1.1.7] - 2021-07-07
### Changed
- 整型忽略非数字部分

## [1.1.6] - 2021-07-07
### Changed
- 整理目录

## [1.1.5] - 2021-07-07
### Changed
- 修订方法名

## [1.1.3] - 2021-07-07
### Changed
- 去掉没有引用的package
- 增加忽略异常关闭I/O
- 调整基础工具的package
- 增加安全的字符串转数值的方法
- 增加数值转字符串的方法

## [1.1.2] - 2021-06-29
### Changed
- add AOP
- 修改golang版本
- bou.ke/monkey的目录结构
- 调整package路径

## [1.1.1] - 2020-08-01
### Changed
- add encoding

## [1.1.0] - 2020-08-01
### Changed
- fix golang version

## [1.0.28] - 2020-03-30
### Changed
- fix version

## [1.0.27] - 2019-08-17
### Changed
- fix init

## [1.0.26] - 2019-08-17
### Changed
- fix init

## [1.0.25] - 2019-08-17
### Changed
- fix MDC

## [1.0.24] - 2019-08-14
### Changed
- fix logs

## [1.0.23] - 2019-08-10
### Changed
- fix package

## [1.0.22] - 2019-08-10
### Changed
- fix package

## [1.0.21] - 2019-08-10
### Changed
- add 增加滑动窗口式的waitgroup

## [1.0.20] - 2019-05-09
### Changed
- fix fastjson data

## [1.0.19] - 2019-05-09
### Changed
- 调整slf4g包名

## [1.0.18] - 2019-05-09
### Changed
- 格式化代码
- 调整slf4g包名

## [1.0.17] - 2019-04-23
### Changed
- fix package
- 增加测试中文
- fix
- fix sign
- Merge branch 'v1.0.x' of https://github.com/mymmsc/gox into v1.0.x
- fix 备份历史文件的日期

## [1.0.16] - 2019-04-14
### Changed
- fix go test

## [1.0.15] - 2019-04-14
### Changed
- add github.com/ddliu/go-httpclient

## [1.0.14] - 2019-04-14
### Changed
- add github.com/robfig/cron

## [1.0.13] - 2019-04-13
### Changed
- fix package

## [1.0.12] - 2019-04-13
### Changed
- add GetString

## [1.0.11] - 2019-04-13
### Changed
- add fastjson

## [1.0.10] - 2019-04-12
### Changed
- add filetype
- delete filetype
- fix timestamp

## [1.0.9] - 2019-04-05
### Changed
- fix util error

## [1.0.8] - 2019-04-05
### Changed
- 修改仓库地址
- fix travis
- fix travis
- fix util
- fix package
- fix package
- fix
- fix package
- fix README
- add github.com/emirpasic/gods README
- fix package

## [1.0.7] - 2019-03-31
### Changed
- 忽略go test测试结果
- 调整go module依赖
- change to execute
- 调整测试脚本
- 调整测试脚本
- fix module list

## [1.0.6] - 2019-03-31
### Changed
- add thread local
- 删除废弃的代码

## [1.0.5] - 2019-03-30
### Changed
- 删除golang.org/x/sys的引用

## [1.0.4] - 2019-03-30
### Changed
- add logger

## [1.0.3] - 2019-03-30
### Changed
- add redis pool

## [1.0.2] - 2019-03-24
### Changed
- add testing

## [1.0.1] - 2019-03-24
### Changed
- fix version
- fix version

## [1.0.0] - 2019-03-24
### Changed
- Initial commit
- add gitingore
- add support go module
- add QR terminal
- 修改二维码路径
- add go-xorm mysql template
- add go-xorm 反转数据库结构 用法
- add travis-ci
- fix travis
- fix
- fix
- fix
- fix
- fix golang version
- fix travis
- fix travis
- fix travis
- fix license
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- fix codecov
- add test code
- add test code
- fix codecov
- fix codecov
- fix license
- fix codecov


[Unreleased]: https://gitee.com/quant1x/gox.git/compare/v1.24.7...HEAD
[1.24.7]: https://gitee.com/quant1x/gox.git/compare/v1.24.6...v1.24.7
[1.24.6]: https://gitee.com/quant1x/gox.git/compare/v1.24.5...v1.24.6
[1.24.5]: https://gitee.com/quant1x/gox.git/compare/v1.24.4...v1.24.5
[1.24.4]: https://gitee.com/quant1x/gox.git/compare/v1.24.3...v1.24.4
[1.24.3]: https://gitee.com/quant1x/gox.git/compare/v1.24.2...v1.24.3
[1.24.2]: https://gitee.com/quant1x/gox.git/compare/v1.24.1...v1.24.2
[1.24.1]: https://gitee.com/quant1x/gox.git/compare/v1.24.0...v1.24.1
[1.24.0]: https://gitee.com/quant1x/gox.git/compare/v1.23.2...v1.24.0
[1.23.2]: https://gitee.com/quant1x/gox.git/compare/v1.23.1...v1.23.2
[1.23.1]: https://gitee.com/quant1x/gox.git/compare/v1.23.0...v1.23.1
[1.23.0]: https://gitee.com/quant1x/gox.git/compare/v1.22.13...v1.23.0
[1.22.13]: https://gitee.com/quant1x/gox.git/compare/v1.22.12...v1.22.13
[1.22.12]: https://gitee.com/quant1x/gox.git/compare/v1.22.11...v1.22.12
[1.22.11]: https://gitee.com/quant1x/gox.git/compare/v1.22.10...v1.22.11
[1.22.10]: https://gitee.com/quant1x/gox.git/compare/v1.22.9...v1.22.10
[1.22.9]: https://gitee.com/quant1x/gox.git/compare/v1.22.8...v1.22.9
[1.22.8]: https://gitee.com/quant1x/gox.git/compare/v1.22.7...v1.22.8
[1.22.7]: https://gitee.com/quant1x/gox.git/compare/v1.22.6...v1.22.7
[1.22.6]: https://gitee.com/quant1x/gox.git/compare/v1.22.5...v1.22.6
[1.22.5]: https://gitee.com/quant1x/gox.git/compare/v1.22.4...v1.22.5
[1.22.4]: https://gitee.com/quant1x/gox.git/compare/v1.22.3...v1.22.4
[1.22.3]: https://gitee.com/quant1x/gox.git/compare/v1.22.2...v1.22.3
[1.22.2]: https://gitee.com/quant1x/gox.git/compare/v1.22.1...v1.22.2
[1.22.1]: https://gitee.com/quant1x/gox.git/compare/v1.22.0...v1.22.1
[1.22.0]: https://gitee.com/quant1x/gox.git/compare/v1.21.9...v1.22.0
[1.21.9]: https://gitee.com/quant1x/gox.git/compare/v1.21.8...v1.21.9
[1.21.8]: https://gitee.com/quant1x/gox.git/compare/v1.21.7...v1.21.8
[1.21.7]: https://gitee.com/quant1x/gox.git/compare/v1.21.6...v1.21.7
[1.21.6]: https://gitee.com/quant1x/gox.git/compare/v1.21.5...v1.21.6
[1.21.5]: https://gitee.com/quant1x/gox.git/compare/v1.21.4...v1.21.5
[1.21.4]: https://gitee.com/quant1x/gox.git/compare/v1.21.3...v1.21.4
[1.21.3]: https://gitee.com/quant1x/gox.git/compare/v1.21.2...v1.21.3
[1.21.2]: https://gitee.com/quant1x/gox.git/compare/v1.21.1...v1.21.2
[1.21.1]: https://gitee.com/quant1x/gox.git/compare/v1.21.0...v1.21.1
[1.21.0]: https://gitee.com/quant1x/gox.git/compare/v1.20.9...v1.21.0
[1.20.9]: https://gitee.com/quant1x/gox.git/compare/v1.20.8...v1.20.9
[1.20.8]: https://gitee.com/quant1x/gox.git/compare/v1.20.7...v1.20.8
[1.20.7]: https://gitee.com/quant1x/gox.git/compare/v1.20.6...v1.20.7
[1.20.6]: https://gitee.com/quant1x/gox.git/compare/v1.20.5...v1.20.6
[1.20.5]: https://gitee.com/quant1x/gox.git/compare/v1.20.4...v1.20.5
[1.20.4]: https://gitee.com/quant1x/gox.git/compare/v1.20.3...v1.20.4
[1.20.3]: https://gitee.com/quant1x/gox.git/compare/v1.20.2...v1.20.3
[1.20.2]: https://gitee.com/quant1x/gox.git/compare/v1.20.1...v1.20.2
[1.20.1]: https://gitee.com/quant1x/gox.git/compare/v1.20.0...v1.20.1
[1.20.0]: https://gitee.com/quant1x/gox.git/compare/v1.19.9...v1.20.0
[1.19.9]: https://gitee.com/quant1x/gox.git/compare/v1.19.8...v1.19.9
[1.19.8]: https://gitee.com/quant1x/gox.git/compare/v1.19.7...v1.19.8
[1.19.7]: https://gitee.com/quant1x/gox.git/compare/v1.19.6...v1.19.7
[1.19.6]: https://gitee.com/quant1x/gox.git/compare/v1.19.5...v1.19.6
[1.19.5]: https://gitee.com/quant1x/gox.git/compare/v1.19.4...v1.19.5
[1.19.4]: https://gitee.com/quant1x/gox.git/compare/v1.19.3...v1.19.4
[1.19.3]: https://gitee.com/quant1x/gox.git/compare/v1.19.2...v1.19.3
[1.19.2]: https://gitee.com/quant1x/gox.git/compare/v1.19.1...v1.19.2
[1.19.1]: https://gitee.com/quant1x/gox.git/compare/v1.19.0...v1.19.1
[1.19.0]: https://gitee.com/quant1x/gox.git/compare/v1.18.9...v1.19.0
[1.18.9]: https://gitee.com/quant1x/gox.git/compare/v1.18.8...v1.18.9
[1.18.8]: https://gitee.com/quant1x/gox.git/compare/v1.18.7...v1.18.8
[1.18.7]: https://gitee.com/quant1x/gox.git/compare/v1.18.6...v1.18.7
[1.18.6]: https://gitee.com/quant1x/gox.git/compare/v1.18.5...v1.18.6
[1.18.5]: https://gitee.com/quant1x/gox.git/compare/v1.18.4...v1.18.5
[1.18.4]: https://gitee.com/quant1x/gox.git/compare/v1.18.3...v1.18.4
[1.18.3]: https://gitee.com/quant1x/gox.git/compare/v1.18.2...v1.18.3
[1.18.2]: https://gitee.com/quant1x/gox.git/compare/v1.18.1...v1.18.2
[1.18.1]: https://gitee.com/quant1x/gox.git/compare/v1.18.0...v1.18.1
[1.18.0]: https://gitee.com/quant1x/gox.git/compare/v1.17.9...v1.18.0
[1.17.9]: https://gitee.com/quant1x/gox.git/compare/v1.17.8...v1.17.9
[1.17.8]: https://gitee.com/quant1x/gox.git/compare/v1.17.7...v1.17.8
[1.17.7]: https://gitee.com/quant1x/gox.git/compare/v1.17.6...v1.17.7
[1.17.6]: https://gitee.com/quant1x/gox.git/compare/v1.17.5...v1.17.6
[1.17.5]: https://gitee.com/quant1x/gox.git/compare/v1.17.4...v1.17.5
[1.17.4]: https://gitee.com/quant1x/gox.git/compare/v1.17.3...v1.17.4
[1.17.3]: https://gitee.com/quant1x/gox.git/compare/v1.17.2...v1.17.3
[1.17.2]: https://gitee.com/quant1x/gox.git/compare/v1.17.1...v1.17.2
[1.17.1]: https://gitee.com/quant1x/gox.git/compare/v1.17.0...v1.17.1
[1.17.0]: https://gitee.com/quant1x/gox.git/compare/v1.16.9...v1.17.0
[1.16.9]: https://gitee.com/quant1x/gox.git/compare/v1.16.8...v1.16.9
[1.16.8]: https://gitee.com/quant1x/gox.git/compare/v1.16.7...v1.16.8
[1.16.7]: https://gitee.com/quant1x/gox.git/compare/v1.16.6...v1.16.7
[1.16.6]: https://gitee.com/quant1x/gox.git/compare/v1.16.5...v1.16.6
[1.16.5]: https://gitee.com/quant1x/gox.git/compare/v1.16.4...v1.16.5
[1.16.4]: https://gitee.com/quant1x/gox.git/compare/v1.16.3...v1.16.4
[1.16.3]: https://gitee.com/quant1x/gox.git/compare/v1.16.2...v1.16.3
[1.16.2]: https://gitee.com/quant1x/gox.git/compare/v1.16.1...v1.16.2
[1.16.1]: https://gitee.com/quant1x/gox.git/compare/v1.16.0...v1.16.1
[1.16.0]: https://gitee.com/quant1x/gox.git/compare/v1.15.9...v1.16.0
[1.15.9]: https://gitee.com/quant1x/gox.git/compare/v1.15.8...v1.15.9
[1.15.8]: https://gitee.com/quant1x/gox.git/compare/v1.15.7...v1.15.8
[1.15.7]: https://gitee.com/quant1x/gox.git/compare/v1.15.6...v1.15.7
[1.15.6]: https://gitee.com/quant1x/gox.git/compare/v1.15.5...v1.15.6
[1.15.5]: https://gitee.com/quant1x/gox.git/compare/v1.15.4...v1.15.5
[1.15.4]: https://gitee.com/quant1x/gox.git/compare/v1.15.3...v1.15.4
[1.15.3]: https://gitee.com/quant1x/gox.git/compare/v1.15.2...v1.15.3
[1.15.2]: https://gitee.com/quant1x/gox.git/compare/v1.15.1...v1.15.2
[1.15.1]: https://gitee.com/quant1x/gox.git/compare/v1.15.0...v1.15.1
[1.15.0]: https://gitee.com/quant1x/gox.git/compare/v1.14.9...v1.15.0
[1.14.9]: https://gitee.com/quant1x/gox.git/compare/v1.14.8...v1.14.9
[1.14.8]: https://gitee.com/quant1x/gox.git/compare/v1.14.7...v1.14.8
[1.14.7]: https://gitee.com/quant1x/gox.git/compare/v1.14.6...v1.14.7
[1.14.6]: https://gitee.com/quant1x/gox.git/compare/v1.14.5...v1.14.6
[1.14.5]: https://gitee.com/quant1x/gox.git/compare/v1.14.4...v1.14.5
[1.14.4]: https://gitee.com/quant1x/gox.git/compare/v1.14.3...v1.14.4
[1.14.3]: https://gitee.com/quant1x/gox.git/compare/v1.14.2...v1.14.3
[1.14.2]: https://gitee.com/quant1x/gox.git/compare/v1.14.1...v1.14.2
[1.14.1]: https://gitee.com/quant1x/gox.git/compare/v1.14.0...v1.14.1
[1.14.0]: https://gitee.com/quant1x/gox.git/compare/v1.13.9...v1.14.0
[1.13.9]: https://gitee.com/quant1x/gox.git/compare/v1.13.8...v1.13.9
[1.13.8]: https://gitee.com/quant1x/gox.git/compare/v1.13.7...v1.13.8
[1.13.7]: https://gitee.com/quant1x/gox.git/compare/v1.13.6...v1.13.7
[1.13.6]: https://gitee.com/quant1x/gox.git/compare/v1.13.5...v1.13.6
[1.13.5]: https://gitee.com/quant1x/gox.git/compare/v1.13.4...v1.13.5
[1.13.4]: https://gitee.com/quant1x/gox.git/compare/v1.13.3...v1.13.4
[1.13.3]: https://gitee.com/quant1x/gox.git/compare/v1.13.2...v1.13.3
[1.13.2]: https://gitee.com/quant1x/gox.git/compare/v1.13.1...v1.13.2
[1.13.1]: https://gitee.com/quant1x/gox.git/compare/v1.13.0...v1.13.1
[1.13.0]: https://gitee.com/quant1x/gox.git/compare/v1.12.9...v1.13.0
[1.12.9]: https://gitee.com/quant1x/gox.git/compare/v1.12.8...v1.12.9
[1.12.8]: https://gitee.com/quant1x/gox.git/compare/v1.12.7...v1.12.8
[1.12.7]: https://gitee.com/quant1x/gox.git/compare/v1.12.6...v1.12.7
[1.12.6]: https://gitee.com/quant1x/gox.git/compare/v1.12.5...v1.12.6
[1.12.5]: https://gitee.com/quant1x/gox.git/compare/v1.12.4...v1.12.5
[1.12.4]: https://gitee.com/quant1x/gox.git/compare/v1.12.3...v1.12.4
[1.12.3]: https://gitee.com/quant1x/gox.git/compare/v1.12.2...v1.12.3
[1.12.2]: https://gitee.com/quant1x/gox.git/compare/v1.12.1...v1.12.2
[1.12.1]: https://gitee.com/quant1x/gox.git/compare/v1.12.0...v1.12.1
[1.12.0]: https://gitee.com/quant1x/gox.git/compare/v1.11.9...v1.12.0
[1.11.9]: https://gitee.com/quant1x/gox.git/compare/v1.11.8...v1.11.9
[1.11.8]: https://gitee.com/quant1x/gox.git/compare/v1.11.7...v1.11.8
[1.11.7]: https://gitee.com/quant1x/gox.git/compare/v1.11.6...v1.11.7
[1.11.6]: https://gitee.com/quant1x/gox.git/compare/v1.11.5...v1.11.6
[1.11.5]: https://gitee.com/quant1x/gox.git/compare/v1.11.4...v1.11.5
[1.11.4]: https://gitee.com/quant1x/gox.git/compare/v1.11.3...v1.11.4
[1.11.3]: https://gitee.com/quant1x/gox.git/compare/v1.11.2...v1.11.3
[1.11.2]: https://gitee.com/quant1x/gox.git/compare/v1.11.1...v1.11.2
[1.11.1]: https://gitee.com/quant1x/gox.git/compare/v1.11.0...v1.11.1
[1.11.0]: https://gitee.com/quant1x/gox.git/compare/v1.10.9...v1.11.0
[1.10.9]: https://gitee.com/quant1x/gox.git/compare/v1.10.8...v1.10.9
[1.10.8]: https://gitee.com/quant1x/gox.git/compare/v1.10.7...v1.10.8
[1.10.7]: https://gitee.com/quant1x/gox.git/compare/v1.10.6...v1.10.7
[1.10.6]: https://gitee.com/quant1x/gox.git/compare/v1.10.5...v1.10.6
[1.10.5]: https://gitee.com/quant1x/gox.git/compare/v1.10.4...v1.10.5
[1.10.4]: https://gitee.com/quant1x/gox.git/compare/v1.10.3...v1.10.4
[1.10.3]: https://gitee.com/quant1x/gox.git/compare/v1.10.2...v1.10.3
[1.10.2]: https://gitee.com/quant1x/gox.git/compare/v1.10.1...v1.10.2
[1.10.1]: https://gitee.com/quant1x/gox.git/compare/v1.10.0...v1.10.1
[1.10.0]: https://gitee.com/quant1x/gox.git/compare/v1.9.9...v1.10.0
[1.9.9]: https://gitee.com/quant1x/gox.git/compare/v1.9.8...v1.9.9
[1.9.8]: https://gitee.com/quant1x/gox.git/compare/v1.9.7...v1.9.8
[1.9.7]: https://gitee.com/quant1x/gox.git/compare/v1.9.6...v1.9.7
[1.9.6]: https://gitee.com/quant1x/gox.git/compare/v1.9.5...v1.9.6
[1.9.5]: https://gitee.com/quant1x/gox.git/compare/v1.9.4...v1.9.5
[1.9.4]: https://gitee.com/quant1x/gox.git/compare/v1.9.3...v1.9.4
[1.9.3]: https://gitee.com/quant1x/gox.git/compare/v1.9.2...v1.9.3
[1.9.2]: https://gitee.com/quant1x/gox.git/compare/v1.9.1...v1.9.2
[1.9.1]: https://gitee.com/quant1x/gox.git/compare/v1.9.0...v1.9.1
[1.9.0]: https://gitee.com/quant1x/gox.git/compare/v1.8.9...v1.9.0
[1.8.9]: https://gitee.com/quant1x/gox.git/compare/v1.8.8...v1.8.9
[1.8.8]: https://gitee.com/quant1x/gox.git/compare/v1.8.7...v1.8.8
[1.8.7]: https://gitee.com/quant1x/gox.git/compare/v1.8.6...v1.8.7
[1.8.6]: https://gitee.com/quant1x/gox.git/compare/v1.8.5...v1.8.6
[1.8.5]: https://gitee.com/quant1x/gox.git/compare/v1.8.4...v1.8.5
[1.8.4]: https://gitee.com/quant1x/gox.git/compare/v1.8.3...v1.8.4
[1.8.3]: https://gitee.com/quant1x/gox.git/compare/v1.8.2...v1.8.3
[1.8.2]: https://gitee.com/quant1x/gox.git/compare/v1.8.1...v1.8.2
[1.8.1]: https://gitee.com/quant1x/gox.git/compare/v1.8.0...v1.8.1
[1.8.0]: https://gitee.com/quant1x/gox.git/compare/v1.7.9...v1.8.0
[1.7.9]: https://gitee.com/quant1x/gox.git/compare/v1.7.8...v1.7.9
[1.7.8]: https://gitee.com/quant1x/gox.git/compare/v1.7.7...v1.7.8
[1.7.7]: https://gitee.com/quant1x/gox.git/compare/v1.7.6...v1.7.7
[1.7.6]: https://gitee.com/quant1x/gox.git/compare/v1.7.5...v1.7.6
[1.7.5]: https://gitee.com/quant1x/gox.git/compare/v1.7.4...v1.7.5
[1.7.4]: https://gitee.com/quant1x/gox.git/compare/v1.7.3...v1.7.4
[1.7.3]: https://gitee.com/quant1x/gox.git/compare/v1.7.2...v1.7.3
[1.7.2]: https://gitee.com/quant1x/gox.git/compare/v1.7.1...v1.7.2
[1.7.1]: https://gitee.com/quant1x/gox.git/compare/v1.7.0...v1.7.1
[1.7.0]: https://gitee.com/quant1x/gox.git/compare/v1.6.9...v1.7.0
[1.6.9]: https://gitee.com/quant1x/gox.git/compare/v1.6.8...v1.6.9
[1.6.8]: https://gitee.com/quant1x/gox.git/compare/v1.6.7...v1.6.8
[1.6.7]: https://gitee.com/quant1x/gox.git/compare/v1.6.6...v1.6.7
[1.6.6]: https://gitee.com/quant1x/gox.git/compare/v1.6.5...v1.6.6
[1.6.5]: https://gitee.com/quant1x/gox.git/compare/v1.6.4...v1.6.5
[1.6.4]: https://gitee.com/quant1x/gox.git/compare/v1.6.3...v1.6.4
[1.6.3]: https://gitee.com/quant1x/gox.git/compare/v1.6.2...v1.6.3
[1.6.2]: https://gitee.com/quant1x/gox.git/compare/v1.6.1...v1.6.2
[1.6.1]: https://gitee.com/quant1x/gox.git/compare/v1.6.0...v1.6.1
[1.6.0]: https://gitee.com/quant1x/gox.git/compare/v1.5.1...v1.6.0
[1.5.1]: https://gitee.com/quant1x/gox.git/compare/v1.5.0...v1.5.1
[1.5.0]: https://gitee.com/quant1x/gox.git/compare/v1.3.33...v1.5.0
[1.3.33]: https://gitee.com/quant1x/gox.git/compare/v1.3.32...v1.3.33
[1.3.32]: https://gitee.com/quant1x/gox.git/compare/v1.3.31...v1.3.32
[1.3.31]: https://gitee.com/quant1x/gox.git/compare/v1.3.30...v1.3.31
[1.3.30]: https://gitee.com/quant1x/gox.git/compare/v1.3.29...v1.3.30
[1.3.29]: https://gitee.com/quant1x/gox.git/compare/v1.3.28...v1.3.29
[1.3.28]: https://gitee.com/quant1x/gox.git/compare/v1.3.27...v1.3.28
[1.3.27]: https://gitee.com/quant1x/gox.git/compare/v1.3.26...v1.3.27
[1.3.26]: https://gitee.com/quant1x/gox.git/compare/v1.3.25...v1.3.26
[1.3.25]: https://gitee.com/quant1x/gox.git/compare/v1.3.24...v1.3.25
[1.3.24]: https://gitee.com/quant1x/gox.git/compare/v1.3.23...v1.3.24
[1.3.23]: https://gitee.com/quant1x/gox.git/compare/v1.3.22...v1.3.23
[1.3.22]: https://gitee.com/quant1x/gox.git/compare/v1.3.21...v1.3.22
[1.3.21]: https://gitee.com/quant1x/gox.git/compare/v1.3.20...v1.3.21
[1.3.20]: https://gitee.com/quant1x/gox.git/compare/v1.3.19...v1.3.20
[1.3.19]: https://gitee.com/quant1x/gox.git/compare/v1.3.18...v1.3.19
[1.3.18]: https://gitee.com/quant1x/gox.git/compare/v1.3.17...v1.3.18
[1.3.17]: https://gitee.com/quant1x/gox.git/compare/v1.3.16...v1.3.17
[1.3.16]: https://gitee.com/quant1x/gox.git/compare/v1.3.15...v1.3.16
[1.3.15]: https://gitee.com/quant1x/gox.git/compare/v1.3.14...v1.3.15
[1.3.14]: https://gitee.com/quant1x/gox.git/compare/v1.3.13...v1.3.14
[1.3.13]: https://gitee.com/quant1x/gox.git/compare/v1.3.12...v1.3.13
[1.3.12]: https://gitee.com/quant1x/gox.git/compare/v1.3.11...v1.3.12
[1.3.11]: https://gitee.com/quant1x/gox.git/compare/v1.3.10...v1.3.11
[1.3.10]: https://gitee.com/quant1x/gox.git/compare/v1.3.9...v1.3.10
[1.3.9]: https://gitee.com/quant1x/gox.git/compare/v1.3.8...v1.3.9
[1.3.8]: https://gitee.com/quant1x/gox.git/compare/v1.3.7...v1.3.8
[1.3.7]: https://gitee.com/quant1x/gox.git/compare/v1.3.6...v1.3.7
[1.3.6]: https://gitee.com/quant1x/gox.git/compare/v1.3.5...v1.3.6
[1.3.5]: https://gitee.com/quant1x/gox.git/compare/v1.3.4...v1.3.5
[1.3.4]: https://gitee.com/quant1x/gox.git/compare/v1.3.3...v1.3.4
[1.3.3]: https://gitee.com/quant1x/gox.git/compare/v1.3.2...v1.3.3
[1.3.2]: https://gitee.com/quant1x/gox.git/compare/v1.3.1...v1.3.2
[1.3.1]: https://gitee.com/quant1x/gox.git/compare/v1.3.0...v1.3.1
[1.3.0]: https://gitee.com/quant1x/gox.git/compare/v1.2.7...v1.3.0
[1.2.7]: https://gitee.com/quant1x/gox.git/compare/v1.2.6...v1.2.7
[1.2.6]: https://gitee.com/quant1x/gox.git/compare/v1.2.5...v1.2.6
[1.2.5]: https://gitee.com/quant1x/gox.git/compare/v1.2.4...v1.2.5
[1.2.4]: https://gitee.com/quant1x/gox.git/compare/v1.2.3...v1.2.4
[1.2.3]: https://gitee.com/quant1x/gox.git/compare/v1.2.2...v1.2.3
[1.2.2]: https://gitee.com/quant1x/gox.git/compare/v1.2.1...v1.2.2
[1.2.1]: https://gitee.com/quant1x/gox.git/compare/v1.2.0...v1.2.1
[1.2.0]: https://gitee.com/quant1x/gox.git/compare/v1.1.21...v1.2.0
[1.1.21]: https://gitee.com/quant1x/gox.git/compare/v1.1.20...v1.1.21
[1.1.20]: https://gitee.com/quant1x/gox.git/compare/v1.1.19...v1.1.20
[1.1.19]: https://gitee.com/quant1x/gox.git/compare/v1.1.18...v1.1.19
[1.1.18]: https://gitee.com/quant1x/gox.git/compare/v1.1.17...v1.1.18
[1.1.17]: https://gitee.com/quant1x/gox.git/compare/v1.1.16...v1.1.17
[1.1.16]: https://gitee.com/quant1x/gox.git/compare/v1.1.15...v1.1.16
[1.1.15]: https://gitee.com/quant1x/gox.git/compare/v1.1.14...v1.1.15
[1.1.14]: https://gitee.com/quant1x/gox.git/compare/v1.1.13...v1.1.14
[1.1.13]: https://gitee.com/quant1x/gox.git/compare/v1.1.12...v1.1.13
[1.1.12]: https://gitee.com/quant1x/gox.git/compare/v1.1.11...v1.1.12
[1.1.11]: https://gitee.com/quant1x/gox.git/compare/v1.1.10...v1.1.11
[1.1.10]: https://gitee.com/quant1x/gox.git/compare/v1.1.9...v1.1.10
[1.1.9]: https://gitee.com/quant1x/gox.git/compare/v1.1.8...v1.1.9
[1.1.8]: https://gitee.com/quant1x/gox.git/compare/v1.1.7...v1.1.8
[1.1.7]: https://gitee.com/quant1x/gox.git/compare/v1.1.6...v1.1.7
[1.1.6]: https://gitee.com/quant1x/gox.git/compare/v1.1.5...v1.1.6
[1.1.5]: https://gitee.com/quant1x/gox.git/compare/v1.1.3...v1.1.5
[1.1.3]: https://gitee.com/quant1x/gox.git/compare/v1.1.2...v1.1.3
[1.1.2]: https://gitee.com/quant1x/gox.git/compare/v1.1.1...v1.1.2
[1.1.1]: https://gitee.com/quant1x/gox.git/compare/v1.1.0...v1.1.1
[1.1.0]: https://gitee.com/quant1x/gox.git/compare/v1.0.28...v1.1.0
[1.0.28]: https://gitee.com/quant1x/gox.git/compare/v1.0.27...v1.0.28
[1.0.27]: https://gitee.com/quant1x/gox.git/compare/v1.0.26...v1.0.27
[1.0.26]: https://gitee.com/quant1x/gox.git/compare/v1.0.25...v1.0.26
[1.0.25]: https://gitee.com/quant1x/gox.git/compare/v1.0.24...v1.0.25
[1.0.24]: https://gitee.com/quant1x/gox.git/compare/v1.0.23...v1.0.24
[1.0.23]: https://gitee.com/quant1x/gox.git/compare/v1.0.22...v1.0.23
[1.0.22]: https://gitee.com/quant1x/gox.git/compare/v1.0.21...v1.0.22
[1.0.21]: https://gitee.com/quant1x/gox.git/compare/v1.0.20...v1.0.21
[1.0.20]: https://gitee.com/quant1x/gox.git/compare/v1.0.19...v1.0.20
[1.0.19]: https://gitee.com/quant1x/gox.git/compare/v1.0.18...v1.0.19
[1.0.18]: https://gitee.com/quant1x/gox.git/compare/v1.0.17...v1.0.18
[1.0.17]: https://gitee.com/quant1x/gox.git/compare/v1.0.16...v1.0.17
[1.0.16]: https://gitee.com/quant1x/gox.git/compare/v1.0.15...v1.0.16
[1.0.15]: https://gitee.com/quant1x/gox.git/compare/v1.0.14...v1.0.15
[1.0.14]: https://gitee.com/quant1x/gox.git/compare/v1.0.13...v1.0.14
[1.0.13]: https://gitee.com/quant1x/gox.git/compare/v1.0.12...v1.0.13
[1.0.12]: https://gitee.com/quant1x/gox.git/compare/v1.0.11...v1.0.12
[1.0.11]: https://gitee.com/quant1x/gox.git/compare/v1.0.10...v1.0.11
[1.0.10]: https://gitee.com/quant1x/gox.git/compare/v1.0.9...v1.0.10
[1.0.9]: https://gitee.com/quant1x/gox.git/compare/v1.0.8...v1.0.9
[1.0.8]: https://gitee.com/quant1x/gox.git/compare/v1.0.7...v1.0.8
[1.0.7]: https://gitee.com/quant1x/gox.git/compare/v1.0.6...v1.0.7
[1.0.6]: https://gitee.com/quant1x/gox.git/compare/v1.0.5...v1.0.6
[1.0.5]: https://gitee.com/quant1x/gox.git/compare/v1.0.4...v1.0.5
[1.0.4]: https://gitee.com/quant1x/gox.git/compare/v1.0.3...v1.0.4
[1.0.3]: https://gitee.com/quant1x/gox.git/compare/v1.0.2...v1.0.3
[1.0.2]: https://gitee.com/quant1x/gox.git/compare/v1.0.1...v1.0.2
[1.0.1]: https://gitee.com/quant1x/gox.git/compare/v1.0.0...v1.0.1

[1.0.0]: https://gitee.com/quant1x/gox.git/releases/tag/v1.0.0
