package constants

//Error Defineds
const (
	ErrSuccess    = 0
	ErrFailed     = 1
	ErrUnknowCmd  = 2
	ErrParamWrong = 3
)

var (
	//ErrorDescription error description
	ErrorDescription = map[int]string{
		ErrSuccess:    "success",
		ErrFailed:     "failed",
		ErrUnknowCmd:  "unknow cmd",
		ErrParamWrong: "param wrong",
	}
)

//Module name define
const (
	ModGate = "modGate"
	ModDB   = "modDB"
)

//modue cmd define
const (
	CmdDBSaveWord = "dbCmdSaveWord"
	CmdDBTopWords = "dbCmdTopWords"
)
