# Simply todo in GO

Une API RESTFUL simple vous permettant d'organiser vos projets personnels.

<p align="center">
  <img 
        alt="GO" 
        width="200" 
        height="200"
        src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg"
        style="margin-left: auto; margin-right: auto"
    />
</p>

### Installation:

To start the project :

```bash
git clonne https://github.com/Dramane-dev/todolist-api-go
cd todolist-api
go run .
```

### User Endpoints

```ruby
GET     http://localhost:8080/api/users
GET     http://localhost:8080/api/user/:userId
POST    http://localhost:8080/api/user/signup
POST    http://localhost:8080/api/user/signin
PATCH   http://localhost:8080/api/user/:userId
DELETE  http://localhost:8080/api/user/:userId
```

### Project Endpoints

```ruby
GET     http://localhost:8080/api/projects
GET     http://localhost:8080/api/projects/:userId
GET     http://localhost:8080/api/projects/:projectId
POST    http://localhost:8080/api/project/:userId
PATCH   http://localhost:8080/api/project/:projectId
DELETE  http://localhost:8080/api/project/:projectId
```

### Task Endpoints

```ruby
GET     http://localhost:8080/api/tasks
GET     http://localhost:8080/api/tasks/:projectId
GET     http://localhost:8080/api/tasks/:taskId
POST    http://localhost:8080/api/task/:projectId
PATCH   http://localhost:8080/api/task/:taskId
DELETE  http://localhost:8080/api/task/:taskId
```

### Attachment Endpoints

```ruby
GET     http://localhost:8080/api/attachments
GET     http://localhost:8080/api/attachments/:projectId
GET     http://localhost:8080/api/attachment/:attachmentId
POST    http://localhost:8080/api/attachment/:projectId
DELETE  http://localhost:8080/api/attachment/:attachmentId
```

### Payment Endpoints

```ruby
POST    http://localhost:8080/api/subscribe/:userId
DELETE  http://localhost:8080/api/unsubscribe/:subscriptionId
```


### Datas structure

Users inside this API represented by :

Fields |Types
-------|-----
**userId**| string
**lastname**| string
**firstname**| string
**email**| string
**mailConfirmed**| boolean
**Projects**| []Project

Each users have a Projects array inside this API are represented by :

Fields |Types
-------|-----
**projectId**| string
**projectName**| string
**projectDescription**| string
**userId**| string
**mailConfirmed**| boolean
**Tasks**| []Project
**Attachments**| []Attachments
**Tasks**| []Tasks

Each users have a Tasks array inside this API are represented by :

Fields |Types
-------|-----
**taskId**| string
**name**| string
**description**| string
**status**| string
**createdAt**| string
**projectId**| []string

Each users have a Attachments array inside this API are represented by :

Fields |Types
-------|-----
**attachmentId**| string
**fileName**| string
**fileType**| string
**filePath**| string
**projectId**| string

Each users have a Subscription array inside this API are represented by :

Fields |Types
-------|-----
**subscriptionId**| string
**name**| string
**description**| string
**amount**| int64
**userId**| string

### JSON Datas example

```javascript
{
    "users": [
        {
            "userId": "08b0579c-fbfa-4e9b-92c6-dd7812c3e795",
            "lastname": "KAMISS0K0",
            "firstname": "dramane",
            "email": "dramane@gmail.com",
            "mailConfirmed": false,
            "Projects": [
                {
                    "projectId": "5e66ecca-b24e-4a88-9d52-df00f0d2684d",
                    "projectName": "RESTFUL API in GO",
                    "projectDescription": "Créer une api restful en GO",
                    "userId": "08b0579c-fbfa-4e9b-92c6-dd7812c3e795",
                    "Tasks": [
                        {
                          "taskId": "43f7cb4b-cc23-482b-aaca-2cc13178f481",
                          "name": "Upload Image",
                          "description": "Ajouter une fonctionnalité permettant d'upload des fichiers.",
                          "status": "todo",
                          "createdAt": "2022-05-23 17:15:21",
                          "projectId": "5e66ecca-b24e-4a88-9d52-df00f0d2684d"
                        },
                        {
                          "taskId": "96324bcb-cd2b-4c84-b1db-d07b7c908c29",
                          "name": "Send image by email",
                          "description": "Ajouter une fonctionnalité permettant d'envoyer une image par mail.",
                          "status": "todo",
                          "createdAt": "2022-05-23 17:27:42",
                          "projectId": "5e66ecca-b24e-4a88-9d52-df00f0d2684d"
                        }
                    ],
                    "Attachments": [
                        {
                          "attachmentId": "ATC095e6e35-a384-40ad-bd8c-1cf2d073974b",
                          "fileName": "pikachu.png",
                          "fileType": "image/png",
                          "filePath": "./uploads/pikachu.png",
                          "projectId": "5e66ecca-b24e-4a88-9d52-df00f0d2684d"
                        },
                        {
                          "attachmentId": "ATC210b1ec5-357f-482f-bc18-c1b2383decfd",
                          "fileName": "pikachu.png",
                          "fileType": "image/png",
                          "filePath": "./uploads/pikachu.png",
                          "projectId": "5e66ecca-b24e-4a88-9d52-df00f0d2684d"
                        },
                        {
                          "attachmentId": "ATC5a5e75bf-6003-4512-8912-45436df82e82",
                          "fileName": "pikachu.png",
                          "fileType": "image/png",
                          "filePath": "./uploads/pikachu.png",
                          "projectId": "5e66ecca-b24e-4a88-9d52-df00f0d2684d"
                        }
                    ]
                },
            ],
            "Subscription": {
              "subscriptionId": "SUBecf2a0bb-c2d7-41f3-9bf4-2fe50d21d444",
              "name": "Abonnement Premium",
              "description": "Un abonnement vous donnant accès à l'intégralité des fonctionnalité présente sur Simply Todo !",
              "amount": 20000,
              "userId": "08b0579c-fbfa-4e9b-92c6-dd7812c3e795"
            }
        }
    ]
}
```
