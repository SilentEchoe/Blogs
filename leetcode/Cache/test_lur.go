package main

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

func GetFromSource(getter Getter, key string) []byte {
	buf, err := getter.Get(key)
	if err == nil {
		return buf
	}
	return nil
}

func main() {
	// GetterFunc 类型的函数作为参数
	GetFromSource(GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	}), "hello")
}
