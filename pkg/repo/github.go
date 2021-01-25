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
	"path/filepath"
	"time"
)

type Repo struct {
	Url        string
	Branch     string
	ConfigFile string
	Since      time.Time
}

func CleanDiskStorage(directory string) (error, error) {
	d, err := os.Open(directory)
	logging.CheckIfError(err)

	defer d.Close()
	names, err := d.Readdirnames(-1)
	logging.CheckIfError(err)

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(directory, name))
		logging.CheckIfError(err)
	}
	return nil, nil
}

func DiskStorage(directory string, url string){
	token := os.Getenv("GITHUB_TOKEN")
	user := os.Getenv("GITHUB_USER")

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              &http.BasicAuth{user, token},
	})
	logging.CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	logging.CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	logging.CheckIfError(err)
	log.Print(fmt.Sprintf("AGChash: %s", commit.Hash))
}

func Pull(directory string) {
	token := os.Getenv("GITHUB_TOKEN")
	user := os.Getenv("GITHUB_USER")

	r, err := git.PlainOpen(directory)
    logging.CheckIfError(err)

	w, err := r.Worktree()
	logging.CheckIfError(err)

	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth:       &http.BasicAuth{user, token},
		Force:      true,
	})
	log.Print(fmt.Sprintf("AGCrepo: %s", err))

	ref, err := r.Head()
	logging.CheckIfError(err)

	commit, err := r.CommitObject(ref.Hash())
	logging.CheckIfError(err)
	log.Print(fmt.Sprintf("AGChash: %s", commit.Hash))
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
	logging.CheckIfError(err)

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
	//TODO: log this action
	logging.CheckIfError(err)
	err = cIter.ForEach(func(c *object.Commit) error {
		hash = append(hash, c.Hash.String())
		return nil
	})
	//TODO: log this action
	log.Print(fmt.Sprintf("worker%d: ID:%s returned hash %s \n", worker, jobId, hash))

}

func NewClone(url string, branch string, configFile string, since time.Time) *Repo {

	return &Repo{
		Url:        url,
		Branch:     branch,
		ConfigFile: configFile,
		Since:      since,
	}
}
