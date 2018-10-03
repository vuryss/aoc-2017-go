package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Prog struct {
	name string
	weight int
	children []string
}

type Node struct {
	name string
	weight int
	totalWeight int
	children []Node
}

func buildNode(programs map[string]Prog, node string) Node {
	program := programs[node]
	children := make([]Node, 0)
	totalWeight := program.weight

	if len(program.children) > 0 {
		for i := range program.children {
			node := buildNode(programs, program.children[i])
			children = append(children, node)
			totalWeight += node.totalWeight
		}
	}

	return Node{
		program.name,
		program.weight,
		totalWeight,
		children,
	}
}

func findNormalizedWeight(node Node, weight int) int {
	differentNode := Node{}
	normalWeight := 0
	found := false

	for i := range node.children {
		diff := 0
		for j := range node.children {
			if i == j {
				continue
			}

			if node.children[i].totalWeight != node.children[j].totalWeight {
				normalWeight = node.children[i].weight - (node.children[i].totalWeight - node.children[j].totalWeight)
				diff++
			}
		}

		if diff > 1 {
			differentNode = node.children[i]
			found = true
			break
		}
	}

	// We've reached the different node
	if !found {
		return weight
	}

	return findNormalizedWeight(differentNode, normalWeight)
}

func part1(input string) {
	start := time.Now()
	data := strings.Split(input, "\n")
	re := regexp.MustCompile(`([a-z]+).*\((\d+)\)(?:.*?([a-z][a-z\s,]+))?`)
	programs := make(map[string]Prog)
	allChildren := make(map[string]bool)

	for i := range data {
		matches := re.FindStringSubmatch(data[i])
		weight, _ := strconv.Atoi(matches[2])
		children := make([]string, 0)

		if len(matches) > 3 {
			children = strings.Split(matches[3], ", ")

			for i := range children {
				allChildren[children[i]] = true
			}
		}

		programs[matches[1]] = Prog{matches[1], weight, children}
	}

	// Find root program
	root := ""

	for i := range programs {
		if _, exists := allChildren[programs[i].name]; exists {
			continue
		}

		root = programs[i].name
		log.Printf("Root program: %v", programs[i].name)
		break
	}

	// Build tree
	rootNode := buildNode(programs, root)

	normalWeight := findNormalizedWeight(rootNode, rootNode.totalWeight)

	log.Printf("Normal weight of imbalanced node: %v", normalWeight)

	log.Printf("Execution time: %v", time.Since(start))
}

func main() {
	input := getInput("day-7")
	part1(input)
}
