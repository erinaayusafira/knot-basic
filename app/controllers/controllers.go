package controllers

import "fmt"
import "net/http"
import "time"
import "html/template"
import "MVCKnotbasic/app/models"

func GetData(w http.ResponseWriter, r *http.Request){
	var date = models.Datetime{
		Now : time.Now().Format("2006-01-02T15:04:05Z"),
	}

	data := map[string]string{
		"Time" : date.Now,
	}
	var t, err = template.ParseFiles("views/view.html")
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	t.Execute(w, data)
}
