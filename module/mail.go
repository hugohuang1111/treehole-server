package module

const (
	//PayloadKeySession session
	PayloadKeySession = "session"
	//PayloadKeyRecvData recv data
	PayloadKeyRecvData = "recvData"
	//PayloadKeySendData send data
	PayloadKeySendData = "sendData"
	//PayloadKeyCmd cmd
	PayloadKeyCmd = "cmd"
	//PayloadKeyConnectID connect id
	PayloadKeyConnectID = "connectID"
	//PayloadKeyUserID user id
	PayloadKeyUserID = "userID"
	//PayloadKeyRoomID room id
	PayloadKeyRoomID = "roomID"
	//PayloadKeyCustom custom
	PayloadKeyCustom = "custom"
	//PayloadKeyError error vaule
	PayloadKeyError = "error"
	//PayloadKeyDescription error description
	PayloadKeyDescription = "description"
)

const (
	//MailTypeUnknow unknow
	MailTypeUnknow = "unknow"
	//MailTypeClient client
	MailTypeClient = "client"
	//MailTypeDisconnect disconnect
	MailTypeDisconnect = "disconnect"
	//MailTypeSend send
	MailTypeSend = "send"
	//MailTypeNormal normal
	MailTypeNormal = "normal"
)

//Mail module event
type Mail struct {
	Sender  string
	Recver  string
	Type    string
	Payload map[string]interface{}
}

//GetCmd get cmd
func GetCmd(mail *Mail) string {
	return GetPayloadValueString(mail, PayloadKeyCmd)
}

//GetSendData get send data
func GetSendData(mail *Mail) interface{} {
	return GetPayloadValue(mail, PayloadKeySendData)
}

//GetConnectID get connect id
func GetConnectID(mail *Mail) uint64 {
	return GetPayloadValueUint64(mail, PayloadKeyConnectID)
}

//GetPayloadValueString get payload string value
func GetPayloadValueString(mail *Mail, key string) string {
	v := GetPayloadValue(mail, key)

	if vstring, ok := v.(string); ok {
		return vstring
	}

	return ""
}

//GetPayloadValueInt get payload int value
func GetPayloadValueInt(mail *Mail, key string) int {
	v := GetPayloadValue(mail, key)

	if vint, ok := v.(int); ok {
		return vint
	}

	return 0
}

//GetPayloadValueInt64 get payload int64 value
func GetPayloadValueInt64(mail *Mail, key string) int64 {
	v := GetPayloadValue(mail, key)

	if vint64, ok := v.(int64); ok {
		return vint64
	}

	return 0
}

//GetPayloadValueUint64 get payload uint64 value
func GetPayloadValueUint64(mail *Mail, key string) uint64 {
	v := GetPayloadValue(mail, key)

	if vuint64, ok := v.(uint64); ok {
		return vuint64
	}

	return 0
}

//GetPayloadValue get payload interface value
func GetPayloadValue(mail *Mail, key string) interface{} {
	if v, ok := mail.Payload[key]; ok {
		return v
	}

	return nil
}
