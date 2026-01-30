package redis

type Redis interface {
	SaveUserLoginCache(sessionKey string) (key string, err error)
}
