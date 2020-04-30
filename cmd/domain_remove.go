package cmd

import (
	"fmt"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
)

var domainRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "delete", "destroy"},
	Short:   "Create a new domain",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			fmt.Printf("Unable to create a Civo API Client: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		domain, err := client.FindDNSDomain(args[0])
		if err != nil {
			fmt.Printf("Unable to find domain for your search: %s\n", aurora.Red(err))
			os.Exit(1)
		}

		_, err = client.DeleteDNSDomain(domain)

		ow := utility.NewOutputWriterWithMap(map[string]string{"ID": domain.ID, "Name": domain.Name})

		switch outputFormat {
		case "json":
			ow.WriteSingleObjectJSON()
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			fmt.Printf("The domain called %s with ID %s was delete\n", aurora.Green(domain.Name), aurora.Green(domain.ID))
		}
	},
}
