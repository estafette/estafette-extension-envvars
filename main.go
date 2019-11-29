package main

import (
	"os"
	"runtime"
	"sort"
	"strings"

	foundation "github.com/estafette/estafette-foundation"
	"github.com/rs/zerolog/log"
)

var (
	appgroup  string
	app       string
	version   string
	branch    string
	revision  string
	buildDate string
	goVersion = runtime.Version()
)

func main() {

	// init log format from envvar ESTAFETTE_LOG_FORMAT
	foundation.InitLoggingFromEnv(appgroup, app, version, branch, revision, buildDate)

	log.Info().Msg("All available estafette environment variables; the _DNS_SAFE suffixed ones can be used to set dns labels. Since leading digits are not allowed some of them are empty.")

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

			log.Info().Msgf("%v: %v", envvarName, envvarValue)
		}
	}
}
