.PHONY: run
run:
	@go run cmd/bot/main.go

CM = "some fix"
.PHONY: git
git:
	@git add *
	@git commit -m "$(CM)"
	@git push origin master