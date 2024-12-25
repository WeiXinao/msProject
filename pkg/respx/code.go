package respx

var (
	OK            = NewError(200, "success")
	IllegalMobile = NewError(2001, "手机号不合法")
)
