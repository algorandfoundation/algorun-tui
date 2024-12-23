package catchup

import (
	"context"
	"github.com/algorandfoundation/algorun-tui/api"
	"github.com/algorandfoundation/algorun-tui/cmd/utils"
	"github.com/algorandfoundation/algorun-tui/internal/algod"
	"github.com/algorandfoundation/algorun-tui/ui/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var stopCmdShort = "Stop a fast catchup"
var stopCmdLong = lipgloss.JoinVertical(
	lipgloss.Left,
	style.Purple(style.BANNER),
	"",
	style.Bold(stopCmdShort),
	"",
	style.BoldUnderline("Overview:"),
	"Stop an active Fast-Catchup. This will abort the catchup process if one has started",
	"",
	style.Yellow.Render("Note: Not all networks support Fast-Catchup."),
)

// stopCmd is a Cobra command used to check the node's sync status and initiate a fast catchup when necessary.
var stopCmd = utils.WithAlgodFlags(&cobra.Command{
	Use:          "stop",
	Short:        stopCmdShort,
	Long:         stopCmdLong,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		err := utils.InitConfig()
		if err != nil {
			log.Fatal(err)
		}
		endpoint := viper.GetString("algod-endpoint")
		token := viper.GetString("algod-token")
		if endpoint == "" {
			log.Fatal("algod-endpoint is required")
		}
		if token == "" {
			log.Fatal("algod-token is required")
		}

		ctx := context.Background()
		httpPkg := new(api.HttpPkg)
		client, err := algod.GetClient(endpoint, token)
		cobra.CheckErr(err)

		status, response, err := algod.NewStatus(ctx, client, httpPkg)
		utils.WithInvalidResponsesExplanations(err, response, cmd.UsageString())
		if status.State != algod.FastCatchupState || status.Catchpoint == nil || *status.Catchpoint == "" {
			log.Fatal(style.Red.Render("Node is not in fast catchup state."))
		}

		msg, _, err := algod.AbortCatchup(ctx, client, *status.Catchpoint)
		if err != nil {
			log.Fatal(err)
		}
		log.Info(style.Green.Render("Catchpoint Message: " + msg))

	},
}, &endpoint, &token)

func init() {
	stopCmd.Flags().BoolVarP(&force, "force", "f", false, style.Yellow.Render("forcefully catchup the node"))
}
