package goPrint

type Ends struct {
	Start string
	End string
}

func (m *Bar)SetEnds(start,end string)  {
	ends:=Ends{Start:start,End:end}
	m.progressEnds=ends
}
