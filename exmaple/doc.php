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
		$result = array();

		$result['shipment'] = array("200", "500", "1000");
		$result['shipping_time'] = array("10", '20', '30');

		return response()->json(['result'=>$result, 'status'=>'200']);
	}
}