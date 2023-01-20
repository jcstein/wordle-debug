/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "wordle.wordle";

export interface Wordle {
  index: string;
  word: string;
  submitter: string;
}

function createBaseWordle(): Wordle {
  return { index: "", word: "", submitter: "" };
}

export const Wordle = {
  encode(message: Wordle, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.word !== "") {
      writer.uint32(18).string(message.word);
    }
    if (message.submitter !== "") {
      writer.uint32(26).string(message.submitter);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Wordle {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWordle();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.word = reader.string();
          break;
        case 3:
          message.submitter = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Wordle {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      word: isSet(object.word) ? String(object.word) : "",
      submitter: isSet(object.submitter) ? String(object.submitter) : "",
    };
  },

  toJSON(message: Wordle): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.word !== undefined && (obj.word = message.word);
    message.submitter !== undefined && (obj.submitter = message.submitter);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Wordle>, I>>(object: I): Wordle {
    const message = createBaseWordle();
    message.index = object.index ?? "";
    message.word = object.word ?? "";
    message.submitter = object.submitter ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
