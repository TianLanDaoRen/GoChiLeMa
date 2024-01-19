package markov

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

type MarkovModel struct {
	Events     []string
	PMat       *mat.Dense
	EventCnt   int
	EventIdMap map[string]int
	IdEventMap map[int]string
}

func NewMarkovModel(events []string) *MarkovModel {
	totEvents, eventIdMap, idEventMap := hashEvents(events)
	// Init empty pmat
	pmat := mat.NewDense(totEvents, totEvents, nil)
	// Calc pmat
	calcPmat(events, eventIdMap, pmat)
	return &MarkovModel{
		Events:     events,
		PMat:       pmat,
		EventCnt:   totEvents,
		EventIdMap: eventIdMap,
		IdEventMap: idEventMap,
	}
}

func (m *MarkovModel) GetNextPiN(N int) *mat.Dense {
	// Pi(0)
	pi0 := make([]float64, m.EventCnt)
	lasEvent := m.Events[len(m.Events)-1]
	lasEventId := m.EventIdMap[lasEvent]
	pi0[lasEventId] = 1
	pi0Vec := mat.NewDense(1, m.EventCnt, pi0)
	// Pi(1) = Pi(0) * PMat
	ansVec := mat.NewDense(1, m.EventCnt, nil)
	ansVec.Mul(pi0Vec, m.PMat)
	// Pi(N) = Pi(N-1) * PMat
	for i := 1; i < N; i++ {
		ansVec.Mul(ansVec, m.PMat)
	}
	return ansVec
}

func (m *MarkovModel) GetLimitPi() *mat.Dense {
	// A X = B A = P.T
	a := mat.DenseCopyOf(m.PMat.T())
	fmt.Println(a)
	// A = -P
	minus1 := make([]float64, m.EventCnt*m.EventCnt)
	for i := 0; i < m.EventCnt*m.EventCnt; i++ {
		minus1[i] = -1
	}
	minus1Mat := mat.NewDense(m.EventCnt, m.EventCnt, minus1)
	a.MulElem(a, minus1Mat)
	fmt.Println(a)
	// A = -P + Eye
	eyeMat := mat.NewDense(m.EventCnt, m.EventCnt, nil)
	for i := 0; i < m.EventCnt; i++ {
		eyeMat.Set(i, i, 1)
	}
	a.Add(a, eyeMat)
	fmt.Println(a)
	// B (n * 1) 列向量
	b := mat.NewDense(m.EventCnt, 1, nil)
	fmt.Println("b", b)
	// 求 InvA
	invA := mat.NewDense(m.EventCnt, m.EventCnt, nil)
	err := invA.Inverse(a)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("inva", invA)
	// 求 X = invA * B
	x := mat.NewDense(m.EventCnt, 1, nil)
	x.Mul(invA, b)
	return x
}

func (m *MarkovModel) GetNextPiNEvent(N int) map[string]float64 {
	vec := m.GetNextPiN(N).RowView(0)
	events := map[string]float64{}
	for i := 0; i < m.EventCnt; i++ {
		events[m.IdEventMap[i]] = vec.At(i, 0)
	}
	return events
}

func hashEvents(events []string) (int, map[string]int, map[int]string) {
	seen := make(map[string]bool) // 创建一个空 map
	eventIdMap := make(map[string]int)
	idEventMap := make(map[int]string)
	count := 0                     // 初始化计数为 0
	for _, event := range events { // 遍历 list
		if !seen[event] { // 如果元素不在 map 中
			seen[event] = true // 将其加入 map
			eventIdMap[event] = count
			idEventMap[count] = event
			count++ // 计数加一
		}
	}
	return count, eventIdMap, idEventMap
}

func calcPmat(events []string, eventsIdMap map[string]int, pmat *mat.Dense) {
	totEvents, _ := pmat.Dims()
	N := mat.NewDense(totEvents, totEvents, nil)
	M := make([]float64, totEvents)
	// By default, each food will at least transfer to others once
	for i := 0; i < totEvents; i++ {
		M[i] = float64(totEvents)
	}
	lasEvent := events[0]
	// Get Nij and Mi
	for idx, event := range events {
		if idx == 0 {
			continue
		}
		lasId, nowId := eventsIdMap[lasEvent], eventsIdMap[event]
		M[lasId]++
		N.Set(lasId, nowId, N.At(lasId, nowId)+1)
		lasEvent = event
	}
	// Loop
	for i := 0; i < totEvents; i++ {
		for j := 0; j < totEvents; j++ {
			pmat.Set(i, j, (N.At(i, j)+1)/M[i])
		}
	}
}
