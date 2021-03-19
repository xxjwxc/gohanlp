package hanlp

type hanlp struct {
	opts Options
}

// HanLPClient 新建客户端
func HanLPClient(opts ...Option) *hanlp {
	options := Options{ // 默认值
		URL:      "https://www.hanlp.com/api",
		Language: "zh",
	}

	for _, f := range opts { // 执行自定义option
		f(&options)
	}

	return &hanlp{
		opts: options,
	}
}

func (h *hanlp) Parse()
