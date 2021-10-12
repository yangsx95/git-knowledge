# git-knowledge

- web框架： [echo](https://github.com/labstack/echo)
- 日志框架： [zap](https://github.com/uber-go/zap)
- ini配置解析： [gopkg.in/ini.v1](http://gopkg.in/ini.v1)
- github api客户端：[go-github](https://github.com/google/go-github)

# 原型/功能列表

## 基础功能

- [ ] 使用github登录，并获取相应的权限与用户信息
  - [ ] 读取仓库的权限
  - [ ] 获取头像，id
- [ ] 首页展示space空间
  - [ ] 支持空间的排序
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