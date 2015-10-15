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
	 * @exmaple {"status":"200","result":[{"first_letter":"B","cities":[{"name":"\u5317\u4eac"}]},{"first_letter":"C","cities":[{"name":"\u91cd\u5e86"},{"name":"\u6210\u90fd"},{"name":"\u957f\u6c99"}]},{"first_letter":"D","cities":[{"name":"\u5927\u8fde"}]},{"first_letter":"G","cities":[{"name":"\u5e7f\u5dde"}]},{"first_letter":"H","cities":[{"name":"\u676d\u5dde"},{"name":"\u54c8\u5c14\u6ee8"}]},{"first_letter":"J","cities":[{"name":"\u6d4e\u5357"}]},{"first_letter":"K","cities":[{"name":"\u6606\u660e"}]},{"first_letter":"N","cities":[{"name":"\u5357\u4eac"},{"name":"\u5b81\u6ce2"}]},{"first_letter":"Q","cities":[{"name":"\u9752\u5c9b"}]},{"first_letter":"S","cities":[{"name":"\u4e0a\u6d77"},{"name":"\u6df1\u5733"},{"name":"\u82cf\u5dde"},{"name":"\u6c88\u9633"},{"name":"\u77f3\u5bb6\u5e84"}]},{"first_letter":"T","cities":[{"name":"\u5929\u6d25"},{"name":"\u592a\u539f"}]},{"first_letter":"W","cities":[{"name":"\u6b66\u6c49"}]},{"first_letter":"X","cities":[{"name":"\u897f\u5b89"},{"name":"\u53a6\u95e8"}]},{"first_letter":"Z","cities":[{"name":"\u90d1\u5dde"}]}]}
	 *
	 */
	public function GetCityList() {

		//code ...
	}

}