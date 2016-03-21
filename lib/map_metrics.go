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

type MapHistogram map[string]*Histogram

var (
	reportHist string
)

func (mh MapHistogram) Init(report string) {
	reportHist = report
}

func (mh MapHistogram) Add(r *Result) {
	v, ok := mh[r.Url]
	if ok {
		v.Add(r)
	} else {
		var h Histogram
		h.Buckets.UnmarshalText([]byte(reportHist[4:]));
		h.Add(r)
		mh[r.Url] = &h
	}
}
