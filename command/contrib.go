package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/naoty/contrib/util"
)

func Contrib(c *cli.Context) {
	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	names, err := getContributors(c.Args())
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	os.Exit(0)
}

func getContributors(filenames []string) ([]string, error) {
	var err error
	contributors := make(map[string]int)

	for _, filename := range filenames {
		names, err := gitLog(filename)
		if err != nil {
			continue
		}

		for _, name := range names {
			contributors[name] += 1
		}
	}

	if err != nil {
		return []string{}, err
	}

	vs := util.NewValSorter(contributors)
	vs.Sort()

	return vs.Keys, nil
}

func gitLog(filename string) ([]string, error) {
	// Print only auther's name
	cmd := exec.Command("git", "log", "--pretty=format:%an", filename)
	out, err := cmd.CombinedOutput()
	str := string(out)
	str = strings.Trim(str, "\n")
	return strings.Split(str, "\n"), err
}