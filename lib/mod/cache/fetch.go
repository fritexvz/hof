package cache

import (
	"fmt"
	"os"
	"strings"

	googithub "github.com/google/go-github/v30/github"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"

	"github.com/hofstadter-io/hof/lib/yagu"
	"github.com/hofstadter-io/hof/lib/yagu/repos/github"
)

func Fetch(lang, mod, ver string) (err error) {
	flds := strings.Split(mod, "/")
	remote := flds[0]
	owner := flds[1]
	repo := flds[2]
	tag := ver

	dir := Outdir(lang, remote, owner, repo, tag)

	_, err = os.Lstat(dir)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" && err.Error() != "no such file or directory" {
			return err
		}
		// not found
		fetch(lang, mod, ver)
	}

	// else we have it already
	// fmt.Println("Found in cache")

	return nil
}

func fetch(lang, mod, ver string) error {
	flds := strings.Split(mod, "/")
	remote := flds[0]
	owner := flds[1]
	repo := flds[2]
	tag := ver

	switch remote {
	case "github.com":
		return fetchGitHub(lang, owner, repo, tag)

	default:
		return fmt.Errorf("Unknown remote: %q in %s", remote, mod)
	}
}

func fetchGitHub(lang, owner, repo, tag string) (err error) {
	FS := memfs.New()

	if tag == "v0.0.0" {
		err = fetchGitHubBranch(FS, lang, owner, repo, "")
	} else {
		err = fetchGitHubTag(FS, lang, owner, repo, tag)
	}
	if err != nil {
		return fmt.Errorf("While fetching from github\n%w\n", err)
	}

	/*
	fmt.Println("filelist:")
	files, err := yagu.BillyGetFilelist(FS)
	if err != nil {
		return fmt.Errorf("While getting filelist\n%w\n", err)
	}

	for _, f := range files {
		fmt.Println(" -", f.Name())
	}

	fmt.Println("Writing...", )
	*/
	err = Write(lang, "github.com", owner, repo, tag, FS)
	if err != nil {
		return fmt.Errorf("While writing to cache\n%w\n", err)
	}

	return nil
}
func fetchGitHubBranch(FS billy.Filesystem, lang, owner, repo, branch string) error {
	client, err := github.NewClient()
	if err != nil {
		return err
	}

	// TODO find and set default branch
	if branch == "" {
		branch = "master"
		r, err := github.GetRepo(client, owner, repo)
		if err != nil {
			return err
		}

		fmt.Printf("%#+v\n", *r)
	}

	// fmt.Println("Fetch github BRANCH", lang, owner, repo, branch)

	zReader, err := github.FetchBranchZip(client, branch)
	if err != nil {
		return fmt.Errorf("While fetching branch zipfile\n%w\n", err)
	}

	err = yagu.BillyLoadFromZip(zReader, FS, true)
	if err != nil {
		return fmt.Errorf("While reading branch zipfile\n%w\n", err)
	}

	return nil
}
func fetchGitHubTag(FS billy.Filesystem, lang, owner, repo, tag string) error {
	// fmt.Println("Fetch github TAG", lang, owner, repo, tag)
	client, err := github.NewClient()
	if err != nil {
		return err
	}

	tags, err := github.GetTags(client, owner, repo)
	if err != nil {
		return err
	}

	// The tag we are looking for
	var T *googithub.RepositoryTag
	for _, t := range tags {
		if tag != "" && tag == *t.Name {
			T = t
			// fmt.Printf("FOUND  %v\n", *t.Name)
		}
	}
	if T == nil {
		return fmt.Errorf("Did not find tag %q for 'https://github.com/%s/%s' @%s", tag, owner, repo, tag)
	}

	zReader, err := github.FetchTagZip(client, T)
	if err != nil {
		return fmt.Errorf("While fetching tag zipfile\n%w\n", err)
	}

	err = yagu.BillyLoadFromZip(zReader, FS, true)
	if err != nil {
		return fmt.Errorf("While reading tag zipfile\n%w\n", err)
	}

	return nil
}
