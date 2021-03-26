package hanlp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/xxjwxc/public/myhttp"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/tools"
)

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

// Parse 解析获取
/*
opts 是临时参数部分
*/
func (h *hanlp) Parse(text string, opts ...Option) (string, error) {
	options := h.opts
	for _, f := range opts { // 执行自定义option
		f(&options)
	}

	req := &HanReq{
		Text:      text,             // 文本
		Language:  options.Language, // 语言(zh,mnt)
		Tasks:     options.Tasks,    // 任务列表()
		SkipTasks: options.SkipTasks,
	}
	b, err := myhttp.PostHeader(options.URL+"/parse", tools.JSONDecode(req), getHeader(options))
	if err != nil {
		mylog.Error(err)
		return "", err
	}

	return string(b), nil
}

// Parse 解析获取
/*
opts 是临时参数部分
*/
func (h *hanlp) ParseObj(text string, opts ...Option) (*HanResp, error) {
	options := h.opts
	for _, f := range opts { // 执行自定义option
		f(&options)
	}

	req := &HanReq{
		Text:      text,             // 文本
		Language:  options.Language, // 语言(zh,mnt)
		Tasks:     options.Tasks,    // 任务列表()
		SkipTasks: options.SkipTasks,
	}
	b, err := myhttp.PostHeader(options.URL+"/parse", tools.JSONDecode(req), getHeader(options))
	if err != nil {
		mylog.Error(err)
		return nil, err
	}

	return marshalHanResp(b)
}

// ParseAny 解析获取
/*
opts 是临时参数部分
*/
func (h *hanlp) ParseAny(text string, resp interface{}, opts ...Option) error {
	reqType := reflect.TypeOf(resp)
	if reqType.Kind() != reflect.Ptr {
		return fmt.Errorf("req type not a pointer:%v", reqType)
	}

	options := h.opts
	for _, f := range opts { // 执行自定义option
		f(&options)
	}

	req := &HanReq{
		Text:      text,             // 文本
		Language:  options.Language, // 语言(zh,mnt)
		Tasks:     options.Tasks,    // 任务列表()
		SkipTasks: options.SkipTasks,
	}
	b, err := myhttp.PostHeader(options.URL+"/parse", tools.JSONDecode(req), getHeader(options))
	if err != nil {
		mylog.Error(err)
		return err
	}

	switch v := resp.(type) {
	case *string:
		*v = string(b)
	case *[]byte:
		*v = b
	case *HanResp:
		tmp, e := marshalHanResp(b)
		*v, err = *tmp, e
	default:
		err = json.Unmarshal(b, v)
	}

	if err != nil {
		return err
	}

	return nil
}

// 解析obj
func marshalHanResp(b []byte) (*HanResp, error) {
	var hr hanResp
	err := json.Unmarshal(b, &hr)
	if err != nil {
		mylog.Error(err)
		return nil, err
	}
	resp := &HanResp{
		TokFine:   hr.TokFine,
		TokCoarse: hr.TokCoarse,
		PosCtb:    hr.PosCtb,
		PosPku:    hr.PosPku,
		Pos863:    hr.Pos863,
	}

	// ner/pku
	for _, v := range hr.NerPku {
		var tmp []NerTuple
		for _, v1 := range v {
			switch t := v1.(type) {
			case []interface{}:
				{
					tmp = append(tmp, NerTuple{
						Entity: t[0].(string),       // 实体
						Type:   t[1].(string),       // 类型
						Begin:  int(t[2].(float64)), // 开始点
						End:    int(t[3].(float64)),
					})
				}
			default:
				mylog.Error("%v : not unmarshal", t)
			}
		}
		resp.NerPku = append(resp.NerPku, tmp)
	}
	// ----------end

	// NerPku:    hr.NerPku,
	// NerMsra      [][]interface{}   `json:"ner/msra"`      // 命名实体识别 https://hanlp.hankcs.com/docs/annotations/ner/msra.html
	// NerOntonotes [][]interface{}   `json:"ner/ontonotes"` // 命名实体识别 https://hanlp.hankcs.com/docs/annotations/ner/ontonotes.html
	// Srl          [][][]interface{} `json:"srl"`           // 语义角色标注 其中谓词被标记为pred https://hanlp.hankcs.com/docs/annotations/srl/index.html
	// Dep          [][]interface{}   `json:"dep"`           // 依存句法分析 https://hanlp.hankcs.com/docs/annotations/dep/index.html
	// Sdp          [][]interface{}   `json:"sdp"`           // 语义依存分析 https://hanlp.hankcs.com/docs/annotations/sdp/index.html
	// Con          []interface{}
	return nil, nil
}

func getHeader(opts Options) http.Header {
	header := make(http.Header)
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	header.Add("Authorization", "Basic "+opts.Auth)
	return header
}
