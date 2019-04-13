# cron

Documentation here: https://godoc.org/github.com/robfig/cron

Cron 的定时任务表达式
```javascript
   # ┌───────────── second (0 - 59)
   # │ ┌───────────── min (0 - 59)
   # │ │ ┌────────────── hour (0 - 23)
   # │ │ │ ┌─────────────── day of month (1 - 31)
   # │ │ │ │ ┌──────────────── month (1 - 12)
   # │ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
   # │ │ │ │ │ │         Saturday, or use names; 7 is also Sunday)
   # │ │ │ │ │ │
   # │ │ │ │ │ │
   # * * * * * * command to execute
```

时间表达式例子
```
每隔 5 秒执行一次：*/5 * * * * ?
每隔 1 分钟执行一次：0 */1 * * * ?
每天 23 点执行一次：0 0 23 * * ?
每天凌晨 1 点执行一次：0 0 1 * * ?
每月 1 号凌晨 1 点执行一次：0 0 1 1 * ?
在 26 分、29 分、33 分执行一次：0 26,29,33 * * * ?
每天的 0 点、13 点、18 点、21 点都执行一次：0 0 0,13,18,21 * * ?
```