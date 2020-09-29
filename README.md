# Amoeba 使用说明

Amoeba是动态进行json转换的组件，通过配置进行数据格式转换，可以减少接口的编码，适用于有较多接口适配需求的业务场景。

## 快速入门
可以参考 rule/ mock/ 下文件。
1. 定制规则，可以放在 rule 内。也可以用其他方式自定义规则。将自定义规则函数传入 amoeba.Start即可。
2. 配置路由，实现反向代理，不同框架有不同实现。router.Any("/amoeba/*proxyPath", amoeba.ReverseProxy)。


## 原理
Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组。
1. 支持7种输出类型："integer"、"boolean"、"float"、"map"、"array"、 "object"、"string"。
2. 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致。
3. 设置Value获取数据，其获取格式为：internal_class_info.name（内部使用 gjson 进行获取数据）。
    3.1 对叶子节点，设置Value。
    3.2 对map/array，设置Value，对map需要设置key

## 注意事项：
1. Array/Map 可以对应 Array|Map，且必须对应两者之一。
2. Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内。
3. Items 是 内不支持不同类型，及Items进行描述时是 Object。
4. 默认直接支持 int\bool转换，integer\string\bool 转换。

## string 结点扩展操作
范围：String 叶子节点，以及 Map里 Key
使用：使用 Golang  fmt.Sprintf(格式化样式, 参数列表…)样式 来支持扩展，如："66-%s, class_id"，注意只能使用 %s。


## 聚合层服务 OR 底层使用
1. Amoeba 可以有效支持微服务之间的依赖倒置。
2. 建议建议底层服务使用，保持聚合层的稳定，由底层服务维护，等同于接口放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。
