// @generated by protobuf-ts 2.7.0 with parameter force_disable_services,add_pb_suffix
// @generated from protobuf file "std/secrets/provider.proto" (package "foundation.std.secrets", syntax proto3)
// tslint:disable
//
// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation
//
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { DeferredInvocationSource } from "../types/invocation_pb";
/**
 * @generated from protobuf message foundation.std.secrets.Secrets
 */
export interface Secrets {
    /**
     * @generated from protobuf field: repeated foundation.std.secrets.Secret secret = 1;
     */
    secret: Secret[];
}
/**
 * @generated from protobuf message foundation.std.secrets.Secret
 */
export interface Secret {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * If specified, a secret is generated according to the specification, rather than being user-specified.
     *
     * @generated from protobuf field: foundation.std.secrets.GenerateSpecification generate = 3;
     */
    generate?: GenerateSpecification;
    /**
     * @generated from protobuf field: foundation.std.types.DeferredInvocationSource initialize_with = 4;
     */
    initializeWith?: DeferredInvocationSource;
    /**
     * By default secrets are required.
     *
     * @generated from protobuf field: bool optional = 5;
     */
    optional: boolean;
    /**
     * @generated from protobuf field: string experimental_mount_as_env_var = 6;
     */
    experimentalMountAsEnvVar: string;
}
/**
 * @generated from protobuf message foundation.std.secrets.GenerateSpecification
 */
export interface GenerateSpecification {
    /**
     * @generated from protobuf field: string unique_id = 1;
     */
    uniqueId: string; // If not set, will default to the package name of the caller.
    /**
     * @generated from protobuf field: int32 random_byte_count = 2;
     */
    randomByteCount: number;
    /**
     * @generated from protobuf field: foundation.std.secrets.GenerateSpecification.Format format = 3;
     */
    format: GenerateSpecification_Format;
}
/**
 * @generated from protobuf enum foundation.std.secrets.GenerateSpecification.Format
 */
export enum GenerateSpecification_Format {
    /**
     * Defaults to base64.
     *
     * @generated from protobuf enum value: FORMAT_UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: FORMAT_BASE64 = 1;
     */
    BASE64 = 1,
    /**
     * @generated from protobuf enum value: FORMAT_BASE32 = 2;
     */
    BASE32 = 2
}
/**
 * @generated from protobuf message foundation.std.secrets.SecretDevMap
 */
export interface SecretDevMap {
    /**
     * @generated from protobuf field: repeated foundation.std.secrets.SecretDevMap.Configure configure = 1;
     */
    configure: SecretDevMap_Configure[];
}
/**
 * @generated from protobuf message foundation.std.secrets.SecretDevMap.Configure
 */
export interface SecretDevMap_Configure {
    /**
     * @generated from protobuf field: string package_name = 1;
     */
    packageName: string;
    /**
     * @generated from protobuf field: repeated foundation.std.secrets.SecretDevMap.SecretSpec secret = 2;
     */
    secret: SecretDevMap_SecretSpec[];
}
/**
 * @generated from protobuf message foundation.std.secrets.SecretDevMap.SecretSpec
 */
export interface SecretDevMap_SecretSpec {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: string from_path = 2;
     */
    fromPath: string;
    /**
     * @generated from protobuf field: string value = 3;
     */
    value: string;
    /**
     * Runtime-specific specification. E.g. in Kubernetes would be actual secret name stored in Kubernetes.
     *
     * @generated from protobuf field: string resource_name = 4;
     */
    resourceName: string;
}
/**
 * @generated from protobuf message foundation.std.secrets.Value
 */
export interface Value {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: string path = 2;
     */
    path: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Secrets$Type extends MessageType<Secrets> {
    constructor() {
        super("foundation.std.secrets.Secrets", [
            { no: 1, name: "secret", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Secret }
        ]);
    }
    create(value?: PartialMessage<Secrets>): Secrets {
        const message = { secret: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Secrets>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Secrets): Secrets {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated foundation.std.secrets.Secret secret */ 1:
                    message.secret.push(Secret.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Secrets, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated foundation.std.secrets.Secret secret = 1; */
        for (let i = 0; i < message.secret.length; i++)
            Secret.internalBinaryWrite(message.secret[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.Secrets
 */
export const Secrets = new Secrets$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Secret$Type extends MessageType<Secret> {
    constructor() {
        super("foundation.std.secrets.Secret", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "generate", kind: "message", T: () => GenerateSpecification, options: { "foundation.std.proto.provision_only": true } },
            { no: 4, name: "initialize_with", kind: "message", T: () => DeferredInvocationSource, options: { "foundation.std.proto.provision_only": true } },
            { no: 5, name: "optional", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "experimental_mount_as_env_var", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Secret>): Secret {
        const message = { name: "", optional: false, experimentalMountAsEnvVar: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Secret>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Secret): Secret {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* foundation.std.secrets.GenerateSpecification generate */ 3:
                    message.generate = GenerateSpecification.internalBinaryRead(reader, reader.uint32(), options, message.generate);
                    break;
                case /* foundation.std.types.DeferredInvocationSource initialize_with */ 4:
                    message.initializeWith = DeferredInvocationSource.internalBinaryRead(reader, reader.uint32(), options, message.initializeWith);
                    break;
                case /* bool optional */ 5:
                    message.optional = reader.bool();
                    break;
                case /* string experimental_mount_as_env_var */ 6:
                    message.experimentalMountAsEnvVar = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Secret, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* foundation.std.secrets.GenerateSpecification generate = 3; */
        if (message.generate)
            GenerateSpecification.internalBinaryWrite(message.generate, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* foundation.std.types.DeferredInvocationSource initialize_with = 4; */
        if (message.initializeWith)
            DeferredInvocationSource.internalBinaryWrite(message.initializeWith, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* bool optional = 5; */
        if (message.optional !== false)
            writer.tag(5, WireType.Varint).bool(message.optional);
        /* string experimental_mount_as_env_var = 6; */
        if (message.experimentalMountAsEnvVar !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.experimentalMountAsEnvVar);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.Secret
 */
export const Secret = new Secret$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GenerateSpecification$Type extends MessageType<GenerateSpecification> {
    constructor() {
        super("foundation.std.secrets.GenerateSpecification", [
            { no: 1, name: "unique_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "random_byte_count", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "format", kind: "enum", T: () => ["foundation.std.secrets.GenerateSpecification.Format", GenerateSpecification_Format, "FORMAT_"] }
        ]);
    }
    create(value?: PartialMessage<GenerateSpecification>): GenerateSpecification {
        const message = { uniqueId: "", randomByteCount: 0, format: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GenerateSpecification>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GenerateSpecification): GenerateSpecification {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string unique_id */ 1:
                    message.uniqueId = reader.string();
                    break;
                case /* int32 random_byte_count */ 2:
                    message.randomByteCount = reader.int32();
                    break;
                case /* foundation.std.secrets.GenerateSpecification.Format format */ 3:
                    message.format = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GenerateSpecification, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string unique_id = 1; */
        if (message.uniqueId !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.uniqueId);
        /* int32 random_byte_count = 2; */
        if (message.randomByteCount !== 0)
            writer.tag(2, WireType.Varint).int32(message.randomByteCount);
        /* foundation.std.secrets.GenerateSpecification.Format format = 3; */
        if (message.format !== 0)
            writer.tag(3, WireType.Varint).int32(message.format);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.GenerateSpecification
 */
export const GenerateSpecification = new GenerateSpecification$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SecretDevMap$Type extends MessageType<SecretDevMap> {
    constructor() {
        super("foundation.std.secrets.SecretDevMap", [
            { no: 1, name: "configure", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => SecretDevMap_Configure }
        ]);
    }
    create(value?: PartialMessage<SecretDevMap>): SecretDevMap {
        const message = { configure: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SecretDevMap>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SecretDevMap): SecretDevMap {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated foundation.std.secrets.SecretDevMap.Configure configure */ 1:
                    message.configure.push(SecretDevMap_Configure.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SecretDevMap, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated foundation.std.secrets.SecretDevMap.Configure configure = 1; */
        for (let i = 0; i < message.configure.length; i++)
            SecretDevMap_Configure.internalBinaryWrite(message.configure[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.SecretDevMap
 */
export const SecretDevMap = new SecretDevMap$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SecretDevMap_Configure$Type extends MessageType<SecretDevMap_Configure> {
    constructor() {
        super("foundation.std.secrets.SecretDevMap.Configure", [
            { no: 1, name: "package_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "secret", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => SecretDevMap_SecretSpec }
        ]);
    }
    create(value?: PartialMessage<SecretDevMap_Configure>): SecretDevMap_Configure {
        const message = { packageName: "", secret: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SecretDevMap_Configure>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SecretDevMap_Configure): SecretDevMap_Configure {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string package_name */ 1:
                    message.packageName = reader.string();
                    break;
                case /* repeated foundation.std.secrets.SecretDevMap.SecretSpec secret */ 2:
                    message.secret.push(SecretDevMap_SecretSpec.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SecretDevMap_Configure, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string package_name = 1; */
        if (message.packageName !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.packageName);
        /* repeated foundation.std.secrets.SecretDevMap.SecretSpec secret = 2; */
        for (let i = 0; i < message.secret.length; i++)
            SecretDevMap_SecretSpec.internalBinaryWrite(message.secret[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.SecretDevMap.Configure
 */
export const SecretDevMap_Configure = new SecretDevMap_Configure$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SecretDevMap_SecretSpec$Type extends MessageType<SecretDevMap_SecretSpec> {
    constructor() {
        super("foundation.std.secrets.SecretDevMap.SecretSpec", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "from_path", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "value", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "resource_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SecretDevMap_SecretSpec>): SecretDevMap_SecretSpec {
        const message = { name: "", fromPath: "", value: "", resourceName: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SecretDevMap_SecretSpec>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SecretDevMap_SecretSpec): SecretDevMap_SecretSpec {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string from_path */ 2:
                    message.fromPath = reader.string();
                    break;
                case /* string value */ 3:
                    message.value = reader.string();
                    break;
                case /* string resource_name */ 4:
                    message.resourceName = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SecretDevMap_SecretSpec, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string from_path = 2; */
        if (message.fromPath !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.fromPath);
        /* string value = 3; */
        if (message.value !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.value);
        /* string resource_name = 4; */
        if (message.resourceName !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.resourceName);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.SecretDevMap.SecretSpec
 */
export const SecretDevMap_SecretSpec = new SecretDevMap_SecretSpec$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Value$Type extends MessageType<Value> {
    constructor() {
        super("foundation.std.secrets.Value", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "path", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Value>): Value {
        const message = { name: "", path: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Value>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Value): Value {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string path */ 2:
                    message.path = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Value, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string path = 2; */
        if (message.path !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.path);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.std.secrets.Value
 */
export const Value = new Value$Type();
