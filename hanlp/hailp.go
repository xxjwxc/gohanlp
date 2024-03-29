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
		Language:  options.Language, // 语言(zh,mnl)
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
		Language:  options.Language, // 语言(zh,mnl)
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
		Language:  options.Language, // 语言(zh,mnl)
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

	// ner/pku 命名实体识别
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

	// ner/msra 命名实体识别
	for _, v := range hr.NerMsra {
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
		resp.NerMsra = append(resp.NerMsra, tmp)
	}
	// ----------end

	// ner/ontonotes 命名实体识别
	for _, v := range hr.NerOntonotes {
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
		resp.NerOntonotes = append(resp.NerOntonotes, tmp)
	}
	// ----------end

	// srl 语义角色标注
	for _, v := range hr.Srl {
		var tmp [][]SrlTuple
		for _, v1 := range v {
			var tmp1 []SrlTuple
			for _, v2 := range v1 {
				switch t := v2.(type) {
				case []interface{}:
					{
						tmp1 = append(tmp1, SrlTuple{
							ArgPred: t[0].(string),       // 参数
							Label:   t[1].(string),       // 标签
							Begin:   int(t[2].(float64)), // 开始点
							End:     int(t[3].(float64)),
						})
					}
				default:
					mylog.Error("%v : not unmarshal", t)
				}
			}
			tmp = append(tmp, tmp1)
		}
		resp.Srl = append(resp.Srl, tmp)
	}
	// -------------end

	// dep 依存句法分析
	for _, v := range hr.Dep {
		var tmp []DepTuple
		for _, v1 := range v {
			switch t := v1.(type) {
			case []interface{}:
				{
					tmp = append(tmp, DepTuple{
						Head:     int(t[0].(float64)), // 头
						Relation: t[1].(string),       // 关系
					})
				}
			default:
				mylog.Error("%v : not unmarshal", t)
			}
		}
		resp.Dep = append(resp.Dep, tmp)
	}
	// ------------end
	// sdp 语义依存分析
	for _, v := range hr.Sdp {
		var tmp [][]DepTuple
		for _, v1 := range v {
			var tmp1 []DepTuple
			for _, v2 := range v1 {
				switch t := v2.(type) {
				case []interface{}:
					{
						tmp1 = append(tmp1, DepTuple{
							Head:     int(t[0].(float64)), // 开始点
							Relation: t[1].(string),       // 实体
						})
					}
				default:
					mylog.Error("%v : not unmarshal", t)
				}
			}
			tmp = append(tmp, tmp1)
		}
		resp.Sdp = append(resp.Sdp, tmp)
	}
	// ------------end
	// Con
	resp.Con = dealCon(hr.Con)
	// ------------end

	return resp, nil
}

func getHeader(opts Options) http.Header {
	header := make(http.Header)
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")
	if len(opts.Auth) > 0 {
		header.Add("Authorization", "Basic "+opts.Auth)
	}
	return header
}

func dealCon(info []interface{}) (re []ConTuple) {
	if len(info) == 0 {
		return nil
	}

	switch t := info[0].(type) {
	case string:
		{
			tmp1 := ConTuple{
				Key: t,
			}
			if len(info) == 2 {
				tmp1.Value = dealCon(info[1].([]interface{}))
			}
			// else { // 理论上不存在
			// 	fmt.Println(info)
			// }
			re = append(re, tmp1)
		}
	case []interface{}:
		{
			for _, t1 := range info {
				tmp1 := ConTuple{}
				tmp1.Value = dealCon(t1.([]interface{}))
				re = append(re, tmp1)
			}
		}
	}

	return re
}

// KeyphraseExtraction 关键词提取
/*
opts 是临时参数部分
*/
func (h *hanlp) KeyphraseExtraction(text string, opts ...Option) (map[string]float64, error) {
	options := h.opts
	for _, f := range opts { // 执行自定义option
		f(&options)
	}

	req := &HanReq{
		Text:      text,             // 文本
		Language:  options.Language, // 语言(zh,mnl)
		Tasks:     options.Tasks,    // 任务列表()
		SkipTasks: options.SkipTasks,
	}
	b, err := myhttp.PostHeader(options.URL+"/keyphrase_extraction", tools.JSONDecode(req), getHeader(options))
	if err != nil {
		mylog.Error(err)
		return nil, err
	}

	mp := make(map[string]float64)
	tools.JSONEncode(string(b), &mp)
	return mp, err
}
