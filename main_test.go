package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanBlocks(t *testing.T) {

	test_cases := []string{"./test_data/ScanBlocksTestDataTrailingNewline", "./test_data/ScanBlocksTestDataNoTrailingNewline"}
	outputs := make([]string, len(test_cases))
	for i, s := range test_cases {
		t.Run(s, func(t *testing.T) {
			f, err := os.Open(s)
			assert.NoError(t, err)
			blockScanner := bufio.NewScanner(f)
			blockScanner.Split(ScanBlock)

			counter := 0
			for blockScanner.Scan() {
				counter += 1
				assert.NoError(t, blockScanner.Err(), "count is %d %s", counter, s)

				outputs[i] += blockScanner.Text()
			}
			assert.NoError(t, blockScanner.Err(), "%s", s)
			// assert.Equal(t, 3, counter, "test data should only have three blocks presents")
			f.Close()
		})
	}
	assert.Equal(t, outputs[0], outputs[1])
}
