package wx

type app struct {
	appid  string
	secret string
	mch_id string
	key    string

	access_token string
	ticket       string
}

func New(appid, secret, mch_id, key string) app {
	return app{appid: appid, secret: secret, mch_id: mch_id, key: key}
}
