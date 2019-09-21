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
		wantErr   bool
	}{
		{
			name: "Timestamp test - valid timestamp",
			args: args{
				timestampBytes: []byte{234, 36, 205, 74, 116, 212, 209, 1},
			},
			timestamp: TimeStamp{},
			want:      "2016-07-02T15:13:30Z",
			wantErr:   false,
		},
		{
			name: "Timestamp test - null timestamp",
			args: args{
				timestampBytes: []byte{},
			},
			timestamp: TimeStamp{},
			want:      "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.timestamp.Parse(tt.args.timestampBytes)
			ts := time.Time(tt.timestamp).Format("2006-01-02T15:04:05Z")
			if ts != tt.want && (err != nil) != tt.wantErr {
				t.Errorf("TimeStamp.Parse() = %v, want = %v", ts, tt.want)
			}
		})
	}
}
