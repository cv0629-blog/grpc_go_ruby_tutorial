this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'pb')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'hello_world_services_pb'

def main
  stub = Greeter::Stub.new('localhost:8080', :this_channel_is_insecure)

  begin
    message = stub.hello(HelloRequest.new(name: 'Yuta Nakamura'))
    puts message.to_json

  rescue GRPC::BadStatus => e
    abort "ERROR: #{e.message}"
  end
end

main