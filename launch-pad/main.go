package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
)

// Basic example of how to checkout a specific commit.
func main() {
	url := "https://github.com/apollo-command-and-service-module/orbit.git"
	//branch := "main"
	directory := "/Users/mimontpe/orbit"

	// Clone the given repository to the given directory
	Info("git clone %s %s", url, directory)
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:           url,
		ReferenceName: plumbing.ReferenceName("refs/heads/qa"),
	})
	CheckIfError(err)

	// Create a new branch to the current HEAD
	Info("git branch my-branch")

	headRef, err := r.Head()
	CheckIfError(err)

	fmt.Println(headRef)

	// Create a new plumbing.HashReference object with the name of the branch
	// and the hash from the HEAD. The reference name should be a full reference
	// name and not an abbreviated one, as is used on the git cli.
	//
	// For tags we should use `refs/tags/%s` instead of `refs/heads/%s` used
	// for branches.
	//ref := plumbing.NewHashReference("refs/heads/my-branch", headRef.Hash())

	// The created reference is saved in the storage.
	//err = r.Storer.SetReference(ref)
	//CheckIfError(err)

	// Or deleted from it.
	//Info("git branch -D my-branch")
	//err = r.Storer.RemoveReference(ref.Name())
	//CheckIfError(err)

	//ref := plumbing.NewHashReference("refs/heads/qa", headRef.Hash())

}
