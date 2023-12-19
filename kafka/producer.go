package kafka

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/dustin/go-humanize"
	"github.com/zp857/util/sliceutil"
	"github.com/zp857/util/structutil"
	"go.uber.org/zap"
)

type Producer struct {
	urls               []string
	config             *sarama.Config
	logger             *zap.SugaredLogger
	noStatusDebugPrint bool
	ignoreTopics       []string
}

func NewProducer(urls []string, statusDebugPrint bool, ignoreTopicList []string) *Producer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	return &Producer{
		urls:               urls,
		config:             config,
		logger:             zap.L().Named("[kafka-producer]").Sugar(),
		noStatusDebugPrint: !statusDebugPrint,
		ignoreTopics:       ignoreTopicList,
	}
}

// WithLogger sets the logger to use for error logging.
func (p *Producer) WithLogger(logger *zap.Logger) {
	p.logger = logger.Sugar()
}

func (p *Producer) SendJSON(topic string, obj any) {
	producer, err := sarama.NewSyncProducer(p.urls, p.config)
	if err != nil {
		p.logger.Errorf("NewSyncProducer err: %v", err)
		return
	}
	defer producer.Close()
	producerMessage := &sarama.ProducerMessage{}
	producerMessage.Topic = topic
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(obj)
	if err != nil {
		p.logger.Errorf("json.Marshal err: %v", err)
		return
	}
	producerMessage.Value = sarama.StringEncoder(jsonBytes)
	partition, offset, err := producer.SendMessage(producerMessage)
	if err != nil {
		p.logger.Errorf("SendMessage err: %v", err)
		p.logger.Errorf("message bytes: %v", humanize.Bytes(uint64(producerMessage.Value.Length())))
		//_ = fileutil.WriteFile("result.json", structutil.JsonMarshalIndent(obj))
		return
	}
	if p.noStatusDebugPrint {
		// 如果不是状态上报，记录请求
		if sliceutil.Contain(p.ignoreTopics, topic) {
			return
		}
	}
	p.logger.Infof("[%v] partition=%v offset=%v msg.bytes=%v request:\n%v", topic, partition, offset, humanize.Bytes(uint64(producerMessage.Value.Length())), structutil.JsonMarshalIndent(obj))
}
