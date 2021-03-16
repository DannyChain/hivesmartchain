// source: tendermint/types/canonical.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var gogoproto_gogo_pb = require('../../gogoproto/gogo_pb.js');
goog.object.extend(proto, gogoproto_gogo_pb);
var tendermint_types_types_pb = require('../../tendermint/types/types_pb.js');
goog.object.extend(proto, tendermint_types_types_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.tendermint.types.CanonicalBlockID', null, global);
goog.exportSymbol('proto.tendermint.types.CanonicalPartSetHeader', null, global);
goog.exportSymbol('proto.tendermint.types.CanonicalProposal', null, global);
goog.exportSymbol('proto.tendermint.types.CanonicalVote', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tendermint.types.CanonicalBlockID = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.types.CanonicalBlockID, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.types.CanonicalBlockID.displayName = 'proto.tendermint.types.CanonicalBlockID';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tendermint.types.CanonicalPartSetHeader = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.types.CanonicalPartSetHeader, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.types.CanonicalPartSetHeader.displayName = 'proto.tendermint.types.CanonicalPartSetHeader';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tendermint.types.CanonicalProposal = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.types.CanonicalProposal, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.types.CanonicalProposal.displayName = 'proto.tendermint.types.CanonicalProposal';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tendermint.types.CanonicalVote = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tendermint.types.CanonicalVote, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tendermint.types.CanonicalVote.displayName = 'proto.tendermint.types.CanonicalVote';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tendermint.types.CanonicalBlockID.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.types.CanonicalBlockID.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.types.CanonicalBlockID} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalBlockID.toObject = function(includeInstance, msg) {
  var f, obj = {
    hash: msg.getHash_asB64(),
    partSetHeader: (f = msg.getPartSetHeader()) && proto.tendermint.types.CanonicalPartSetHeader.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tendermint.types.CanonicalBlockID}
 */
proto.tendermint.types.CanonicalBlockID.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.types.CanonicalBlockID;
  return proto.tendermint.types.CanonicalBlockID.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.types.CanonicalBlockID} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.types.CanonicalBlockID}
 */
proto.tendermint.types.CanonicalBlockID.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setHash(value);
      break;
    case 2:
      var value = new proto.tendermint.types.CanonicalPartSetHeader;
      reader.readMessage(value,proto.tendermint.types.CanonicalPartSetHeader.deserializeBinaryFromReader);
      msg.setPartSetHeader(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tendermint.types.CanonicalBlockID.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.types.CanonicalBlockID.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.types.CanonicalBlockID} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalBlockID.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHash_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getPartSetHeader();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.tendermint.types.CanonicalPartSetHeader.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes hash = 1;
 * @return {!(string|Uint8Array)}
 */
proto.tendermint.types.CanonicalBlockID.prototype.getHash = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes hash = 1;
 * This is a type-conversion wrapper around `getHash()`
 * @return {string}
 */
proto.tendermint.types.CanonicalBlockID.prototype.getHash_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getHash()));
};


/**
 * optional bytes hash = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getHash()`
 * @return {!Uint8Array}
 */
proto.tendermint.types.CanonicalBlockID.prototype.getHash_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getHash()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.tendermint.types.CanonicalBlockID} returns this
 */
proto.tendermint.types.CanonicalBlockID.prototype.setHash = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional CanonicalPartSetHeader part_set_header = 2;
 * @return {?proto.tendermint.types.CanonicalPartSetHeader}
 */
proto.tendermint.types.CanonicalBlockID.prototype.getPartSetHeader = function() {
  return /** @type{?proto.tendermint.types.CanonicalPartSetHeader} */ (
    jspb.Message.getWrapperField(this, proto.tendermint.types.CanonicalPartSetHeader, 2));
};


/**
 * @param {?proto.tendermint.types.CanonicalPartSetHeader|undefined} value
 * @return {!proto.tendermint.types.CanonicalBlockID} returns this
*/
proto.tendermint.types.CanonicalBlockID.prototype.setPartSetHeader = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.types.CanonicalBlockID} returns this
 */
proto.tendermint.types.CanonicalBlockID.prototype.clearPartSetHeader = function() {
  return this.setPartSetHeader(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.types.CanonicalBlockID.prototype.hasPartSetHeader = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.types.CanonicalPartSetHeader.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.types.CanonicalPartSetHeader} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalPartSetHeader.toObject = function(includeInstance, msg) {
  var f, obj = {
    total: jspb.Message.getFieldWithDefault(msg, 1, 0),
    hash: msg.getHash_asB64()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tendermint.types.CanonicalPartSetHeader}
 */
proto.tendermint.types.CanonicalPartSetHeader.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.types.CanonicalPartSetHeader;
  return proto.tendermint.types.CanonicalPartSetHeader.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.types.CanonicalPartSetHeader} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.types.CanonicalPartSetHeader}
 */
proto.tendermint.types.CanonicalPartSetHeader.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setTotal(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setHash(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.types.CanonicalPartSetHeader.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.types.CanonicalPartSetHeader} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalPartSetHeader.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTotal();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = message.getHash_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
};


/**
 * optional uint32 total = 1;
 * @return {number}
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.getTotal = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.tendermint.types.CanonicalPartSetHeader} returns this
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.setTotal = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional bytes hash = 2;
 * @return {!(string|Uint8Array)}
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.getHash = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes hash = 2;
 * This is a type-conversion wrapper around `getHash()`
 * @return {string}
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.getHash_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getHash()));
};


/**
 * optional bytes hash = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getHash()`
 * @return {!Uint8Array}
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.getHash_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getHash()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.tendermint.types.CanonicalPartSetHeader} returns this
 */
proto.tendermint.types.CanonicalPartSetHeader.prototype.setHash = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tendermint.types.CanonicalProposal.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.types.CanonicalProposal.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.types.CanonicalProposal} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalProposal.toObject = function(includeInstance, msg) {
  var f, obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    height: jspb.Message.getFieldWithDefault(msg, 2, 0),
    round: jspb.Message.getFieldWithDefault(msg, 3, 0),
    polRound: jspb.Message.getFieldWithDefault(msg, 4, 0),
    blockId: (f = msg.getBlockId()) && proto.tendermint.types.CanonicalBlockID.toObject(includeInstance, f),
    timestamp: (f = msg.getTimestamp()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    chainId: jspb.Message.getFieldWithDefault(msg, 7, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tendermint.types.CanonicalProposal}
 */
proto.tendermint.types.CanonicalProposal.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.types.CanonicalProposal;
  return proto.tendermint.types.CanonicalProposal.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.types.CanonicalProposal} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.types.CanonicalProposal}
 */
proto.tendermint.types.CanonicalProposal.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.tendermint.types.SignedMsgType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setHeight(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setRound(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPolRound(value);
      break;
    case 5:
      var value = new proto.tendermint.types.CanonicalBlockID;
      reader.readMessage(value,proto.tendermint.types.CanonicalBlockID.deserializeBinaryFromReader);
      msg.setBlockId(value);
      break;
    case 6:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTimestamp(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setChainId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tendermint.types.CanonicalProposal.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.types.CanonicalProposal.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.types.CanonicalProposal} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalProposal.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getHeight();
  if (f !== 0) {
    writer.writeSfixed64(
      2,
      f
    );
  }
  f = message.getRound();
  if (f !== 0) {
    writer.writeSfixed64(
      3,
      f
    );
  }
  f = message.getPolRound();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getBlockId();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.tendermint.types.CanonicalBlockID.serializeBinaryToWriter
    );
  }
  f = message.getTimestamp();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getChainId();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
};


/**
 * optional SignedMsgType type = 1;
 * @return {!proto.tendermint.types.SignedMsgType}
 */
proto.tendermint.types.CanonicalProposal.prototype.getType = function() {
  return /** @type {!proto.tendermint.types.SignedMsgType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.tendermint.types.SignedMsgType} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional sfixed64 height = 2;
 * @return {number}
 */
proto.tendermint.types.CanonicalProposal.prototype.getHeight = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.setHeight = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional sfixed64 round = 3;
 * @return {number}
 */
proto.tendermint.types.CanonicalProposal.prototype.getRound = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.setRound = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int64 pol_round = 4;
 * @return {number}
 */
proto.tendermint.types.CanonicalProposal.prototype.getPolRound = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.setPolRound = function(value) {
  return jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional CanonicalBlockID block_id = 5;
 * @return {?proto.tendermint.types.CanonicalBlockID}
 */
proto.tendermint.types.CanonicalProposal.prototype.getBlockId = function() {
  return /** @type{?proto.tendermint.types.CanonicalBlockID} */ (
    jspb.Message.getWrapperField(this, proto.tendermint.types.CanonicalBlockID, 5));
};


/**
 * @param {?proto.tendermint.types.CanonicalBlockID|undefined} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
*/
proto.tendermint.types.CanonicalProposal.prototype.setBlockId = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.clearBlockId = function() {
  return this.setBlockId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.types.CanonicalProposal.prototype.hasBlockId = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional google.protobuf.Timestamp timestamp = 6;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.tendermint.types.CanonicalProposal.prototype.getTimestamp = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 6));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
*/
proto.tendermint.types.CanonicalProposal.prototype.setTimestamp = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.clearTimestamp = function() {
  return this.setTimestamp(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.types.CanonicalProposal.prototype.hasTimestamp = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional string chain_id = 7;
 * @return {string}
 */
proto.tendermint.types.CanonicalProposal.prototype.getChainId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.tendermint.types.CanonicalProposal} returns this
 */
proto.tendermint.types.CanonicalProposal.prototype.setChainId = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tendermint.types.CanonicalVote.prototype.toObject = function(opt_includeInstance) {
  return proto.tendermint.types.CanonicalVote.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tendermint.types.CanonicalVote} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalVote.toObject = function(includeInstance, msg) {
  var f, obj = {
    type: jspb.Message.getFieldWithDefault(msg, 1, 0),
    height: jspb.Message.getFieldWithDefault(msg, 2, 0),
    round: jspb.Message.getFieldWithDefault(msg, 3, 0),
    blockId: (f = msg.getBlockId()) && proto.tendermint.types.CanonicalBlockID.toObject(includeInstance, f),
    timestamp: (f = msg.getTimestamp()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    chainId: jspb.Message.getFieldWithDefault(msg, 6, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tendermint.types.CanonicalVote}
 */
proto.tendermint.types.CanonicalVote.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tendermint.types.CanonicalVote;
  return proto.tendermint.types.CanonicalVote.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tendermint.types.CanonicalVote} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tendermint.types.CanonicalVote}
 */
proto.tendermint.types.CanonicalVote.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.tendermint.types.SignedMsgType} */ (reader.readEnum());
      msg.setType(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setHeight(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setRound(value);
      break;
    case 4:
      var value = new proto.tendermint.types.CanonicalBlockID;
      reader.readMessage(value,proto.tendermint.types.CanonicalBlockID.deserializeBinaryFromReader);
      msg.setBlockId(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTimestamp(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setChainId(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tendermint.types.CanonicalVote.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tendermint.types.CanonicalVote.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tendermint.types.CanonicalVote} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tendermint.types.CanonicalVote.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getType();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getHeight();
  if (f !== 0) {
    writer.writeSfixed64(
      2,
      f
    );
  }
  f = message.getRound();
  if (f !== 0) {
    writer.writeSfixed64(
      3,
      f
    );
  }
  f = message.getBlockId();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.tendermint.types.CanonicalBlockID.serializeBinaryToWriter
    );
  }
  f = message.getTimestamp();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getChainId();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
};


/**
 * optional SignedMsgType type = 1;
 * @return {!proto.tendermint.types.SignedMsgType}
 */
proto.tendermint.types.CanonicalVote.prototype.getType = function() {
  return /** @type {!proto.tendermint.types.SignedMsgType} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {!proto.tendermint.types.SignedMsgType} value
 * @return {!proto.tendermint.types.CanonicalVote} returns this
 */
proto.tendermint.types.CanonicalVote.prototype.setType = function(value) {
  return jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional sfixed64 height = 2;
 * @return {number}
 */
proto.tendermint.types.CanonicalVote.prototype.getHeight = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.tendermint.types.CanonicalVote} returns this
 */
proto.tendermint.types.CanonicalVote.prototype.setHeight = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional sfixed64 round = 3;
 * @return {number}
 */
proto.tendermint.types.CanonicalVote.prototype.getRound = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.tendermint.types.CanonicalVote} returns this
 */
proto.tendermint.types.CanonicalVote.prototype.setRound = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional CanonicalBlockID block_id = 4;
 * @return {?proto.tendermint.types.CanonicalBlockID}
 */
proto.tendermint.types.CanonicalVote.prototype.getBlockId = function() {
  return /** @type{?proto.tendermint.types.CanonicalBlockID} */ (
    jspb.Message.getWrapperField(this, proto.tendermint.types.CanonicalBlockID, 4));
};


/**
 * @param {?proto.tendermint.types.CanonicalBlockID|undefined} value
 * @return {!proto.tendermint.types.CanonicalVote} returns this
*/
proto.tendermint.types.CanonicalVote.prototype.setBlockId = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.types.CanonicalVote} returns this
 */
proto.tendermint.types.CanonicalVote.prototype.clearBlockId = function() {
  return this.setBlockId(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.types.CanonicalVote.prototype.hasBlockId = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.Timestamp timestamp = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.tendermint.types.CanonicalVote.prototype.getTimestamp = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.tendermint.types.CanonicalVote} returns this
*/
proto.tendermint.types.CanonicalVote.prototype.setTimestamp = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.tendermint.types.CanonicalVote} returns this
 */
proto.tendermint.types.CanonicalVote.prototype.clearTimestamp = function() {
  return this.setTimestamp(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.tendermint.types.CanonicalVote.prototype.hasTimestamp = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string chain_id = 6;
 * @return {string}
 */
proto.tendermint.types.CanonicalVote.prototype.getChainId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.tendermint.types.CanonicalVote} returns this
 */
proto.tendermint.types.CanonicalVote.prototype.setChainId = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


goog.object.extend(exports, proto.tendermint.types);
