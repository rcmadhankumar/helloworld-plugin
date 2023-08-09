// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"github.com/spf13/cobra"
	"fmt"

	"github.com/vmware-tanzu/tanzu-framework/packageclients/pkg/packagedatamodel"
)

var repoOp = packagedatamodel.NewRepositoryOptions()

var testCmd = &cobra.Command{
	Use:               "test command",
	Short:             "Repository operations",
	Args:  cobra.NoArgs,
	Long:              `Add, list, get or delete a package repository for Tanzu packages. A package repository is a collection of packages that are grouped together into an imgpkg bundle.`,
	RunE: testHelloworld,
	SilenceUsage: true,
}

func init() {
	fmt.Println("Hello world app initialization !")
}

func testHelloworld(cmd *cobra.Command, _ []string) error{
	fmt.Println("Hello world!")
	return nil
}