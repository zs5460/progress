package progress

import (
	"fmt"
	"strings"
)

type Progress struct {
	percent int
	current int
	total   int
	rate    string
	graph   string
}

func New(total int) *Progress {
	p := new(Progress)
	p.current = 0
	p.total = total
	p.graph = ">"
	p.percent = p.GetPercent()
	return p
}

func (p *Progress) GetPercent() int {
	return int(float32(p.current) / float32(p.total) * 100)
}

func (p *Progress) Add(i int) {

	p.current += i

	if p.current > p.total {
		return
	}

	p.percent = p.GetPercent()
	p.rate = strings.Repeat(p.graph, p.percent/2)

	fmt.Printf("\r[%-50s]%4d%% %8d/%d", p.rate, p.percent, p.current, p.total)

	if p.current == p.total {
		p.Done()
	}
}

func (p *Progress) SetGraph(s string) {
	p.graph = s
}

func (p *Progress) Incr() {
	p.Add(1)
}

func (p *Progress) Done() {
	fmt.Println()
}
