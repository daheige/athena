// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var hello_pb = require('./hello_pb.js');
// var google_api_annotations_pb = require('./google/api/annotations_pb.js');

function serialize_App_Grpc_Hello_BatchUsersReply(arg) {
  if (!(arg instanceof hello_pb.BatchUsersReply)) {
    throw new Error('Expected argument of type App.Grpc.Hello.BatchUsersReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_BatchUsersReply(buffer_arg) {
  return hello_pb.BatchUsersReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_BatchUsersReq(arg) {
  if (!(arg instanceof hello_pb.BatchUsersReq)) {
    throw new Error('Expected argument of type App.Grpc.Hello.BatchUsersReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_BatchUsersReq(buffer_arg) {
  return hello_pb.BatchUsersReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_HelloReply(arg) {
  if (!(arg instanceof hello_pb.HelloReply)) {
    throw new Error('Expected argument of type App.Grpc.Hello.HelloReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_HelloReply(buffer_arg) {
  return hello_pb.HelloReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_HelloReq(arg) {
  if (!(arg instanceof hello_pb.HelloReq)) {
    throw new Error('Expected argument of type App.Grpc.Hello.HelloReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_HelloReq(buffer_arg) {
  return hello_pb.HelloReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_InfoReply(arg) {
  if (!(arg instanceof hello_pb.InfoReply)) {
    throw new Error('Expected argument of type App.Grpc.Hello.InfoReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_InfoReply(buffer_arg) {
  return hello_pb.InfoReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_App_Grpc_Hello_InfoReq(arg) {
  if (!(arg instanceof hello_pb.InfoReq)) {
    throw new Error('Expected argument of type App.Grpc.Hello.InfoReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_App_Grpc_Hello_InfoReq(buffer_arg) {
  return hello_pb.InfoReq.deserializeBinary(new Uint8Array(buffer_arg));
}


// service 定义开放调用的服务
var GreeterServiceService = exports.GreeterServiceService = {
  sayHello: {
    path: '/App.Grpc.Hello.GreeterService/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.HelloReq,
    responseType: hello_pb.HelloReply,
    requestSerialize: serialize_App_Grpc_Hello_HelloReq,
    requestDeserialize: deserialize_App_Grpc_Hello_HelloReq,
    responseSerialize: serialize_App_Grpc_Hello_HelloReply,
    responseDeserialize: deserialize_App_Grpc_Hello_HelloReply,
  },
  info: {
    path: '/App.Grpc.Hello.GreeterService/Info',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.InfoReq,
    responseType: hello_pb.InfoReply,
    requestSerialize: serialize_App_Grpc_Hello_InfoReq,
    requestDeserialize: deserialize_App_Grpc_Hello_InfoReq,
    responseSerialize: serialize_App_Grpc_Hello_InfoReply,
    responseDeserialize: deserialize_App_Grpc_Hello_InfoReply,
  },
  batchUsers: {
    path: '/App.Grpc.Hello.GreeterService/BatchUsers',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.BatchUsersReq,
    responseType: hello_pb.BatchUsersReply,
    requestSerialize: serialize_App_Grpc_Hello_BatchUsersReq,
    requestDeserialize: deserialize_App_Grpc_Hello_BatchUsersReq,
    responseSerialize: serialize_App_Grpc_Hello_BatchUsersReply,
    responseDeserialize: deserialize_App_Grpc_Hello_BatchUsersReply,
  },
};

exports.GreeterServiceClient = grpc.makeGenericClientConstructor(GreeterServiceService);
