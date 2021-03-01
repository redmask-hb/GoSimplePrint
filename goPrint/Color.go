package goPrint

type BarColor struct {
	Graph int
	Back int
	Ratio int
	Percent int
	Notice int
}

func (m *Bar)SetProgressGraphColor(color int)  {
	m.Color.Graph=color
}

func (m *Bar)SetColor(color BarColor)  {
	m.Color=color
}

func (m *Bar)SetBackColor(color int)  {
	m.Color.Back=color
}

func (m *Bar)SetRatioColor(color int)  {
	m.Color.Ratio=color
}

func (m *Bar)SetPercentColor(color int)  {
	m.Color.Percent=color
}

func (m *Bar)SetNoticeColor(color int)  {
	m.Color.Notice=color
}

