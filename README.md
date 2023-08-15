# go-next-monorepo

Compile the frontend using next (SSG), and then put it into the Go backend to compile into a single executable file.

# feature

-   output single executable file
-   single port in dev mode
-   hot reload when file change in dev mode
-   small docker image size(~14MB)

## frontend

**Notice** For SSG, do not use next api route

-   nextjs
-   ts
-   eslint
-   tailwindcss
-   pnpm

## backend

-   gin
-   go embed

# Usage

```
$ make backend # start backend dev server(hotreload with nodemon, you need install in global)
$ make frontend # start frontend dev server
$ make build # build frontend, embed into backend server and build into a single executable file
```

# TODO

-   add Action to build automatically
-   ~~docker image ?~~
