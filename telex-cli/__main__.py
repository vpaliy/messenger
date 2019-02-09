#!/usr/bin/env python

import asyncio
import websockets
import argparse
import json


async def send(websocket, message, channel):
  await websocket.send(json.dumps({
    'event': 'send',
    'action': dict(channel=channel, content=message)
  }))


async def join(websocket, group):
  await websocket.send(json.dumps({
    'event': 'join',
    'action': dict(channel=group)
  }))


async def run(uri):
  async with websockets.connect(uri) as websocket:
    print(f'Connection has been established: {uri}')
    await join(websocket, 'Telex')
    while True:
      message = input('Enter your message\n')
      await send(websocket, message, 'Telex')
      raw = await websocket.recv()
      print(f'Received response:\n{raw}')


def main():
  parser = argparse.ArgumentParser()
  parser.add_argument(
    '--host',
    help='Address of the Telex server',
    default='localhost:8080'
  )
  parser.add_argument(
    '--endpoint',
    help='Endpoint to which we are connecting',
    default='ws'
  )
  parser.add_argument(
    '--token',
    help='JWT token used for authentication',
    default=None
  )

  args = parser.parse_args()
  uri = f'ws://{args.host}/{args.endpoint}'
  asyncio.get_event_loop().run_until_complete(run(uri))

if __name__ == '__main__':
  main()
