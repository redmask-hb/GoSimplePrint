# GoSimplePrint
A simple terminal progress bar printing with golang.
Golang 屏幕进度条打印

使用：`go get -u github.com/redmask-hb/GoSimplePrint/goPrint`

例子：

```
func main()  {
	bar:=goPrint.NewBar(20)
	bar.SetNotice("进度条：")
	bar.SetGraph(">")
	for i:=1;i<=20;i++  {
		bar.PrintBar(i)
		time.Sleep(time.Second)
	}
	bar.PrintEnd("Finish!")
}
```


效果：

进度条：[>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>] 100%  	    20/20

Finish!

GoPrint是用golang写的一个简单进度条打印功能包。你只需要简单传入当前值和总值就可以实现进度条功能。颜色，样式，进度图块，等，可以自己根据需要定义，不需要的组成部分可以随意隐藏。

1. 初始化

```
bar:=goPrint.NewBar(totalValue)  //totalValue是进度最后的终止值
```

2、打印
```
bar.PrintBar(curValue)  //curValue 是进度的当前值，在for循环中，动态将当前值传入到bar对象
```

3、结束

```
bar.PrintEnd(tip) //tip为可选参数  
```
当进度条完成后，可以使用PrintEnd() 给出一个结束提示，可以传入一个要打印的内容，也可以不传。不传即为在进度条最后加一个"\n"

4、配置参数

![输入图片说明](https://images.gitee.com/uploads/images/2020/1210/172004_1688d385_8040941.jpeg "QQ图片20201210165306.jpg")

进度条各组成部分命名规则如上图所示

（1）设置

```
bar.SetGraph(graph string)  //设置进度条图块，默认"█"
bar.SetNotice(notice string) //设置进度条前的提示内容，可选
bar.SetProgressGraphTotal(totalGraph int) //设置进度条完成所填图块的总数，默认：50
bar.SetEnds(start,end string)  //设置进度条端块，默认为 srart:"[",end:"]"
bar.SetBackGraph(graph string) //设置背面被进度条替换的图形，默认" "
```
（2）隐藏

```
bar.HideProgressBar()   //隐藏进度条
bar.HidePercent()   //隐藏百分比
bar.HideRatio()    //隐藏比值
```

（3）颜色设置 （由于windows的cmd颜色打印，需要调用系统API，目前win操作系统下颜色设置无效）
```
bar.SetColor(color BarColor) //配置所有的颜色选项
bar.SetBackColor(color int) //设置进度条背景色
bar.SetRatioColor(color int) //设置比值颜色
bar.SetPercentColor(color int) //设置百分比颜色
bar.SetNoticeColor(color int) //设置提示颜色
```

颜色为int数据类型，对比下表配置（配置背景色用背景色值，前景色用前景色值）


| 前景色 | 背景色 | 颜色  |
|-----|-----|-----|
| 30  | 40  | 黑色  |
| 31  | 41  | 红色  |
| 32  | 42  | 绿色  |
| 33  | 43  | 黄色  |
| 34  | 44  | 蓝色  |
| 35  | 45  | 紫红色 |
| 36  | 46  | 青蓝色 |
| 37  | 47  | 白色  |


改动：

v1.1.0

1、修正了背景颜色设置无效的BUG。

2、增加了颜色预置定义 

```
bar.SetNoticeColor(goPrint.FontColor.Yellow)
bar.SetBackColor(goPrint.BackColor.Red)
```
3、修正了win打印错位的问题。

