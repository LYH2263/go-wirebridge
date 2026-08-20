package frame

import "encoding/hex"

// DecodeHex 从 hex 字符串解码帧。
func DecodeHex(s string, maxFrame int) (Frame, error) {
	raw, err := hex.DecodeString(s)
	if err != nil {
		return Frame{}, err
	}
	return Decode(raw, maxFrame, nil)
}

// EncodeToHex 编码为 hex。
func EncodeToHex(f Frame, maxFrame int) (string, error) {
	raw, err := Encode(f, maxFrame, nil)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(raw), nil
}
