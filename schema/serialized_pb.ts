// @generated by protobuf-ts 2.7.0 with parameter force_disable_services,add_pb_suffix,force_exclude_all_options
// @generated from protobuf file "schema/serialized.proto" (package "foundation.schema", syntax proto3)
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
import { SerializedProgram } from "./definition_pb";
import { IngressFragment } from "./networking_pb";
import { Server } from "./schema_pb";
import { Stack } from "./schema_pb";
import { Environment } from "./schema_pb";
/**
 * @generated from protobuf message foundation.schema.DeployPlan
 */
export interface DeployPlan {
    /**
     * @generated from protobuf field: foundation.schema.Environment environment = 7;
     */
    environment?: Environment;
    /**
     * @generated from protobuf field: foundation.schema.Stack stack = 1;
     */
    stack?: Stack;
    /**
     * @generated from protobuf field: repeated foundation.schema.Server focus_server = 2;
     */
    focusServer: Server[];
    /**
     * @generated from protobuf field: repeated string rel_location = 3;
     */
    relLocation: string[];
    /**
     * @generated from protobuf field: repeated foundation.schema.IngressFragment ingress_fragment = 4;
     */
    ingressFragment: IngressFragment[];
    /**
     * @generated from protobuf field: repeated string hints = 5;
     */
    hints: string[];
    /**
     * @generated from protobuf field: foundation.schema.SerializedProgram program = 6;
     */
    program?: SerializedProgram;
}
// @generated message type with reflection information, may provide speed optimized methods
class DeployPlan$Type extends MessageType<DeployPlan> {
    constructor() {
        super("foundation.schema.DeployPlan", [
            { no: 7, name: "environment", kind: "message", T: () => Environment },
            { no: 1, name: "stack", kind: "message", T: () => Stack },
            { no: 2, name: "focus_server", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Server },
            { no: 3, name: "rel_location", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "ingress_fragment", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => IngressFragment },
            { no: 5, name: "hints", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "program", kind: "message", T: () => SerializedProgram }
        ]);
    }
    create(value?: PartialMessage<DeployPlan>): DeployPlan {
        const message = { focusServer: [], relLocation: [], ingressFragment: [], hints: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<DeployPlan>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeployPlan): DeployPlan {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* foundation.schema.Environment environment */ 7:
                    message.environment = Environment.internalBinaryRead(reader, reader.uint32(), options, message.environment);
                    break;
                case /* foundation.schema.Stack stack */ 1:
                    message.stack = Stack.internalBinaryRead(reader, reader.uint32(), options, message.stack);
                    break;
                case /* repeated foundation.schema.Server focus_server */ 2:
                    message.focusServer.push(Server.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated string rel_location */ 3:
                    message.relLocation.push(reader.string());
                    break;
                case /* repeated foundation.schema.IngressFragment ingress_fragment */ 4:
                    message.ingressFragment.push(IngressFragment.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated string hints */ 5:
                    message.hints.push(reader.string());
                    break;
                case /* foundation.schema.SerializedProgram program */ 6:
                    message.program = SerializedProgram.internalBinaryRead(reader, reader.uint32(), options, message.program);
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
    internalBinaryWrite(message: DeployPlan, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* foundation.schema.Environment environment = 7; */
        if (message.environment)
            Environment.internalBinaryWrite(message.environment, writer.tag(7, WireType.LengthDelimited).fork(), options).join();
        /* foundation.schema.Stack stack = 1; */
        if (message.stack)
            Stack.internalBinaryWrite(message.stack, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated foundation.schema.Server focus_server = 2; */
        for (let i = 0; i < message.focusServer.length; i++)
            Server.internalBinaryWrite(message.focusServer[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated string rel_location = 3; */
        for (let i = 0; i < message.relLocation.length; i++)
            writer.tag(3, WireType.LengthDelimited).string(message.relLocation[i]);
        /* repeated foundation.schema.IngressFragment ingress_fragment = 4; */
        for (let i = 0; i < message.ingressFragment.length; i++)
            IngressFragment.internalBinaryWrite(message.ingressFragment[i], writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* repeated string hints = 5; */
        for (let i = 0; i < message.hints.length; i++)
            writer.tag(5, WireType.LengthDelimited).string(message.hints[i]);
        /* foundation.schema.SerializedProgram program = 6; */
        if (message.program)
            SerializedProgram.internalBinaryWrite(message.program, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message foundation.schema.DeployPlan
 */
export const DeployPlan = new DeployPlan$Type();
