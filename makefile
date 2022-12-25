entgo-init:
	go run entgo.io/ent/cmd/ent init $(SCHEMA)
.PHONY:entgo-init