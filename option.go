package randomstr

type Option func(*Randstr)

func WithDigit() Option {
	return func(r *Randstr) {
		r.rs = append(r.rs, []rune("0123456789")...)
	}
}

func WithLowerCase() Option {
	return func(r *Randstr) {
		r.rs = append(r.rs, []rune("abcdefghijklmnopqrstuvwxyz")...)
	}
}

func WithUpperCase() Option {
	return func(r *Randstr) {
		r.rs = append(r.rs, []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")...)
	}
}

func WithLetter() Option {
	return func(r *Randstr) {
		r.rs = append(r.rs, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")...)
	}
}

func WithSpechar() Option {
	return func(r *Randstr) {
		r.rs = append(r.rs, []rune(`!"#$%&'()*+,-./:;<>?@[\]^_{|}~`)...)
	}
}

func WithCustom(s string) Option {
	return func(r *Randstr) {
		r.rs = append(r.rs, []rune(s)...)
	}
}
