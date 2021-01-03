package repo

import (
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
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

	r, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:        x.Url,
		Auth:       &http.BasicAuth{user, token},
		RemoteName: x.Branch,
	})
	if err == nil {
		ref, err := r.Head()

		currentTime := time.Now()

		cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &x.Since, Until: &currentTime})
		pkg.CheckIfError(err, worker, jobId)

		err = cIter.ForEach(func(c *object.Commit) error {
			hash = append(hash, c.Hash.String())
			return nil
		})
		pkg.CheckIfError(err, worker, jobId)
		pkg.Info("worker%d: ID:%s hash %s \n", worker, jobId, hash)
	} else {
		pkg.CheckIfError(err, worker, jobId)
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
