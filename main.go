package main

import (
	"fmt"
	"github.com/spf13/viper"
	log "github.com/eris-ltd/eris-logger"
	"github.com/eris-ltd/playingWithViper/definitions"
)

func main() {
	log.SetLevel(log.WarnLevel)
	log.Warn("Loading eris-pm Package Definition File.")
	do = definitions.NowDo()
	var pkg = definitions.BlankJobs()
	var epmJobs = viper.New()

	epmJobs.AddConfigPath(".")
	epmJobs.SetConfigName("epm")
	epmJobs.SetConfigType("yaml")

	// load file
	if err := epmJobs.ReadInConfig(); err != nil {
		log.Warn(fmt.Errorf("Sorry, the marmots were unable to load the eris-pm jobs file. Please check your path.\nERROR =>\t\t\t%v", err))
	}

	// marshall file
	if err := epmJobs.Unmarshal(pkg); err != nil {
		log.Warn(fmt.Errorf("Sorry, the marmots could not figure that eris-pm jobs file out.\nPlease check your epm.yaml is properly formatted.\n"))
	}

	for _,job := range pkg.Jobs {
		beginJob(job, do)
	}
}

func beginJob(job *definitions.JobsCommon, do *definitions.Do) {
	job.PreProcess(do)
	a, err := job.Execute(do)
	log.Warn(a)
}