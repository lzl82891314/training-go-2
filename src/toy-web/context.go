package toy_web

type IContext interface {
	Json(status int, v interface{}, m string) error
	Ok(v interface{}) error
	NotFound(m string) error

	QueryInt(key string, def int) int
	QueryStr(key string, def string) string
	QueryArr(key string, def []string) []string
	QueryAll() map[string][]string

	FormInt(key string, def int) int
	FormStr(key string, def string) string
	FormArr(key string, def []string) []string
	FormAll() map[string][]string
}
