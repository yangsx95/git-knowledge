# git-knowledge

- web框架： [echo](https://github.com/labstack/echo)
- 日志框架： [zap](https://github.com/uber-go/zap)
- ini配置解析： [gopkg.in/ini.v1](http://gopkg.in/ini.v1)
- github api客户端：[go-github](https://github.com/google/go-github)

# 原型/功能列表

概览： 
1. 一个用户可包含多个Space空间
2. 一个Space空间对应多个域名，以及一个 git-knowledge-config 配置仓库（名称可指定）
3. 配置仓库主要包含Space的配置 `.gitknowledge`
  ```ini
  ```
## 基础功能
- 注册
  - [x] 通过注册页面直接注册
    - [ ] 发送账户邮箱验证
    - [ ] 邮箱验证
    - [ ] 生成随机头像(头像源)
  - [ ] 通过Github登录并获取用户信息并直接注册
    - [ ] 自动生成userid以及password，并发送邮箱
    - [ ] 如果用户尚未修改初始密码，则发送消息通知提示用户修改密码
- 登录
  - [x] 通过ID登录
  - [x] 通过Email登录
  - [ ] 校验账号邮箱是否激活，如果未激活提示激活，阻止登录，并提示
    - [ ] Github注册登录无需校验邮箱
- 首页
  - [ ] 展示space空间列表
  - [ ] 增加添加空间的按钮
  - [ ] 支持空间的删除
  - [ ] 支持空间的创建
- [ ] 创建space
  - [ ] 关联指定的github仓库以及相应的分支
  - [ ] space的名称、图标
  - [ ] 读取仓库信息
    - 待定：是否需要clone仓库
    - 待定：是否要将仓库缓存到数据库中
    - 待定：是否直接调用github API展示内容
- [ ] Markdown渲染
  - [ ] 图片是否要进行缓存，还是直接调用github
  - [ ] SUMMARY.md 文件解析
  - [ ] SUMMARY.md 扩展
  - 待定：图片是否要进行缓存，还是直接使用github地址
- [ ] 编辑功能
  - [ ] 在线修改SUMMARY.md，完成排序、新增、删除、隐藏
  - [ ] 如何将修改提交至github
  - [ ] 如果出现冲突时该如何解决

## 扩展功能

- [ ] 支持Gitee
- [ ] 读Bilibili视频

## 架构原型

- `conf`，配置组件，负责读取ini配置文件为结构体`conf.Config`
  - 通过`conf.InitConfig()` 初始化组件
    - 读取配置并创建全局配置结构体对象
    - 提供默认值
    - 提供值的校验
      - 值的类型是否合法
      - 枚举值是否合法
      - 路径是否存在
      - 路径是否有权限
    - 校验不通过会直接退出程序
  - 通过导出方法`conf.GetConfig()`获取配置文件结构体
- `logger`，日志组件，负责初始化全局`Logger`
  - 通过`logger.InitLogger()` 初始化组件
    - 目的地
    - 日志格式
  - 通过导出的函数提供打印日志的功能，比如：`logger.Info()`
- `helper` 帮助
  - `util` 工具函数

## 相关资源

- [echo 中文](http://echo.topgoer.com)