package ecode

func InitEcodeText() {
	texts := map[Code]string{}

	// 通用ecode
	texts[ServerErr] = "服务器未知错误"
	texts[NetworkErr] = "网络异常"
	texts[ParamWrong] = "参数无效"
	texts[FuncNotImpl] = "功能未实现"
	texts[DaoOperationErr] = "Dao操作错误"

	Register(texts)
}
