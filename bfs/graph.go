package main

/* This sample undirected graph comes from
  "Introduction to Algorithms - Third Edition" 22.2 BFS

  V = 8 (node count)
  E = 9 (edge count)
  define undirected graph G(V,E) as below (`s` is the source node):

r - s   t - u
|   | / |   |
v   w - x - y
*/

type adjacencyListGraph map[string][]string

var adjListGraph = adjacencyListGraph{
	"r": {"s", "v"},
	"s": {"r", "w"},
	"t": {"u", "w", "x"},
	"u": {"t", "y"},
	"v": {"r"},
	"w": {"s", "t", "x"},
	"x": {"t", "w", "y"},
	"y": {"u", "x"},
}
