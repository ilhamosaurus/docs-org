.PHONY: tailwind-dev
tailwind-dev:
	npm run tailwind

.PHONY: templ-dev
templ-dev:
	templ generate --watch

.PHONY: dev
dev:
	air