# pangolin

# 项目介绍
pangolin是第三届青训营的搜索引擎项目。由于我们组员都是go初学者，所以是边学边写，加上人员调整和时间问题，暂时只实现了简单的搜索引擎功能，但基本完成要求。此后我们还将不断扩展pangolin的功能。

## 项目灵感

pangolin参考了[GoFound](https://github.com/newpanjing/gofound)开源项目的一些实现

## 项目实现点

* 支持文本搜索
* 支持用户过滤词查询
* 支持搜索结果分页展示
* 简单的关联度算法

## 项目亮点

* 项目中实现并使用AhoCorasick和DoubleArrayTrie两种结合的数据结构实现存储和倒排索引。
* 排序使用第三届青训营《数据结构与算法》课程中学到的pdqSort
* 实现并使用红黑树保证输入的数据有序
* 实现并使用前缀树实现过滤关键词的功能
* 使用Roaring Bitmap参与关联度算法，和它的并集功能

## 使用技术

* Go
* Mysql
* jieba分词
* Gin框架

## 技术说明

- Double Array Trie是TRIE树的一种变形，它是在保证TRIE树检索速度的前提下，提高空间利用率而提出的一种数据结构，本质上是一个确定有限自动机(deterministic finite automaton，简称DFA)。而AC自动机(Aho-Corasick automaton)算法于1975年产生于贝尔实验室，是一种用于解决多模式匹配问题的经典算法。两者结合能迅速查找用户输入的query对应的Doc，而且内存还很低，是本项目的最大亮点。

- 文档存储在MySql。

- Gin框架对路由进行注册管理

# 结构说明

## pangolin主体结构：

```
pangolin/
├── core
│   ├── AhoCorasickDoubleArrayTrie
│   │   ├── AhoCorasickDoubleArrayTrie.go
│   │   ├── Aho_test.go
│   │   ├── Builder.go
│   │   ├── Hit.go
│   │   ├── Index.go
│   │   └── State.go
│   ├── association
│   │   └── BM25.go
│   ├── engine.go
│   ├── engine_test.go
│   ├── global.go
│   ├── model
│   │   ├── doc.go
│   │   └── search.go
│   ├── MoreLikeThis
│   │   └── moreLikeThis.go
│   ├── pageSplit
│   │   ├── pageSplit.go
│   │   └── pageSplit_test.go
│   ├── sort
│   │   ├── pdqsort.go
│   │   └── pdqsort_test.go
│   ├── split
│   │   ├── dict
│   │   │   └── dictionary.txt
│   │   ├── split.go
│   │   └── split_test.go
│   ├── storage
│   │   └── main.go
│   ├── util
│   │   ├── maputils.go
│   │   └── utils.go
│   └── wordsFilter
│       ├── node.go
│       ├── words_filter.go
│       ├── words_filter_test.go
│       └── words_test.txt
├── dao
│   ├── dao_test.go
│   ├── db
│   │   ├── Douban.go
│   │   └── init.go
│   └── init.go
├── data
│   ├── e32a16d4-a10c-158d-a14a-dcf7abd54f6d.data
│   └── wukong50k_release.csv
├── datastructure
│   ├── queue
│   │   └── queue.go
│   ├── rbTree
│   │   ├── rbTree.go
│   │   └── rbTree_test.go
│   └── trie
│       ├── trie.go
│       └── trie_test.go
├── doc
│   └── image
│       ├── target.png
│       └── task.png
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.en.md
├── README.md
├── redis
│   ├── simple.go
│   └── simple_test.go
├── tests
│   └── http
│       └── query.http
└── web
    ├── controller
    │   ├── base.go
    │   ├── README.md
    │   ├── response.go
    │   └── services.go
    ├── result.go
    ├── router
    │   ├── base.go
    │   ├── README.md
    │   └── router.go
    └── service
        ├── base.go
        └── README.md

28 directories, 59 files

```

## API简介

- 查询：/api/query    

  - 请求方式：post

  - 请求参数：

  - | "query"     | string 用户输入的查询             |
    | ----------- | --------------------------------- |
    | "page"      | Int 分页查询的第page页            |
    | "limit"     | Int 分页查询的每页limit结果       |
    | "sensitive" | String 用户可以自定义的关键词过滤 |

  - 返回值：json字符串，查询的结果

您可以在pangolin\tests\http\query.http进行测试此项

# 作者&鸣谢

## 本项目人员分配：

黄滨 -- 负责项目开发统筹和文档编写

占文星 -- 负责项目Douban数据查询和部分索引的开发

谢荣飞 -- 负责分页功能和关键词过滤

朱骁 -- 负责Gin框架路由注册与管理部分



## 参考鸣谢：gofound开源项目，CSDN，知乎等