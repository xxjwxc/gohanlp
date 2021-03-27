# gohanlp
中文分词 词性标注 命名实体识别 依存句法分析 语义依存分析 新词发现 关键词短语提取 自动摘要 文本分类聚类 拼音简繁转换 自然语言处理


## [HanLP](https://github.com/hankcs/HanLP) 的golang 接口
- 在线轻量级RESTful API
- 仅数KB，适合敏捷开发、移动APP等场景。服务器算力有限，匿名用户配额较少

## 使用方式

### 安装
```
go get -u github.com/xxjwxc/gohanlp@master

```
#### 使用

#### 申请auth认证

https://bbs.hanlp.com/t/hanlp2-1-restful-api/53
#### 文本形式

```
client := hanlp.HanLPClient(hanlp.WithAuth("你申请到的auth")) // auth不填则匿名
s, _ := client.Parse("2021年HanLPv2.1为生产环境带来次世代最先进的多语种NLP技术。阿婆主来到北京立方庭参观自然语义科技公司。",hanlp.WithLanguage("zh"))
fmt.Println(s)
```

#### 对象形式

```
client := hanlp.HanLPClient(hanlp.WithAuth("你申请到的auth")) // auth不填则匿名
resp, _ := client.ParseObj("2021年HanLPv2.1为生产环境带来次世代最先进的多语种NLP技术。阿婆主来到北京立方庭参观自然语义科技公司。",hanlp.WithLanguage("zh"))
fmt.Println(resp)
```


#### 更多调用API 请查看
[options](https://github.com/xxjwxc/gohanlp/blob/main/hanlp/option.go)

#### 更多信息

[English](https://github.com/hankcs/HanLP/tree/master) | [文档](https://hanlp.hankcs.com/docs/) |  [1.x版](https://github.com/hankcs/HanLP/tree/1.x) | [论坛](https://bbs.hankcs.com/) | [docker](https://github.com/wangedison/hanlp-jupyterlab-docker) | [▶️在线运行](https://play.hanlp.ml/)

面向生产环境的多语种自然语言处理工具包，基于PyTorch和TensorFlow 2.x双引擎，目标是普及落地最前沿的NLP技术。HanLP具备功能完善、性能高效、架构清晰、语料时新、可自定义的特点。

借助世界上最大的多语种语料库，HanLP2.1支持包括简繁中英日俄法德在内的104种语言上的10种联合任务：**分词**（粗分、细分2个标准，强制、合并、校正3种[词典模式](https://github.com/hankcs/HanLP/blob/master/plugins/hanlp_demo/hanlp_demo/zh/demo_custom_dict.py)）、**词性标注**（PKU、863、CTB、UD四套词性规范）、**命名实体识别**（PKU、MSRA、OntoNotes三套规范）、**依存句法分析**（SD、UD规范）、**成分句法分析**、**语义依存分析**（SemEval16、DM、PAS、PSD四套规范）、**语义角色标注**、**词干提取**、**词法语法特征提取**、**抽象意义表示**（AMR）。

量体裁衣，HanLP提供**RESTful**和**native**两种API，分别面向轻量级和海量级两种场景。无论何种API何种语言，HanLP接口在语义上保持一致，在代码上坚持开源。

