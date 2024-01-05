package loader

import (
	"encoding/json"
	"io"
	"os"

	"github.com/connoraubry/graph-cycle/tools/graph"
	log "github.com/sirupsen/logrus"
)

func LoadFromJSON(filename string) []graph.Connection {
	f, err := os.Open(filename)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	var c []graph.Connection
	bytes, _ := io.ReadAll(f)
	json.Unmarshal(bytes, &c)

	return c
}

func SaveToJSON(filename string, cnxList []graph.Connection) {

	bytes, err := json.MarshalIndent(cnxList, "", "  ")
	if err != nil {
		log.Error("Error marshalling cnxList: ", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Error(err)
	}
	defer f.Close()

	f.Write(bytes)
}
