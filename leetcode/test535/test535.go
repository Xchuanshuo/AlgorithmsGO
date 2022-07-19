package test535

import "math/rand"

const str = "abcdefghijkmnopqrstuvwxyz0123456789"

type Codec struct {
	mp map[string]string
}

func Constructor() Codec {
	return Codec{mp: map[string]string{}}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	var bytes []byte
	for i := 0; i < 6; i++ {
		bytes = append(bytes, str[rand.Intn(len(str))])
	}
	var t = string(bytes)
	this.mp[t] = longUrl
	return "http://tinyurl.com/" + t
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	var last = shortUrl[19:]
	return this.mp[last]
}
