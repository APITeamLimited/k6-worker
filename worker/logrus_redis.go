package worker

import (
	"context"

	"github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

// RedisHook to sends logs to Redis server
type RedisHook struct {
	ctx      context.Context
	client   *redis.Client
	jobId    string
	workerId string
}

// NewRedisHook creates a hook to be added to an instance of logger
func NewRedisHook(client *redis.Client, ctx context.Context, jobId string, workerId string) (*RedisHook, error) {
	return &RedisHook{
		client:   client,
		ctx:      ctx,
		jobId:    jobId,
		workerId: workerId,
	}, nil

}

// Fire is called when a log event is fired.
func (hook *RedisHook) Fire(entry *logrus.Entry) error {
	go dispatchMessage(hook.ctx, hook.client, hook.jobId, hook.workerId, entry.Message, messageType)
	return nil
}

// Levels returns the available logging levels.
func (hook *RedisHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.TraceLevel,
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
