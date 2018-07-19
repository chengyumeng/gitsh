package author

import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
)

type Option struct {
	Email string
	NewName string
	NewEmail string
}

var (
	option Option
	sh string = `git filter-branch -f --commit-filter '
      if [ "$GIT_AUTHOR_EMAIL" = "%s" ];
      then
		GIT_AUTHOR_NAME="%s";
		GIT_AUTHOR_EMAIL="%s";
      fi
      if [ "$GIT_COMMITTER_EMAIL" = "%s" ];
      then
        GIT_COMMITTER_NAME="%s";
        GIT_COMMITTER_EMAIL="%s";
      fi
        git commit-tree "$@";
      '  --tag-name-filter cat -- --branches --tags`
	AuthorCmd = &cobra.Command{
		Use:     "author",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			ex := fmt.Sprintf(sh,option.Email,option.NewName,option.NewEmail,option.Email,option.NewName,option.NewEmail)
			c := exec.Command(ex)
			fmt.Println(ex)
			c.Output()
		},
	}
)

func init() {
	flags := AuthorCmd.Flags()

	flags.StringVarP(&option.Email, "email", "e", "", "原 email")
	flags.StringVarP(&option.NewName, "new-name", "", "", "新 name")
	flags.StringVarP(&option.NewEmail, "new-email", "", "", "新 email")

}
