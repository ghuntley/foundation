// This file was automatically generated.

import { DependencyGraph, Initializer } from "@namespacelabs/foundation";
import * as i0 from "@namespacelabs.dev-foundation/languages-nodejs-testdata-services-simple/deps.fn";
import * as i1 from "@namespacelabs.dev-foundation/languages-nodejs-testdata-services-simplehttp/deps.fn";
import * as i2 from "@namespacelabs.dev-foundation/languages-nodejs-testdata-services-numberformatter/deps.fn";
import * as i3 from "@namespacelabs.dev-foundation/languages-nodejs-testdata-services-postuser/deps.fn";
import * as i4 from "@namespacelabs.dev-foundation/std-nodejs-monitoring-tracing/deps.fn";
import * as i5 from "@namespacelabs.dev-foundation/std-nodejs-monitoring-tracing-jaeger/deps.fn";
import * as i6 from "@namespacelabs.dev-foundation/std-nodejs-http/deps.fn";
import * as i7 from "@namespacelabs.dev-foundation/std-nodejs-monitoring-tracing-fastify/deps.fn";

import {provideGrpcRegistrar, GrpcServer} from "@namespacelabs.dev-foundation/std-nodejs-grpc/impl"
import {provideHttpServer, HttpServerImpl} from "@namespacelabs.dev-foundation/std-nodejs-http/impl"

// Returns a list of initialization errors.
const wireServices = async (graph: DependencyGraph): Promise<unknown[]> => {
	const errors: unknown[] = [];
	try {
		await i0.wireService(i0.Package.instantiateDeps(graph));
	} catch (e) {
		errors.push(e);
	}
	try {
		await i1.wireService(i1.Package.instantiateDeps(graph));
	} catch (e) {
		errors.push(e);
	}
	try {
		await i2.wireService(i2.Package.instantiateDeps(graph));
	} catch (e) {
		errors.push(e);
	}
	try {
		await i3.wireService(i3.Package.instantiateDeps(graph));
	} catch (e) {
		errors.push(e);
	}
	return errors;
};

const TransitiveInitializers: Initializer[] = [
	...i4.TransitiveInitializers,
	...i5.TransitiveInitializers,
	...i6.TransitiveInitializers,
	...i7.TransitiveInitializers,
	...i0.TransitiveInitializers,
	...i1.TransitiveInitializers,
	...i2.TransitiveInitializers,
	...i3.TransitiveInitializers,
];

async function main() {
	const graph = new DependencyGraph();
	await graph.runInitializers(TransitiveInitializers);
	const errors = await wireServices(graph);
	if (errors.length > 0) {
		errors.forEach((e) => console.error(e));
		console.error("%d services failed to start.", errors.length);
		process.exit(1);
	}

	(provideGrpcRegistrar() as GrpcServer).start();
	((await provideHttpServer()) as HttpServerImpl).start();
}

main();
