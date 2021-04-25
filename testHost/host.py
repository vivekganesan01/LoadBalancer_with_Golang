#!/usr/bin/env python3

"""
Flask API server
Note: make sure to pass server-name and port
python3 testHost/host.py "server-1" "6000"  
"""

from flask import Flask
import sys
app = Flask(__name__)


server_name = sys.argv[1]
port = sys.argv[2]

@app.route("/")
def hello():
    return "Hello from {server_name}:{port} !".format(server_name=server_name, port=port)

if __name__ == '__main__':
    app.run(port=int(port))
