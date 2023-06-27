package base64

import (
	"encoding/base64"
	"gitee.com/quant1x/gox/api"
	"golang.org/x/exp/slices"
	"math/rand"
	"time"
)

const (
	BASE64CHAR = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

// Base64GetEncoder 生成一个新的base64的码表
func Base64GetEncoder() string {
	encoder := slices.Clone(api.String2Bytes(BASE64CHAR))
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(encoder), func(i, j int) {
		encoder[i], encoder[j] = encoder[j], encoder[i]
	})
	return api.Bytes2String(encoder)
}

func getBase64Encoder(seed int64) string {
	encoder := slices.Clone(api.String2Bytes(BASE64CHAR))
	rand.Seed(seed)
	rand.Shuffle(len(encoder), func(i, j int) {
		encoder[i], encoder[j] = encoder[j], encoder[i]
	})
	return api.Bytes2String(encoder)
}

type PseudoBase64 struct {
	base64Chars string
	encoding    *base64.Encoding
}

func NewBase64(encoder string) PseudoBase64 {
	return PseudoBase64{
		base64Chars: encoder,
		encoding:    base64.NewEncoding(encoder),
	}
}

func NewPseudoBase64(seed int64) PseudoBase64 {
	encoder := getBase64Encoder(seed)
	return PseudoBase64{
		base64Chars: encoder,
		encoding:    base64.NewEncoding(encoder),
	}
}

func (codec PseudoBase64) Encode(s string) string {
	return codec.encoding.EncodeToString(api.String2Bytes(s))
}

func (codec PseudoBase64) Decode(s string) (string, error) {
	data, err := codec.encoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return api.Bytes2String(data), nil
}
