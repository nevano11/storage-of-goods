.DEFAULT_GOAL: go/up

include make/docker.mk

include make/migrate.mk

include make/golang.mk