package goprint

import (
	"fmt"
	"strings"
)

func (m *Bar) PrintBar(currValue int) {
	m.SetCurrentValue(currValue)
	printStr := "\r" + fmt.Sprintf(" %c[0;0;%vm%v%c[0m",0x1B,m.Color.Notice,m.Notice,0x1B)
	if m.IsShowBar {
		printStr+=m.ProgressPrintString()
	}
	if m.IsShowPercent {
		printStr+=m.PercentPrintString()
	}
	if m.IsShowRatio {
		printStr+=m.RatioPrintString()
	}
	fmt.Print(printStr)
}

func (m *Bar)ProgressPrintString()string{
	back:=m.GetBackString()
	printStr:=m.ProgressEnds.Start
	printStr+=strings.Replace(back,m.BackGraph,m.ProgressGraph,m.CurrentPrintGraphNumber())
	printStr+=m.ProgressEnds.End
	return fmt.Sprintf("%c[1;%v;%vm%s%c[0m", 0x1B,m.Color.Back,m.Color.Graph,printStr,0x1B)
}

func (m *Bar)PercentPrintString() string {
	return fmt.Sprintf(" %c[0;0;%vm%v%%%c[0m",0x1B,m.Color.Percent,m.CurrentRate,0x1B)
}

func (m *Bar)RatioPrintString() string {
	return fmt.Sprintf(" %c[0;0;%vm\t%v/%v%c[0m",0x1B,m.Color.Ratio,m.CurrentValue,m.TotalValue,0x1B)
}

func (m *Bar)PrintEnd(tip ... string)  {
	fmt.Printf("\n")
	if len(tip)>0 {
		fmt.Println(tip[0])
	}
}

func (m *Bar) GetBackString()  string {
	res:=""
	for i := 0; i < m.ProgressGraphTotal; i++ {
		res+= m.BackGraph
	}
	return res
}
