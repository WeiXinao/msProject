package respx

var (
	OK                            = NewError(200, "success")
	ErrIllegalInput               = NewError(400, "非法输入")
	ErrInternalServer             = NewError(500, "系统内部错误")
	IllegalMobile                 = NewError(401001, "手机号不合法")
	IllegalEmail                  = NewError(401002, "邮箱不合法")
	InconsistentPwdAndConfirm     = NewError(401003, "两次输入密码不一致")
	ErrEmailDuplicated            = NewError(401004, "邮箱已存在")
	ErrAccountDuplicated          = NewError(401005, "账号已存在")
	ErrMobileDuplicated           = NewError(401006, "手机号已存在")
	ErrVerifyCaptchaFail          = NewError(401007, "验证码错误")
	ErrAccountOrPasswordIncorrect = NewError(401008, "用户名或密码错误")
)
