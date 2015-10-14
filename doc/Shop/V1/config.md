

> 接口名称

获取配送时间和配送范围可选列表

> 调用地址

Shop/V1/config



> 调用方式

GET

>参数

无


> 返回类型

JSON

> 返回结果

```
{
  "shipment": [
    1,
    2,
    3,
    4,
    5
  ],
  "shipping_time": [
    1,
    2,
    3,
    4,
    5
  ]
}
```


> 返回字段说明

字段名|变量名|类型|描述
---|---|---|---
配送范围|shipment|array+int|无
配送时间|shipping_time|array+int|单位为分




> 作者

nopsky  <cnnopsky@gmail.com>
