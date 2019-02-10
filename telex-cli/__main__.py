#!/usr/bin/env python

import asyncio
import websockets
import argparse
import json
import time

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


async def receiver(websocket):
  while True:
    message = await websocket.recv()
    print(f'Received:\n{message}')


async def spinner(websocket):
  await join(websocket, 'Telex')
  while True:
    time.sleep(5)
    message = input('Enter your message\n')
    await send(websocket, message, 'Telex')
    message = await websocket.recv()
    print(f'Received:\n{message}')


async def run(uri):
  async with websockets.connect(uri) as websocket:
    print(f'Connection has been established: {uri}')
    spinner_task = asyncio.ensure_future(
            spinner(websocket))
    receiver_task = asyncio.ensure_future(
        receiver(websocket))
    done, pending = await asyncio.wait(
        [receiver_task, spinner_task],
        return_when=asyncio.ALL_COMPLETED,
    )
    for task in pending:
        task.cancel()


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
