package respx

var (
	OK                            = NewError(200, "success")
	ErrIllegalInput               = NewError(400, "非法输入")
	ErrInternalServer             = NewError(500, "系统内部错误")
	ErrNotHasAuthority = NewError(401000, "没有权限，请联系管理员")
	IllegalMobile                 = NewError(401001, "手机号不合法")
	IllegalEmail                  = NewError(401002, "邮箱不合法")
	InconsistentPwdAndConfirm     = NewError(401003, "两次输入密码不一致")
	ErrEmailDuplicated            = NewError(401004, "邮箱已存在")
	ErrAccountDuplicated          = NewError(401005, "账号已存在")
	ErrMobileDuplicated           = NewError(401006, "手机号已存在")
	ErrVerifyCaptchaFail          = NewError(401007, "验证码错误")
	ErrAccountOrPasswordIncorrect = NewError(401008, "用户名或密码错误")
	ErrNotMember = NewError(401009, "不是项目成员，无操作权限")
	ErrNotOwner = NewError(401009, "不是项目所有者，无操作权限")
)

var (
	ErrProjectAlreadyDeleted = NewError(402001, "项目已删除")
)

var (
	ErrEmptyTaskName = NewError(403001, "任务标题不能为空")
	ErrTaskStageNotExists = NewError(403002, "任务步骤不存在")
)

var (
)