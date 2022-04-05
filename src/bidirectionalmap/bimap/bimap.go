package bimap

import "fmt"

type Bimap struct {
	forward, inverse map[string]string
}

func (bi *Bimap) Store(key string, value string) {
	if bi.forward == nil {
		bi.forward = make(map[string]string)
		bi.inverse = make(map[string]string)
	}
	k, exists := bi.inverse[value]
	if exists { // value is already associated with k
		delete(bi.forward, k)
	}
	v, exists := bi.forward[key]
	if exists { // key is already associated with v
		delete(bi.inverse, v)
	}
	bi.forward[key] = value
	bi.inverse[value] = key
}

func (bi *Bimap) LookupValue(key string) (string, bool) {
	v, ok := bi.forward[key]
	return v, ok
}

func (bi *Bimap) LookupKey(value string) (string, bool) {
	k, ok := bi.inverse[value]
	return k, ok
}

func (bi *Bimap) String() string {
	return fmt.Sprintf("bi%v", bi.forward)
}
