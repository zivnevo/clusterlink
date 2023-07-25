package subcommand

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	cmdutil "github.ibm.com/mbg-agent/cmd/util"
	"github.ibm.com/mbg-agent/pkg/admin"
	"github.ibm.com/mbg-agent/pkg/api"
)

// importCreateOptions is the command line options for 'create import'
type importCreateOptions struct {
	myID string
	name string
	host string
	port uint16
}

// ImportCreateCmd - create an imported service.
func ImportCreateCmd() *cobra.Command {
	o := importCreateOptions{}
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Create an imported service",
		Long:  `Create an imported service that can be bounded to other peers`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run()
		}}

	o.addFlags(cmd.Flags())
	cmdutil.MarkFlagsRequired(cmd, []string{"name", "port"})

	return cmd
}

// addFlags registers flags for the CLI.
func (o *importCreateOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.myID, "myid", "", "gwctl ID")
	fs.StringVar(&o.name, "name", "", "Imported service name")
	fs.StringVar(&o.host, "host", "", "Imported service endpoint (IP/DNS), if unspecified, uses the service name")
	fs.Uint16Var(&o.port, "port", 0, "Imported service port")
}

// run performs the execution of the 'create import' subcommand
func (o *importCreateOptions) run() error {
	g, err := admin.GetClientFromID(o.myID)
	if err != nil {
		return err
	}

	err = g.CreateImportService(api.Import{
		Name: o.name,
		Spec: api.ImportSpec{
			Service: api.Endpoint{
				Host: o.host,
				Port: o.port},
		},
	})
	if err != nil {
		return err
	}

	fmt.Printf("Imported service created successfully\n")
	return nil
}

// importDeleteOptions is the command line options for 'delete import'
type importDeleteOptions struct {
	myID string
	name string
}

// ImportDeleteCmd - delete an imported service command.
func ImportDeleteCmd() *cobra.Command {
	o := importDeleteOptions{}
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Delete an imported service",
		Long:  `Delete an imported service`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run()
		},
	}

	o.addFlags(cmd.Flags())
	cmdutil.MarkFlagsRequired(cmd, []string{"name"})

	return cmd
}

// addFlags registers flags for the CLI.
func (o *importDeleteOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.myID, "myid", "", "gwctl ID")
	fs.StringVar(&o.name, "name", "", "Imported service name")
}

// run performs the execution of the 'delete import' subcommand
func (o *importDeleteOptions) run() error {
	g, err := admin.GetClientFromID(o.myID)
	if err != nil {
		return err
	}

	err = g.DeleteImportService(api.Import{Name: o.name})
	if err != nil {
		return err
	}

	fmt.Printf("Imported service was deleted successfully\n")
	return nil
}

// importGetOptions is the command line options for 'get import'
type importGetOptions struct {
	myID string
	name string
}

// ImportGetCmd - get imported service command.
func ImportGetCmd() *cobra.Command {
	o := importGetOptions{}
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Get an imported service",
		Long:  `Get an imported service`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return o.run()
		},
	}
	o.addFlags(cmd.Flags())

	return cmd
}

// addFlags registers flags for the CLI.I.
func (o *importGetOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.myID, "myid", "", "gwctl ID")
	fs.StringVar(&o.name, "name", "", "Imported service name. If empty gets all imported services.")
}

// run performs the execution of the 'get import' subcommand
func (o *importGetOptions) run() error {
	g, err := admin.GetClientFromID(o.myID)
	if err != nil {
		return err
	}

	if o.name == "" {
		sArr, err := g.GetImportServices()
		if err != nil {
			return err
		}
		fmt.Printf("Imported services:\n")
		for i, s := range sArr {
			fmt.Printf("%d. Imported Name: %s. Endpoint %v\n", i+1, s.Name, s.Spec.Service)
		}
	} else {
		iSvc, err := g.GetImportService(api.Import{Name: o.name})
		if err != nil {
			return err
		}
		fmt.Printf("Imported Name: %s. Endpoint %v\n", iSvc.Name, iSvc.Spec.Service)
	}

	return nil
}