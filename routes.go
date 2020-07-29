package main

import (
	"fmt"
	"github.com/joram/jsnek/api"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	games = map[string][]api.SnakeRequest{
		"example": []api.SnakeRequest{{
			Game: api.Game{ID: "123"},
			Turn: 1,
			You:  exampleBoard.Snakes[0],
		}},
	}
)

func Start(res http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	respond(res, api.StartResponse{
		APIVersion: 1,
		Author: "John Oram",
		Color: "#75CEDD",
		Head: "silly",
		Tail: "fat-rattle",
	})
}

func Move(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	start := time.Now()
	sr := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &sr)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	_, exists := games[sr.Game.ID]
	if !exists {
		games[sr.Game.ID] = []api.SnakeRequest{}
	}
	games[sr.Game.ID] = append(games[sr.Game.ID], sr)

	response := api.MoveResponse{Move: move(sr)}
	delta := time.Since(start)
	fmt.Printf("took %s\n", delta)
	respond(res, response)
}

func End(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	sr := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &sr)
	if err != nil {
		log.Printf("Bad end request: %v", err)
	}
	games[sr.Game.ID] = append(games[sr.Game.ID], sr)

	//b, err := json.Marshal(games[sr.Game.ID])
	//if err != nil {
	//	log.Printf("Bad end request: %v", err)
	//}
	//content := string(b)

	//timeStr := time.Now().Format("2006-01")
	//util.WriteToS3("jsnek", fmt.Sprintf("%s/%s.json", timeStr, sr.Game.ID), content)
	return
}

func Ping(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	return
}

func Debug(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	gameID := params.ByName("gameID")
	if gameID == "" {
		const tpl = `<!DOCTYPE html><html><body>
		{{range .Games}}<div><a href="/debug/{{ . }}">{{ . }}</a></div>{{else}}<div><strong>no games</strong></div>{{end}}
	</body></html>`
		t, err := template.New("webpage").Parse(tpl)
		if err != nil {
			log.Fatal(err)
		}

		gamesNames := []string{}
		for n, _ := range games {
			gamesNames = append(gamesNames, n)
		}
		renderContext := struct{ Games []string }{
			Games: gamesNames,
		}
		t.Execute(res, renderContext)
		return
	}

	const tpl2 = `<!DOCTYPE html><html>
	<head>
      <script
  		src="https://code.jquery.com/jquery-3.3.1.min.js"
  		integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8="
  		crossorigin="anonymous"></script>
	  <script src="/static/boards.js"></script>
	  <link href="/static/boards.css" rel="stylesheet" type="text/css" media="all">
	</head>
	<body>
      <div id="board"></div>
	  <script>
        requests = {{ .SnakeRequests }};	
		$( document ).ready(function() {
    	  render_board(0, requests);
		});
	  </script>
	</body>
  </html>`
	t, err := template.New("webpage").Parse(tpl2)
	if err != nil {
		log.Fatal(err)
	}

	renderContext := struct{ SnakeRequests []api.SnakeRequest }{
		SnakeRequests: games[gameID],
	}
	t.Execute(res, renderContext)
}
