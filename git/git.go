package git

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var regex *regexp.Regexp

func init() {
	regex, _ = regexp.Compile("^\\d+")
}

func issueNoByBranch(branch string) (r int, e error) {
	strs := strings.Split(branch, "/")
	if len(strs) < 2 {
		e = errors.New("not an issue branch")
		return
	}
	last := strs[len(strs)-1]
	numstr := regex.FindString(last)
	if numstr == "" {
		e = errors.New("no issue number found")
	}
	res, err := strconv.Atoi(numstr)
	if err != nil {
		panic("unpredicted error")
	}
	r = res
	return
}

// CurrentIssueNo returns an issue number of the current git branch.
func CurrentIssueNo() (r int, e error) {
	branch := CurrentBranchName()
	issue, err := issueNoByBranch(branch)
	if err != nil {
		e = errors.New("not in a issue branch")
		return
	}
	r = issue
	return
}

// CurrentBranchName returns the current git branch name.
func CurrentBranchName() string {
	name, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatalln("not in git repository")
	}
	return string(name)
}

// CommitWithTemplate shows commit screen with the specified template.
func CommitWithTemplate(message string) error {
	cmd := exec.Command("git", "commit", "-m", message, "-e")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// CreateNewBranch creates and checkouts new branch.
func CreateNewBranch(name string) error {
	cmd := exec.Command("git", "checkout", "-b", name)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// CreateNewIssueBranch creates and checkouts new isuse branch.
func CreateNewIssueBranch(issueNo int) {
	if err := CreateNewBranch("issue/" + strconv.Itoa(issueNo)); err != nil {
		log.Fatalln("not in git repository")
	}
}
