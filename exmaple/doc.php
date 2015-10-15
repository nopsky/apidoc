<?php
namespace App\Http\Controllers\V1;


use App\Http\Controllers\Controller;
use App\Http\Models\City;

//@class 获取城市列表 city
class CityController extends Controller {

	/**
	 * @name 获取城市列表 /V1/city/list
	 *
	 * @method get
	 *
	 * @return_type json
	 *
	 * @return_param 首字母 first_letter string
	 *
	 * @return_param 城市列表 cities array
	 *
	 * @return_param 城市名称 cityies.[0].name string
	 *
	 * @exmaple {"status":"200","result":[{"first_letter":"B","cities":[{"name":"\u5317\u4eac"}]}]}
	 */
	public function GetCityList() {

		//code ...
	}

}