import * as $protobuf from "protobufjs";
/** Namespace pb. */
export namespace pb {

    /** Properties of a Telemetry. */
    interface ITelemetry {

        /** Telemetry DeviceID */
        DeviceID?: (number|null);

        /** Telemetry ActionTime */
        ActionTime?: (google.protobuf.ITimestamp|null);

        /** Telemetry ID */
        ID?: (number|null);

        /** Telemetry Value */
        Value?: (number|null);
    }

    /** Represents a Telemetry. */
    class Telemetry implements ITelemetry {

        /**
         * Constructs a new Telemetry.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb.ITelemetry);

        /** Telemetry DeviceID. */
        public DeviceID: number;

        /** Telemetry ActionTime. */
        public ActionTime?: (google.protobuf.ITimestamp|null);

        /** Telemetry ID. */
        public ID: number;

        /** Telemetry Value. */
        public Value: number;

        /**
         * Creates a new Telemetry instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Telemetry instance
         */
        public static create(properties?: pb.ITelemetry): pb.Telemetry;

        /**
         * Encodes the specified Telemetry message. Does not implicitly {@link pb.Telemetry.verify|verify} messages.
         * @param message Telemetry message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb.ITelemetry, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Telemetry message, length delimited. Does not implicitly {@link pb.Telemetry.verify|verify} messages.
         * @param message Telemetry message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb.ITelemetry, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Telemetry message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Telemetry
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb.Telemetry;

        /**
         * Decodes a Telemetry message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Telemetry
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb.Telemetry;

        /**
         * Verifies a Telemetry message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Telemetry message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Telemetry
         */
        public static fromObject(object: { [k: string]: any }): pb.Telemetry;

        /**
         * Creates a plain object from a Telemetry message. Also converts values to other types if specified.
         * @param message Telemetry
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb.Telemetry, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Telemetry to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }

    /** Properties of a Health. */
    interface IHealth {

        /** Health DeviceID */
        DeviceID?: (number|null);

        /** Health ActionTime */
        ActionTime?: (google.protobuf.ITimestamp|null);

        /** Health Value */
        Value?: (number|null);
    }

    /** Represents a Health. */
    class Health implements IHealth {

        /**
         * Constructs a new Health.
         * @param [properties] Properties to set
         */
        constructor(properties?: pb.IHealth);

        /** Health DeviceID. */
        public DeviceID: number;

        /** Health ActionTime. */
        public ActionTime?: (google.protobuf.ITimestamp|null);

        /** Health Value. */
        public Value: number;

        /**
         * Creates a new Health instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Health instance
         */
        public static create(properties?: pb.IHealth): pb.Health;

        /**
         * Encodes the specified Health message. Does not implicitly {@link pb.Health.verify|verify} messages.
         * @param message Health message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: pb.IHealth, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Health message, length delimited. Does not implicitly {@link pb.Health.verify|verify} messages.
         * @param message Health message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: pb.IHealth, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Health message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Health
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): pb.Health;

        /**
         * Decodes a Health message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Health
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): pb.Health;

        /**
         * Verifies a Health message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Health message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Health
         */
        public static fromObject(object: { [k: string]: any }): pb.Health;

        /**
         * Creates a plain object from a Health message. Also converts values to other types if specified.
         * @param message Health
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: pb.Health, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Health to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };
    }
}

/** Namespace google. */
export namespace google {

    /** Namespace protobuf. */
    namespace protobuf {

        /** Properties of a Timestamp. */
        interface ITimestamp {

            /** Timestamp seconds */
            seconds?: (number|Long|null);

            /** Timestamp nanos */
            nanos?: (number|null);
        }

        /** Represents a Timestamp. */
        class Timestamp implements ITimestamp {

            /**
             * Constructs a new Timestamp.
             * @param [properties] Properties to set
             */
            constructor(properties?: google.protobuf.ITimestamp);

            /** Timestamp seconds. */
            public seconds: (number|Long);

            /** Timestamp nanos. */
            public nanos: number;

            /**
             * Creates a new Timestamp instance using the specified properties.
             * @param [properties] Properties to set
             * @returns Timestamp instance
             */
            public static create(properties?: google.protobuf.ITimestamp): google.protobuf.Timestamp;

            /**
             * Encodes the specified Timestamp message. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @param message Timestamp message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encode(message: google.protobuf.ITimestamp, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @param message Timestamp message or plain object to encode
             * @param [writer] Writer to encode to
             * @returns Writer
             */
            public static encodeDelimited(message: google.protobuf.ITimestamp, writer?: $protobuf.Writer): $protobuf.Writer;

            /**
             * Decodes a Timestamp message from the specified reader or buffer.
             * @param reader Reader or buffer to decode from
             * @param [length] Message length if known beforehand
             * @returns Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): google.protobuf.Timestamp;

            /**
             * Decodes a Timestamp message from the specified reader or buffer, length delimited.
             * @param reader Reader or buffer to decode from
             * @returns Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): google.protobuf.Timestamp;

            /**
             * Verifies a Timestamp message.
             * @param message Plain object to verify
             * @returns `null` if valid, otherwise the reason why it is not
             */
            public static verify(message: { [k: string]: any }): (string|null);

            /**
             * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
             * @param object Plain object
             * @returns Timestamp
             */
            public static fromObject(object: { [k: string]: any }): google.protobuf.Timestamp;

            /**
             * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
             * @param message Timestamp
             * @param [options] Conversion options
             * @returns Plain object
             */
            public static toObject(message: google.protobuf.Timestamp, options?: $protobuf.IConversionOptions): { [k: string]: any };

            /**
             * Converts this Timestamp to JSON.
             * @returns JSON object
             */
            public toJSON(): { [k: string]: any };
        }
    }
}
