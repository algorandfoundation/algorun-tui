package node

import (
	"github.com/algorandfoundation/algorun-tui/internal/algod"
	"github.com/algorandfoundation/algorun-tui/ui/style"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"os"
)

const UpgradeMsg = "Upgrading Algod"

var upgradeCmd = &cobra.Command{
	Use:              "upgrade",
	Short:            "Upgrade Algod",
	Long:             "Upgrade Algod (if installed with package manager).",
	SilenceUsage:     true,
	PersistentPreRun: NeedsToBeStopped,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: get expected version and check if update is required
		log.Info(style.Green.Render(UpgradeMsg + " vX.X.X"))
		// Warn user for prompt
		log.Warn(style.Yellow.Render(SudoWarningMsg))
		// TODO: Check Version from S3 against the local binary
		err := algod.Update()
		if err != nil {
			log.Error(err)
		}
		// If it's not running, start the daemon (can happen)
		if !algod.IsRunning() {
			err = algod.Start()
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
		}
	},
}
