package db

type SqlFields struct {
	Username     string
	Password     string
	RootPassword string
	Host         string
	Port         string
	Dbname       string
	Service      string
}

type RedisFields struct {
	Host     string
	Port     string
	Password string
	Dbname   int
	Service  string
}

type KafkaFields struct {
	Brokers       []string
	Group         string
	Service       string
	ServerAddress string
	Topic         string
}
