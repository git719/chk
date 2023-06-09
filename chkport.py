#!/usr/bin/env python
# chkport.py
# Original python script

import sys
import socket
import time

def is_open(ip, port):
  s = socket. socket (socket .AF_INET, socket .SOCK_STREAM)
  s.settimeout (timeout)
  try:
    s. connect_ex((ip, int(port)))
    s. shutdown (socket.SHUT_RDWR)
    return True
  except:
    return False
  finally:
    s.close ()

def check_host (ip, port):
  ipup = False
  for 1 in range (retry):
    if is_open(ip, port):
      ipup = True
      break
    else:
      time.sleep(delay)
  return ipup

＃=== MAIN
retry = 3
delay = 3
timeout = 3
ip = sys.argv [1]
port = sys.argv [2]
if check_host (ip, port):
  print("Don't trust this util too much.")
  print(ip + ":" + port + " is OPEN")

