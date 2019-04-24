package repository

import (
	"basic-api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := models.Articles{
		models.Article{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		models.Article{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	_ = json.NewEncoder(w).Encode(articles)
}

func ReturnArticleById(w http.ResponseWriter, r *http.Request) {
	articles := models.Articles{
		models.Article{Id: 1, Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		models.Article{Id: 2, Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	vars := mux.Vars(r)
	key := vars["id"]
	for _, value := range articles {
		if keyValInt, _ := strconv.Atoi(key); keyValInt == value.Id {
			_ = json.NewEncoder(w).Encode(value)
		}
	}
}
