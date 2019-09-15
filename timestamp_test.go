package GoFor_MFT_Parser

import (
	"encoding/hex"
	"testing"
)

func TestTimeStamp_Parse(t *testing.T) {
	timestampBytes, _ := hex.DecodeString("EA24CD4A74D4D101")
	type args struct {
		timestampBytes []byte
	}
	tests := []struct {
		name      string
		timestamp TimeStamp
		args      args
	}{
		{
			name: "Timestamp test",
			args: args{
				timestampBytes: timestampBytes,
			},
			timestamp: TimeStamp("2016-07-02T15:13:30Z"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
