package cm

type IntProvider interface {
	Int(key string) (int, error)
}

type IntProviderFunc func(key string) (int, error)

func (f IntProviderFunc) Int(key string) (int, error) {
	return f(key)
}

type StringProvider interface {
	String(key string) (string, error)
}

type StringProviderFunc func(key string) (string, error)

func (f StringProviderFunc) String(key string) (string, error) {
	return f(key)
}
