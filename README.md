# Project starter

Ahoy! This is the project starter. I've done some
of the hard parts for you. Good luck.

## Version Control

You'll want *just one code repo* for your team. You should 
use git branches liberally to keep track of your work. I like
using GitHub Issues and Projects to track work. Though, you'll
have to shoehorn it a little bit to make it work for Scrum
strictly.

## Building the project

Run `go build -mod vendor` to compile your project and then
run `./classproject` or `./classproject` on Windows
in order to run the app. Preview the app by visiting
[http://localhost:8080](http://localhost:8080) if you're
running it locally. Or, preview using the URL provided
GitHub CodeSpaces if you're running there, natch.

## Deploying to Render, Heroku, Fly.io, or whatever

You'll need to deploy your app to "production" somewhere.
I don't care where you deploy it, as long as that URL is
available from Yale's networks. You could deploy it to
Fly.io, Render, Heroku, DigitalOcean, whatever. I think
you're going to find Render the easiest.

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

## Automatic reload

If you want your app to reload every time you make a
change you can do the following. First
install reflex with `go get github.com/cespare/reflex`.

Then, run

```
~/go/bin/reflex -d fancy -r'\.go' -r'\.gohtml' -s -R vendor. -- go run *.go
```

or something like that. Look at the reflex documentation. Automatic
reload is pretty rad while developing. As a general rule, developers
want to let computers do what computers are good at (tasks that can be automated)
so that they, the developers, can focus on what they are good at: the
logic of the product.

## Other info

Information about the class final project is distributed between
a few places and I apologize for this. You can find information
about the project in the following places:

- This page (which you're likely looking at in your own repo)
- The sprint-1 assignment page:
  [656](https://www.656.mba/#assignments/project-sprint-1)
- The "about" repo for the class:
  [https://github.com/yale-mgt-656-fall-2023/about/blob/master/class-project.md](https://github.com/yale-mgt-656-fall-2023/about/blob/master/class-project.md)
- The online grading: <http://grading.656.mba/new/project>
- The reference solution
  [http://project.solutions.656.mba/](http://project.solutions.656.mba/).
  This is probably the most useful thing. If you're having difficulty passing
  some test on the grading code, please look at my code (feel free!) and 
  make sure you're using the kinds of HTML attributes that the grading
  code expects.
- My comments in class
