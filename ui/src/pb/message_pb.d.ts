// package: pb
// file: message.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Telemetry extends jspb.Message {
  getDeviceid(): number;
  setDeviceid(value: number): void;

  hasActiontime(): boolean;
  clearActiontime(): void;
  getActiontime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setActiontime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getId(): number;
  setId(value: number): void;

  getValue(): number;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Telemetry.AsObject;
  static toObject(includeInstance: boolean, msg: Telemetry): Telemetry.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Telemetry, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Telemetry;
  static deserializeBinaryFromReader(message: Telemetry, reader: jspb.BinaryReader): Telemetry;
}

export namespace Telemetry {
  export type AsObject = {
    deviceid: number,
    actiontime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    id: number,
    value: number,
  }
}

