package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"photo-upload-service/internal/models"
	"time"
)

func (p *Publisher) PublishPhoto(ctx context.Context, id uuid.UUID) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	//TODO: добавить методы и данные о мониторе после реализации бд и остальных энгдпоинтов
	message := models.EvaluationMessage{
		PhotoID: id,
		Monitor: models.MonitorData{ID: 1},
		Method:  "Mock method",
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	var publishCtx context.Context
	var cancel context.CancelFunc

	if _, ok := ctx.Deadline(); !ok {
		publishCtx, cancel = context.WithTimeout(ctx, 30*time.Second)
		defer cancel()
	} else {
		publishCtx = ctx
	}

	err = p.channel.PublishWithContext(
		publishCtx,
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         jsonData,
			DeliveryMode: amqp.Persistent,
			MessageId:    id.String(),
			Timestamp:    time.Now(),
			Headers: amqp.Table{
				"photo_id": id.String(),
			},
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish photo as base64: %w", err)
	}

	return nil
}
