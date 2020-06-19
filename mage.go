//+build mage

package main

import (
	"github.com/aevea/magefiles"
	"github.com/magefile/mage/sh"
)

func Test() error {
	err := sh.RunV("sh", "./testdata/setup-test-repos.sh")

	if err != nil {
		return err
	}

	return magefiles.Test()
}

func GoModTidy() error {
	return magefiles.GoModTidy()
}
