package goPrint

const defaultCount = 50

type Bar struct {
	totalValue         int
	currentValue       int
	currentRate        int
	progressGraph      string
	progressGraphTotal int
	backGraph          string
	isShowPercent      bool
	isShowRatio        bool
	isShowBar          bool
	notice             string
	progressEnds       Ends
	color              BarColor
}

func NewBar(totalVale int) *Bar {
	res := &Bar{
		totalValue:         totalVale,
		currentValue:       0,
		progressGraphTotal: defaultCount,
		progressGraph:      "█",
		backGraph:          " ",
		progressEnds:       Ends{Start: "[", End: "]"},
		isShowPercent:      true,
		isShowRatio:        true,
		isShowBar:          true,
	}
	res.CountCurrentRate()
	return res
}

func (m *Bar) CountCurrentRate() {
	if m.currentValue == 0 {
		m.currentRate = 0
	} else {
		m.currentRate = m.currentValue * 100 / m.totalValue
	}
}

func (m *Bar)SetGraph(graph string)  {
	m.progressGraph=graph
}

func (m *Bar) SetCurrentValue(value int) {
	m.currentValue = value
	m.CountCurrentRate()
}

func (m *Bar) CurrentPrintGraphNumber() int {
	if m.currentRate == 100 {
		return m.progressGraphTotal
	} else {
		return int(float64(m.currentRate) * (float64(m.progressGraphTotal) / float64(100)))
	}
}

func (m *Bar) HideProgressBar() {
	m.isShowBar = false
}

func (m *Bar) HidePercent() {
	m.isShowPercent = false
}

func (m *Bar) HideRatio() {
	m.isShowRatio = false
}

func (m *Bar) SetNotice(notice string) {
	m.notice = notice
}

func (m *Bar)SetProgressGraphTotal(totalGraph int)  {
	m.SetProgressGraphTotal(totalGraph)
}

func (m *Bar)SetBackGraph(graph string)  {
	m.backGraph=graph
}
