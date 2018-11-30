
# TuringML APIs
This is the RESTful APIs of TuringML

Table of Contents

1. [Get all the nodes based on the playground ID](#playgrounds)
1. [Get all the playgrounds based on the user ID](#users)

<a name="playgrounds"></a>

## playgrounds

| Specification | Value |
|-----|-----|
| Resource Path | /playgrounds |
| API Version | 0.0.1 |
| BasePath for the API | https://api.turingml.org/ |
| Consumes | application/json |
| Produces | application/json |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /playgrounds/\{id\}/nodes | [GET](#GetNodes) | Get all the nodes based on the playground ID |
| /playgrounds/\{id\}/nodes/\{nodeId\} | [GET](#GetNode) | Get a single node given the ID |
| /playgrounds/\{id\}/nodes/\{nodeId\} | [POST](#CreateNode) | Create a single node in the playground |
| /playgrounds/\{id\}/nodes/\{nodeId\} | [PUT](#UpdateNode) | Updates a single node in the playground given the ID |
| /playgrounds/\{id\}/nodes/\{nodeId\} | [PUT](#DeleteNode) | Deletes a single node in the playground given the ID |



<a name="GetNodes"></a>

#### API: /playgrounds/\{id\}/nodes (GET)


Get all the nodes based on the playground ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Node](#github.com.turing-ml.turing-api.api.models.Node) |  |
| 500 | string | string | Internal Server Error |


<a name="GetNode"></a>

#### API: /playgrounds/\{id\}/nodes/\{nodeId\} (GET)


Get a single node given the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| nodeId | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Node](#github.com.turing-ml.turing-api.api.models.Node) |  |
| 500 | string | string | Internal Server Error |


<a name="CreateNode"></a>

#### API: /playgrounds/\{id\}/nodes/\{nodeId\} (POST)


Create a single node in the playground



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Node](#github.com.turing-ml.turing-api.api.models.Node) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdateNode"></a>

#### API: /playgrounds/\{id\}/nodes/\{nodeId\} (PUT)


Updates a single node in the playground given the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| nodeId | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | node updated |
| 500 | string | string | Internal Server Error |


<a name="DeleteNode"></a>

#### API: /playgrounds/\{id\}/nodes/\{nodeId\} (PUT)


Deletes a single node in the playground given the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| nodeId | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | node deleted |
| 500 | string | string | Internal Server Error |




### Models

<a name="github.com.turing-ml.turing-api.api.models.Node"></a>

#### Node

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| UpdatedAt | Time |  |
| active | bool |  |
| configId | int |  |
| name | string |  |
| playgroundId | github.com.turing-ml.turing-api.api.models.Playground |  |
| type | string |  |
| x | float64 |  |
| y | float64 |  |

<a name="github.com.turing-ml.turing-api.api.models.Playground"></a>

#### Playground

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| UpdatedAt | Time |  |
| name | string |  |
| userId | string |  |


<a name="users"></a>

## users

| Specification | Value |
|-----|-----|
| Resource Path | /users |
| API Version | 0.0.1 |
| BasePath for the API | https://api.turingml.org/ |
| Consumes | application/json |
| Produces | application/json |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /users/\{id\}/playgrounds | [GET](#GetPlaygrounds) | Get all the playgrounds based on the user ID |
| /users/\{user_id\}/playgrounds/\{playground_id\} | [GET](#GetPlayground) | Get a specific playground based on the ID in input of a single user |
| /users/\{user_id\}/playgrounds/\{playground_id\} | [PUT](#UpdatePlayground) | Update the playground of a specific user |
| /users/\{user_id\}/playgrounds/\{playground_id\} | [DELETE](#DeletePlayground) | Delete the playground of a specific user |
| /users/\{user_id\}/playgrounds | [POST](#CreatePlayground) | Create a new playground based on the parameters in input |



<a name="GetPlaygrounds"></a>

#### API: /users/\{id\}/playgrounds (GET)


Get all the playgrounds based on the user ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The user ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Playground](#github.com.turing-ml.turing-api.api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="GetPlayground"></a>

#### API: /users/\{user_id\}/playgrounds/\{playground_id\} (GET)


Get a specific playground based on the ID in input of a single user



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| playground_id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Playground](#github.com.turing-ml.turing-api.api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdatePlayground"></a>

#### API: /users/\{user_id\}/playgrounds/\{playground_id\} (PUT)


Update the playground of a specific user



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| playground_id | path | string | The playground ID | Yes |
| name | query | string | The playground name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | playground updated |
| 500 | string | string | Internal Server Error |


<a name="DeletePlayground"></a>

#### API: /users/\{user_id\}/playgrounds/\{playground_id\} (DELETE)


Delete the playground of a specific user



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| playground_id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | playground deleted |
| 500 | string | string | Internal Server Error |


<a name="CreatePlayground"></a>

#### API: /users/\{user_id\}/playgrounds (POST)


Create a new playground based on the parameters in input



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| name | query | string | The playground name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Playground](#github.com.turing-ml.turing-api.api.models.Playground) |  |
| 500 | string | string | Internal Server Error |




### Models

<a name="github.com.turing-ml.turing-api.api.models.Playground"></a>

#### Playground

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| UpdatedAt | Time |  |
| name | string |  |
| userId | string |  |


