package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var downloadURL = ""

type Node struct {
	XMLName xml.Name `xml:"node"`
	Id      string   `xml:"id"`
	Name    string   `xml:"name"`
}

type Edge struct {
	XMLName    xml.Name `xml:"edge"`
	Id         string   `xml:"id"`
	FromString string   `xml:"from"`
	FromNode   *Node
	ToString   string `xml:"to"`
	ToNode     *Node
	Cost       float32 `xml:"cost"`
}

type Graph struct {
	XMLName xml.Name `xml:"graph"`
	Id      string   `xml:"id"`
	Name    string   `xml:"name"`
	Nodes   []*Node  `xml:"nodes"`
	Edges   []*Edge  `xml:"edges"`
}

//parse xml

func main() {
	xmlFile, err := os.Create("output.txt")
	defer xmlFile.Close()

	resp, err := http.Get(downloadURL)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("bad status: %s", resp.Status))
	}

	_, err = io.Copy(xmlFile, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	parsedGraph, err := parseXML(xmlFile)
	if err != nil {
		panic("unable to parse xml")
	}
	fmt.Printf("%v", parsedGraph)

	_, err = cleanXML(parsedGraph)
	if err != nil {
		panic("unable to clean graph data")
	}
}

func parseXML(xmlFile *os.File) (*Graph, error) {
	fmt.Printf("parsing file \n")
	bytes, _ := io.ReadAll(xmlFile)
	//fmt.Printf("%v", bytes)

	var graph Graph
	xml.Unmarshal(bytes, &graph)
	fmt.Printf("Graph parsed; name %v \n", graph.Name)

	return &graph, nil
}

func cleanXML(input *Graph) (*Graph, error) {
	if input.Id == "" || input.Name == "" {
		return nil, errors.New("graph is invalid; either Id or Name is empty")
	}
	graph := &Graph{
		Id:    input.Id,
		Name:  input.Name,
		Edges: []*Edge{},
	}
	nodeNames := map[string]*Node{}
	newNodes := []*Node{}

	//add node to graph if it's unique
	for _, node := range input.Nodes {
		if _, ok := nodeNames[node.Name]; !ok {
			newNodes = append(newNodes, node)
			nodeNames[node.Name] = node
		}
	}
	if len(newNodes) == 0 {
		return nil, errors.New("graph is invalid; nodes are empty")
	}
	graph.Nodes = newNodes

	//add edges to graph
	newEdges := []*Edge{}
	for _, edge := range input.Edges {
		newEdge := &Edge{
			Id:         edge.Id,
			FromString: edge.FromString,
			ToString:   edge.ToString,
		}
		if edge.Cost < 0 {
			newEdge.Cost = 0
		} else {
			newEdge.Cost = edge.Cost
		}
		if val, ok := nodeNames[edge.FromString]; ok {
			newEdge.FromNode = val
		} else {
			continue
		}
		if val, ok := nodeNames[edge.ToString]; ok {
			newEdge.ToNode = val
		} else {
			continue
		}
		newEdges = append(newEdges, newEdge)

	}

	return graph, nil
}
