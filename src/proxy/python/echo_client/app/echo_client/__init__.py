import grpc

from echo_client.rpc.message import message_pb2
from echo_client.rpc.message import message_pb2_grpc


def main():
    # open a gRPC channel
    channel = grpc.insecure_channel('localhost:50051')
    # create a stub (client)
    stub = message_pb2_grpc.EchoMessageStub(channel)

    response = stub.echoMessage(message_pb2.MessageRequest(title="Hello?", content="Hello, Server!"))
    print("server", response)

