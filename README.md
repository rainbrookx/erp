# README - ERP

> ⚠️ 本项目仅为个人学习练手项目，**非官方版本**，不提供任何生产保障，用于技术学习与简历展示

## 项目说明

本项目基于开源 ERP 项目二次改造：

- 原始上游项目名称：管伊佳ERP
- 原始上游仓库：<https://gitee.com/jishenghua/JSH_ERP>
- 基准 commit：2026-10-06, 26124e31babdbcda411b051ac8df4ad0c6721b1a, 26124e3

改造和保留内容：

【前端】

- 完整保留上游前端源码、页面、组件、静态资源
- 修改前端项目的根目录名，从 `jshERP-web` 改为 `erp-web-vue2`
- 添加了部分便于使用 Visual Studio Code 的配置：`.eslintrc.js`、`.nvmrc`、`jsconfig.json`

【后端】

- 移除上游全部 Java 后端源码，修改后端项目的根目录名，从 `jshERP-boot` 改为 `erp-api`
- 使用 Golang 重新实现后端服务，对齐原有前端接口协议

## 协议 License

本项目整体采用 **Apache License 2.0**

- 上游原有前端源码版权归原作者所有
- 本人新增的 Golang 后端源码同样以 Apache-2.0 许可开源

完整协议文本见根目录 `LICENSE` 文件

## 技术栈

- 前端：和上游保持一致
- 后端：Golang、Gin、GORM

## 项目结构

## 启动方式

## 学习收获

- golang 默认0值
- 数据库 bool 用 bit(1) 的坑
- 原始 Java 代码有太多的“隐式依赖”

## Java 的隐式依赖（缺点）

- 自定义的分页：直接用无传参 `PageUtils.startPage()` 配合 `RequestContextHolder` 获取分页参数
- 自定义的多租户插件：直接用 `RequestContextHolder` 的 header 参数中获取 token，解析出用户和租户
