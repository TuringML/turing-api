
# TuringML APIs
This is the RESTful APIs of TuringML

Table of Contents

1. [Get all the playgrounds in the database](#playground)
1. [Create a new user](#user)
1. [Get all the playgrounds based on the user ID](#users)

<a name="playground"></a>

## playground

| Specification | Value |
|-----|-----|
| Resource Path | /playground |
| API Version | 0.0.1 |
| BasePath for the API | https://api.turingml.org/ |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /playgrounds | [GET](#GetAllPlaygrounds) | Get all the playgrounds in the database |



<a name="GetAllPlaygrounds"></a>

#### API: /playgrounds (GET)


Get all the playgrounds in the database



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Playground](#github.com.turing-ml.turing-api.models.Playground) |  |
| 500 | string | string | Internal Server Error |




### Models

<a name="github.com.turing-ml.turing-api.models.Playground"></a>

#### Playground

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| createdAt | Time |  |
| id | gopkg.in.mgo.v2.bson.ObjectId |  |
| pipelineSchema | string |  |
| updatedAt | Time |  |
| userId | string |  |

<a name="gopkg.in.mgo.v2.bson.ObjectId"></a>

#### ObjectId

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|


<a name="user"></a>

## user

| Specification | Value |
|-----|-----|
| Resource Path | /user |
| API Version | 0.0.1 |
| BasePath for the API | https://api.turingml.org/ |
| Consumes | application/json |
| Produces |  |



### Operations


| Resource Path | Operation | Description |
|-----|-----|-----|
| /users/ | [POST](#CreateUser) | Create a new user |
| /users/\{id\} | [PUT](#UpdateUser) | Update the user specified by the ID |
| /users/\{id\} | [DELETE](#DeleteUser) | Delete the user specified by the ID |



<a name="CreateUser"></a>

#### API: /users/ (POST)


Create a new user



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| name | query | string | Name of the user | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [User](#github.com.turing-ml.turing-api.models.User) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdateUser"></a>

#### API: /users/\{id\} (PUT)


Update the user specified by the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | ID of the user - the MongoDB Object ID | Yes |
| name | query | string | New name of the user | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | user updated |
| 500 | string | string | Internal Server Error |


<a name="DeleteUser"></a>

#### API: /users/\{id\} (DELETE)


Delete the user specified by the ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | ID of the user - the MongoDB Object ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | string | string | user deleted |
| 500 | string | string | Internal Server Error |




### Models

<a name="github.com.turing-ml.turing-api.models.User"></a>

#### User

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| createdAt | Time |  |
| id | gopkg.in.mgo.v2.bson.ObjectId |  |
| name | string |  |
| updatedAt | Time |  |

<a name="gopkg.in.mgo.v2.bson.ObjectId"></a>

#### ObjectId

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|


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
| /users | [GET](#GetUsers) | Get all the users |
| /users/\{id\} | [GET](#GetUser) | Get user object by ID |



<a name="GetPlaygrounds"></a>

#### API: /users/\{id\}/playgrounds (GET)


Get all the playgrounds based on the user ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | The user ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [Playground](#github.com.turing-ml.turing-api.models.Playground) |  |
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
| 200 | object | [Playground](#github.com.turing-ml.turing-api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="UpdatePlayground"></a>

#### API: /users/\{user_id\}/playgrounds/\{playground_id\} (PUT)


Update the playground of a specific user



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| user_id | path | string | The user ID | Yes |
| playground_id | path | string | The playground ID | Yes |
| pipeline_schema | query | string | The playground pipeline schema as JSON string to update | Yes |


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
| pipeline_schema | query | string | The playground pipeline schema as JSON string | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [Playground](#github.com.turing-ml.turing-api.models.Playground) |  |
| 500 | string | string | Internal Server Error |


<a name="GetUsers"></a>

#### API: /users (GET)


Get all the users



| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | array | [User](#github.com.turing-ml.turing-api.models.User) |  |


<a name="GetUser"></a>

#### API: /users/\{id\} (GET)


Get user object by ID



| Param Name | Param Type | Data Type | Description | Required? |
|-----|-----|-----|-----|-----|
| id | path | string | Some ID | Yes |


| Code | Type | Model | Message |
|-----|-----|-----|-----|
| 200 | object | [User](#github.com.turing-ml.turing-api.models.User) |  |
| 400 | string | string | ID is mandatory for this endpoint |
| 500 | string | string | Internal Server Error |




### Models

<a name="github.com.turing-ml.turing-api.models.Playground"></a>

#### Playground

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| createdAt | Time |  |
| id | gopkg.in.mgo.v2.bson.ObjectId |  |
| pipelineSchema | string |  |
| updatedAt | Time |  |
| userId | string |  |

<a name="github.com.turing-ml.turing-api.models.User"></a>

#### User

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|
| createdAt | Time |  |
| id | gopkg.in.mgo.v2.bson.ObjectId |  |
| name | string |  |
| updatedAt | Time |  |

<a name="gopkg.in.mgo.v2.bson.ObjectId"></a>

#### ObjectId

| Field Name (alphabetical) | Field Type | Description |
|-----|-----|-----|


