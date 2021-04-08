package notify

import (
	"os"
	"testing"
)

func TestNotify_Send(t *testing.T) {
	type fields struct {
		config *Config
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test pushover notify",
			fields{config: &Config{
				Platform: Platform("pushover"),
				Token:    os.Getenv("PUSHOVER_TOKEN"),
				Channel:  os.Getenv("PUSHOVER_CHANNEL"),
			}},
			args{msg: "test case"},
		},
		{
			"test slack notify",
			fields{config: &Config{
				Platform: Platform("slack"),
				Token:    os.Getenv("SLACK_TOKEN"),
				Channel:  os.Getenv("SLACK_CHANNEL"),
			}},
			args{msg: "test case"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notify{
				config: tt.fields.config,
			}
			err := n.Send(tt.args.msg)
			if err != nil {
				t.Errorf(err.Error())
			}
		})
	}
}