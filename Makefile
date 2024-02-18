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


# @$(call check,"executable","the target that will not work, no comma is allowed")
define check
	@command -v $(1) > /dev/null 2>&1 && echo "detected $(1)" || echo "$(1) is NOT installed, the '$(2)' target will not work"
	
endef

doctor:
	@$(call check,tmux,dev)
	@$(call check,nodemon,devBackend)
	@$(call check,go,depBackend and buildBackend)
	@$(call check,pnpm,depFrontend and devFrontend and buildFrontend)
	@$(call check,docker,buildDist)
	@$(call check,prettier,format)
	@$(call check,bash,buildFrontend)

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
	cd ./frontend/ && sh build.sh

buildBackend:
	rm -rf ./backend/static/
	cp -r ./frontend/out/ ./backend/static/
	cd ./backend/ && bash build.sh

format:
	cd frontend && prettier --write src 
	cd backend && gofmt -w .

clean:
	rm -rf ./main
	rm -rf ./backend/static/
	rm -rf ./frontend/out/ ./frontend/node_modules/ ./frontend/.next/
	mkdir ./backend/static/

.PHONY: devBackend devFrontend build buildBackend buildFrontend clean dep depFrontend depBackend
