package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sort"
	"strconv"
)

type minheap []int

func (h minheap) Len() int { return len(h) }
func (h minheap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h minheap) Top() int {
	return h[0]
}

type maxheap []int

func (h maxheap) Len() int { return len(h) }
func (h maxheap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxheap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h maxheap) Top() int {
	return h[0]
}

type RequestStat struct {
	RequestIdentifier string
	Called            int
	Times             []float64
	maxh              *maxheap
	minh              *minheap
	sum               int
}

func NewRequestStat() *RequestStat {
	r := &RequestStat{
		maxh: &maxheap{},
		minh: &minheap{},
	}
	heap.Init(r.maxh)
	heap.Init(r.minh)
	return r
}

func (r *RequestStat) Add(num int) {
	// Add to one of the heaps
	if r.maxh.Len() == 0 || num <= r.maxh.Top() {
		heap.Push(r.maxh, num)
	} else {
		heap.Push(r.minh, num)
	}

	// Balance heaps
	if r.maxh.Len() > r.minh.Len()+1 {
		heap.Push(r.minh, heap.Pop(r.maxh))
	} else if r.minh.Len() > r.maxh.Len() {
		heap.Push(r.maxh, heap.Pop(r.minh))
	}
	r.sum += num
}

func (r *RequestStat) Median() float64 {
	if r.maxh.Len() > r.minh.Len() {
		return float64(r.maxh.Top())
	} else if r.maxh.Len() < r.minh.Len() {
		return float64(r.minh.Top())
	}

	return float64(r.maxh.Top()+r.minh.Top()) / 2
}

func (r *RequestStat) Mean() float64 {
	return float64(r.sum) / float64(r.maxh.Len()+r.minh.Len())
}

func main() {
	fileURL := "https://gist.githubusercontent.com/bss/6dbc7d4d6d2860c7ecded3d21098076a/raw/244045d24337e342e35b85ec1924bca8425fce2e/sample.small.log"
	resp, err := http.Get(fileURL)
	if err != nil {
		log.Fatalf("Error fetching log file: %v", err)
	}
	defer resp.Body.Close()

	stats := map[string]*RequestStat{}
	scanner := bufio.NewScanner(resp.Body)

	//r := regexp.MustCompile(`method=(?P<method>\w+) path=(?P<path>/api/users/\d+(?:/\w+)?) .* connect=(?P<connect>\d+)ms service=(?P<service>\d+)ms`)
	r := regexp.MustCompile(`method=(?P<method>\w+) path=(?P<path>\/api\/users\/\d+(?:\/\w+)?) .* connect=(?P<connect>\d+)ms service=(?P<service>\d+)ms`)
	//r := regexp.MustCompile(`method=(?P<method>\w+) path=(?P<path>\/api\/users\/\d+\/\w+) .* connect=(?P<connect>\d+)ms service=(?P<service>\d+)ms`)

	for scanner.Scan() {
		ss := scanner.Text()
		matches := r.FindStringSubmatch(ss)
		if len(matches) == 0 {
			continue
		}

		method, path, connect, service := matches[1], matches[2], matches[3], matches[4]
		identifier := fmt.Sprintf("%s %s", method, path)

		if _, ok := stats[identifier]; !ok {
			stats[identifier] = &RequestStat{
				RequestIdentifier: identifier,
				Called:            0,
			}
		}

		connectTime, _ := strconv.Atoi(connect)
		serviceTime, _ := strconv.Atoi(service)
		totalTime := float64(connectTime + serviceTime)

		stats[identifier].Called++
		stats[identifier].Times = append(stats[identifier].Times, totalTime)
	}

	results := []map[string]interface{}{}
	for _, stat := range stats {
		sort.Float64s(stat.Times)
		median := stat.Times[len(stat.Times)/2]
		if len(stat.Times)%2 == 0 {
			median = (stat.Times[len(stat.Times)/2-1] + median) / 2
		}

		var sum float64
		for _, v := range stat.Times {
			sum += v
		}

		mean := sum / float64(len(stat.Times))

		results = append(results, map[string]interface{}{
			"request_identifier":   stat.RequestIdentifier,
			"called":               stat.Called,
			"response_time_mean":   mean,
			"response_time_median": median,
		})
	}

	for _, r := range results {
		fmt.Println(r)
	}
}
