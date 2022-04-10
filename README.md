### Functional Specification


The goal of this project is to create a website where users can upload short videoclips that they can share very quickly.
Think of this like [imgur](https://imgur.com/) but for video files (without the search/exploration part). 

Features:

* Unique users
    * Website should be usable without creating an account (cookie-based unique ID or something similar).
    * Users should be able to create accounts & log in.
* Signed in users should be able to:
    * See a list of all videos they have uploaded.
    * Manually delete videos.
* When a video is uploaded, the backend should use [FFmpeg](https://www.ffmpeg.org/) to process it (transcoding to viable format)
    * I'd like to make it possible to cut the video as well (for example, making the video only contain seconds [5-25] in a 30 second long clip).
    * Uploaded videos generate a unique url that anyone can use to watch the video.

### Technical Specification

The client will be created with [create-react-app](https://reactjs.org/docs/create-a-new-react-app.html) and TypeScript will be used.
The goal is to have a simple single page app for the users.
I've decided to use React since it is relatively simple, and I think that is more than powerful enough for my idea.

For the backend, [gin](https://github.com/gin-gonic/gin) wll be used.
I've worked with Golang before, but not much at all doing web server stuff.
Gin looks like a stable, popular and simple framework and should make it simpler to build a REST API compared to only using the standard library.
Here, I think the simplest would be to also serve the static react build from the go server, which should work nicely as I'd only need one running service to serve the entire website.

For the transcoding FFmpeg will be used.
I'm thinking of doing this in one of two ways, either through so Go library with ffmpeg bindings (I found several decently looking libs on GitHub), or through running FFmpeg directly through the cli on the server (or maybe even spawning a docker container to run the jobs, depending on how complex it is).

Data will be stored using [sqlite](https://www.sqlite.org/index.html).
The video files will be stored on disk, and the database will have entries containing the path to each respective file.
This seems to be the best way to do this.
I decided to use sqlite because of its simplicity and serverless nature.

I also want to look at using GitHub Actions.
This is because I haven't used used any CI/CD tools before and am very keen on learning how to do so.
I'm thinking something like running tests/coverage on the backend on each push, and something that builds the frontend automatically and moves the build to the correct directory in the backend and pushes it.
If time permits, I'd also like to look at GitLab CI/CD.

For deployment, I'm probably just going to host everything myself at home.
