package common

import (
	"errors"
	"sync"
	"time"
)

/*
* 1                                               42           52             64
* +-----------------------------------------------+------------+---------------+
* | timestamp(ms)                                 | workerid   | sequence      |
* +-----------------------------------------------+------------+---------------+
* | 0000000000 0000000000 0000000000 0000000000 0 | 0000000000 | 0000000000 00 |
* +-----------------------------------------------+------------+---------------+
*
* 1. 41-bit time truncate (in milliseconds), note that this is the difference of the time truncate (current time truncate - start time truncate)。It can be used for about 70 years: (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69
* 2. 10 data machine bits, which can be deployed on 1024 nodes
* 3. 12 bit sequence, count in millisecond, same machine, same time intercept concurrent 4096 serial Numbers
 */

const (
	twepoch        = int64(1483228800000)             // Start time cut-off (2017-01-01)
	workeridBits   = uint(10)                         // The number of bits occupied by the machine id
	sequenceBits   = uint(12)                         // The number of bits occupied by a sequence
	workeridMax    = int64(-1 ^ (-1 << workeridBits)) // Maximum number of machine ids supported
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) //
	workeridShift  = sequenceBits                     // Left shift number of machine id
	timestampShift = sequenceBits + workeridBits      // Time stamp left shift number
)

// A Snowflake struct holds the basic information needed for a snowflake generator worker
type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

// NewNode returns a new snowflake worker that can be used to generate snowflake IDs
func NewSnowflake(workerid int64) *Snowflake {
	if workerid < 0 || workerid > workeridMax {
		Logger.ErrorfPanic(errors.New("workerid must be between 0 and 1023"), "invalid worker id. id: %d", workerid)
	}

	return &Snowflake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}
}

// Generate creates and returns a unique snowflake ID
func (s *Snowflake) Generate() int64 {
	s.Lock()

	now := time.Now().UnixNano() / 1000000

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask

		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	r := int64((now-twepoch)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))

	s.Unlock()
	return r
}
