package git

import (
	"errors"
	"fmt"

	"github.com/go-git/go-billy/v5"
	bmemory "github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Repository allows for easy access of files in a remote git repo
type Repository struct {
	store *memory.Storage
	billy.Filesystem
	repo     *git.Repository
	worktree *git.Worktree
}

// Checkout sets the Repository's commit
func (r *Repository) Checkout(commit string) error {
	branch := ""
	ref := fmt.Sprintf("%s:refs/head", commit)
	hash := plumbing.NewHash(commit)
	if hash.IsZero() {
		branch = fmt.Sprintf("refs/heads/%s", commit)
		ref = fmt.Sprintf("refs/heads/%[1]s:refs/heads/%[1]s", commit)
	}

	if err := r.repo.Fetch(&git.FetchOptions{
		RemoteName: "origin",
		RefSpecs:   []config.RefSpec{config.RefSpec(ref)},
		Depth:      1,
	}); err != nil {
		if !errors.Is(err, git.NoErrAlreadyUpToDate) {
			return fmt.Errorf("could not fetch repository: %w", err)
		}
	}

	if err := r.worktree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(branch),
		Hash:   hash,
	}); err != nil {
		return fmt.Errorf("could not checkout commit: %w", err)
	}

	return nil
}

// Hash returns the hash currently checked out
func (r *Repository) Hash() (string, error) {
	h, err := r.repo.Head()
	if err != nil {
		return "", fmt.Errorf("could not get HEAD: %w", err)
	}

	return h.Hash().String(), nil
}

// New returns a new Repository with the given URL and commit hash
func New(repoURL, commit string) (*Repository, error) {
	store := memory.NewStorage()
	fs := bmemory.New()

	repo, err := git.Init(store, fs)
	if err != nil {
		return nil, fmt.Errorf("could not init repo: %w", err)
	}

	if _, err = repo.CreateRemote(
		&config.RemoteConfig{
			Name: "origin",
			URLs: []string{repoURL},
		},
	); err != nil {
		return nil, fmt.Errorf("could not create remote: %w", err)
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("could not get worktree: %w", err)
	}

	r := &Repository{store: store, Filesystem: fs, repo: repo, worktree: worktree}

	if err = r.Checkout(commit); err != nil {
		return nil, err
	}

	return r, nil
}
