package session

//Mgr session mgr
type Mgr struct {
	sessionMap map[string]Session
}

//FindOrCreate find session, if not exist create
func (mgr Mgr) FindOrCreate(session string) Session {
	if sess, exist := mgr.sessionMap[session]; exist {
		return sess
	}
	sess := Session{}
	sess.info = make(map[string]interface{})
	mgr.sessionMap[session] = sess

	return sess
}

//Delete delete session
func (mgr Mgr) Delete(session string) {
	delete(mgr.sessionMap, session)
}

//Set set session key value
func (mgr Mgr) Set(session, key string, value interface{}) {
	mgr.FindOrCreate(session).Set(key, value)
}

//Get get session key value
func (mgr Mgr) Get(session, key string) interface{} {
	return mgr.FindOrCreate(session).Get(key)
}
