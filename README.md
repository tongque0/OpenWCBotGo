# OpenWCBotGo
一个基于openwechat SDK的微信机器人，欢迎开发。

## 配置指南

在开始使用之前，您需要根据 `config-template.yaml` 文件创建一个 `config.yaml` 文件，并填写适合您环境的配置值。

1. 复制 `config-template.yaml` 到 `config.yaml`。
2. 修改 `config.yaml` 文件中的配置项以满足您的环境需求。

## 目录结构

本项目的目录结构如下所示：
```
OpenWCBotGo
├─ main.go                  //入口文件
├─ README.md
├─ Task                     //定时任务
│  ├─ taskinit.go
│  └─ TaskManager.go
├─ LLM                      //大模型先关
│  └─ openai.go
├─ instance                 //openwechat实例
│  ├─ auth.go
│  ├─ init.go
│  └─ reciver.go
└─ config                   //配置先关
│  └─config.go
├─ .gitignore
├─ config-template.yaml     //配置模板
├─ config.yaml              //主配置文件
├─ go.mod
├─ go.sum
```
### 目录说明

- `config`: 包含项目运行所需的配置文件。
    - `config.yaml`: 主配置文件，存储项目运行的环境参数。

- `instance`: 包含与项目实例相关的代码，如微信机器人的初始化和管理。

- `LLM`: 逻辑层或模型层代码，负责处理业务逻辑和数据模型。

- `Task`: 包含定时任务相关的代码，用于执行计划内的任务。


## 快速开始

非开发者直接下载可执行文件，配置config.yaml即可运行





