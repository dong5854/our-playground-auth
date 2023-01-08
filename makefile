entgo-init:
	go run entgo.io/ent/cmd/ent init $(SCHEMA)
.PHONY:entgo-init

entgo-generate:
	go generate ./ent
.PHONY:entgo-generate
