# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: hello_world.proto for package ''

require 'grpc'
require 'hello_world_pb'

module Greeter
  # HelloRequest を受け取って HelloResponse を返すメソッドの定義
  class Service

    include GRPC::GenericService

    self.marshal_class_method = :encode
    self.unmarshal_class_method = :decode
    self.service_name = 'Greeter'

    rpc :Hello, ::HelloRequest, ::HelloResponse
  end

  Stub = Service.rpc_stub_class
end