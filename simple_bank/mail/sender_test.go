package mail

import (
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {

	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EMAILEmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>hello world</h1>
	<p>This is a test message from <a href="http://zfefef.com"> gdsgfd</a></p>
	`
	to := []string{"dsffgdfgfs@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)

}
