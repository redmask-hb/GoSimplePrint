package goPrint

const DefaultCount = 50

type Bar struct {
	TotalValue         int
	CurrentValue       int
	CurrentRate        int
	ProgressGraph      string
	ProgressGraphTotal int
	BackGraph          string
	IsShowPercent      bool
	IsShowRatio        bool
	IsShowBar          bool
	Notice             string
	ProgressEnds       Ends
	Color              BarColor
}

func NewBar(totalVale int) *Bar {
	res := &Bar{
		TotalValue:         totalVale,
		CurrentValue:       0,
		ProgressGraphTotal: DefaultCount,
		ProgressGraph:      "█",
		BackGraph:          " ",
		ProgressEnds:       Ends{Start: "[", End: "]"},
		IsShowPercent:      true,
		IsShowRatio:        true,
		IsShowBar:          true,
	}
	res.CountCurrentRate()
	return res
}

func (m *Bar) CountCurrentRate() {
	if m.CurrentValue == 0 {
		m.CurrentRate = 0
	} else {
		m.CurrentRate = m.CurrentValue * 100 / m.TotalValue
	}
}

func (m *Bar)SetGraph(graph string)  {
	m.ProgressGraph=graph
}

func (m *Bar) SetCurrentValue(value int) {
	m.CurrentValue = value
	m.CountCurrentRate()
}

func (m *Bar) CurrentPrintGraphNumber() int {
	if m.CurrentRate == 100 {
		return m.ProgressGraphTotal
	} else {
		return int(float64(m.CurrentRate) * (float64(m.ProgressGraphTotal) / float64(100)))
	}
}

func (m *Bar) HideProgressBar() {
	m.IsShowBar = false
}

func (m *Bar) HidePercent() {
	m.IsShowPercent = false
}

func (m *Bar) HideRatio() {
	m.IsShowRatio = false
}

func (m *Bar) SetNotice(notice string) {
	m.Notice = notice
}

func (m *Bar)SetProgressGraphTotal(totalGraph int)  {
	m.SetProgressGraphTotal(totalGraph)
}

func (m *Bar)SetBackGraph(graph string)  {
	m.BackGraph=graph
}
