# API 接口文档生成工具



#目前支持的注释参数与格式要求如下:

```

@name [name] [route] [desc]

@author [username] [email]

@method [method]
	
@input_param [name] [varName] [varType] [isMust] [varValue] [desc]

@return_type [name]

@return_param [name] [varName] [varType] [parent] [desc] [varValue]
	
```

#Demo

```
<?php
namespace App\Http\Controllers\Shop\V1;
use App\Http\Controllers\Controller;

class ConfigController extends Controller {
	/**
	 * @name 获取配送时间和配送范围可选列表 Shop/V1/config
	 * @author nopsky cnnopsky@gmail.com
	 * @method get
	 * @return_type json
	 * @return_param 配送范围 shipment array+int
	 * @return_param 配送时间 shipping_time array+int null 单位为分
	 */
	public function config() {
		//code....
	}
}

```

1.`@name` : 用来描述方法的名称等

方法名称|请求方法URI|方法说明(可选)
---|---|---

#####Demo
> 
> @name 获取用户信息 user/GetUserInfo 获取用户信息
	

2.`@author` : 接口开发者

姓名 | 联系方式
--- | ---
#####Demo
> @author nopsky cnnopsky@gmail.com

3.`@method` 接口访问方式

| 访问方式 |
| --- |
#####Demo
> @method post

4.`@input_param` : 提交的参数

字段名|变量名|类型|必填(可选)|示例值(可选)|描述(可选)
---|---|---|---|---|---
#####Demo
> @input_param 用户UID uid int 1 100 用户的唯一ID

必填的值为 : `0` or `1`

支持的类型有:`int`, `string`, `float`, `bool`, `url`, `array`, `file`

5.`@return_type` : 返回的格式

|返回格式|
|---|
#####Demo
> @return_type json

目前只支持`JSON`

6.`@return_param` : 返回的参数

字段名|变量名|类型|参数所属|描述（可选)|示例值(可选)
---|---|---|---|---|---

#####Demo
> @return\_param 配送时间 shipping_time array+int null 时间单位为(分)

> @return\_param 配送时间 shipping_time array+int userInfo.username 时间单位为(分)

参数所属:根据层级，使用`.`进行组装，例如:

```
{
	"list" : [
		{
			"userInfo": {
				"uid" : 1,
				"username" : nopsky
			}
		}	
	]

}
```
>>字段uid的参数所属为:list.userInfo

>>当为null时，表示在返回结构中属于顶级

支持的类型为:`int`, `string`, `float`, `url`, `bool`, `array+int`(整型数组), `array+string`(字符串数组), `array+url`(字符串数组), `array+float`(浮点型数组)