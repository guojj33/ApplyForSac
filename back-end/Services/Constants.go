package Services

type LoginStatusType int

const (
	LoginStatus_Out   LoginStatusType = 0
	LoginStatus_User  LoginStatusType = 1
	LoginStatus_Admin LoginStatusType = 2
)
