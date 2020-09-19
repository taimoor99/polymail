package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"polymail/app/controller"
	"polymail/app/routes"
	"polymail/app/services"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	goc, err := services.GetSession(ctx)
	if err != nil {
		fmt.Println("[mongo] mongodb connection error")
		panic(err)
	}
	defer goc.Disconnect(ctx)

	userDB := services.DbClient(goc)
	ctrl := controller.DraftMailRepository(userDB)
	r := routes.NewRouter(ctrl)

	fmt.Println("server started on PORT 8070")
	http.ListenAndServe(":" + os.Getenv("PORT"), r)
}
