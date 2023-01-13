package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/KunalAnkur/todo-app/helper"
	mongomodel "github.com/KunalAnkur/todo-app/model/mongo_model"
	mysqlmodel "github.com/KunalAnkur/todo-app/model/sql_model"
	"github.com/KunalAnkur/todo-app/router"
)

func init() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose your database: \n1. MongoDB\n2. MySql")

	input, _ := reader.ReadString('\n')
	index, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 0)

	helper.DATABASE_INDEX = int(index)
	switch index {
	case 1:
		mongomodel.MongoConnect()
	case 2:
		mysqlmodel.MySqlConnect()
	}
}

func main() {

	r := router.Router()

	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
