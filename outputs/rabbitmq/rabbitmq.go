package fileout

import (
	"encoding/json"
	"time"

	"github.com/cvaldemar/libbeat/common"
	"github.com/cvaldemar/libbeat/logp"
	"github.com/cvaldemar/libbeat/outputs"
	"github.com/streadway/amqp"
)

type RabbitMQOutput struct {
	Conn amqp.Connection
	Url string
	connected bool
}

func (out *RabbitMQOutput) Init(beat string, config outputs.MothershipConfig, topology_expire int) error {
	return nil
}

func (out *RabbitMQOutput) PublishIPs(name string, localAddrs []string) error {
	// not supported by this output type
	return nil
}

func (out *RabbitMQOutput) GetNameByIP(ip string) string {
	// not supported by this output type
	return ""
}

func (out *RabbitMQOutput) PublishEvent(ts time.Time, event common.MapStr) error {

	return nil
}

func (out *RabbitMQOutput) Connect() error {
	var err error
	out.Conn, err = amqp.Dial(out.Url)
	if err != nil {
		return err
	}
	out.connected = true

	return nil
}