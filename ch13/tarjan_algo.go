///main package has examples shown
// in Go Data Structures and algorithms book
package main

// importing bytes, encoding/json,fmt, os, io, log and sync  packages
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

//Tarjan Algo
func TarjanAlgo(graph DirGraph) [][]Identifier {
	var data *GraphData
	data = CreateGraphData()

	var vertex Identifier
	for vertex = range graph.GetGraphNodes() {
		var ok bool
		if _, ok = data.indices[vertex]; !ok {
			tarjanAlgo(graph, vertex, data)
		}
	}
	return data.output
}

//GraphData Class
type GraphData struct {
	mutex    sync.Mutex
	globalId int
	indices  map[Identifier]int
	lowLinks map[Identifier]int
	idArray  []Identifier
	idMap    map[Identifier]struct{}
	output   [][]Identifier
}

// CreateGraphData method
func CreateGraphData() *GraphData {
	return &GraphData{
		globalId: 0,
		indices:  make(map[Identifier]int),
		lowLinks: make(map[Identifier]int),
		idArray:  []Identifier{},
		idMap:    make(map[Identifier]struct{}),
		output:   [][]Identifier{},
	}
}

// tarjanAlgo method
func tarjanAlgo(
	graph DirGraph,
	id Identifier,
	data *GraphData,
) {

	data.mutex.Lock()
	data.indices[id] = data.globalId
	data.lowLinks[id] = data.globalId
	data.globalId++

	data.idArray = append(data.idArray, id)
	data.idMap[id] = struct{}{}

	data.mutex.Unlock()

	var cmap map[Identifier]GraphNode
	var err error
	cmap, err = graph.GetTargetNodes(id)
	if err != nil {
		panic(err)
	}
	var w Identifier
	for w = range cmap {

		var ok bool
		if _, ok = data.indices[w]; !ok {

			tarjanAlgo(graph, w, data)

			data.lowLinks[id] = min(data.lowLinks[id], data.lowLinks[w])

		} else if _, ok = data.idMap[w]; ok {

			data.lowLinks[id] = min(data.lowLinks[id], data.indices[w])
		}
	}

	data.mutex.Lock()
	defer data.mutex.Unlock()

	if data.lowLinks[id] == data.indices[id] {
		var component []Identifier
		component = []Identifier{}

		for {

			var u Identifier
			u = data.idArray[len(data.idArray)-1]
			data.idArray = data.idArray[: len(data.idArray)-1 : len(data.idArray)-1]
			delete(data.idMap, u)

			component = append(component, u)

			if u == id {
				data.output = append(data.output, component)
				break
			}
		}
	}
}

// min method
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//Identifier interface
type Identifier interface {
	IDString() string
}

// IdentifierStr
type IdentifierStr string

func (str IdentifierStr) IDString() string {
	return string(str)
}

//Graph Node Interface
type GraphNode interface {
	GetIdentifier() Identifier
	GetString() string
}

// Node class
type Node struct {
	identifier string
}

// variable uint64
var nodeCnt uint64

// CreateNode method
func CreateNode(idstr string) *Node {
	return &Node{
		identifier: idstr,
	}
}

// GetIdentifier method
func (node *Node) GetIdentifier() Identifier {
	return IdentifierStr(node.identifier)
}

// GetString method
func (node *Node) GetString() string {
	return node.identifier
}

// GraphEdge interface
type GraphEdge interface {
	GetSourceNode() GraphNode
	GetTargetNode() GraphNode
	GetWeight() float64
	GetStringFormat() string
}

// Edge class
type Edge struct {
	source GraphNode
	target GraphNode
	weight float64
}

// CreateGraphEdge method
func CreateGraphEdge(sourceNode GraphNode, targetNode GraphNode, weight float64) GraphEdge {
	return &Edge{
		source: sourceNode,
		target: targetNode,
		weight: weight,
	}
}

// GetSourceNode method
func (edge *Edge) GetSourceNode() GraphNode {
	return edge.source
}

// GetTargetNode method
func (edge *Edge) GetTargetNode() GraphNode {
	return edge.target
}

// GetWeight method
func (edge *Edge) GetWeight() float64 {
	return edge.weight
}

// GetStringFormat method
func (edge *Edge) GetStringFormat() string {
	return fmt.Sprintf("%s -- %.3f -→ %s\n", edge.source, edge.weight, edge.target)
}

// EdgesArray
type EdgesArray []Edge

func (edgesArray EdgesArray) Len() int { return len(edgesArray) }
func (edgesArray EdgesArray) Less(i int, j int) bool {
	return edgesArray[i].GetWeight() < edgesArray[j].GetWeight()
}
func (edgesArray EdgesArray) Swap(i int, j int) {
	edgesArray[i], edgesArray[j] = edgesArray[j], edgesArray[i]
}

// DirGraph Interface
type DirGraph interface {
	Initialise()
	GetGraphNodeCount() int
	GetGraphNode(id Identifier) GraphNode
	GetGraphNodes() map[Identifier]GraphNode
	AddGraphNode(node GraphNode) bool
	DeleteGraphNode(id Identifier) bool
	AddGraphEdge(id1 Identifier, id2 Identifier, weight float64) error
	ReplaceGraphEdge(id1 Identifier, id2 Identifier, weight float64) error
	DeleteGraphEdge(id1 Identifier, id2 Identifier) error
	GetWeight(id1 Identifier, id2 Identifier) (float64, error)
	GetSourceNodes(id Identifier) (map[Identifier]GraphNode, error)
	GetTargetNodes(id Identifier) (map[Identifier]GraphNode, error)
	GetStringNotation() string
}

// Graph class
type Graph struct {
	mutex               sync.RWMutex
	nodeMap             map[Identifier]GraphNode
	nodeToSourceNodeMap map[Identifier]map[Identifier]float64
	nodeToTargetNodeMap map[Identifier]map[Identifier]float64
}

// CreateGraph method
func CreateDirGraph() *Graph {
	return &Graph{
		nodeMap:             make(map[Identifier]GraphNode),
		nodeToSourceNodeMap: make(map[Identifier]map[Identifier]float64),
		nodeToTargetNodeMap: make(map[Identifier]map[Identifier]float64),
	}
}

// Create Graph method
func CreateGraph() DirGraph {
	return CreateDirGraph()
}

// Initialise method
func (graph *Graph) Initialise() {
	graph.nodeMap = make(map[Identifier]GraphNode)
	graph.nodeToSourceNodeMap = make(map[Identifier]map[Identifier]float64)
	graph.nodeToTargetNodeMap = make(map[Identifier]map[Identifier]float64)
}

// GetGraphNodeCount method
func (graph *Graph) GetGraphNodeCount() int {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()

	return len(graph.nodeMap)
}

// GetGraphNode method
func (graph *Graph) GetGraphNode(id Identifier) GraphNode {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()

	return graph.nodeMap[id]
}

// GetGraphNodes method
func (graph *Graph) GetGraphNodes() map[Identifier]GraphNode {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()

	return graph.nodeMap
}

// CheckNotSafeExistID method
func (graph *Graph) CheckNotSafeExistID(id Identifier) bool {
	var ok bool
	_, ok = graph.nodeMap[id]
	return ok
}

//addGraphNode method
func (graph *Graph) AddGraphNode(node GraphNode) bool {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()

	if graph.CheckNotSafeExistID(node.GetIdentifier()) {
		return false
	}
	var id Identifier
	id = node.GetIdentifier()
	graph.nodeMap[id] = node
	return true
}

//DeleteGraphNode method
func (graph *Graph) DeleteGraphNode(id Identifier) bool {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()

	if !graph.CheckNotSafeExistID(id) {
		return false
	}

	delete(graph.nodeMap, id)

	delete(graph.nodeToTargetNodeMap, id)
	var smap map[Identifier]float64
	for _, smap = range graph.nodeToTargetNodeMap {
		delete(smap, id)
	}

	delete(graph.nodeToSourceNodeMap, id)
	for _, smap = range graph.nodeToSourceNodeMap {
		delete(smap, id)
	}

	return true
}

// AddGraphEdge method
func (graph *Graph) AddGraphEdge(id1 Identifier, id2 Identifier, weight float64) error {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()

	if !graph.CheckNotSafeExistID(id1) {
		return fmt.Errorf("%s is not in the graph.", id1)
	}
	if !graph.CheckNotSafeExistID(id2) {
		return fmt.Errorf("%s is not in the graph.", id2)
	}
	var ok bool
	if _, ok = graph.nodeToSourceNodeMap[id1]; ok {
		var v float64
		var ok2 bool
		if v, ok2 = graph.nodeToTargetNodeMap[id1][id2]; ok2 {
			graph.nodeToTargetNodeMap[id1][id2] = v + weight
		} else {
			graph.nodeToTargetNodeMap[id1][id2] = weight
		}
	} else {
		var tmap map[Identifier]float64
		tmap = make(map[Identifier]float64)
		tmap[id2] = weight
		graph.nodeToTargetNodeMap[id1] = tmap
	}
	if _, ok = graph.nodeToSourceNodeMap[id2]; ok {
		var ok2 bool
		var v float64
		if v, ok2 = graph.nodeToSourceNodeMap[id2][id1]; ok2 {
			graph.nodeToSourceNodeMap[id2][id1] = v + weight
		} else {
			graph.nodeToSourceNodeMap[id2][id1] = weight
		}
	} else {
		var tmap map[Identifier]float64
		tmap = make(map[Identifier]float64)
		tmap[id1] = weight
		graph.nodeToSourc
eNodeMap[id2] = tmap
	}

	return nil
}

// ReplaceEdge method
func (graph *Graph) ReplaceGraphEdge(id1 Identifier, id2 Identifier, weight float64) error {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()

	if !graph.CheckNotSafeExistID(id1) {
		return fmt.Errorf("%s is not in the graph.", id1)
	}
	if !graph.CheckNotSafeExistID(id2) {
		return fmt.Errorf("%s is not  in the graph.", id2)
	}
	var ok bool
	if _, ok = graph.nodeToTargetNodeMap[id1]; ok {
		graph.nodeToTargetNodeMap[id1][id2] = weight
	} else {
		var tmap map[Identifier]float64
		tmap = make(map[Identifier]float64)
		tmap[id2] = weight
		graph.nodeToTargetNodeMap[id1] = tmap
	}
	if _, ok = graph.nodeToSourceNodeMap[id2]; ok {
		graph.nodeToSourceNodeMap[id2][id1] = weight
	} else {
		var tmap map[Identifier]float64
		tmap = make(map[Identifier]float64)
		tmap[id1] = weight
		graph.nodeToSourceNodeMap[id2] = tmap
	}
	return nil
}

// DeleteGraphEdge method
func (graph *Graph) DeleteGraphEdge(id1 Identifier, id2 Identifier) error {
	graph.mutex.Lock()
	defer graph.mutex.Unlock()

	if !graph.CheckNotSafeExistID(id1) {
		return fmt.Errorf("%s is not in the graph.", id1)
	}
	if !graph.CheckNotSafeExistID(id2) {
		return fmt.Errorf("%s is not in the graph.", id2)
	}
	var ok bool
	if _, ok = graph.nodeToTargetNodeMap[id1]; ok {
		if _, ok = graph.nodeToTargetNodeMap[id1][id2]; ok {
			delete(graph.nodeToTargetNodeMap[id1], id2)
		}
	}
	if _, ok = graph.nodeToSourceNodeMap[id2]; ok {
		if _, ok = graph.nodeToSourceNodeMap[id2][id1]; ok {
			delete(graph.nodeToSourceNodeMap[id2], id1)
		}
	}
	return nil
}

// GetWeight method
func (graph *Graph) GetWeight(id1 Identifier, id2 Identifier) (float64, error) {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()

	if !graph.CheckNotSafeExistID(id1) {
		return 0, fmt.Errorf("%s is not in the graph.", id1)
	}
	if !graph.CheckNotSafeExistID(id2) {
		return 0, fmt.Errorf("%s is not in the graph.", id2)
	}
	var ok bool
	if _, ok = graph.nodeToTargetNodeMap[id1]; ok {
		var v float64
		if v, ok = graph.nodeToTargetNodeMap[id1][id2]; ok {
			return v, nil
		}
	}
	return 0.0, fmt.Errorf(" no edge from %s to %s", id1, id2)
}

//GetSourceNodes method
func (graph *Graph) GetSourceNodes(id Identifier) (map[Identifier]GraphNode, error) {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()

	if !graph.CheckNotSafeExistID(id) {
		return nil, fmt.Errorf("%s is not in the graph.", id)
	}
	var rs map[Identifier]GraphNode
	rs = make(map[Identifier]GraphNode)
	var ok bool
	if _, ok = graph.nodeToSourceNodeMap[id]; ok {
		var n Identifier
		for n = range graph.nodeToSourceNodeMap[id] {
			rs[n] = graph.nodeMap[n]
		}
	}
	return rs, nil
}

//GetTargetNodes method
func (graph *Graph) GetTargetNodes(id Identifier) (map[Identifier]GraphNode, error) {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()

	if !graph.CheckNotSafeExistID(id) {
		return nil, fmt.Errorf("%s is not in the graph.", id)
	}
	var rs map[Identifier]GraphNode
	rs = make(map[Identifier]GraphNode)
	var ok bool
	if _, ok = graph.nodeToTargetNodeMap[id]; ok {
		var n Identifier
		for n = range graph.nodeToTargetNodeMap[id] {
			rs[n] = graph.nodeMap[n]
		}
	}
	return rs, nil
}

// GetStringNotation method
func (graph *Graph) GetStringNotation() string {
	graph.mutex.RLock()
	defer graph.mutex.RUnlock()
	var buf *bytes.Buffer
	buf = new(bytes.Buffer)
	var id1 Identifier
	var nd1 GraphNode
	for id1, nd1 = range graph.nodeMap {
		var nmap map[Identifier]GraphNode
		nmap, _ = graph.GetTargetNodes(id1)
		var id2 Identifier
		var nd2 GraphNode
		for id2, nd2 = range nmap {
			var weight float64
			weight, _ = graph.GetWeight(id1, id2)
			fmt.Fprintf(buf, "%s -- %.3f -→ %s\n", nd1, weight, nd2)
		}
	}
	return buf.String()
}

// CreateGraphFromJSON method
func CreateGraphFromJSON(reader io.Reader, graphID string) (DirGraph, error) {
	var js map[string]map[string]map[string]float64
	js = make(map[string]map[string]map[string]float64)

	var dec *json.Decoder
	dec = json.NewDecoder(reader)
	for {
		var err error
		if err = dec.Decode(&js); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	var ok bool
	if _, ok = js[graphID]; !ok {
		return nil, fmt.Errorf("%s is not there", graphID)
	}
	var gmap map[string]map[string]float64
	gmap = js[graphID]

	var graph DirGraph
	graph = CreateDirGraph()
	var id1 string
	var mm map[string]float64
	for id1, mm = range gmap {
		var nd1 GraphNode
		nd1 = graph.GetGraphNode(IdentifierStr(id1))
		if nd1 == nil {
			nd1 = CreateNode(id1)
			var ok bool
			if ok = graph.AddGraphNode(nd1); !ok {
				return nil, fmt.Errorf("%s is present", nd1)
			}
		}
		var id2 string
		var weight float64
		for id2, weight = range mm {
			var nd2 GraphNode
			nd2 = graph.GetGraphNode(IdentifierStr(id2))
			if nd2 == nil {
				nd2 = CreateNode(id2)
				var ok bool
				if ok = graph.AddGraphNode(nd2); !ok {
					return nil, fmt.Errorf("%s is present", nd2)
				}
			}
			graph.ReplaceGraphEdge(nd1.GetIdentifier(), nd2.GetIdentifier(), weight)
		}
	}

	return graph, nil
}

// main method
func main() {
	f, err := os.Open("input.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var graph DirGraph
	graph, err = CreateGraphFromJSON(f, "graph_21")
	if err != nil {
		panic(err)
	}
	var strongComponents [][]Identifier
	strongComponents = TarjanAlgo(graph)
	if len(strongComponents) != 4 {
		log.Fatalf("Expected 4 Actual %v", strongComponents)
	}
	fmt.Println("Tarjan graph_21:", strongComponents)
}
