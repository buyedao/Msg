```
println(msg.SkyBlue().Bold().TranslationMsg("hello"))
```


### 字体颜色设置
* func (m *Msg) Black() *Msg 
* func (m *Msg) Red() *Msg 
* func (m *Msg) Green() *Msg 
* func (m *Msg) Yellow() *Msg 
* func (m *Msg) Blue() *Msg 
* func (m *Msg) Purple() *Msg 
* func (m *Msg) SkyBlue() *Msg 
* func (m *Msg) White() *Msg 

### 背景颜色设置
* func (m *Msg) BackgrounBlack() *Msg 
* func (m *Msg) BackgrounRed() *Msg 
* func (m *Msg) BackgrounGreen() *Msg 
* func (m *Msg) BackgrounYellow() *Msg 
* func (m *Msg) BackgrounBlue() *Msg 
* func (m *Msg) BackgrounPurple() *Msg 
* func (m *Msg) BackgrounSkyBlue() *Msg 
* func (m *Msg) BackgrounWhite() *Msg 

### 字体样式设置
* func (m *Msg) Bold()*Msg

### 根据设置转换字符串显示方式
* func (m *Msg) TranslationMsg(msg string) string

