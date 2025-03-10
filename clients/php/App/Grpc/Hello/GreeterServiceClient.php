<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\Grpc\Hello;

/**
 * service 定义开放调用的服务
 */
class GreeterServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \App\Grpc\Hello\HelloReq $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function SayHello(\App\Grpc\Hello\HelloReq $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/App.Grpc.Hello.GreeterService/SayHello',
        $argument,
        ['\App\Grpc\Hello\HelloReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \App\Grpc\Hello\InfoReq $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Info(\App\Grpc\Hello\InfoReq $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/App.Grpc.Hello.GreeterService/Info',
        $argument,
        ['\App\Grpc\Hello\InfoReply', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \App\Grpc\Hello\BatchUsersReq $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function BatchUsers(\App\Grpc\Hello\BatchUsersReq $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/App.Grpc.Hello.GreeterService/BatchUsers',
        $argument,
        ['\App\Grpc\Hello\BatchUsersReply', 'decode'],
        $metadata, $options);
    }

}
