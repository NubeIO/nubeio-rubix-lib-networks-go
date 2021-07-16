# rubix-lib-rest-go

setup, clone the repo

```
go mod download
```

````
go mod <command> [arguments]

The commands are:

        download    download modules to local cache
        edit        edit go.mod from tools or scripts
        graph       print module requirement graph
        init        initialize new module in current directory
        tidy        add missing and remove unused modules
        vendor      make vendored copy of dependencies
        verify      verify dependencies have expected content
        why         explain why packages or modules are needed

```

## Filtering, Search and Pagination
UGin has it's own filtering, search and pagination system. You just need to use these parameters.

**Query parameters:**
```
/networks/?Limit=2
/networks/?Offset=0
/networks/?Sort=ID
/networks/?Order=DESC
/networks/?Search=hello
```
Full: **http://localhost:1920/api/networks/?Limit=25&Offset=0&Sort=ID&Order=DESC&Search=hello**

## Middlewares
### 1. Logger and Recovery Middlewares
Gin has 2 important built-in middlewares: **Logger** and **Recovery**. UGin calls these two in default.
```
router := gin.Default()
```

This is same with the following lines.
```
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
```

### 2. CORS Middleware
CORS is important for API's and UGin has it's own CORS middleware in **include/middleware.go**. CORS middleware is called with the code below.
```
router.Use(include.CORS())
```
There is also a good repo for this: https://github.com/gin-contrib/cors

### 3. BasicAuth Middleware
Almost every API needs a protected area. Gin has **BasicAuth** middleware for protecting routes. Basic Auth is an authorization type that requires a verified username and password to access a data resource. In UGin, you can find an example for a basic auth. To access these protected routes, you need to add **Basic Authorization credentials** in your requests. If you try to reach these endpoints from browser, you should see a window prompting you for username and password.

```
authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
    "username": "password",
}))

// /admin/dashboard endpoint is now protected
authorized.GET("/dashboard", controller.Dashboard)
```


## build for ARM 7 (BBB)

```
sudo apt-get install g++-arm-linux-gnueabi 
sudo apt-get install g++-arm-linux-gnueabihf
```
As its using sqlite more detail is needed in the cross compile
```
env GOOS=linux GOARCH=arm CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc GOARM=7 go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension`

```

```
env GOOS=linux GOARCH=arm GOARM=7 go build 
```
to run change permissions if needed
```
chmod +rwxrwxrwx go-lora-decoder 
or
chmod 777 go-lora-decoder 
```

