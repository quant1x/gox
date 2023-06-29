# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]

## [1.8.6] - 2023-06-30
### Changed
- 增加日期切换功能, 默认不开启.

## [1.8.5] - 2023-06-29
### Changed
- 没必要加锁, 加锁是个多余的操作.

## [1.8.4] - 2023-06-29
### Changed
- 重置计数器加锁.

## [1.8.3] - 2023-06-27
### Changed
- 修复chanel阻塞的bug.

## [1.8.2] - 2023-06-27
### Changed
- 修复死锁的bug.

## [1.8.1] - 2023-06-27
### Changed
- 连接池增加关闭所有链接的方法.

## [1.8.0] - 2023-06-27
### Changed
- 增加base64算法.

## [1.7.9] - 2023-06-27
### Changed
- 调整package.

## [1.7.8] - 2023-06-27
### Changed
- 增加CPU序列号的获取函数.

## [1.7.7] - 2023-06-26
### Changed
- 新增项目内的消息队列.

## [1.7.6] - 2023-06-24
### Changed
- 季度函数增加财报季的返回值.

## [1.7.5] - 2023-06-21
### Changed
- 新增一个定时调度任务, 回调函数不会并发执行.

## [1.7.4] - 2023-06-21
### Changed
- 增加计算指定日期的季度开始和结束时间函数.

## [1.7.3] - 2023-06-17
### Changed
- 修复字符串数据反射结构体存在数组长度和字段数量不匹配导致数据缺失的bug.

## [1.7.2] - 2023-06-16
### Changed
- 新增结构体tag的反射缓存.

## [1.7.1] - 2023-06-16
### Changed
- 更新依赖库.

## [1.7.0] - 2023-06-14
### Changed
- 新增slice range函数.

## [1.6.9] - 2023-06-14
### Changed
- SliceUnique函数第一个参数约束为指针.

## [1.6.8] - 2023-06-14
### Changed
- 新增文件系统相关的检测函数.

## [1.6.7] - 2023-06-14
### Changed
- 新增slice和csv文件互转的函数.

## [1.6.6] - 2023-06-14
### Changed
- 新增MultiOnce组件, 可以重置的Sync.Once.

## [1.6.5] - 2023-06-13
### Changed
- 新增排序和去重两个新函数.

## [1.6.4] - 2023-06-13
### Changed
- 微调字符串数组转结构体的函数.
- 修订部分警告信息.

## [1.6.3] - 2023-06-11
### Changed
- 修订Copy的两个入参, 限制必须是指针.

## [1.6.2] - 2023-06-07
### Changed
- 修订部分util工具库.

## [1.6.1] - 2023-06-06
### Changed
- 新增四舍五入, 日期时间等函数.

## [1.6.0] - 2023-06-03
### Changed
- 调整vek目录为num.

## [1.5.1] - 2023-05-13
### Changed
- 收录vek汇编工具库.

## [1.5.0] - 2023-05-13
### Changed
- 迁移代码仓库到gitee.

## [1.3.33] - 2023-05-12
### Changed
- 更新依赖库版本号.

## [1.3.32] - 2023-05-12
### Changed
- 修订http客户端header中accept字段错误的bug.

## [1.3.31] - 2023-05-11
### Changed
- 调整进度条检测机制.

## [1.3.30] - 2023-05-10
### Changed
- 更新依赖库版本号.

## [1.3.29] - 2023-05-07
### Changed
- 调整CheckFilepath参数名.

## [1.3.28] - 2023-05-07
### Changed
- 调整array tag反射机制代码.
- 删除早期的测试代码.

## [1.3.27] - 2023-05-07
### Changed
- 调整git仓库同步脚本.

## [1.3.26] - 2023-05-06
### Changed
- 增加切片唯一性排序.

## [1.3.25] - 2023-04-26
### Changed
- 调整源文件名.
- Update changelog.

## [1.3.24] - 2023-04-26
### Changed
- 增加slice 过滤函数.

## [1.3.23] - 2023-04-24
### Changed
- Treemap加锁.

## [1.3.22] - 2023-04-23
### Changed
- 调整bar代码.
- 调整切片反转函数测试代码.

## [1.3.21] - 2023-04-23
### Changed
- 增加切片反转函数.

## [1.3.20] - 2023-04-23
### Changed
- 修复计时的bug.
- 提升bar更新速度.

## [1.3.19] - 2023-04-23
### Changed
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
- Merge branch 'master' of https://gitee.com/quant1x/gox.
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
- Merge branch 'v1.0.x' of https://gitee.com/quant1x/gox into v1.0.x.
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

[Unreleased]: https://gitee.com/quant1x/gox/compare/v1.8.6...HEAD

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
