package consumer

import (
	"crm_system/internal/auth/controller/consumer/dto"
	"crm_system/pkg/auth/logger"
	"encoding/json"
	"github.com/IBM/sarama"
)

type UserVerificationCallback struct {
	logger *logger.Logger
}

func NewUserVerificationCallback(logger *logger.Logger) *UserVerificationCallback {
	return &UserVerificationCallback{logger: logger}
}

func (c *UserVerificationCallback) Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError) {
	for {
		select {
		case msg := <-message:
			var userCode dto.UserCode

			err := json.Unmarshal(msg.Value, &userCode)
			if err != nil {
				c.logger.Error("failed to unmarshall record value err: %v", err)
			} else {
				c.logger.Info("user code: %s", userCode)
			}
		case err := <-error:
			c.logger.Error("failed consume err: %v", err)
		}
	}
}
