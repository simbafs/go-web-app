npm=pnpm

help:
	@echo "Available commands:"
	@echo "Development mode"
	@echo "  dev           - Start development servers for frontend and backend"
	@echo "  devBackend    - Start the backend development server"
	@echo "  devFrontend   - Start the frontend development server"
	@echo "Dependence"
	@echo "  dep           - Install dependencies for frontend and backend"
	@echo "  depBackend    - Install backend dependencies"
	@echo "  depFrontend   - Install frontend dependencies"
	@echo "Build"
	@echo "  build         - Build both frontend and backend"
	@echo "  buildFrontend - Build frontend"
	@echo "  buildBackend  - Build backend"
	@echo "  buildDist     - Build dist from docker"
	@echo "Misc"
	@echo "  doctor        - Check tools"
	@echo "  clean         - Clean generated files"
	@echo "  format        - Format code"

doctor:
	@command -v tmux >/dev/null 2>&1 && echo "tmux is installed" || echo "tmux is NOT installed, the 'dev' target will not work"
	@command -v nodemon >/dev/null 2>&1 && echo "nodemon is installed" || echo "nodemon is NOT installed, the 'dev' and 'devFrontend' targets will not work"

dev: 
	tmux split-window -h make devFrontend
	make devBackend

dep: depBackend depFrontend

depBackend:
	cd ./backend/ && go mod download
	mkdir -p ./backend/static/ 
	touch ./backend/static/.gitkeep

depFrontend:
	cd ./frontend/ && $(npm) install

devBackend: 
	cd ./backend/ && nodemon -e go --watch './**/*.go' --signal SIGTERM --exec 'go' run . 

devFrontend:
	cd ./frontend/ && $(npm) run dev

build: buildFrontend buildBackend

buildDist:
	docker build --output out .

buildFrontend:
	cd ./frontend/ && NODE_ENV=production $(npm) run build

buildBackend:
	rm -rf ./backend/static/
	mv ./frontend/out/ ./backend/static/
	cd ./backend/ && ./build.sh

format:
	cd frontend && prettier --write src 
	cd backend && gofmt -w .

clean:
	rm -rf ./main
	rm -rf ./backend/static/
	rm -rf ./frontend/out/ ./frontend/node_modules/ ./frontend/.next/
	mkdir ./backend/static/

.PHONY: devBackend devFrontend build buildBackend buildFrontend clean dep depFrontend depBackend
