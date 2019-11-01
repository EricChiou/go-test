package main

import (
	"fmt"

	"github.com/EricChiou/goroutinepool"
	"github.com/EricChiou/httprouter"
)

func main() {
	// test goroutine pool
	// testGoroutinepool()

	// test http router
	testHttprouter()
}

func testGoroutinepool() {
	pool := goroutinepool.New(5) // new gorotine pool has 5 gorotine
	for i := 0; i < 1000; i++ {
		pool.Add()
		go func(i int) {
			fmt.Println("gorotine ", i)
			pool.Done()
		}(i)
	}
	pool.Wait()
}

func testHttprouter() {
	httprouter.Get("/", handler)
	httprouter.Get("/path", handler)
	httprouter.Get("/path/id/path2", handler)
	httprouter.Get("/path/path/path2", handler)

	// path parameter
	httprouter.Get("/path/:id/path", handlerParam)
	httprouter.Get("/:id/path", handlerParam)
	httprouter.Get("/path/path/path2/:id", handlerParam)

	// duplicate path
	// httprouter.Get("/path/path", handler)

	// invalid character, only accept 0-9, a-z, A-Z
	// httprouter.Get("/path/&", handler)
	// httprouter.Get("/path/:!", handler)

	// wrong format
	// httprouter.Get("path/path", handler) // should start with "/"
	// httprouter.Get("/path/path/", handler) // should not end with "/"
	// httprouter.Get("/path//path", handler)

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// start http server
	if err := httprouter.ListenAndServe(":6200"); err != nil {
		panic(err)
	}

	// start https server
	// if err := httprouter.ListenAndServeTLS(":6200", "cert file path...", "key file path..."); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}

func handler(context *httprouter.Context) {
	fmt.Fprintf(context.Ctx, "url path: %s", string(context.Ctx.Path()))
}

func handlerParam(context *httprouter.Context) {
	id, _ := context.GetPathParam("id")
	fmt.Fprintf(context.Ctx, "url path: %s\nid: %s", string(context.Ctx.Path()), id)
}
