package Models

type ApplyStatusType int

const (
	ApplyStatus_Applying ApplyStatusType = 0 //正在申请
	ApplyStatus_Canceled ApplyStatusType = 1 //取消申请
)

type ReviewStatusType int

const (
	ReviewStatus_Waiting  ReviewStatusType = 0 //等待审核
	ReviewStatus_Accepted ReviewStatusType = 1 //通过审核
	ReviewStatus_Rejected ReviewStatusType = 2 //未通过审核
)

type CheckStatusType int

const (
	CheckStatus_Waiting CheckStatusType = 0 //还未签到
	CheckStatus_Normal  CheckStatusType = 1 //准时签入/出
	CheckStatus_Late    CheckStatusType = 2 //迟到签入/出
)

type TimeDurCompType int

const (
	TimeDurComp_Earlier     TimeDurCompType = 0
	TimeDurComp_Overlapping TimeDurCompType = 1
	TimeDurComp_Later       TimeDurCompType = 2
)

type AccountType int

const (
	AccountType_User     AccountType = 0
	AccountType_Admin    AccountType = 1
	AccountType_NotExist AccountType = 2
)
