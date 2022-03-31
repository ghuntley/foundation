import (
	"namespacelabs.dev/foundation/std/fn"
)

server: fn.#OpaqueServer & {
	id:   "422eajpp5jt8pjwfp2vrq5f0ce"
	name: "postgresql"

	isStateful: true

	binary: {
		image: "postgres:14.0"
	}

	import: [
		"namespacelabs.dev/foundation/universe/db/postgres/server/creds",	
		"namespacelabs.dev/foundation/universe/db/postgres/server/data",
	]

	service: "postgres": {
		containerPort: 5432
		metadata: protocol: "tcp"
	}
}
