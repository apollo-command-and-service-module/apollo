package repo

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/apollo-command-and-service-module/apollo/pkg"
	"time"
)

type Repo struct {
	Url        string
	Branch     string
	ConfigFile string
	Since      time.Time
}

func (x Repo) ReadIntoMemory(){
	fs := memfs.New()
	var hash []string

	r, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:        x.Url,
		RemoteName: x.Branch,
	})
	pkg.CheckIfError(err)

	ref, err := r.Head()
	pkg.CheckIfError(err)

	currentTime := time.Now()

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &x.Since, Until: &currentTime})
	pkg.CheckIfError(err)

	err = cIter.ForEach(func(c *object.Commit) error {
		hash = append(hash, c.Hash.String())
		return nil
	})
	pkg.CheckIfError(err)
	pkg.Info("hash: %s\n", hash)

}

func NewClone(url string, branch string, configFile string, since time.Time) *Repo {

	return &Repo{
		Url:        url,
		Branch:     branch,
		ConfigFile: configFile,
		Since:      since,
	}
}

