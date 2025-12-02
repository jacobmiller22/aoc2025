package pconfig

import (
	"flag"
	"log"
)

type ProblemConfig struct {
	InputPath string
	Part      int
}

func Parse() ProblemConfig {
	var pcfg ProblemConfig

	flag.StringVar(&pcfg.InputPath, "inputpath", "", "Problem part")
	flag.IntVar(&pcfg.Part, "part", 1, "Problem part")

	flag.Parse()

	if pcfg.InputPath == "" {
		log.Fatal("pconfig: Invalid inputpath passed")
	}

	if pcfg.Part == 0 {
		log.Fatal("pconfig: Invalid part passed")
	}

	return pcfg
}
