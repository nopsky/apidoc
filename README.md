# API 接口文档生成工具

## 下载

```
go get github.com/nopsky/apidoc
```

## 安装
```	
make
```

## 使用

```
apidoc -doc doc -path exmaple -suffix .php
```

## 特点
```

1.自动根据类路径自动生成文档路径

2.自动生成API接口文档的索引文件

3.自动生成接口的返回结果的示例说明

4.支持在文件的任何地方进行接口文档的补充

```

### 说明

```
目前只支持`/*...*/`，`//`的注释方式

```

#目前支持的注释参数与格式要求如下:

```

@class [name] [classPath]

@name [name] [route] [desc]

@author [username] [email]

@method [method]
	
@input_param [name] [varName] [varType] [isMust] [varValue] [desc]

@return_type [name]

@return_param [name] [varName] [varType] [parent] [desc] [varValue]

@exmaple [json]
	
```

#Demo

```
<?php
namespace App\Http\Controllers\Shop\V1;
use App\Http\Controllers\Controller;

//@class 获取配置信息	Config
class ConfigController extends Controller {

	/**
	 * @name  配送时间和配送范围 Shop/V1/config 获取配送时间和配送范围可选列表
	 * @author nopsky cnnopsky@gmail.com
	 * @method get
	 * @return_type json
	 * @return_param 配送范围 shipment array+int
	 * @return_param 配送时间 shipping_time array+int null 单位为分
	 */
	public function config() {
		$result = array();

		$result['shipment'] = array("200", "500", "1000");
		$result['shipping_time'] = array("10", '20', '30');

		return response()->json(['result'=>$result, 'status'=>'200']);
	}
}

```

1.`@class` : 用来归类接口

名称|类路径
---|---

#####Demo
> 
> @name 用户相关 user

接口文档的索引文件会按照`类名称`进行分组
接口文档会根据`类路径`进行生成

2.`@name` : 用来描述方法的名称等

方法名称|请求方法URI|方法说明(可选)
---|---|---

#####Demo
> 
> @name 获取城市列表 /V1/city/list
	

3.`@author` : 接口开发者

姓名 | 联系方式
--- | ---
#####Demo
> @author nopsky cnnopsky@gmail.com

4.`@method` 接口访问方式

| 访问方式 |
| --- |
#####Demo
> @method get

5.`@input_param` : 提交的参数

字段名|变量名|类型|必填(可选)|示例值(可选)|描述(可选)
---|---|---|---|---|---
#####Demo
> @input_param 用户UID uid int 1 100 用户的唯一ID

必填的值为 : `0` or `1`

支持的类型有:`int`, `string`, `float`, `bool`, `url`, `array`, `file`, 程序会根据类型自动生成示例值

6.`@return_type` : 返回的格式

|返回格式|
|---|
#####Demo
> @return_type json

目前只支持`JSON`

7.`@return_param` : 返回的参数

字段名|变量名|类型|描述（可选)
---|---|---|---|---|---

#####Demo
> @return\_param 首字母 first_letter string

> @return\_param 城市列表 cities array

> @return\_param 城市名称 cityies.[0].name string
