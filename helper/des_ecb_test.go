package helper

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestTripleEcbDesEncrypt(t *testing.T) {
	b, err := TripleEcbDesEncrypt([]byte("1"), []byte("4B8202A3BB3B4D650306BEE15DF35E13F70D2D704A4E4F8C"))

	data := base64.StdEncoding.EncodeToString(b)

	fmt.Println(data)
	fmt.Println(err)
}

func TestTripleEcbDesDecrypt(t *testing.T) {

	b, _ := base64.StdEncoding.DecodeString("0swe7EwSDGnUfZCnE7oYIXDwbs8M80qGJuJg52yutPY8sZFl5gnDo7Gk09+MaHz5uwBf4nb8f9Q+aYxUMTL2kYDrnkqwSe+BwhJ5UmbXabbjSbKOCDOi3+BUJX2lJT5N4tp/YZ4u0P9tI5P1QGWMV+RGoCfy3taLpq4FdKqFK7KmOqKgvURIfrm3PS4kBPFNLQiXTeFWhTI0HRQAaWtYnQJmRh+BBp5cpjqioL1ESH44FwY0UhqIzXH4IM7d70FjH2JW2FV7b3JClVaDOsJMYX0EOLERm10acQr8f2O9Ukbzi17i7B+GH9n5f5PZ9ntC0nCYgWJij+t8kADzy3wxcJp9wEGxtg57tyiwcTRckPu8xLbEGetunhnl7cxAdwF57M4DXSrGTibpvOYFFIXCRbqp47Cdwq9GP4Xdpb5NDRE=")

	b, err := TripleEcbDesDecrypt(b, []byte("9b15ede720a940c0a44939e252c5b5ee"))

	fmt.Println(string(b))
	fmt.Println(err)
}
