package check_go

import (
	"bytes"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"golang.org/x/crypto/sha3"
)

func hashKey(userAgent string, XForwarded string) string {
	hasher := sha3.Sum256([]byte(userAgent + ":" + XForwarded))

	return string(hasher[:])
}

func bloomKey(userAgent string, XForwarded string) bool {
	filter := bloom.New(10000, 3)
	var buf bytes.Buffer
	bytesWritten, err := filter.WriteTo(&buf)
	if err != nil {
		fmt.Println(err.Error())
	}
	filter.Add([]byte("2504fda9c530eb977444kdr7e4b833291fe2b9428f4ffenw38a8511236fa3ad7"))
	filter.Add([]byte("2504fda9c530eb977444jd4564b833291fe2b9428f4f7aa028andhsu36fa3ad7"))
	filter.Add([]byte("2504fda9c530hdyfr74445537e4b833291fe2b9428f4f7aa028a85184hyfa3ad7"))
	var g bloom.BloomFilter
	bytesRead, err := g.ReadFrom(&buf)
	if err != nil {
		fmt.Println("read from err")
		fmt.Println(err.Error())
	} else {
		fmt.Println(bytesRead)
	}
	if bytesWritten != bytesRead {
		fmt.Println("not =")
	}
	return filter.Test([]byte("2504fda9c530eb9774445537e4b833291fe2b9428f4f7aa028a8511236fa3ad7"))
}
