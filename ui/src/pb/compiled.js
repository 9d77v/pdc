/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader,
    $Writer = $protobuf.Writer,
    $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.pb = (function () {

    /**
     * Namespace pb.
     * @exports pb
     * @namespace
     */
    var pb = {};

    pb.Telemetry = (function () {

        /**
         * Properties of a Telemetry.
         * @memberof pb
         * @interface ITelemetry
         * @property {number|null} [DeviceID] Telemetry DeviceID
         * @property {google.protobuf.ITimestamp|null} [ActionTime] Telemetry ActionTime
         * @property {number|null} [ID] Telemetry ID
         * @property {number|null} [Value] Telemetry Value
         */

        /**
         * Constructs a new Telemetry.
         * @memberof pb
         * @classdesc Represents a Telemetry.
         * @implements ITelemetry
         * @constructor
         * @param {pb.ITelemetry=} [properties] Properties to set
         */
        function Telemetry(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Telemetry DeviceID.
         * @member {number} DeviceID
         * @memberof pb.Telemetry
         * @instance
         */
        Telemetry.prototype.DeviceID = 0;

        /**
         * Telemetry ActionTime.
         * @member {google.protobuf.ITimestamp|null|undefined} ActionTime
         * @memberof pb.Telemetry
         * @instance
         */
        Telemetry.prototype.ActionTime = null;

        /**
         * Telemetry ID.
         * @member {number} ID
         * @memberof pb.Telemetry
         * @instance
         */
        Telemetry.prototype.ID = 0;

        /**
         * Telemetry Value.
         * @member {number} Value
         * @memberof pb.Telemetry
         * @instance
         */
        Telemetry.prototype.Value = 0;

        /**
         * Creates a new Telemetry instance using the specified properties.
         * @function create
         * @memberof pb.Telemetry
         * @static
         * @param {pb.ITelemetry=} [properties] Properties to set
         * @returns {pb.Telemetry} Telemetry instance
         */
        Telemetry.create = function create(properties) {
            return new Telemetry(properties);
        };

        /**
         * Encodes the specified Telemetry message. Does not implicitly {@link pb.Telemetry.verify|verify} messages.
         * @function encode
         * @memberof pb.Telemetry
         * @static
         * @param {pb.ITelemetry} message Telemetry message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Telemetry.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.DeviceID != null && Object.hasOwnProperty.call(message, "DeviceID"))
                writer.uint32( /* id 1, wireType 0 =*/ 8).uint32(message.DeviceID);
            if (message.ActionTime != null && Object.hasOwnProperty.call(message, "ActionTime"))
                $root.google.protobuf.Timestamp.encode(message.ActionTime, writer.uint32( /* id 2, wireType 2 =*/ 18).fork()).ldelim();
            if (message.ID != null && Object.hasOwnProperty.call(message, "ID"))
                writer.uint32( /* id 3, wireType 0 =*/ 24).uint32(message.ID);
            if (message.Value != null && Object.hasOwnProperty.call(message, "Value"))
                writer.uint32( /* id 4, wireType 1 =*/ 33).double(message.Value);
            return writer;
        };

        /**
         * Encodes the specified Telemetry message, length delimited. Does not implicitly {@link pb.Telemetry.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb.Telemetry
         * @static
         * @param {pb.ITelemetry} message Telemetry message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Telemetry.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Telemetry message from the specified reader or buffer.
         * @function decode
         * @memberof pb.Telemetry
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb.Telemetry} Telemetry
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Telemetry.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length,
                message = new $root.pb.Telemetry();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                    case 1:
                        message.DeviceID = reader.uint32();
                        break;
                    case 2:
                        message.ActionTime = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                        break;
                    case 3:
                        message.ID = reader.uint32();
                        break;
                    case 4:
                        message.Value = reader.double();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                }
            }
            return message;
        };

        /**
         * Decodes a Telemetry message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb.Telemetry
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb.Telemetry} Telemetry
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Telemetry.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Telemetry message.
         * @function verify
         * @memberof pb.Telemetry
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Telemetry.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.DeviceID != null && message.hasOwnProperty("DeviceID"))
                if (!$util.isInteger(message.DeviceID))
                    return "DeviceID: integer expected";
            if (message.ActionTime != null && message.hasOwnProperty("ActionTime")) {
                var error = $root.google.protobuf.Timestamp.verify(message.ActionTime);
                if (error)
                    return "ActionTime." + error;
            }
            if (message.ID != null && message.hasOwnProperty("ID"))
                if (!$util.isInteger(message.ID))
                    return "ID: integer expected";
            if (message.Value != null && message.hasOwnProperty("Value"))
                if (typeof message.Value !== "number")
                    return "Value: number expected";
            return null;
        };

        /**
         * Creates a Telemetry message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb.Telemetry
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb.Telemetry} Telemetry
         */
        Telemetry.fromObject = function fromObject(object) {
            if (object instanceof $root.pb.Telemetry)
                return object;
            var message = new $root.pb.Telemetry();
            if (object.DeviceID != null)
                message.DeviceID = object.DeviceID >>> 0;
            if (object.ActionTime != null) {
                if (typeof object.ActionTime !== "object")
                    throw TypeError(".pb.Telemetry.ActionTime: object expected");
                message.ActionTime = $root.google.protobuf.Timestamp.fromObject(object.ActionTime);
            }
            if (object.ID != null)
                message.ID = object.ID >>> 0;
            if (object.Value != null)
                message.Value = Number(object.Value);
            return message;
        };

        /**
         * Creates a plain object from a Telemetry message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb.Telemetry
         * @static
         * @param {pb.Telemetry} message Telemetry
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Telemetry.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.DeviceID = 0;
                object.ActionTime = null;
                object.ID = 0;
                object.Value = 0;
            }
            if (message.DeviceID != null && message.hasOwnProperty("DeviceID"))
                object.DeviceID = message.DeviceID;
            if (message.ActionTime != null && message.hasOwnProperty("ActionTime"))
                object.ActionTime = $root.google.protobuf.Timestamp.toObject(message.ActionTime, options);
            if (message.ID != null && message.hasOwnProperty("ID"))
                object.ID = message.ID;
            if (message.Value != null && message.hasOwnProperty("Value"))
                object.Value = options.json && !isFinite(message.Value) ? String(message.Value) : message.Value;
            return object;
        };

        /**
         * Converts this Telemetry to JSON.
         * @function toJSON
         * @memberof pb.Telemetry
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Telemetry.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Telemetry;
    })();

    pb.Health = (function () {

        /**
         * Properties of a Health.
         * @memberof pb
         * @interface IHealth
         * @property {number|null} [DeviceID] Health DeviceID
         * @property {google.protobuf.ITimestamp|null} [ActionTime] Health ActionTime
         * @property {number|null} [Value] Health Value
         */

        /**
         * Constructs a new Health.
         * @memberof pb
         * @classdesc Represents a Health.
         * @implements IHealth
         * @constructor
         * @param {pb.IHealth=} [properties] Properties to set
         */
        function Health(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Health DeviceID.
         * @member {number} DeviceID
         * @memberof pb.Health
         * @instance
         */
        Health.prototype.DeviceID = 0;

        /**
         * Health ActionTime.
         * @member {google.protobuf.ITimestamp|null|undefined} ActionTime
         * @memberof pb.Health
         * @instance
         */
        Health.prototype.ActionTime = null;

        /**
         * Health Value.
         * @member {number} Value
         * @memberof pb.Health
         * @instance
         */
        Health.prototype.Value = 0;

        /**
         * Creates a new Health instance using the specified properties.
         * @function create
         * @memberof pb.Health
         * @static
         * @param {pb.IHealth=} [properties] Properties to set
         * @returns {pb.Health} Health instance
         */
        Health.create = function create(properties) {
            return new Health(properties);
        };

        /**
         * Encodes the specified Health message. Does not implicitly {@link pb.Health.verify|verify} messages.
         * @function encode
         * @memberof pb.Health
         * @static
         * @param {pb.IHealth} message Health message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Health.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.DeviceID != null && Object.hasOwnProperty.call(message, "DeviceID"))
                writer.uint32( /* id 1, wireType 0 =*/ 8).uint32(message.DeviceID);
            if (message.ActionTime != null && Object.hasOwnProperty.call(message, "ActionTime"))
                $root.google.protobuf.Timestamp.encode(message.ActionTime, writer.uint32( /* id 2, wireType 2 =*/ 18).fork()).ldelim();
            if (message.Value != null && Object.hasOwnProperty.call(message, "Value"))
                writer.uint32( /* id 3, wireType 0 =*/ 24).uint32(message.Value);
            return writer;
        };

        /**
         * Encodes the specified Health message, length delimited. Does not implicitly {@link pb.Health.verify|verify} messages.
         * @function encodeDelimited
         * @memberof pb.Health
         * @static
         * @param {pb.IHealth} message Health message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Health.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Health message from the specified reader or buffer.
         * @function decode
         * @memberof pb.Health
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {pb.Health} Health
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Health.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length,
                message = new $root.pb.Health();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                    case 1:
                        message.DeviceID = reader.uint32();
                        break;
                    case 2:
                        message.ActionTime = $root.google.protobuf.Timestamp.decode(reader, reader.uint32());
                        break;
                    case 3:
                        message.Value = reader.uint32();
                        break;
                    default:
                        reader.skipType(tag & 7);
                        break;
                }
            }
            return message;
        };

        /**
         * Decodes a Health message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof pb.Health
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {pb.Health} Health
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Health.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Health message.
         * @function verify
         * @memberof pb.Health
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Health.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.DeviceID != null && message.hasOwnProperty("DeviceID"))
                if (!$util.isInteger(message.DeviceID))
                    return "DeviceID: integer expected";
            if (message.ActionTime != null && message.hasOwnProperty("ActionTime")) {
                var error = $root.google.protobuf.Timestamp.verify(message.ActionTime);
                if (error)
                    return "ActionTime." + error;
            }
            if (message.Value != null && message.hasOwnProperty("Value"))
                if (!$util.isInteger(message.Value))
                    return "Value: integer expected";
            return null;
        };

        /**
         * Creates a Health message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof pb.Health
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {pb.Health} Health
         */
        Health.fromObject = function fromObject(object) {
            if (object instanceof $root.pb.Health)
                return object;
            var message = new $root.pb.Health();
            if (object.DeviceID != null)
                message.DeviceID = object.DeviceID >>> 0;
            if (object.ActionTime != null) {
                if (typeof object.ActionTime !== "object")
                    throw TypeError(".pb.Health.ActionTime: object expected");
                message.ActionTime = $root.google.protobuf.Timestamp.fromObject(object.ActionTime);
            }
            if (object.Value != null)
                message.Value = object.Value >>> 0;
            return message;
        };

        /**
         * Creates a plain object from a Health message. Also converts values to other types if specified.
         * @function toObject
         * @memberof pb.Health
         * @static
         * @param {pb.Health} message Health
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Health.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.DeviceID = 0;
                object.ActionTime = null;
                object.Value = 0;
            }
            if (message.DeviceID != null && message.hasOwnProperty("DeviceID"))
                object.DeviceID = message.DeviceID;
            if (message.ActionTime != null && message.hasOwnProperty("ActionTime"))
                object.ActionTime = $root.google.protobuf.Timestamp.toObject(message.ActionTime, options);
            if (message.Value != null && message.hasOwnProperty("Value"))
                object.Value = message.Value;
            return object;
        };

        /**
         * Converts this Health to JSON.
         * @function toJSON
         * @memberof pb.Health
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Health.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Health;
    })();

    return pb;
})();

$root.google = (function () {

    /**
     * Namespace google.
     * @exports google
     * @namespace
     */
    var google = {};

    google.protobuf = (function () {

        /**
         * Namespace protobuf.
         * @memberof google
         * @namespace
         */
        var protobuf = {};

        protobuf.Timestamp = (function () {

            /**
             * Properties of a Timestamp.
             * @memberof google.protobuf
             * @interface ITimestamp
             * @property {number|Long|null} [seconds] Timestamp seconds
             * @property {number|null} [nanos] Timestamp nanos
             */

            /**
             * Constructs a new Timestamp.
             * @memberof google.protobuf
             * @classdesc Represents a Timestamp.
             * @implements ITimestamp
             * @constructor
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             */
            function Timestamp(properties) {
                if (properties)
                    for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                        if (properties[keys[i]] != null)
                            this[keys[i]] = properties[keys[i]];
            }

            /**
             * Timestamp seconds.
             * @member {number|Long} seconds
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.seconds = $util.Long ? $util.Long.fromBits(0, 0, false) : 0;

            /**
             * Timestamp nanos.
             * @member {number} nanos
             * @memberof google.protobuf.Timestamp
             * @instance
             */
            Timestamp.prototype.nanos = 0;

            /**
             * Creates a new Timestamp instance using the specified properties.
             * @function create
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp=} [properties] Properties to set
             * @returns {google.protobuf.Timestamp} Timestamp instance
             */
            Timestamp.create = function create(properties) {
                return new Timestamp(properties);
            };

            /**
             * Encodes the specified Timestamp message. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encode = function encode(message, writer) {
                if (!writer)
                    writer = $Writer.create();
                if (message.seconds != null && Object.hasOwnProperty.call(message, "seconds"))
                    writer.uint32( /* id 1, wireType 0 =*/ 8).int64(message.seconds);
                if (message.nanos != null && Object.hasOwnProperty.call(message, "nanos"))
                    writer.uint32( /* id 2, wireType 0 =*/ 16).int32(message.nanos);
                return writer;
            };

            /**
             * Encodes the specified Timestamp message, length delimited. Does not implicitly {@link google.protobuf.Timestamp.verify|verify} messages.
             * @function encodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.ITimestamp} message Timestamp message or plain object to encode
             * @param {$protobuf.Writer} [writer] Writer to encode to
             * @returns {$protobuf.Writer} Writer
             */
            Timestamp.encodeDelimited = function encodeDelimited(message, writer) {
                return this.encode(message, writer).ldelim();
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer.
             * @function decode
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @param {number} [length] Message length if known beforehand
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decode = function decode(reader, length) {
                if (!(reader instanceof $Reader))
                    reader = $Reader.create(reader);
                var end = length === undefined ? reader.len : reader.pos + length,
                    message = new $root.google.protobuf.Timestamp();
                while (reader.pos < end) {
                    var tag = reader.uint32();
                    switch (tag >>> 3) {
                        case 1:
                            message.seconds = reader.int64();
                            break;
                        case 2:
                            message.nanos = reader.int32();
                            break;
                        default:
                            reader.skipType(tag & 7);
                            break;
                    }
                }
                return message;
            };

            /**
             * Decodes a Timestamp message from the specified reader or buffer, length delimited.
             * @function decodeDelimited
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
             * @returns {google.protobuf.Timestamp} Timestamp
             * @throws {Error} If the payload is not a reader or valid buffer
             * @throws {$protobuf.util.ProtocolError} If required fields are missing
             */
            Timestamp.decodeDelimited = function decodeDelimited(reader) {
                if (!(reader instanceof $Reader))
                    reader = new $Reader(reader);
                return this.decode(reader, reader.uint32());
            };

            /**
             * Verifies a Timestamp message.
             * @function verify
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} message Plain object to verify
             * @returns {string|null} `null` if valid, otherwise the reason why it is not
             */
            Timestamp.verify = function verify(message) {
                if (typeof message !== "object" || message === null)
                    return "object expected";
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (!$util.isInteger(message.seconds) && !(message.seconds && $util.isInteger(message.seconds.low) && $util.isInteger(message.seconds.high)))
                        return "seconds: integer|Long expected";
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    if (!$util.isInteger(message.nanos))
                        return "nanos: integer expected";
                return null;
            };

            /**
             * Creates a Timestamp message from a plain object. Also converts values to their respective internal types.
             * @function fromObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {Object.<string,*>} object Plain object
             * @returns {google.protobuf.Timestamp} Timestamp
             */
            Timestamp.fromObject = function fromObject(object) {
                if (object instanceof $root.google.protobuf.Timestamp)
                    return object;
                var message = new $root.google.protobuf.Timestamp();
                if (object.seconds != null)
                    if ($util.Long)
                        (message.seconds = $util.Long.fromValue(object.seconds)).unsigned = false;
                    else if (typeof object.seconds === "string")
                    message.seconds = parseInt(object.seconds, 10);
                else if (typeof object.seconds === "number")
                    message.seconds = object.seconds;
                else if (typeof object.seconds === "object")
                    message.seconds = new $util.LongBits(object.seconds.low >>> 0, object.seconds.high >>> 0).toNumber();
                if (object.nanos != null)
                    message.nanos = object.nanos | 0;
                return message;
            };

            /**
             * Creates a plain object from a Timestamp message. Also converts values to other types if specified.
             * @function toObject
             * @memberof google.protobuf.Timestamp
             * @static
             * @param {google.protobuf.Timestamp} message Timestamp
             * @param {$protobuf.IConversionOptions} [options] Conversion options
             * @returns {Object.<string,*>} Plain object
             */
            Timestamp.toObject = function toObject(message, options) {
                if (!options)
                    options = {};
                var object = {};
                if (options.defaults) {
                    if ($util.Long) {
                        var long = new $util.Long(0, 0, false);
                        object.seconds = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                    } else
                        object.seconds = options.longs === String ? "0" : 0;
                    object.nanos = 0;
                }
                if (message.seconds != null && message.hasOwnProperty("seconds"))
                    if (typeof message.seconds === "number")
                        object.seconds = options.longs === String ? String(message.seconds) : message.seconds;
                    else
                        object.seconds = options.longs === String ? $util.Long.prototype.toString.call(message.seconds) : options.longs === Number ? new $util.LongBits(message.seconds.low >>> 0, message.seconds.high >>> 0).toNumber() : message.seconds;
                if (message.nanos != null && message.hasOwnProperty("nanos"))
                    object.nanos = message.nanos;
                return object;
            };

            /**
             * Converts this Timestamp to JSON.
             * @function toJSON
             * @memberof google.protobuf.Timestamp
             * @instance
             * @returns {Object.<string,*>} JSON object
             */
            Timestamp.prototype.toJSON = function toJSON() {
                return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
            };

            return Timestamp;
        })();

        return protobuf;
    })();

    return google;
})();

module.exports = $root;