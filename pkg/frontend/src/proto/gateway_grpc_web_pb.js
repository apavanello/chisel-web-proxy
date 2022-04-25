/**
 * @fileoverview gRPC-Web generated client stub for gateway
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.gateway = require('./gateway_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.gateway.ProxyServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.gateway.ProxyServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.gateway.ConnectRequest,
 *   !proto.gateway.ConnectResponse>}
 */
const methodDescriptor_ProxyService_Connect = new grpc.web.MethodDescriptor(
  '/gateway.ProxyService/Connect',
  grpc.web.MethodType.UNARY,
  proto.gateway.ConnectRequest,
  proto.gateway.ConnectResponse,
  /**
   * @param {!proto.gateway.ConnectRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.gateway.ConnectResponse.deserializeBinary
);


/**
 * @param {!proto.gateway.ConnectRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.gateway.ConnectResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gateway.ConnectResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gateway.ProxyServiceClient.prototype.connect =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gateway.ProxyService/Connect',
      request,
      metadata || {},
      methodDescriptor_ProxyService_Connect,
      callback);
};


/**
 * @param {!proto.gateway.ConnectRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gateway.ConnectResponse>}
 *     Promise that resolves to the response
 */
proto.gateway.ProxyServicePromiseClient.prototype.connect =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/gateway.ProxyService/Connect',
      request,
      metadata || {},
      methodDescriptor_ProxyService_Connect);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.gateway.DisconnectRequest,
 *   !proto.gateway.DisconnectResponse>}
 */
const methodDescriptor_ProxyService_Disconnect = new grpc.web.MethodDescriptor(
  '/gateway.ProxyService/Disconnect',
  grpc.web.MethodType.UNARY,
  proto.gateway.DisconnectRequest,
  proto.gateway.DisconnectResponse,
  /**
   * @param {!proto.gateway.DisconnectRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.gateway.DisconnectResponse.deserializeBinary
);


/**
 * @param {!proto.gateway.DisconnectRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.gateway.DisconnectResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gateway.DisconnectResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gateway.ProxyServiceClient.prototype.disconnect =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gateway.ProxyService/Disconnect',
      request,
      metadata || {},
      methodDescriptor_ProxyService_Disconnect,
      callback);
};


/**
 * @param {!proto.gateway.DisconnectRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gateway.DisconnectResponse>}
 *     Promise that resolves to the response
 */
proto.gateway.ProxyServicePromiseClient.prototype.disconnect =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/gateway.ProxyService/Disconnect',
      request,
      metadata || {},
      methodDescriptor_ProxyService_Disconnect);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.gateway.StatusRequest,
 *   !proto.gateway.StatusResponse>}
 */
const methodDescriptor_ProxyService_Status = new grpc.web.MethodDescriptor(
  '/gateway.ProxyService/Status',
  grpc.web.MethodType.UNARY,
  proto.gateway.StatusRequest,
  proto.gateway.StatusResponse,
  /**
   * @param {!proto.gateway.StatusRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.gateway.StatusResponse.deserializeBinary
);


/**
 * @param {!proto.gateway.StatusRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.gateway.StatusResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gateway.StatusResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gateway.ProxyServiceClient.prototype.status =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gateway.ProxyService/Status',
      request,
      metadata || {},
      methodDescriptor_ProxyService_Status,
      callback);
};


/**
 * @param {!proto.gateway.StatusRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gateway.StatusResponse>}
 *     Promise that resolves to the response
 */
proto.gateway.ProxyServicePromiseClient.prototype.status =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/gateway.ProxyService/Status',
      request,
      metadata || {},
      methodDescriptor_ProxyService_Status);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.gateway.PreLoadServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.gateway.PreLoadServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.gateway.PreLoadRequest,
 *   !proto.gateway.PreLoadResponse>}
 */
const methodDescriptor_PreLoadService_PreLoad = new grpc.web.MethodDescriptor(
  '/gateway.PreLoadService/PreLoad',
  grpc.web.MethodType.UNARY,
  proto.gateway.PreLoadRequest,
  proto.gateway.PreLoadResponse,
  /**
   * @param {!proto.gateway.PreLoadRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.gateway.PreLoadResponse.deserializeBinary
);


/**
 * @param {!proto.gateway.PreLoadRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.gateway.PreLoadResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gateway.PreLoadResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gateway.PreLoadServiceClient.prototype.preLoad =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gateway.PreLoadService/PreLoad',
      request,
      metadata || {},
      methodDescriptor_PreLoadService_PreLoad,
      callback);
};


/**
 * @param {!proto.gateway.PreLoadRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gateway.PreLoadResponse>}
 *     Promise that resolves to the response
 */
proto.gateway.PreLoadServicePromiseClient.prototype.preLoad =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/gateway.PreLoadService/PreLoad',
      request,
      metadata || {},
      methodDescriptor_PreLoadService_PreLoad);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.gateway.HostsRequest,
 *   !proto.gateway.HostsResponse>}
 */
const methodDescriptor_PreLoadService_GetHosts = new grpc.web.MethodDescriptor(
  '/gateway.PreLoadService/GetHosts',
  grpc.web.MethodType.UNARY,
  proto.gateway.HostsRequest,
  proto.gateway.HostsResponse,
  /**
   * @param {!proto.gateway.HostsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.gateway.HostsResponse.deserializeBinary
);


/**
 * @param {!proto.gateway.HostsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.gateway.HostsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.gateway.HostsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.gateway.PreLoadServiceClient.prototype.getHosts =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/gateway.PreLoadService/GetHosts',
      request,
      metadata || {},
      methodDescriptor_PreLoadService_GetHosts,
      callback);
};


/**
 * @param {!proto.gateway.HostsRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.gateway.HostsResponse>}
 *     Promise that resolves to the response
 */
proto.gateway.PreLoadServicePromiseClient.prototype.getHosts =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/gateway.PreLoadService/GetHosts',
      request,
      metadata || {},
      methodDescriptor_PreLoadService_GetHosts);
};


module.exports = proto.gateway;

