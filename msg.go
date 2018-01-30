package Msg

import (
	"fmt"
	"strings"
)

type Msg struct {
	color      string
	background string
	bold string
}

func (m *Msg) Black() *Msg {
	m.color = "30"
	return m
}
func (m *Msg) Red() *Msg {
	m.color = "31"
	return m
}

func (m *Msg) Green() *Msg {
	m.color = "32"
	return m
}
func (m *Msg) Yellow() *Msg {
	m.color = "33"
	return m
}
func (m *Msg) Blue() *Msg {
	m.color = "34"
	return m
}
func (m *Msg) Purple() *Msg {
	m.color = "35"
	return m
}

func (m *Msg) SkyBlue() *Msg {
	m.color = "36"
	return m
}
func (m *Msg) White() *Msg {
	m.color = "37"
	return m
}

func (m *Msg) BackgrounBlack() *Msg {
	m.background = "40"
	return m
}
func (m *Msg) BackgrounRed() *Msg {
	m.background = "41"
	return m
}
func (m *Msg) BackgrounGreen() *Msg {
	m.background = "42"
	return m
}
func (m *Msg) BackgrounYellow() *Msg {
	m.background = "43"
	return m
}
func (m *Msg) BackgrounBlue() *Msg {
	m.background = "44"
	return m
}
func (m *Msg) BackgrounPurple() *Msg {
	m.background = "45"
	return m
}

func (m *Msg) BackgrounSkyBlue() *Msg {
	m.background = "46"
	return m
}
func (m *Msg) BackgrounWhite() *Msg {
	m.background = "47"
	return m
}
func (m *Msg) Bold()*Msg{
	m.bold = "1"
	return m
}
func (m *Msg) TranslationMsg(msg string) string {
	strStyle := make([]string,0)
	if m.bold != ""{
		strStyle = append(strStyle,m.bold)
	}
	if m.background != ""{
		strStyle = append(strStyle,m.background)
	}
	if m.color != ""{
		strStyle = append(strStyle,m.color)
	}
	strMsg := strings.Join(strStyle,";")
	if strMsg != ""{
		return fmt.Sprintf("\033[%sm%s\033[0m",strMsg,msg)
	}
	return msg;
}

