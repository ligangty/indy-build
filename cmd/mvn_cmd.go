package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ligangty/indy-build/process"
)

var indyURLMvn, gitURLMvn, tagMvn, branchMvn, buildNameMvn string

var mvnCmd = &cobra.Command{
	Use:   "maven",
	Short: "do maven build against indy",
	Long:  "maven build against indy, includes build, folo, promote",
	Run: func(cmd *cobra.Command, vArgs []string) {
		mvnArgs := &baseArgs{
			indyURL:   indyURLMvn,
			gitURL:    gitURLMvn,
			tag:       tagMvn,
			branch:    branchMvn,
			buildName: buildNameMvn,
		}
		readyToRun := true
		checkout, checkoutType, validC := getCheckout(mvnArgs)
		validV := validateBaseArgs(mvnArgs)
		validPrepare := process.CheckPrerequisites(CMD_MVN)
		readyToRun = validC && validV && validPrepare
		indyURL, gitURL, buildName := mvnArgs.indyURL, mvnArgs.gitURL, mvnArgs.buildName
		if readyToRun {
			process.RunBuild(indyURL, gitURL, checkoutType, checkout, process.TYPE_MVN, buildName)
		}
	},
}

func init() {
	mvnCmd.Flags().StringVarP(&indyURLMvn, "indy_url", "i", "", "indy url.")
	mvnCmd.Flags().StringVarP(&gitURLMvn, "gitURL", "g", "", "project git.")
	mvnCmd.Flags().StringVarP(&tagMvn, "tag", "t", "", "project git tag to build")
	mvnCmd.Flags().StringVarP(&branchMvn, "branch", "b", "", "project git branch to build.")
	mvnCmd.Flags().StringVarP(&buildNameMvn, "buildName", "n", "", "build name.")

	mvnCmd.MarkFlagRequired("indy_url")
	mvnCmd.MarkFlagRequired("buildName")
	mvnCmd.MarkFlagRequired("gitURL")
}
