package session

//Session session
type Session struct {
	info map[string]interface{}
}

//Set set session value
func (s Session) Set(key string, val interface{}) {
	s.info[key] = val
}

//Get get session value
func (s Session) Get(key string) interface{} {
	if v, exist := s.info[key]; exist {
		return v
	}

	return nil
}
