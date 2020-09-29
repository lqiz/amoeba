package amoeba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestMap2Map(t *testing.T) {
	inputFile := "history_class_api_resp.json"
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Panicf("failed to read the input file with error")
	}

	const jsonSource = `{"0-qq335":{"id":9527,"origin_id":7890,"product_type":800,"class_id":600,"cc_id":505,"title":"Amoeba是动态进行json转换的组件","logo":"https://xxx3.dudu.com","url":"https://xxx3.dudu.com/url","summary":"aaaa","mold":3,"push_content":"","publish_time":3587900000,"push_time":0,"push_status":0,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_content":"在高一层运用熟练了，再把那一层当做默认信息，再高一层，功夫得这么练。","share_switch":3,"aba_id":9090909,"status":3,"create_time":3586900000,"update_time":3596900000,"clc":0,"is_free_try":false,"ft":false,"order_num":9,"ab_article_token":"","product_id":808,"lgd":"33373","lgt":"article","audio":[{"article_id":404,"audio_id":8080,"audio":{"id":9527,"media_id":"aaaaaaaaaaaaaaaa","title":"Amoeba是动态进行json转换的组件","icon":"https://cdn2.dudu.com/audio/source_icon.png","share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","duration":502,"ctime":"","utime":"","audio_type":66,"share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","from_id":99,"source_name":"Amoeba 使用说明","url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","article_id":404,"article_title":"String 类型扩展，可以直接对其进行简单操作拼接等","token":"XXXXXXXX","avi_version":3,"mp3_play_url":"https://cdn2.dudu.com/audio/mp3.m4a","etag":"xxx","is_limit_free":false,"button_type":0,"aba_id":9090909,"class":600,"class_article_id":999,"class_cc_id":200,"class_article_share_switch":3,"source_intro":"可以有效支持微服务之间的依赖倒置","logo":"https://xxx3.dudu.com","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","ll_id":33,"avatar":"https://xxx3.dudu.com/avatar/qq/2.png","ll_img":"xxxx","ll_uid":256,"view_type":3,"ps_status":355,"used_avi":3,"a_media_id":0}}],"favorite":false,"share_id":"xxxx","share_url":"https://xxx3.dudu.com/share/qq","share":"https://xxx3.dudu.com/share","is_read":false,"recommend_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","internal_class_info":{"id":9527,"product_id":808,"product_type":800,"view_type":3,"ps_status":355,"status":3,"name":"amoeba","formal_article_count":336,"enid":"kkkkkkkkk","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","news":"第三季已完整上线。第四季已于qqqq年3月3日正式上线，欢迎试听试读。","highlight":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内， Array/Map 可以对应 Array|Map，且必须对应两者之一， 默认直接支持 intbool转换，integerstringbool 转换","notice":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内","price":39900,"logo":"https://xxx3.dudu.com","index_img_某let":"","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","phase_num":337,"has_cc":3,"start_time":0,"end_time":0,"ll_id":33,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","manage_user":"110","order_num":9,"url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","price_desc":"讲","last_update":3575336000,"is_finished":3,"create_time":3586900000,"update_time":3596900000,"is_no_audio":0,"publish_time":3587900000,"free_article_ids":null,"intro_article_ids":[83336],"learn_user_count":0,"share_url":"https://xxx3.dudu.com/share/qq","lgd":"48","lgt":"column","ll_info":{"id":9527,"name":"amoeba","nick":"奇正","avatar":"https://xxx3.dudu.com/avatar/qq/2.png","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","title":"Amoeba是动态进行json转换的组件","simg":"https://xxx3.dudu.com/img/qq/2.png","info":"","status":3,"data_updated":3586970038,"create_time":3586900000,"update_time":3596900000,"user":512,"lgd":"33","lgt":"ll"},"last_article_title":"","is_subscribe":3,"last_read_article_title":"","presale_url":"","pms":false,"mini_share_img":"","dimg":"","est":512512,"edt":876,"corner_img":"","mtc":5,"trial":[7878,7879,86538],"m_id":"","video_cover":""},"m_id":0,"video_status":0,"cover_img":"","tnd":""},"0-qq336":{"id":9528,"origin_id":7890,"product_type":800,"class_id":601,"cc_id":505,"title":"Amoeba是动态进行json转换的组件","logo":"https://xxx3.dudu.com","url":"https://xxx3.dudu.com/url","summary":"aaaa","mold":3,"push_content":"","publish_time":3587900000,"push_time":0,"push_status":0,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_content":"在高一层运用熟练了，再把那一层当做默认信息，再高一层，功夫得这么练。","share_switch":3,"aba_id":9090909,"status":3,"create_time":3586900000,"update_time":3596900000,"clc":0,"is_free_try":false,"ft":false,"order_num":9,"ab_article_token":"","product_id":808,"lgd":"33373","lgt":"article","audio":[{"article_id":404,"audio_id":8080,"audio":{"id":9527,"media_id":"aaaaaaaaaaaaaaaa","title":"Amoeba是动态进行json转换的组件","icon":"https://cdn2.dudu.com/audio/source_icon.png","share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","duration":502,"ctime":"","utime":"","audio_type":66,"share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","from_id":99,"source_name":"Amoeba 使用说明","url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","article_id":404,"article_title":"String 类型扩展，可以直接对其进行简单操作拼接等","token":"XXXXXXXX","avi_version":3,"mp3_play_url":"https://cdn2.dudu.com/audio/mp3.m4a","etag":"xxx","is_limit_free":false,"button_type":0,"aba_id":9090909,"class":600,"class_article_id":999,"class_cc_id":200,"class_article_share_switch":3,"source_intro":"可以有效支持微服务之间的依赖倒置","logo":"https://xxx3.dudu.com","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","ll_id":33,"avatar":"https://xxx3.dudu.com/avatar/qq/2.png","ll_img":"xxxx","ll_uid":256,"view_type":3,"ps_status":355,"used_avi":3,"a_media_id":0}}],"favorite":false,"share_id":"xxxx","share_url":"https://xxx3.dudu.com/share/qq","share":"https://xxx3.dudu.com/share","is_read":false,"recommend_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","internal_class_info":{"id":9527,"product_id":808,"product_type":800,"view_type":3,"ps_status":355,"status":3,"name":"amoeba","formal_article_count":336,"enid":"kkkkkkkkk","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","news":"第三季已完整上线。第四季已于qqqq年3月3日正式上线，欢迎试听试读。","highlight":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内， Array/Map 可以对应 Array|Map，且必须对应两者之一， 默认直接支持 intbool转换，integerstringbool 转换","notice":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内","price":39900,"logo":"https://xxx3.dudu.com","index_img_某let":"","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","phase_num":337,"has_cc":3,"start_time":0,"end_time":0,"ll_id":33,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","manage_user":"110","order_num":9,"url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","price_desc":"讲","last_update":3575336000,"is_finished":3,"create_time":3586900000,"update_time":3596900000,"is_no_audio":0,"publish_time":3587900000,"free_article_ids":null,"intro_article_ids":[83336],"learn_user_count":0,"share_url":"https://xxx3.dudu.com/share/qq","lgd":"48","lgt":"column","ll_info":{"id":9527,"name":"amoeba","nick":"奇正","avatar":"https://xxx3.dudu.com/avatar/qq/2.png","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","title":"Amoeba是动态进行json转换的组件","simg":"https://xxx3.dudu.com/img/qq/2.png","info":"","status":3,"data_updated":3586970038,"create_time":3586900000,"update_time":3596900000,"user":512,"lgd":"33","lgt":"ll"},"last_article_title":"","is_subscribe":3,"last_read_article_title":"","presale_url":"","pms":false,"mini_share_img":"","dimg":"","est":512512,"edt":876,"corner_img":"","mtc":5,"trial":[7878,7879,86538],"m_id":"","video_cover":""},"m_id":0,"video_status":0,"cover_img":"","tnd":""},"0-qq337":{"id":9527,"origin_id":7890,"product_type":800,"class_id":602,"cc_id":505,"title":"Amoeba是动态进行json转换的组件","logo":"https://xxx3.dudu.com","url":"https://xxx3.dudu.com/url","summary":"aaaa","mold":3,"push_content":"","publish_time":3587900000,"push_time":0,"push_status":0,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_content":"在高一层运用熟练了，再把那一层当做默认信息，再高一层，功夫得这么练。","share_switch":3,"aba_id":9090909,"status":3,"create_time":3586900000,"update_time":3596900000,"clc":0,"is_free_try":false,"ft":false,"order_num":9,"ab_article_token":"","product_id":808,"lgd":"33373","lgt":"article","audio":[{"article_id":404,"audio_id":8080,"audio":{"id":9527,"media_id":"aaaaaaaaaaaaaaaa","title":"Amoeba是动态进行json转换的组件","icon":"https://cdn2.dudu.com/audio/source_icon.png","share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","duration":502,"ctime":"","utime":"","audio_type":66,"share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","from_id":99,"source_name":"Amoeba 使用说明","url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","article_id":404,"article_title":"String 类型扩展，可以直接对其进行简单操作拼接等","token":"XXXXXXXX","avi_version":3,"mp3_play_url":"https://cdn2.dudu.com/audio/mp3.m4a","etag":"xxx","is_limit_free":false,"button_type":0,"aba_id":9090909,"class":600,"class_article_id":999,"class_cc_id":200,"class_article_share_switch":3,"source_intro":"可以有效支持微服务之间的依赖倒置","logo":"https://xxx3.dudu.com","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","ll_id":33,"avatar":"https://xxx3.dudu.com/avatar/qq/2.png","ll_img":"xxxx","ll_uid":256,"view_type":3,"ps_status":355,"used_avi":3,"a_media_id":0}}],"favorite":false,"share_id":"xxxx","share_url":"https://xxx3.dudu.com/share/qq","share":"https://xxx3.dudu.com/share","is_read":false,"recommend_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","internal_class_info":{"id":9527,"product_id":808,"product_type":800,"view_type":3,"ps_status":355,"status":3,"name":"amoeba","formal_article_count":336,"enid":"kkkkkkkkk","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","news":"第三季已完整上线。第四季已于qqqq年3月3日正式上线，欢迎试听试读。","highlight":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内， Array/Map 可以对应 Array|Map，且必须对应两者之一， 默认直接支持 intbool转换，integerstringbool 转换","notice":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内","price":39900,"logo":"https://xxx3.dudu.com","index_img_某let":"","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","phase_num":337,"has_cc":3,"start_time":0,"end_time":0,"ll_id":33,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","manage_user":"110","order_num":9,"url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","price_desc":"讲","last_update":3575336000,"is_finished":3,"create_time":3586900000,"update_time":3596900000,"is_no_audio":0,"publish_time":3587900000,"free_article_ids":null,"intro_article_ids":[83336],"learn_user_count":0,"share_url":"https://xxx3.dudu.com/share/qq","lgd":"48","lgt":"column","ll_info":{"id":9527,"name":"amoeba","nick":"奇正","avatar":"https://xxx3.dudu.com/avatar/qq/2.png","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","title":"Amoeba是动态进行json转换的组件","simg":"https://xxx3.dudu.com/img/qq/2.png","info":"","status":3,"data_updated":3586970038,"create_time":3586900000,"update_time":3596900000,"user":512,"lgd":"33","lgt":"ll"},"last_article_title":"","is_subscribe":3,"last_read_article_title":"","presale_url":"","pms":false,"mini_share_img":"","dimg":"","est":512512,"edt":876,"corner_img":"","mtc":5,"trial":[7878,7879,86538],"m_id":"","video_cover":""},"m_id":0,"video_status":0,"cover_img":"","tnd":""},"0-qq338":{"id":9528,"origin_id":7890,"product_type":800,"class_id":603,"cc_id":505,"title":"Amoeba是动态进行json转换的组件","logo":"https://xxx3.dudu.com","url":"https://xxx3.dudu.com/url","summary":"aaaa","mold":3,"push_content":"","publish_time":3587900000,"push_time":0,"push_status":0,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_content":"在高一层运用熟练了，再把那一层当做默认信息，再高一层，功夫得这么练。","share_switch":3,"aba_id":9090909,"status":3,"create_time":3586900000,"update_time":3596900000,"clc":0,"is_free_try":false,"ft":false,"order_num":9,"ab_article_token":"","product_id":808,"lgd":"33373","lgt":"article","audio":[{"article_id":404,"audio_id":8080,"audio":{"id":9528,"media_id":"aaaaaaaaaaaaaaaa","title":"Amoeba是动态进行json转换的组件","icon":"https://cdn2.dudu.com/audio/source_icon.png","share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","duration":502,"ctime":"","utime":"","audio_type":66,"share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","from_id":99,"source_name":"Amoeba 使用说明","url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","article_id":404,"article_title":"String 类型扩展，可以直接对其进行简单操作拼接等","token":"XXXXXXXX","avi_version":3,"mp3_play_url":"https://cdn2.dudu.com/audio/mp3.m4a","etag":"xxx","is_limit_free":false,"button_type":0,"aba_id":9090909,"class":600,"class_article_id":999,"class_cc_id":200,"class_article_share_switch":3,"source_intro":"可以有效支持微服务之间的依赖倒置","logo":"https://xxx3.dudu.com","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","ll_id":33,"avatar":"https://xxx3.dudu.com/avatar/qq/2.png","ll_img":"xxxx","ll_uid":256,"view_type":3,"ps_status":355,"used_avi":3,"a_media_id":0}}],"favorite":false,"share_id":"xxxx","share_url":"https://xxx3.dudu.com/share/qq","share":"https://xxx3.dudu.com/share","is_read":false,"recommend_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","internal_class_info":{"id":9528,"product_id":808,"product_type":800,"view_type":3,"ps_status":355,"status":3,"name":"amoeba","formal_article_count":336,"enid":"kkkkkkkkk","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","news":"第三季已完整上线。第四季已于qqqq年3月3日正式上线，欢迎试听试读。","highlight":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内， Array/Map 可以对应 Array|Map，且必须对应两者之一， 默认直接支持 intbool转换，integerstringbool 转换","notice":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内","price":39900,"logo":"https://xxx3.dudu.com","index_img_某let":"","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","phase_num":337,"has_cc":3,"start_time":0,"end_time":0,"ll_id":33,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","manage_user":"110","order_num":9,"url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","price_desc":"讲","last_update":3575336000,"is_finished":3,"create_time":3586900000,"update_time":3596900000,"is_no_audio":0,"publish_time":3587900000,"free_article_ids":null,"intro_article_ids":[83336],"learn_user_count":0,"share_url":"https://xxx3.dudu.com/share/qq","lgd":"48","lgt":"column","ll_info":{"id":9527,"name":"amoeba","nick":"奇正","avatar":"https://xxx3.dudu.com/avatar/qq/2.png","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","title":"Amoeba是动态进行json转换的组件","simg":"https://xxx3.dudu.com/img/qq/2.png","info":"","status":3,"data_updated":3586970038,"create_time":3586900000,"update_time":3596900000,"user":512,"lgd":"33","lgt":"ll"},"last_article_title":"","is_subscribe":3,"last_read_article_title":"","presale_url":"","pms":false,"mini_share_img":"","dimg":"","est":512512,"edt":876,"corner_img":"","mtc":5,"trial":[7878,7879,86538],"m_id":"","video_cover":""},"m_id":0,"video_status":0,"cover_img":"","tnd":""},"0-qq339":{"id":9529,"origin_id":7890,"product_type":800,"class_id":604,"cc_id":505,"title":"Amoeba是动态进行json转换的组件","logo":"https://xxx3.dudu.com","url":"https://xxx3.dudu.com/url","summary":"aaaa","mold":3,"push_content":"","publish_time":3587900000,"push_time":0,"push_status":0,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_content":"在高一层运用熟练了，再把那一层当做默认信息，再高一层，功夫得这么练。","share_switch":3,"aba_id":9090909,"status":3,"create_time":3586900000,"update_time":3596900000,"clc":0,"is_free_try":false,"ft":false,"order_num":9,"ab_article_token":"","product_id":808,"lgd":"33373","lgt":"article","audio":[{"article_id":404,"audio_id":8080,"audio":{"id":9527,"media_id":"aaaaaaaaaaaaaaaa","title":"Amoeba是动态进行json转换的组件","icon":"https://cdn2.dudu.com/audio/source_icon.png","share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","duration":502,"ctime":"","utime":"","audio_type":66,"share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","from_id":99,"source_name":"Amoeba 使用说明","url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","article_id":404,"article_title":"String 类型扩展，可以直接对其进行简单操作拼接等","token":"XXXXXXXX","avi_version":3,"mp3_play_url":"https://cdn2.dudu.com/audio/mp3.m4a","etag":"xxx","is_limit_free":false,"button_type":0,"aba_id":9090909,"class":600,"class_article_id":999,"class_cc_id":200,"class_article_share_switch":3,"source_intro":"可以有效支持微服务之间的依赖倒置","logo":"https://xxx3.dudu.com","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","ll_id":33,"avatar":"https://xxx3.dudu.com/avatar/qq/2.png","ll_img":"xxxx","ll_uid":256,"view_type":3,"ps_status":355,"used_avi":3,"a_media_id":0}}],"favorite":false,"share_id":"xxxx","share_url":"https://xxx3.dudu.com/share/qq","share":"https://xxx3.dudu.com/share","is_read":false,"recommend_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","internal_class_info":{"id":9530,"product_id":808,"product_type":800,"view_type":3,"ps_status":355,"status":3,"name":"amoeba","formal_article_count":336,"enid":"kkkkkkkkk","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","news":"第三季已完整上线。第四季已于qqqq年3月3日正式上线，欢迎试听试读。","highlight":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内， Array/Map 可以对应 Array|Map，且必须对应两者之一， 默认直接支持 intbool转换，integerstringbool 转换","notice":"Items 是属于Map/Array内子类型表述所用，其数据解析，限定为该 Map/Array内","price":39900,"logo":"https://xxx3.dudu.com","index_img_某let":"","logo_iphonex":"https://xxx3.dudu.com","player_img":"https://xxx3.dudu.com","phase_num":337,"has_cc":3,"start_time":0,"end_time":0,"ll_id":33,"share_title":"Amoeba是动态进行json转换的组件，其递归Schema，与gjson组件，对json进行解析及重组","share_summary":"1. 接口由底层服务维护，放在底层服务，底层接口变更时，相关接口单测充分覆盖。更容易维护。2. 建议底层服务使用，保持聚合层的稳定性。","manage_user":"110","order_num":9,"url":"https://xxx3.dudu.com/url","url_qr":"https://xxx3.dudu.com/url/qr","price_desc":"讲","last_update":3575336000,"is_finished":3,"create_time":3586900000,"update_time":3596900000,"is_no_audio":0,"publish_time":3587900000,"free_article_ids":null,"intro_article_ids":[83336],"learn_user_count":0,"share_url":"https://xxx3.dudu.com/share/qq","lgd":"48","lgt":"column","ll_info":{"id":9530,"name":"amoeba","nick":"奇正","avatar":"https://xxx3.dudu.com/avatar/qq/2.png","intro":"支持7种输出类型：integer、boolean、float、map、array、 object、string支持7种输出类型：integer、boolean、float、map、array、 object、string； 使用Json Schema，仅使用简单字段，并扩展增加Value、Key字段。并扩展支持Map，和Array里Item作用一致","title":"Amoeba是动态进行json转换的组件","simg":"https://xxx3.dudu.com/img/qq/2.png","info":"","status":3,"data_updated":3586970038,"create_time":3586900000,"update_time":3596900000,"user":512,"lgd":"33","lgt":"ll"},"last_article_title":"","is_subscribe":3,"last_read_article_title":"","presale_url":"","pms":false,"mini_share_img":"","dimg":"","est":512512,"edt":876,"corner_img":"","mtc":5,"trial":[7878,7879,86538],"m_id":"","video_cover":""},"m_id":0,"video_status":0,"cover_img":"","tnd":""}}`

	timeNow := time.Now()
	for i := 0; i < 1; i++ {
		var schema Schema
		err := json.Unmarshal(b, &schema)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		result, err := Marshal(&schema, jsonSource)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		log.Println(string(result))
	}
	duration := time.Now().Sub(timeNow)
	fmt.Printf("----- %+v------", duration)
}
