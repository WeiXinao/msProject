package respx

var (
	OK                        = NewError(200, "success")
	IllegalMobile             = NewError(201001, "手机号不合法")
	IllegalEmail              = NewError(201002, "邮箱不合法")
	InconsistentPwdAndConfirm = NewError(201003, "两次输入密码不一致")
	ErrEmailDuplicated        = NewError(201004, "邮箱已存在")
	ErrAccountDuplicated      = NewError(201005, "账号已存在")
	ErrMobileDuplicated       = NewError(201006, "手机号已存在")
	ErrVerifyCaptchaFail      = NewError(201007, "验证码错误")
)
