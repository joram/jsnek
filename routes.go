package main

import (
	"fmt"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/filters"
	"github.com/joram/jsnek/logic"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var (
	logics = []logic.Responsibility{
		// NO OPTION
		logic.OnlyOneChoice{},
		logic.AvoidHeadOnHead{},
		// ONLY ONE NOT THREATENED CHOICE
		// HUNGRY (health level?) GO FOR FOOD
		logic.GoEatOrthogonal{25},
		// SHORTEST SNAKE GO FOR FOOD
		// POTENTIAL KILL
		// EAT THEIR LUNCH (force them to starve)
		logic.GoMoreRoom{3},
		logic.TrapFood{},
		logic.ValidDirection{},
	}
	directionStrings = map[int]string{
		api.UP: "up",
		api.DOWN: "down",
		api.LEFT: "left",
		api.RIGHT: "right",
		api.UNKNOWN: "WFT!",
	}
	decisionFilters = []filters.DecisionFilter{
		filters.IsUnknownFilter{},
		filters.IsSolidFilter{},
	}
	games = map[string][]api.SnakeRequest{
			"example": []api.SnakeRequest{{
				Game:api.Game{ID:"123"},
				Turn:1,
				You: exampleBoard.Snakes[0],
			}},
	}
)

func Start(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

	for _, l := range logics {
		choice := l.Decision(&sr)
		okChoice := true
		for _, filter := range decisionFilters {
			ok, _ := filter.Allowed(choice, &sr)
			if !ok {
				okChoice = false
				break
			}
		}
		if choice == api.UNKNOWN {
			continue
		}
		if !okChoice {
			println("skipping choice "+directionStrings[choice]+" by "+l.Taunt())
			continue
		}
		fmt.Println(sr.Game.ID, l.Taunt())
		respond(res, api.MoveResponse{
			Move:  directionStrings[choice],
			Taunt: l.Taunt(),
		})
		return
	}

	respond(res, api.MoveResponse{Move: "down"})
}

func End(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	return
}

func Debug(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	gameID := params.ByName("gameID")
	if gameID == ""{
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
		renderContext := struct {Games []string}{
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

	renderContext := struct {SnakeRequests []api.SnakeRequest}{
		SnakeRequests: games[gameID],
	}
	t.Execute(res, renderContext)
}

