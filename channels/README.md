
### GET `channels.info`
**Request**

`token` - JWT token sent in the authorization header (required).

`channel` - name of the channel to which to subscribe (required).

Sample:

```json
{
"channel": "Telex"
}
```

**Request**:
```json
{
"channel": "channel-id",
"name": "channel-name",
"created_at": 1550194115,
"updated_at": 1550194115,
"private": true,
"description": "Channel's description",
"tags": ["A", "B", "C", "D"],
"image": "https://channel-avatar.jpeg",
"archived": false,
"creator": {
  "id": "creator-id",
  "username": "creator-username",
  "email": "email@example.com",
  "fullname": "creator-fullname",
  "image": "https://creator-avatar.jpeg"
},
"members": [
  {
    "id": "participant-id",
    "username": "participant-username",
    "email": "email@example.com",
    "fullname": "participant-fullname",
    "image": "https://participant-avatar.jpeg"
  }
]
}
```

### GET `channels.list`
**Request**

`token` - JWT token sent in the authorization header (required).

**Response**:
```json
{
"user": {
  "id": "participant-id",
  "username": "participant-username",
  "email": "email@example.com",
  "fullname": "participant-fullname",
  "image": "https://participant-avatar.jpeg"
},
"subscriptions": [{
  "id": "subscription-id",
  "channel": "channel-id",
  "created_at": 1550194115,
  "updated_at": 1550194115,
  "private": true,
  "snippet": "First message in that channel",
  "unread": 10
},{
  "id": "subscription-id",
  "channel": "channel-id",
  "created_at": 1550194115,
  "updated_at": 1550194115,
  "private": true,
  "snippet": "First message in that channel",
  "unread": 10
}]
}
```

### POST `channels.create`
**Request**

`token` - JWT token sent in the authorization header (required).

`channel` - name of the channel (required).

`tags` - a list of tags. Default is an empty array.

`image` - an image URL that will server as channel's picture. Default is `null`.

`description` - a description of the channel. Default is `null`.

`type` - type of the channel (dialog, group, general). Default is `group`.

`private` - whether the channel should be private. Default is `false`.

Sample:

```json
{
"channel": "Telex",
"tags": ["chat", "open-source"],
"image": "https://telex-avatar.jpeg",
"description": "An open-source chat server",
"type": "group",
"private": false
}
```

**Response**:
```json
{
"channel": 1,
"name": "Telex",
"created_at": 1550194115,
"updated_at": 1550194115,
"private": false,
"description": "An open-source chat server",
"tags": ["chat", "open-source"],
"image": "https://telex-avatar.jpeg",
"archived": false,
"creator": {
  "id": "creator-id",
  "username": "creator-username",
  "email": "email@example.com",
  "fullname": "creator-fullname",
  "image": "https://creator-avatar.jpeg"
},
"members":[]
}
```

### POST `channels.join`
**Request**

`token` - JWT token sent in the authorization header (required).

`channel` - name of the channel to which to subscribe (required).

Sample:

```json
{
  "channel": "Telex"
}
```

**Response**:
```json
{
  "id": "subscription-id",
  "channel": "channel-id",
  "created_at": 1550194115,
  "updated_at": 1550194115,
  "private": true,
  "snippet": "Welcome to Telex",
  "unread": 1
}
```

### POST `channels.update`
**Request**

`token` - JWT token sent in the authorization header (required).

`channel` - name of the channel (required).

`tags` - a list of tags. Default is an empty array.

`image` - an image URL that will server as channel's picture. Default is `null`.

`description` - a description of the channel. Default is `null`.

`private` - whether the channel should be private. Default is `false`.

Sample:

```json
{
  "channel": "Telex",
  "tags": ["chat", "open-source"],
  "image": "https://telex-avatar.jpeg",
  "description": "An open-source chat server",
  "private": false
}
```

**Response**:
```json
{
  "channel": 1,
  "name": "Telex",
  "created_at": 1550194115,
    "updated_at": 1550194115,
    "private": false,
    "description": "An open-source chat server",
    "tags": ["chat", "open-source"],
    "image": "https://telex-avatar.jpeg",
    "archived": false,
    "creator": {
      "id": "creator-id",
      "username": "creator-username",
      "email": "email@example.com",
      "fullname": "creator-fullname",
      "image": "https://creator-avatar.jpeg"
    },
    "members":[]
  }
  ```
