#!/bin/sh

npm=$(which pnpm)

NODE_ENV=production $npm run build
