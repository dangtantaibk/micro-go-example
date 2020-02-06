package concurrency

type concurrencyFunc func() error

// Concurrency struct
type Concurrency struct {
	f []concurrencyFunc
}

// New concurrency
func New() *Concurrency {
	return &Concurrency{}
}

// Add concurrency
func (cc *Concurrency) Add(f ...concurrencyFunc) *Concurrency {
	cc.f = append(cc.f, f...)
	return cc
}

// Do concurrency
func (cc *Concurrency) Do() error {
	return Do(cc.f...)
}

// Do concurrency
func Do(ccf ...concurrencyFunc) error {
	l := len(ccf)
	if l == 0 {
		return nil
	} else if l == 1 {
		return ccf[0]()
	}
	ec := make(chan error, l)
	for _, _f := range ccf {
		f := _f
		go func() { ec <- f() }()
	}
	for e := range ec {
		if e != nil {
			return e
		}
		l--
		if l == 0 {
			return nil
		}
	}

	return nil

}
