package global

import (
	"github.com/Shopify/sarama"
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	RedisClient         *redis.Client
	MysqlClient         *gorm.DB
	ElasticsearchClient *elastic.Client
	KafkaClient         sarama.SyncProducer
	LoggerClient        *zap.SugaredLogger
)

var (
	Env             string
	RedisHost       string
	RedisPassword   string
	MysqlDsn        string
	MysqlPrefix     string
	AppReadTimeOut  int
	AppWriteTimeOut int
	AppHttpPort     int
	JwtSecret       string
	JwtExpire       int
	LoggerFilePath  string
	KafkaHost       string
	KafkaTopic      string
	ElasticHost     string
	ElasticUsername string
	ElasticPassword string
)
