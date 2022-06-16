package controllers

import (
	"net/http"
	"strconv" 
	solver "github.com/korawichtawan/eightPuzzleSolver"
	"github.com/gin-gonic/gin"
)


func SolverHandler(c *gin.Context) {
	var board  [3][3]int;
	var minMove int;
	var answer []string;
	var err error;
	rawData := make(map[string]interface{});

	c.Request.ParseForm();
    for key, value := range c.Request.PostForm {
		runeKey := []rune(key);
		cellNumber,_ := strconv.Atoi(string(runeKey[len(runeKey) - 1]));
		row := cellNumber / 3;
		col := cellNumber % 3 ;
		valueInt, _ := strconv.Atoi(value[0]);
		board[row][col] = valueInt;
		rawData[key] = value[0];
    }
	minMove,answer,err = solver.Solve(board);
	getFullAnswer(&answer);
	if err != nil {
		rawData["err"] = err;
	} else {
		rawData["minMove"] = strconv.Itoa(minMove) + "moves";
		rawData["answer"] = answer;
	}
	
	c.HTML(http.StatusOK, "index.html", rawData)
}

func ViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getFullAnswer(answer *[]string) {
	for idx,val := range *answer {
		switch val {
		case "t":
			(*answer)[idx] = "top";
		case "l":
			(*answer)[idx] = "left";
		case "b":
			(*answer)[idx] = "bottom";
		case "r":
			(*answer)[idx] = "right";
		}
	}
}