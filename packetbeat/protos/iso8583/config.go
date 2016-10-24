package iso8583

import (
	"github.com/elastic/beats/packetbeat/config"
	"github.com/elastic/beats/packetbeat/protos"
)

type iso8583Config struct {
	config.ProtocolCommon `config:",inline"`
}

var (
	defaultConfig = iso8583Config{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
	}
)
