package hanlp

import "time"

// Options opts define
type Options struct {
	URL      string
	Auth     string
	Language string
	Timeout  time.Time
}

// Option opts list func
type Option func(*Options)

// WithURL 设置hanlp地址
func WithURL(url string) Option {
	return func(o *Options) {
		o.URL = url
	}
}

// WithAuth 设置授权码
func WithAuth(auth string) Option {
	return func(o *Options) {
		o.Auth = auth
	}
}

// WithLanguage 设置语言
func WithLanguage(language string) Option {
	return func(o *Options) {
		o.Language = language
	}
}

// WithTimeout 调用超时设置
func WithTimeout(timeout time.Time) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}
