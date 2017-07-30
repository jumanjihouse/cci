// This code is based on https://github.com/microscaling/microbadger
// For more information please see microbadger.com/api
// (c) 2016 Force12io Ltd
// (c) 2017 pmorgan
//
// Sample usage:
//
// $ microbadger -name jumanjiman/aws
// Image has these labels:
//
// io.github.jumanjiman.build-date: 20170728T1022
// io.github.jumanjiman.ci-build-url: https://circleci.com/gh/jumanjihouse/docker-aws/522
// io.github.jumanjiman.docker.dockerfile: /Dockerfile
// io.github.jumanjiman.license: Apache License 2.0
// io.github.jumanjiman.vcs-ref: 4f9da70
// io.github.jumanjiman.vcs-type: Git
// io.github.jumanjiman.vcs-url: https://github.com/jumanjihouse/docker-aws.git
// io.github.jumanjiman.version: 1.11.127
//
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/microscaling/microbadger/api"
)

func main() {
	var cli = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	namePtr := cli.String("name", "", "name of public docker image")

	err := cli.Parse(os.Args[1:])

	if err == flag.ErrHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *namePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	labels, err := api.GetLabels(*namePtr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Image %s has these labels:\n", *namePtr)
	fmt.Println()
	for key, val := range labels {
		fmt.Printf("%s: %s\n", key, val)
	}
	fmt.Println()
}
