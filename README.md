# go-web-lesson

## lesson01 基于go默认的web
	使用go自带的 net/http 包

## lesson02 gin框架
	基本使用gin框架 学习路由

## lesson05  基于go默认的web
	原始net中的http使用template模板

## lesson06  基于go默认的web
	原始net中的http使用template模板嵌套

## lesson07  基于go默认的web
	原始net中的http使用template模板嵌套

## lesson08  基于go默认的web
	- 原始net中的http使用template模板嵌套
	- 更改标识符
    - 转义符号 如何 去除默认的脱义

## lesson09  基于gin中的模板渲染
	- r.LoadHTMLGlob("templates/**/*")
	- 学会如何解析和渲染模板

## lesson10  基于gin web
	学习如何使用gin 返回json数据
	- 方法1. 使用 gin.H 自己拼接json 其实就是map
	- 方法2. 使用结构体 还要学会灵活使用tag对结构体字段做定制化操作

## lesson26 GORM 删除
    - 软删除 引用了gorm.Model公用结构体的才有
    - 删除记录 根据主键id 
    - 删除记录 根据 where条件 
        db.Debug().Where("age >= ?", 10).Delete(&User{})
    - 物理删除 Unscoped().Where

## lesson25 JWT的生成和创建
    - createToken
    - ParseToken
## lesson26 中间件middleware JWT 
	- 学习如何使用JWT
	- CreateToken
	- ParseToken
	- 加BufferTime 原理

## lesson27 学习casbin
	数据来自mysql  https://github.com/casbin/gorm-adapter
	学习资料来自: https://casbin.org/docs/zh-CN/overview
	- 演示casbin基本使用
	- Adapter基于mysql, 策略存入mysql
	- 如何增加策略 删除策略

## lesson28 学习验证码
    资料来自 // https://github.com/mojocn/base64Captcha
    - 学习常见的我仲验证码 数字 字符串 数学验证码(加减乘除) 中文 语音
    - 学习数字验证码
    - 学会创建验证码和校验验证码 
    - 思考前后端如何交互****
    
## lesson29 debug学习
    资料来自 https://zhuanlan.zhihu.com/p/62610785
    什么是debug: 
    断点调试，英文 breakpoint。用大白话来解释下，断点调试其实就是在程序自动运行的过程中，
    你在代码某一处打上了断点，当程序跑到你设置的断点位置处，
    则会中断下来(等待你的按钮指令 怎么往下走)，此时你可以看到之前运行过的所有程序变量
    - 学会如何启动debug调试程序
    - 如何打断点
    - 断点按钮有哪几几种
    
    重点按钮
    - F8 step over 单步调试（出），不关心函数，只关心函数输出
    - F7 step into 单步调试（近），关心函数内部是什么，可能会进入源代码里面
    - Alt+F9 直接到下一个断点
    
    几个按钮看需求配合使用