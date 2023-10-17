package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
)

// func lineCounter(r io.Reader) (int, error) {
// 	// https://stackoverflow.com/a/24563853
// 	buf := make([]byte, 32*1024)
// 	count := 0
// 	lineSep := []byte{'\n'}

// 	for {
// 		c, err := r.Read(buf)
// 		count += bytes.Count(buf[:c], lineSep)

// 		switch {
// 		case err == io.EOF:
// 			return count, nil

// 		case err != nil:
// 			return count, err
// 		}
// 	}
// }

func main() {
	fmt.Println("Howdy Partner!")
	// file, err := os.Open("/app/adjacency_list.txt")
	// defer file.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// num_lines, err := lineCounter(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// nodes := num_lines + 1
	// fmt.Println("number of nodes", nodes)

	file, err := os.Open("/app/adjacency_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g := simple.NewDirectedGraph()
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		for _, dest_node_str := range strings.Fields(scanner.Text()) {
			dest_node_int, err := strconv.Atoi(dest_node_str)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Adding edge from ", i, " to ", dest_node_int)
			g.SetEdge(simple.Edge{F: simple.Node(i), T: simple.Node(dest_node_int)})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(topo.DirectedCyclesIn(g))
}
