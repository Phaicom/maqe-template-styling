package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/phaicom/maqe-template-styling/models"
)

func ForumEndpoint(w http.ResponseWriter, r *http.Request) {
	forums, err := GetForum()
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(forums)
	if err != nil {
		log.Fatal(err)
	}
}

func GetForum() (*[]models.Forum, error) {
	authors, err := GetAuthors()
	if err != nil {
		return nil, err
	}

	posts, err := GetPosts()
	if err != nil {
		return nil, err
	}

	forums := make([]models.Forum, 0)

	for _, post := range *posts {
		for _, author := range *authors {
			if post.AuthorID == author.ID {
				forum := models.Forum{
					Author: author,
					Post:   post,
				}
				forums = append(forums, forum)
			}
		}
	}

	return &forums, nil
}

func GetAuthors() (*[]models.Author, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://maqe.github.io/json/authors.json", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	authors := make([]models.Author, 0)
	err = json.Unmarshal(body, &authors)
	if err != nil {
		return nil, err
	}

	// log.Printf("GetAuthor %+v successful", authors)
	return &authors, nil
}

func GetPosts() (*[]models.Post, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://maqe.github.io/json/posts.json", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	posts := make([]models.Post, 0)
	err = json.Unmarshal(body, &posts)
	if err != nil {
		return nil, err
	}

	// log.Printf("GetPost %+v successful", posts)
	return &posts, nil
}
