package main

import (
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
)

var (
	version   string
	branch    string
	revision  string
	buildDate string
	goVersion = runtime.Version()
)

func main() {

	// log to stdout and hide timestamp
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// log startup message
	log.Printf("Starting estafette-extension-github-status version %v...\n\n", version)

	log.Printf("All available estafette environment variables; the _DNS_SAFE suffixed ones can be used to set dns labels. Since leading digits are not allowed some of them are empty.\n\n")

	estafetteEnvvars := []string{}

	// get all envvars starting with ESTAFETTE_
	for _, e := range os.Environ() {
		kvPair := strings.SplitN(e, "=", 2)

		if len(kvPair) == 2 {
			envvarName := kvPair[0]

			if strings.HasPrefix(envvarName, "ESTAFETTE_") {
				estafetteEnvvars = append(estafetteEnvvars, e)
			}
		}
	}

	// sort envvars, since they're returned randomly
	sort.Strings(estafetteEnvvars)

	// log all Estafette envvars
	for _, e := range estafetteEnvvars {
		kvPair := strings.SplitN(e, "=", 2)

		if len(kvPair) == 2 {
			envvarName := kvPair[0]
			envvarValue := kvPair[1]

			log.Printf("%v: %v\n", envvarName, envvarValue)
		}
	}
}
