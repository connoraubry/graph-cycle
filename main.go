package main

import (
	"fmt"

	"github.com/connoraubry/graph-cycle/tools/graph"
	loader "github.com/connoraubry/graph-cycle/tools/loaders"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("vim-go")
	logrus.SetLevel(logrus.DebugLevel)

	connections := loader.LoadFromJSON("data/2023-cnx.json")
	g := graph.New()

	logrus.WithField("len", len(connections)).Debug("Adding Connections")
	g.AddConnections(connections)
	logrus.Debug("Connections added")

	logrus.Debug("Evaluating cycles")
	g.EvaluateCycles()
	logrus.Debug("Cycles evaluated")

	maxCycleLen := 0
	for _, cycle := range g.NodeToCycle {
		if len(cycle) > maxCycleLen {
			maxCycleLen = len(cycle)
		}
	}
	logrus.WithField("len", maxCycleLen).Info("Found Max Cycle length")

}
