package worker

import (
	"context"
	"encoding/json"
	"fmt"
	db "simple-bank/db/sqlc"
	"simple-bank/util"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (rtd *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option) error {

	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, payloadByte, opts...)

	info, err := rtd.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("queue", info.Queue).
		Int("max_retry", info.MaxRetry).
		Msg("enqueued task")

	return nil

}

func (rtp *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {

	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal task payload: %w", asynq.SkipRetry)
	}

	user, err := rtp.store.GetUser(ctx, payload.Username)
	if err != nil {
		// if errors.Is(err, db.ErrRecordNotFound) {
		// 	return fmt.Errorf("user doesn't exist: %w", asynq.SkipRetry)
		// }

		return fmt.Errorf("failed to get user: %w", err)
	}

	//send email to user

	verifyEmail, err := rtp.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
		Username:   user.Username,
		Email:      user.Email,
		SecretCode: util.RandomString(32),
	})

	if err != nil {
		return fmt.Errorf("failed to create verify email: %w", err)
	}
	subject := "welcome to simple bank"
	content := `
	<h1>hello world</h1>
	<p>This is a test message from <a href="http://zfefef.com"> gdsgfd</a></p>
	`
	to := []string{verifyEmail.Email}
	err = rtp.mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("email", user.Email).
		Msg("processed task")

	return nil
}
