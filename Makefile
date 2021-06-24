include .env
export $(shell sed 's/=.*//' .env)

PWD=$(shell pwd)

.PHONY: help
help:
	@echo no help

.PHONY: cli
cli:
	zsh
