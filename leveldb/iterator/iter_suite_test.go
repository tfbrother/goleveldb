package iterator_test

import (
	"testing"

	"github.com/tfbrother/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
