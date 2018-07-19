package author

import (
	"github.com/spf13/cobra"
	"fmt"
	"os/exec"
)

type Option struct {
	Author string
	Email string
	NewAuthor string
	NewEmail string
}

var (
	option Option
	sh string = `git filter-branch -f --commit-filter '
      if [ "$GIT_AUTHOR_EMAIL" = "%s"];
      then
		GIT_AUTHOR_NAME="%s";
		GIT_AUTHOR_EMAIL="%s";
		GIT_COMMITTER_NAME="%s";
		GIT_COMMITTER_EMAIL="%s";
        git commit-tree "$@";
      else
		git commit-tree "$@";
      fi'  --tag-name-filter cat -- --branches --tags`
	AuthorCmd = &cobra.Command{
		Use:     "author",
		Aliases: []string{"v"},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf(sh,option.Email,option.NewAuthor,option.NewEmail,option.NewAuthor,option.NewEmail))
			c := exec.Command(fmt.Sprintf(sh,option.Email,option.NewAuthor,option.NewEmail,option.NewAuthor,option.NewEmail))
			c.Output()
		},
	}
)

func init() {
	flags := AuthorCmd.Flags()

	flags.StringVarP(&option.Author, "author", "a", "", "原 author")
	flags.StringVarP(&option.Email, "email", "e", "", "原 email")
	flags.StringVarP(&option.NewAuthor, "new-author", "", "", "新 author")
	flags.StringVarP(&option.NewEmail, "new-email", "", "", "新 email")

}
