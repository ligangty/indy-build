package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ligangty/indy-build/process"
)

var indyURLNpm, gitURLNpm, tagNpm, branchNpm, buildNameNpm string

var npmCmd = &cobra.Command{
	Use:   "npm",
	Short: "do npm build against indy",
	Long:  "npm build against indy, includes build, folo, promote",
	Run: func(cmd *cobra.Command, args []string) {
		npmArgs := &baseArgs{
			indyURL:   indyURLNpm,
			gitURL:    gitURLNpm,
			tag:       tagNpm,
			branch:    branchNpm,
			buildName: buildNameNpm,
		}
		readyToRun := true
		checkout, checkoutType, validC := getCheckout(npmArgs)
		validV := validateBaseArgs(npmArgs)
		validPrepare := process.CheckPrerequisites(CMD_NPM)
		readyToRun = validC && validV && validPrepare
		indyURL, gitURL, buildName := npmArgs.indyURL, npmArgs.gitURL, npmArgs.buildName
		if readyToRun {
			process.RunBuild(indyURL, gitURL, checkoutType, checkout, process.TYPE_NPM, buildName)
		}
	},
}

func init() {
	npmCmd.Flags().StringVarP(&indyURLNpm, "indy_url", "i", "", "indy url.")
	npmCmd.Flags().StringVarP(&gitURLNpm, "gitURL", "g", "", "project git.")
	npmCmd.Flags().StringVarP(&tagNpm, "tag", "t", "", "project git tag to build")
	npmCmd.Flags().StringVarP(&branchNpm, "branch", "b", "", "project git branch to build.")
	npmCmd.Flags().StringVarP(&buildNameNpm, "buildName", "n", "", "build name.")

	npmCmd.MarkFlagRequired("indy_url")
	npmCmd.MarkFlagRequired("buildName")
	npmCmd.MarkFlagRequired("gitURL")
}
