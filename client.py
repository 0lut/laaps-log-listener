import socket
import pickle
import json

HOST = 'localhost'  # Standard loopback interface address (localhost).
PORT = 8081        # Port to listen on (non-privileged ports are > 1023)

with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))

    
    s.sendall(b'Hello, world\n')
    # s.sendall(toSend)
    data = s.recv(1024)
    print('Received', repr(data))
    s.close()
