package GoFor_MFT_Parser

import (
	"testing"
	"time"
)

func TestTimeStamp_Parse(t *testing.T) {
	type args struct {
		timestampBytes []byte
	}
	tests := []struct {
		name      string
		timestamp TimeStamp
		args      args
		want      string
	}{
		{
			name: "Timestamp test",
			args: args{
				timestampBytes: []byte{234, 36, 205, 74, 116, 212, 209, 1},
			},
			timestamp: TimeStamp{},
			want:      "2016-07-02T15:13:30Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.timestamp.Parse(tt.args.timestampBytes)
			ts := time.Time(tt.timestamp).Format("2006-01-02T15:04:05Z")
			if ts != tt.want {
				t.Errorf("TimeStamp.Parse() = %v, want = %v", ts, tt.want)

			}
		})
	}
}
