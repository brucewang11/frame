package service


//业务返回
type CodeModel struct {
	Code int    //结合业务的错误code
	Data interface{} //不同场景下需要返回的数据
	TemplateData map[string]interface{} //如果有模板
}
