# Project Introduction

Withered-horizon is delighted to present it's brand new low cost event management website intended to be used by educational institutions in the US to effectively register new events and RSVP invitees and safely handle invitee information. 

Our objectives in this project:

1) To generate a commercially viable low-cost alternative to similar to sites like Evenbrite
2) To offer appealing user-experience to the intended audience.
3) To adhere to software development best-practices by adopting an Agile framework in our work. 


# Development team:

Huangrui Chu (eager-boar) : Lead Developer

Xincong "Jerry" He (proud-salmon): Senior Developer

Lu Cao (puzzled-hornet): Junior Developer

Abhinav Gokari (alert-butterfly): Product Owner

Edgardo Rios (tender-pigeon): Scrum Master

## Duration: Oct 31st - Dec 11th

## Our Wiki

https://github.com/HuangruiChu/Haxel-Eventbrite/wiki

## Acknowledgement
Thanks to [Professor Kyle L. Jensen](https://github.com/kljensen) for providing Yale School of Management  MGT656 "Management of Software Development" in 2023 Fall. This project is based on the template Kyle provided.
## Building the project
Run `go build -mod vendor` to compile your project and then
run `./classproject` or `./classproject` on Windows
in order to run the app. Preview the app by visiting
[http://localhost:8080](http://localhost:8080) if you're
running it locally. Or, preview using the URL provided
GitHub CodeSpaces if you're running there, natch.

## What is here

| File                      | Role                                                                                                                      |
| ------------------------- | ------------------------------------------------------------------------------------------------------------------------- |
| ./README.md               | This file!                                                                                                                |
| ./server.go               | Entrypoint for the code; contains the `main` function                                                                     |
| ./routes.go               | Maps your URLs to the controllers/handlers for each URL                                                                   |
| ./event_models.go         | Defines your data structure and logic. I put in a few default events.                                                     |
| ./index_controllers.go    | Controllers related to the index (home) page                                                                              |
| ./templates.go            | Sets up your templates from which you generate HTML                                                                       |
| ./templates               | Directory with your templates. You'll need more templates ;P                                                              |
| ./templates/layout.gohtml | The "base" layout for your HTML frontend                                                                                  |
| ./templates/index.gohtml  | The template for the index (home) page                                                                                    |
| ./static.go               | Sets up the static file server (see next entry)                                                                           |
| ./staticfiles             | Directory with our "static" assets like images, css, js                                                                   |
| ./go.mod                  | [Go modules file](https://www.kablamo.com.au/blog/2018/12/10/just-tell-me-how-to-use-go-modules). Lists our dependencies. |
| ./go.sum                  | A "checksum" file that says precisely what versions of our dependencies need to be installed.                             |
| ./vendor                  | A directory containing our dependencies                                                                                   |
