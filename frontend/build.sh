#!/bin/sh

npm=$(which pnpm)

NODE_ENV=production $npm run build &&
cd out &&
	find . -name "*.html" ! -name "index.html" ! -name "404.html" -exec sh -c '
    for i; do
        f="${i%.*}"
        echo $f

        mkdir -p -- "$f"

        mv -- "$i" "$f/index.html"
    done
' sh {} +
