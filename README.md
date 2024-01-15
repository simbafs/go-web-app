# go-next-monorepo

Compile the frontend using next (SSG), and then put it into the Go backend to compile into a single executable file.

# get start

```
$ make dep
$ make dev
$ make build
```

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

# TODO

-   add Action to build automatically
-   ~~docker image ?~~
-   add descripton of make command
