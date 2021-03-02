package goPrint

import (
	"fmt"
	"strings"
)

func (m *Bar) PrintBar(currValue int) {
	m.SetCurrentValue(currValue)
	printStr := "\r" + m.NoticePrintString()
	if m.isShowBar {
		printStr+=m.ProgressPrintString()
	}
	if m.isShowPercent {
		printStr+=m.PercentPrintString()
	}
	if m.isShowRatio {
		printStr+=m.RatioPrintString()
	}
	fmt.Print(printStr)
}

func (m *Bar)NoticePrintString()string {
	if OS=="windows" {
		return m.notice
	}else {
		return fmt.Sprintf(" %c[%vm%v%c[0m",0x1B,m.color.Notice,m.notice,0x1B)
	}
}

func (m *Bar)ProgressPrintString()string{
	back:=m.GetBackString()
	printStr:=m.progressEnds.Start
	printStr+=strings.Replace(back,m.backGraph,m.progressGraph,m.CurrentPrintGraphNumber())
	printStr+=m.progressEnds.End
	if OS=="windows" {
		return printStr
	}else {
		return fmt.Sprintf("%c[%v;%vm%s%c[0m", 0x1B,m.color.Graph,m.color.Back,printStr,0x1B)
	}
}

func (m *Bar)PercentPrintString() string {
	if OS=="windows" {
		return fmt.Sprintf(" %v%%",m.currentRate)
	}
	return fmt.Sprintf(" %c[%vm%v%%%c[0m",0x1B,m.color.Percent,m.currentRate,0x1B)
}

func (m *Bar)RatioPrintString() string {
	if OS=="windows" {
		return fmt.Sprintf(" %v/%v",m.currentValue,m.totalValue)
	}
	return fmt.Sprintf(" %c[%vm\t%v/%v%c[0m",0x1B,m.color.Ratio,m.currentValue,m.totalValue,0x1B)
}

func (m *Bar)PrintEnd(tip ... string)  {
	fmt.Printf("\n")
	if len(tip)>0 {
		fmt.Println(tip[0])
	}
}

func (m *Bar) GetBackString()  string {
	res:=""
	for i := 0; i < m.progressGraphTotal; i++ {
		res+= m.backGraph
	}
	return res
}
