package encodeS

import (
	b64 "encoding/base64"
	"fmt"
)

func EncodingURL(link string, limitation int) string {
	byteLink := []byte(link)
	stop := false
	for i := 0; i < len(byteLink); i++ {
		if !stop {
			byteLink = byteLink[1:]
		}
		if string(byteLink[i]) == "/" {
			stop = true
		}
		if string(byteLink[i]) != "/" && stop {
			break
		}
	}
	fmt.Print(string(byteLink))
	encoded := b64.URLEncoding.EncodeToString([]byte(link))

	return encoded[:limitation]
}
