/*
 * Copyright (c) 2019 Alec Randazzo
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 */

package GoFor_MFT_Parser

import (
	"fmt"
	bin "github.com/AlecRandazzo/BinaryTransforms"
	"time"
)

type TimeStamp time.Time

// Parse a byte slice containing a unix timestamp and convert it to a timestamp string.
func (timestamp *TimeStamp) Parse(timestampBytes []byte) (err error) {

	// verify that we are getting the bytes we need
	if len(timestampBytes) != 8 {
		err = fmt.Errorf("timestamp.parse() received %v bytes, not 8 bytes", len(timestampBytes))
		*timestamp = TimeStamp(time.Time{}) // time.Time nil equivalent
		return
	}

	var delta = time.Date(1970-369, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	// Convert the byte slice to little endian int64 and then convert it to a string
	timestampInt64, err := bin.LittleEndianBinaryToInt64(timestampBytes)
	if err != nil {
		err = fmt.Errorf("failed to parse timestamp bytes, %v", err)
		*timestamp = TimeStamp(time.Time{})
		return
	}

	*timestamp = TimeStamp(time.Unix(0, timestampInt64*100+delta).UTC())

	return
}
