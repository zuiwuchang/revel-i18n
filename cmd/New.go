package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zuiwuchang/revel-i18n/cmd/cmdnew"
	"log"
)

func init() {
	context := &cmdnew.Context{}
	cmd := &cobra.Command{
		Use:   "new",
		Short: "new message file",
		Long: `new message file
	revel-i18n new -v app/views -m messages -l zh-TW
`,
		Run: func(cmd *cobra.Command, args []string) {
			p, e := cmdnew.NewProcessor(context)
			if e != nil {
				log.Fatalln(e)
			}
			e = p.Run()
			if e != nil {
				log.Fatalln(e)
			}
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&context.Src,
		"views", "v",
		"app/views",
		"revel views directory",
	)
	flags.StringVarP(&context.Dist,
		"messages", "m",
		"messages",
		"revel messages directory",
	)
	flags.StringVarP(&context.Locale,
		"locale", "l",
		"zh",
		"locale zh zh-TW zh-HK de ...",
	)
	flags.BoolVarP(&context.Touch,
		"touch", "t",
		false,
		"true (Coverage file) false (Merge file)",
	)
	flags.BoolVar(&context.NotLine,
		"no-line",
		false,
		"if true not write key file:line",
	)
	flags.StringVar(&context.Delimiters,
		"delimiters",
		"{{ }}",
		"go template delimiters",
	)
	rootCmd.AddCommand(cmd)
}
