package main

import (
	"fmt"
	"net/http"

	"github.com/EricChiou/goroutinepool"
	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	// test goroutine pool
	// testGoroutinepool()

	// test http router
	testHttprouter()
	// testHttprouterFasthttp()
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
	hadler := func(context *httprouter.Context) {
		fmt.Fprintf(context.Rep, "url path: %s", context.Req.RequestURI)
	}
	handlerParam := func(context *httprouter.Context) {
		id, _ := context.GetPathParam("id")
		fmt.Fprintf(context.Rep, "url path: %s\nid: %s", context.Req.RequestURI, id)
	}

	httprouter.Get("/", hadler)
	httprouter.Get("/path", hadler)
	httprouter.Get("/path/id/path2", hadler)
	httprouter.Get("/path/path/path2", hadler)

	// path parameter
	httprouter.Get("/path/:id/path", handlerParam)
	httprouter.Get("/:id/path", handlerParam)
	httprouter.Get("/path/path/path2/:id", handlerParam)

	// duplicate path
	// httprouter.Get("/path/path", hadler)

	// invalid character, only accept 0-9, a-z, A-Z
	// httprouter.Get("/path/&", hadler)
	// httprouter.Get("/path/:!", hadler)

	// wrong format
	// httprouter.Get("path/path", hadler) // should start with "/"
	// httprouter.Get("/path/path/", hadler) // should not end with "/"
	// httprouter.Get("/path//path", hadler)

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// net/http http server
	if err := http.ListenAndServe(":6200", httprouter.HTTPHandler()); err != nil {
		panic(err)
	}

	// net/http https server
	// if err := http.ListenAndServeTLS(":6200", "cert file path...", "key file path...", httprouter.HTTPHandler()); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}

func testHttprouterFasthttp() {
	hadler := func(context *httprouter.Context) {
		fmt.Fprintf(context.Ctx, "url path: %s", string(context.Ctx.Path()))
	}
	handlerParam := func(context *httprouter.Context) {
		id, _ := context.GetPathParam("id")
		fmt.Fprintf(context.Ctx, "url path: %s\nid: %s", string(context.Ctx.Path()), id)
	}

	httprouter.Get("/", hadler)
	httprouter.Get("/path", hadler)
	httprouter.Get("/path/id/path2", hadler)
	httprouter.Get("/path/path/path2", hadler)

	// path parameter
	httprouter.Get("/path/:id/path", handlerParam)
	httprouter.Get("/:id/path", handlerParam)
	httprouter.Get("/path/path/path2/:id", handlerParam)

	// duplicate path
	// httprouter.Get("/path/path", hadler)

	// invalid character, only accept 0-9, a-z, A-Z
	// httprouter.Get("/path/&", hadler)
	// httprouter.Get("/path/:!", hadler)

	// wrong format
	// httprouter.Get("path/path", hadler) // should start with "/"
	// httprouter.Get("/path/path/", hadler) // should not end with "/"
	// httprouter.Get("/path//path", hadler)

	// set headers
	httprouter.SetHeader("Access-Control-Allow-Origin", "*")
	httprouter.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS")
	httprouter.SetHeader("Access-Control-Allow-Headers", "Content-Type")

	// fasthttp http server
	if err := fasthttp.ListenAndServe(":6200", httprouter.FasthttpHandler()); err != nil {
		panic(err)
	}

	// fasthttp https server
	// if err := fasthttp.ListenAndServeTLS(":6200", "cert file path...", "key file path...", httprouter.FasthttpHandler()); err != nil {
	// 	panic(err)
	// }

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("start server error: ", err)
		}
	}()
}
