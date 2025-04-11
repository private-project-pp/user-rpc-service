package redis

type Redis interface {
	SaveUserLoginCache(sessionKey, sessionInfo string) (err error)
}
