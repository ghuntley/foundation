package fn

import (
	"namespacelabs.dev/foundation/std/fn:types"
	"namespacelabs.dev/foundation/std/fn:inputs"
)

_#Imports: {
	"import": [...string]
}

_#Node: {
	_#Imports

	instantiate: [#InstanceName]: {
		packageName?:   string
		type?:          string
		typeDefinition: types.#Proto
		with?: {...}
	}

	packageData: [...string]

	requirePersistentStorage?: {
		persistentId: string
		byteCount:    string
		mountPath:    string
	}
}

#Extension: {
	_#Node

	initializeIn?: [...#Framework]

	provides?: #Provides

	#Provides: [X=string]: {
		name: X
		{input: types.#Proto} | {type: types.#Proto}
    availableIn: [#Framework]: {...}
	}
}

#InstanceName: string

#Service: {
	_#Node

	framework: #Framework

	ingress: *"PRIVATE" | "INTERNET_FACING"

	exportService?:        #GrpcService
	exportServicesAsHttp?: bool // XXX move this to the service definition.

	exportHttp?: [...#HttpPath]
}

#GrpcService: types.#Proto

#HttpPath: {
	path:  string
	kind?: string
}

#Framework: "GO" | "GO_GRPC" | "NODEJS_GRPC" | "WEB" | "NODEJS"

#Server: {
	_#Imports

	id:   string
	name: string

	framework: #Framework | "OPAQUE"

	isStateful?: bool

	if framework == "OPAQUE" {
		binary: {
			image: string
		}
	}

	if framework == "OPAQUE" || framework == "NODEJS" {
		service: [string]: #ServiceSpec
	}

	// XXX temporary
	env: [string]: string

	urlmap: [...#UrlMapEntry]

	#ServiceSpec: {
		name?:         string
		containerPort: int
		metadata: {
			kind?:    string
			protocol: string
		}
		internal: *false | true
	}

	#UrlMapEntry: {
		path:    string
		import?: inputs.#Package
	}

	#Naming: {
		withOrg?: string
	}
}

#OpaqueServer: #Server & {
	framework: "OPAQUE"
}

#Image: {
	prebuilt?: string
	src?:      #BuildPlan // XXX validation is done by the Go runtime at the moment.
}

#BuildPlan: {
	buildFile?: string
	imageRoot:  *"." | string
	hermetic:   *false | true
	...
}

#OpaqueBinary: {
	#Image
	command: [...string]
	... // XXX not a real fan of leaving this open; but need to if want extensions to the binary definition.
}

#Args: [string]: string

#WithPackageName: {
	packageName: inputs.#Package
	...
}

_#ConfigureBase: {
	stack?: {
		append: [...#WithPackageName]
	}
	startup?: #Startup
	init?: [...#Init]
	naming?: #Naming
	...

	provisioning?: #Provisioning
	#Provisioning: {
		// XXX add purpose, e.g. contributes startup inputs.
		with?: {
			binary:     inputs.#Package
			args:       #Args
			workingDir: *"/" | string
			mount: [string]: {fromWorkspace: string}
			snapshot: [string]: {fromWorkspace: string}
			noCache:      *false | true
			requiresKeys: *false | true
		}
	}

	#Startup: {
		args?: #Args
		env: [string]: string
	}

	#Init: {
		binary: inputs.#Package
		args:   #Args
	}

	#Naming: {
		withOrg?: string
		*{} | {domainName: string} | {tlsManagedDomainName: string}
	}
}

#Configure: _#ConfigureBase & {
	with?: #Invocation
}

// Deprecated.
#Extend: _#ConfigureBase & {
	provisioning?: {
		with?: #Invocation
	}
}

// XXX add purpose, e.g. contributes startup inputs.
#Invocation: {
	binary:     inputs.#Package
	args:       #Args
	workingDir: *"/" | string
	mount: [string]: {fromWorkspace: string}
	snapshot: [string]: {fromWorkspace: string}
	noCache:      *false | true
	requiresKeys: *false | true
}

#Binary: {
	name?:       string
	repository?: string
	digest?:     string
	from: {
		go_package?:    string
		dockerfile?:    string
		web_build?:     string
		llb_go_binary?: string
	}
}

#Test: {
	name:    string
	binary:  #Binary
	fixture: *{
		sut: inputs.#Package
		serversUnderTest: [sut]
	} | {
		serversUnderTest: [inputs.#Package, ...inputs.#Package]
	}
}
