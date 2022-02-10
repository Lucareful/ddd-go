# ddd-go 基于golang语言的DDD实践

> 域是软件打算处理的主题或问题。应该编写软件以反映领域。

## 初始目录结构
- `domain` 文件夹，我们将在其中存储所有需要的子域.
- `entity` 为它将包含 DDD 方法中所谓的实体.
- `ValueObjects` 是不可变的并且不需要唯一标识符. 
- `aggregate` 聚合是一组实体和值对象的组合.
  - DDD 建议使用工厂来创建复杂的聚合、存储库和服务。
  - 我们将实现一个将创建一个新Customer实例的工厂函数。
  - 这将导致一个名为NewCustomer接受name参数的函数，函数内部发生的事情对于想要初始化新客户的域来说不应该关心
- `repository` 定义用于存储客户的存储库的规则
  - DDD 描述了应该使用存储库来存储和管理聚合。
- `service` 定义用于处理客户的服务
  - 连接业务逻辑和存储库.

## 优化后的目录结构
- `cmd` 项目启动文件.
- `domain` 这里将聚合层和存储库的都合并到每个模块相关目录下，整体结构更加清晰.
- `entity` 将实体层和不可变struct合并到一起，都存放实体等一些结构逻辑.(原文章中是没有这一层，它将三个实体放到根目录下了，于是我抽出一个目录存放它们)
- `services` 服务层按照需求拆分成多个模块，每个模块都包含一个服务相关逻辑。


## DDD 小结

```text
实体——可变的可识别结构。

值对象——不可变的不可识别结构。

聚合— 实体和值对象的组合集，存储在存储库中。

Repository — 存储聚合或其他信息的实现

Factory - 一个构造函数，用于创建复杂对象并使其他领域的开发人员更容易创建新实例

服务- 将业务流程构建在一起的存储库和子服务的集合
```

## 参考文章
- [How To Implement Domain-Driven Design (DDD) in Golang](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)
- [How To Structure DDD in Go](https://programmingpercy.tech/blog/how-to-structure-ddd-in-go/)