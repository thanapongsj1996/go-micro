package main

import (
	"go-micro/cmd/microservice"
	"net/http"
)

func main() {
	ms := microservice.NewMicroservice()
	ms.GET("/test/:id", func(ctx microservice.IContext) error {
		id := ctx.Param("id")
		page := ctx.QueryParam("page")
		ctx.Log("GET: /test/" + id)
		citizen := map[string]interface{}{
			"id":   id,
			"page": page,
		}
		ctx.Response(http.StatusOK, citizen)
		return nil
	})

	defer func(ms *microservice.Microservice) {
		err := ms.Cleanup()
		if err != nil {

		}
	}(ms)

	err := ms.Start()
	if err != nil {
		return
	}
}
