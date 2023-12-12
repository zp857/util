package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/zp857/util/threading"
	"go.uber.org/zap"
)

type Consumer struct {
	urls        []string
	config      *sarama.Config
	TopicMap    map[string]interface{}
	bindings    []HandlerBinding
	middlewares []MiddlewareFunc
	logger      *zap.SugaredLogger
}

func NewConsumer(urls []string, topicMap map[string]interface{}) (consumer *Consumer) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	return &Consumer{
		urls:        urls,
		config:      config,
		TopicMap:    topicMap,
		bindings:    []HandlerBinding{},
		middlewares: []MiddlewareFunc{},
		logger:      zap.L().Named("[kafka-consumer]").Sugar(),
	}
}

// WithLogger sets the logger to use for error logging.
func (c *Consumer) WithLogger(logger *zap.SugaredLogger) {
	c.logger = logger
}

// AddMiddleware will add a ServerMiddleware to the list of middlewares to be
func (c *Consumer) AddMiddleware(m MiddlewareFunc) {
	c.middlewares = append(c.middlewares, m)
}

// Bind will add a HandlerBinding to the list of bindings
func (c *Consumer) Bind(bingding HandlerBinding) {
	c.bindings = append(c.bindings, bingding)
}

func (c *Consumer) GetBindings() []HandlerBinding {
	return c.bindings
}

func (c *Consumer) ListenAndServe() {
	c.StartConsume()
}

func (c *Consumer) StartConsume() {
	// 初始化连接
	consumer := c.conn()
	// 监听消息
	for _, binding := range c.bindings {
		go c.consume(consumer, binding)
	}
}

func (c *Consumer) conn() (customer sarama.Consumer) {
	consumer, err := sarama.NewConsumer(c.urls, c.config)
	if err != nil {
		c.logger.Errorf("初始化消費者失败: %v", err)
		for {
			time.Sleep(3 * time.Second)
			consumer, err = sarama.NewConsumer(c.urls, c.config)
			if err != nil {
				c.logger.Errorf("重试初始化消费者失败: %v", err)
			} else {
				c.logger.Infof("重试初始化消费者成功")
				break
			}
		}
	}
	return consumer
}

func (c *Consumer) consume(consumer sarama.Consumer, bingding HandlerBinding) {
	// 查询所有分区
	partitions, err := consumer.Partitions(bingding.TopicName)
	if err != nil {
		c.logger.Errorf("消费者获取kafka分区列表失败: %v", err)
		for {
			partitions, err = consumer.Partitions(bingding.TopicName)
			if err != nil {
				c.logger.Errorf("重试消费者获取kafka分区列表失败: %v", err)
				time.Sleep(3 * time.Second)
			} else {
				c.logger.Infof("消费者获取kafka分区列表成功")
				break
			}
		}
	}
	// 监听所有分区的消息
	wg := &sync.WaitGroup{}
	for partition := range partitions {
		var destConsumer sarama.PartitionConsumer
		for {
			destConsumer, err = consumer.ConsumePartition(bingding.TopicName, int32(partition), sarama.OffsetNewest)
			if err != nil {
				c.logger.Errorf("消费者尝试消费kafka分区[%v]最新消息失败: %v", int32(partition), err)
				time.Sleep(3 * time.Second)
			} else {
				break
			}
		}
		wg.Add(1)
		// 为每个分区开启一个go协程取值
		go func(sarama.PartitionConsumer) {
			// 阻塞直到有值发送过来，然后继续等待
			for msg := range destConsumer.Messages() {
				// handler
				handler := MiddlewareChain(bingding.HandlerFunc, c.middlewares...)
				ctx := context.Background()
				ctx = context.WithValue(ctx, "topicName", bingding.TopicName)
				threading.GoSafe(func() {
					handler(ctx, msg.Value)
				})
			}
			defer destConsumer.AsyncClose()
			defer wg.Done()
		}(destConsumer)
	}
	wg.Wait()
	err = consumer.Close()
	if err != nil {
		c.logger.Errorf("关闭kafka失败: %v", err)
	}
}
