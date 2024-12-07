import logging
import time
import uuid
from concurrent import futures

import grpc

from echo.rpc.message import message_pb2
from echo.rpc.message import message_pb2_grpc
from echo.rpc.message import message_pb2_grpc

class MessageServicer(message_pb2_grpc.EchoMessageServicer):

    def __init__(self):
        pass

    def echoMessage(self, request, context):
        print("sendMessage:request", request)

        response = message_pb2.MessageResponse(title=request.title)

        print("sendMessage:response", response)
        return response

def server():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    message_pb2_grpc.add_EchoMessageServicer_to_server(MessageServicer(), server)
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    print(f"Starting server. Listening on port {port}.")
    server.wait_for_termination()

if __name__ == "__main__":
    logging.basicConfig()
    server()

