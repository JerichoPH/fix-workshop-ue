package tools

import (
	"fmt"
	"strings"
)

const (
	ORIGINAL = "\033["   // 起始标识符
	FINISHED = "\033[0m" // 终止标识符

	COLOR_BLACK  = "0" // 黑色
	COLOR_RED    = "1" //  红色
	COLOR_GREEN  = "2" // 绿色
	COLOR_YELLOW = "3" // 黄色
	COLOR_BLUE   = "4" // 蓝色
	COLOR_PURPLE = "5" // 紫色
	COLOR_CYAN   = "6" // 青色
	COLOR_WHITE  = "7" // 白色

	STYLE_DEFAULT     = "0" // 终端默认
	STYLE_DARK        = "1" // 变暗
	STYLE_HIGHLIGHT   = "2" // 高亮
	STYLE_ITALIC      = "3" // 倾斜
	STYLE_UNDERLINE   = "4" // 下横线
	STYLE_BLINK       = "5" // 闪烁
	STYLE_INVERSE     = "7" // 反白
	STYLE_INVISIBLE   = "8" // 不可见
	STYLE_DELETE_LINE = "9" // 删除线
)

type StdoutHelper struct {
	content interface{}
	fgColor string
	bgColor string
	style   string
}

// GetContent 获取内容
func (cls *StdoutHelper) GetContent() string {
	return cls.GetContentAndNext("")
}

func (cls *StdoutHelper) GetContentAndNext(next interface{}) string {
	r := fmt.Sprintf("%s%v%s%v", cls.getOriginal(), cls.content, cls.getFinished(), next)
	return r
}

// SetContent 设置内容
func (cls *StdoutHelper) SetContent(content interface{}) *StdoutHelper {
	cls.content = content
	return cls
}

// GetFgColor 获取前景色
func (cls *StdoutHelper) GetFgColor() string {
	return "3" + cls.fgColor
}

// SetFgColor 设置前景色
func (cls *StdoutHelper) SetFgColor(fgColor string) *StdoutHelper {
	cls.fgColor = fgColor
	return cls
}

// GetBgColor 获取背景色
func (cls *StdoutHelper) GetBgColor() string {
	return "4" + cls.bgColor
}

// SetBgColor 设置背景色
func (cls *StdoutHelper) SetBgColor(bgColor string) *StdoutHelper {
	cls.bgColor = bgColor
	return cls
}

// GetStyle 获取样式
func (cls *StdoutHelper) GetStyle() string {
	return cls.style
}

// SetStyle 设置样式
func (cls *StdoutHelper) SetStyle(style string) *StdoutHelper {
	cls.style = style
	return cls
}

// 获取起始标识符
func (cls *StdoutHelper) getOriginal() string {
	var options []string

	if cls.GetFgColor() != "" {
		options = append(options, cls.GetFgColor())
	}
	if cls.GetBgColor() != "" {
		options = append(options, cls.GetBgColor())
	}
	if cls.GetStyle() != "" {
		options = append(options, cls.GetStyle())
	}

	r := ORIGINAL + strings.Join(options, ";") + "m"
	return r
}

// 获取终止标识符
func (cls *StdoutHelper) getFinished() string {
	return FINISHED
}

// StdoutSuccess 成功格式
func StdoutSuccess(content interface{}, style string) *StdoutHelper {
	ins := &StdoutHelper{style: STYLE_HIGHLIGHT, fgColor: COLOR_WHITE, bgColor: COLOR_GREEN, content: content}
	if style != "" {
		ins.SetStyle(style)
	}
	return ins
}

// StdoutInfo 消息
func StdoutInfo(content interface{}, style string) *StdoutHelper {
	ins := &StdoutHelper{style: STYLE_HIGHLIGHT, fgColor: COLOR_WHITE, bgColor: COLOR_CYAN, content: content}
	if style != "" {
		ins.SetStyle(style)
	}
	return ins
}

// StdoutComment 注释
func StdoutComment(content interface{}, style string) *StdoutHelper {
	ins := &StdoutHelper{style: STYLE_HIGHLIGHT, fgColor: COLOR_YELLOW, content: content}
	if style != "" {
		ins.SetStyle(style)
	}
	return ins
}

// StdoutWarning 警告
func StdoutWarning(content interface{}, style string) *StdoutHelper {
	ins := &StdoutHelper{style: STYLE_HIGHLIGHT, fgColor: COLOR_PURPLE, content: content}
	if style != "" {
		ins.SetStyle(style)
	}
	return ins
}

// StdoutWrong 错误
func StdoutWrong(content interface{}, style string) *StdoutHelper {
	ins := &StdoutHelper{style: STYLE_HIGHLIGHT, fgColor: COLOR_WHITE, bgColor: COLOR_RED, content: content}
	if style != "" {
		ins.SetStyle(style)
	}
	return ins
}
