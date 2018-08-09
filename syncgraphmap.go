package syncgraphmap

import (
	"sync"
)

// GraphSingleMutex DataStructure containing urls and connections
// GraphSingleMutex is just a graph
type GraphSingleMutex struct {
	Nodes      map[string]struct{}
	Directions map[string]map[string]struct{}
	sync.Mutex
}

// NewGraphSingleMutex Gives an instance of SiteMap
func NewGraphSingleMutex() *GraphSingleMutex {
	return &GraphSingleMutex{
		Nodes:      make(map[string]struct{}),
		Directions: make(map[string]map[string]struct{}),
	}
}

// AddNode add a new url in the sitemap
// If url already exist, then it returns false
func (s *GraphSingleMutex) AddNode(url string) bool {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.Nodes[url]; ok {
		return false
	}
	s.Nodes[url] = struct{}{}
	return true
}

// AddDirection add a new connection from u to v in the sitemap
// If connection already exist, then it returns true,false
// It also returns false,false if any node doesn't already exists
func (s *GraphSingleMutex) AddDirection(u, v string) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Directions[u]; !ok {
		s.Directions[u] = make(map[string]struct{})
	}
	s.Directions[u][v] = struct{}{}
}

// Nodes map with mutex
type Nodes struct {
	N map[string]struct{}
	sync.Mutex
}

// Directions map with Mutex
type Directions struct {
	D map[string]map[string]struct{}
	sync.Mutex
}

// SiteMapSingleMutex DataStructure containing urls and connections
type GraphDoubleMutex struct {
	Nodes      Nodes
	Directions Directions
}

// New Gives an instance of SiteMap
func NewSiteMapDoubleMutex() *GraphDoubleMutex {
	return &GraphDoubleMutex{
		Nodes: Nodes{N: make(map[string]struct{})},
		Directions: Directions{
			D: make(map[string]map[string]struct{})},
	}
}

// AddNode add a new url in the sitemap
// If url already exist, then it returns false
func (s *GraphDoubleMutex) AddNode(url string) bool {
	s.Nodes.Lock()
	defer s.Nodes.Unlock()
	if _, ok := s.Nodes.N[url]; ok {
		return false
	}
	s.Nodes.N[url] = struct{}{}
	return true
}

// AddDirection add a new connection from u to v in the sitemap
// If connection already exist, then it returns true,false
// It also returns false,false if any node doesn't already exists
func (s *GraphDoubleMutex) AddDirection(u, v string) {
	s.Directions.Lock()
	defer s.Directions.Unlock()

	if _, ok := s.Directions.D[u]; !ok {
		s.Directions.D[u] = make(map[string]struct{})
	}
	s.Directions.D[u][v] = struct{}{}
}

// GraphSTD DataStructure containing urls and connections
// GraphSTD is just a graph
type GraphSTD struct {
	Nodes     *sync.Map
	Direction *sync.Map
}

// NewGraphSTD Gives an instance of SiteMapSTD
func NewGraphSTD() *GraphSTD {
	return &GraphSTD{
		Nodes:     &sync.Map{},
		Direction: &sync.Map{},
	}
}

// AddNode add a new url in the sitemap
// If url already exist, then it returns false
func (s *GraphSTD) AddNode(url string) bool {
	_, ok := s.Nodes.LoadOrStore(url, struct{}{})
	return !ok
}

// AddDirection add a new connection from u to v in the sitemap
// If connection already exist, then it returns true,false
// It also returns false,false if any node doesn't already exists
func (s *GraphSTD) AddDirection(u, v string) {
	val, ok := s.Direction.LoadOrStore(u, &sync.Map{})

	if ok {
		val.(*sync.Map).LoadOrStore(v, struct{}{})
	}
}

// GraphHybrid DataStructure containing urls and connections
// GraphHybrid is just a graph
type GraphHybrid struct {
	Node       *sync.Map
	Directions Directions
}

// NewGraphHybrid Gives an instance of SiteMapSTD
func NewGraphHybrid() *GraphHybrid {
	return &GraphHybrid{
		Node: &sync.Map{},
		Directions: Directions{
			D: make(map[string]map[string]struct{})},
	}
}

// AddNode add a new url in the sitemap
// If url already exist, then it returns false
func (s *GraphHybrid) AddNode(url string) bool {
	_, ok := s.Node.LoadOrStore(url, struct{}{})
	return !ok
}

// AddDirection add a new connection from u to v in the sitemap
// If connection already exist, then it returns true,false
// It also returns false,false if any node doesn't already exists
func (s *GraphHybrid) AddDirection(u, v string) {
	s.Directions.Lock()
	defer s.Directions.Unlock()

	if _, ok := s.Directions.D[u]; !ok {
		s.Directions.D[u] = make(map[string]struct{})
	}
	s.Directions.D[u][v] = struct{}{}
}
