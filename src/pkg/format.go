package pkg

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

var secret = byte(0x5A)

func EncryptCursor(id int64) string {
	idBytes := []byte(fmt.Sprintf("%d", id))
	for i := range idBytes {
		idBytes[i] ^= secret
	}
	return base64.URLEncoding.EncodeToString(idBytes)
}

func DecryptCursor(cursor string) (int64, error) {
	decoded, err := base64.URLEncoding.DecodeString(cursor)
	if err != nil {
		return 0, err
	}
	for i := range decoded {
		decoded[i] ^= secret
	}
	var id int64
	_, err = fmt.Sscanf(string(decoded), "%d", &id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
