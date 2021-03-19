package hanlp

// HanReq hanlp 请求参数
type HanReq struct {
	Text      string   `json:"text,omitempty"`     // 文本
	Language  string   `json:"language,omitempty"` // 语言(zh,mnt)
	Tasks     []string `json:"tasks,omitempty"`    // 任务列表()
	SkipTasks []string `json:"skip_tasks"`
}

// HanResp hanlp 返回参数
type HanResp struct {
}
