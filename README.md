# Astroweather

There's a solar system with some weird coupling between it's planets relative location in space and the meteorological effects of their atmospheres.
This is a REST service to tell them how the weather is gonna be a given day; to be taken as a token of piece from the inhabitants of Earth.

## Setup

To run this program, you need a working installation of Go and Gcloud CLI, and an initialized appengine project to make the deployment. With those requirements, you should be able to run your own version of the service with:

`git checkout github.com/thelmholtz/astroweather`
`gcloud app deploy`

## Try it

You can play with [a quick version of this service](https://astroweather.appspot.com) if you rather not go through the hustle.

