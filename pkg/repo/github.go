package repo

import (
	"fmt"
	"github.com/apollo-command-and-service-module/apollo/pkg/logging"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"log"
	"os"
	"time"
)

type Repo struct {
	Url        string
	Branch     string
	ConfigFile string
	Since      time.Time
}

func (x Repo) ReadIntoMemory(worker int, jobId string) {
	fs := memfs.New()
	var hash []string

	token := os.Getenv("GITHUB_TOKEN")
	user := os.Getenv("GITHUB_USER")

	refpath := fmt.Sprintf("refs/heads/%s", x.Branch)
	r, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:           x.Url,
		ReferenceName: plumbing.ReferenceName(refpath),
		Auth:          &http.BasicAuth{user, token},
		SingleBranch:  true,
	})
	if err != nil {
		//TODO: How we should log this action
		log.Print(fmt.Sprintf("worker%d: ID: %s error %s", worker, jobId, err))
	} else {

		w, err := r.Worktree()
		//TODO: check for error log but don't exit
		logging.CheckIfError(err)

		err = w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName(x.Branch),
		})

		ref, err := r.Head()
		//TODO: check for error log but don't exit
		logging.CheckIfError(err)

		currentTime := time.Now()

		cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &x.Since, Until: &currentTime})
		//TODO: How we should log this action
		logging.CheckIfError(err)
		err = cIter.ForEach(func(c *object.Commit) error {
			hash = append(hash, c.Hash.String())
			return nil
		})
		//TODO: How we should log this action
		log.Print(fmt.Sprintf("worker%d: ID:%s succeeded with hash %s \n", worker, jobId, hash))
	}
}

func NewClone(url string, branch string, configFile string, since time.Time) *Repo {

	return &Repo{
		Url:        url,
		Branch:     branch,
		ConfigFile: configFile,
		Since:      since,
	}
}
