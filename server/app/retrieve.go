package app

import (
	"net/http"
	"strings"
	"github.com/subutai-io/cdn/db"
	"github.com/boltdb/bolt"
	"fmt"
)

type SearchRequest struct {
	fileID    string
	name      string
	owner     string
	token     string
	tags      string
	version   string
	repo      string
	operation string
}

// ParseRequest takes HTTP request and converts it into Request struct
func (r *SearchRequest) ParseRequest(req *http.Request) (err error) {
	r.fileID = req.URL.Query().Get("id")
	r.name = req.URL.Query().Get("name")
	r.owner = req.URL.Query().Get("owner")
	r.token = req.URL.Query().Get("token")
	r.tags = req.URL.Query().Get("tags")
	r.version = req.URL.Query().Get("version")
	r.repo = strings.Split(req.RequestURI, "/")[3] // Splitting /kurjun/rest/repo/func into ["", "kurjun", "rest", "repo" (index: 3), "func"]
	return
}

// BuildQuery constructs the query out of the existing parameters in SearchRequest
func (r *SearchRequest) BuildQuery() (query map[string]string) {
	if r.fileID != "" {
		query["fileID"] = r.fileID
	}
	if r.name != "" {
		query["name"] = r.name
	}
	if r.owner != "" {
		query["owner"] = r.owner
	}
	if r.tags != "" {
		query["tags"] = r.tags
	}
	if r.version != "" {
		query["version"] = r.version
	}
	if r.repo != "" {
		query["repo"] = r.repo
	}
	return
}

type SearchResult struct {

}

func Retrieve(request SearchRequest) []SearchResult {
	query := request.BuildQuery()
	results, err := Search(query)
	if err != nil {

	}
//	if operation == ""
	return results
}

func GetFileInfo(id string) (info map[string]string, err error) {
	info["fileID"] = id
	err = db.DB.View(func(tx *bolt.Tx) error {
		file := tx.Bucket(db.MyBucket).Bucket([]byte(id))
		if file == nil {
			return fmt.Errorf("file %s not found", id)
		}
		owner := file.Bucket([]byte("owner"))
		key, _ := owner.Cursor().First()
		info["owner"] = string(key)
		info["name"] = string(file.Get([]byte("name")))
		repo := file.Bucket([]byte("type"))
		if repo != nil {
			key, _ = repo.Cursor().First()
			info["repo"] = string(key)
		}
		if len(info["repo"]) == 0 {
			return fmt.Errorf("couldn't find repo for file %s", id)
		}
		info["version"] = string(file.Get([]byte("version")))
		info["tags"] = string(file.Get([]byte("tag")))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}

func MatchQuery(file, query map[string]string) bool {
	for key, value := range query {
		if file[key] != value {
			return false
		}
 	}
 	return true
}

func Search(query map[string]string) ([]SearchResult, error) {
	db.DB.View(func(tx *bolt.Tx) error {
		files := tx.Bucket(db.MyBucket)
		files.ForEach(func(k, v []byte) error {
			file, err := GetFileInfo(string(k))
			if err != nil {
				return err
			}
			if MatchQuery(file, query) {

			}
			return nil
		})
		return nil
	})
}