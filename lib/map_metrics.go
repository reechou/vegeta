package vegeta

type MapMetrics map[string]*Metrics

func (mm MapMetrics) Add(r *Result) {
	v, ok := mm[r.Url]
	if ok {
		v.Add(r)
	} else {
		var m Metrics
		m.Add(r)
		mm[r.Url] = &m
	}
}

func (mm MapMetrics) Close() {
	for _, v := range mm {
		v.Close()
	}
}
