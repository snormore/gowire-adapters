package kafka_adapter

import (
        "github.com/snormore/gowire/message"
)

type KafkaConfig struct{}

type KafkaInputter struct{}

func (in KafkaInputter) Transform(rawMessage interface{}) (message.Message, error) {
        return message.Message{"", rawMessage}, nil
}

func (in KafkaInputter) Listen() chan interface{} {
        return nil
}

import (
  "github.com/snormore/gowire"
  "github.com/snormore/gowire/message"
  "sync"
)

type KafkaOutputter struct{}

func (l KafkaOutputter) Start(messages chan message.Message, wg *sync.WaitGroup, t *tomb.Tomb) error {
  return nil
}
