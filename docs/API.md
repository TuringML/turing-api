
# TuringML APIs
This is the RESTful APIs of TuringML

Table of Contents

1. [Get all the fields of a node](#playgrounds)
1. [Update the playground of a specific user](#users)

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
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields | [GET](#GetFields) | Get all the fields of a node |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields | [POST](#CreateField) | Create a single field for a node |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields/\{field_id\} | [GET](#GetField) | Get a single field given the ID |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields/\{field_id\} | [PUT](#UpdateField) | Updates a single field of a node |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields/\{field_id\} | [DELETE](#DeleteField) | Deletes a single field |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/links | [GET](#GetLinks) | Get all the links from/to a node |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/links | [POST](#CreateLink) | Create a single link from a node |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/links/\{link_id\} | [GET](#GetLink) | Get a single link given the ID |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/links/\{link_id\} | [PUT](#UpdateLink) | Updates a single link |
| /playgrounds/\{playground_id\}/nodes/\{node_id\}/links/\{link_id\} | [DELETE](#DeleteLink) | Deletes a single link |
| /playgrounds | [GET](#GetPlaygrounds) | Get all the playgrounds based on the user ID |
| /playgrounds | [POST](#CreatePlayground) | Create a new playground based on the parameters in input |
| /playgrounds/\{playground_id\} | [GET](#GetPlayground) | Get a specific playground based on the ID in input of a single user |
| /playgrounds/\{playground_id\}/nodes | [GET](#GetNodes) | Get all the nodes based on the playground ID |
| /playgrounds/\{playground_id\}/nodes | [POST](#CreateNode) | Create a single node in the playground |
| /playgrounds/\{playground_id\}/nodes/\{node_id\} | [GET](#GetNode) | Get a single node given the ID |
| /playgrounds/\{playground_id\}/nodes/\{node_id\} | [PUT](#UpdateNode) | Updates a single node in the playground given the ID |
| /playgrounds/\{playground_id\}/nodes/\{node_id\} | [PUT](#DeleteNode) | Deletes a single node in the playground given the ID |



<a name="GetFields"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields (GET)


Get all the fields of a node



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Link](#github.com.turing-ml.turing-api.api.models.Link) |  |
| 500 | string | string | Internal Server Error |


<a name="CreateField"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields (POST)


Create a single field for a node



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Field](#github.com.turing-ml.turing-api.api.models.Field) |  |
| 500 | string | string | Internal Server Error |


<a name="GetField"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields/\{field_id\} (GET)


Get a single field given the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |
| field_id | path | string | The field ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Field](#github.com.turing-ml.turing-api.api.models.Field) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdateField"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields/\{field_id\} (PUT)


Updates a single field of a node



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |
| field_id | path | string | The field ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | field updated |
| 500 | string | string | Internal Server Error |


<a name="DeleteField"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/fields/\{field_id\} (DELETE)


Deletes a single field



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |
| field_id | path | string | The field ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | link deleted |
| 500 | string | string | Internal Server Error |


<a name="GetLinks"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/links (GET)


Get all the links from/to a node



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Link](#github.com.turing-ml.turing-api.api.models.Link) |  |
| 500 | string | string | Internal Server Error |


<a name="CreateLink"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/links (POST)


Create a single link from a node



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Link](#github.com.turing-ml.turing-api.api.models.Link) |  |
| 500 | string | string | Internal Server Error |


<a name="GetLink"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/links/\{link_id\} (GET)


Get a single link given the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |
| link_id | path | string | The link ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Link](#github.com.turing-ml.turing-api.api.models.Link) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdateLink"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/links/\{link_id\} (PUT)


Updates a single link



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |
| link_id | path | string | The link ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | link updated |
| 500 | string | string | Internal Server Error |


<a name="DeleteLink"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\}/links/\{link_id\} (DELETE)


Deletes a single link



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| node_id | path | string | The node ID | Yes |
| link_id | path | string | The link ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | link deleted |
| 500 | string | string | Internal Server Error |


<a name="GetPlaygrounds"></a>

#### API: /playgrounds (GET)


Get all the playgrounds based on the user ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The user ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Playground](#github.com.turing-ml.turing-api.api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="CreatePlayground"></a>

#### API: /playgrounds (POST)


Create a new playground based on the parameters in input



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| name | query | string | The playground name | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Playground](#github.com.turing-ml.turing-api.api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="GetPlayground"></a>

#### API: /playgrounds/\{playground_id\} (GET)


Get a specific playground based on the ID in input of a single user



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| playground_id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Playground](#github.com.turing-ml.turing-api.api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="GetNodes"></a>

#### API: /playgrounds/\{playground_id\}/nodes (GET)


Get all the nodes based on the playground ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Node](#github.com.turing-ml.turing-api.api.models.Node) |  |
| 500 | string | string | Internal Server Error |


<a name="CreateNode"></a>

#### API: /playgrounds/\{playground_id\}/nodes (POST)


Create a single node in the playground



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Node](#github.com.turing-ml.turing-api.api.models.Node) |  |
| 500 | string | string | Internal Server Error |


<a name="GetNode"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\} (GET)


Get a single node given the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The playground ID | Yes |
| nodeId | path | string | The node ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Node](#github.com.turing-ml.turing-api.api.models.Node) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdateNode"></a>

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\} (PUT)


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

#### API: /playgrounds/\{playground_id\}/nodes/\{node_id\} (PUT)


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

<a name="github.com.turing-ml.turing-api.api.models.Class"></a>

#### Class

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|

<a name="github.com.turing-ml.turing-api.api.models.Configuration"></a>

#### Configuration

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| UpdatedAt | Time |  |
| blob | github.com.turing-ml.turing-api.api.models.JSON |  |

<a name="github.com.turing-ml.turing-api.api.models.Field"></a>

#### Field

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| Node | github.com.turing-ml.turing-api.api.models.Node |  |
| UpdatedAt | Time |  |
| active | bool |  |
| key_id | int |  |
| key_name | string |  |
| key_primary | bool |  |
| nodeId | uint |  |
| value_class | github.com.turing-ml.turing-api.api.models.Class |  |
| value_delimiter | string |  |
| value_example | string |  |
| value_kind | string |  |
| value_type | string |  |

<a name="github.com.turing-ml.turing-api.api.models.JSON"></a>

#### JSON

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|

<a name="github.com.turing-ml.turing-api.api.models.Link"></a>

#### Link

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| FromField | github.com.turing-ml.turing-api.api.models.Field |  |
| FromNode | github.com.turing-ml.turing-api.api.models.Node |  |
| ID | uint |  |
| ToField | github.com.turing-ml.turing-api.api.models.Field |  |
| ToNode | github.com.turing-ml.turing-api.api.models.Node |  |
| UpdatedAt | Time |  |
| from_field_id | uint |  |
| from_node_id | uint |  |
| to_field_id | uint |  |
| to_node_id | uint |  |

<a name="github.com.turing-ml.turing-api.api.models.Node"></a>

#### Node

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| Configuration | github.com.turing-ml.turing-api.api.models.Configuration |  |
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| Playground | github.com.turing-ml.turing-api.api.models.Playground |  |
| UpdatedAt | Time |  |
| active | bool |  |
| configuration_id | uint |  |
| name | string |  |
| playground_id | uint |  |
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
| User | github.com.turing-ml.turing-api.api.models.User |  |
| name | string |  |
| user_id | string |  |

<a name="github.com.turing-ml.turing-api.api.models.User"></a>

#### User

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| CreatedAt | Time |  |
| DeletedAt | Time |  |
| ID | uint |  |
| UpdatedAt | Time |  |
| name | string |  |


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
| /users/\{user_id\}/playgrounds/\{playground_id\} | [PUT](#UpdatePlayground) | Update the playground of a specific user |
| /users/\{user_id\}/playgrounds/\{playground_id\} | [DELETE](#DeletePlayground) | Delete the playground of a specific user |



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


