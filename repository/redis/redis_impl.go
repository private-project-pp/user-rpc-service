package redis

type redis struct {
}

func SetupRedis() Redis {
	return &redis{}
}
