# gorala-link-shortener

gorala-link-shortener is a super fast and simple link shortener written in Go.
It provides a an API as well as a web frontend for shortining your links.
For persistence it uses SQLite. The project also includes an admin interface to show some statistics.
Under the hood it uses [Echo](https://echo.labstack.com/) + [GORM](https://gorm.io/). 

If you are looking for a minimal solution without much to set up this project might be for you.
Since the project is sosmall and straight forward it might also be a good starting point if you want to learn Go or [Echo](https://echo.labstack.com/).

If you are looking for a more sophisticated version of a link shortener with a broad feature scope I'd suggest you go with something like [polr](https://github.com/cydrobolt/polr)

Usecases such as a private mode or enhanced statistics are currently not included in this project.

## Usage
Simply build the included `Dockerfile` or clone the project and do `go run server.go`.
The basic authentication for the admin area is via environment variables.
Inside the `Dockerfile` you have:

```
ENV link-admin-name=rala
ENV link-admin-pw=#97CJrey1ni6
```

You can adapt those variables to your personal credentials.
If you are running locally via `go run` make sure to also set the environment variables in case you want to access the admin area.

## Example
You can find a demo version of the project [here](https://l.gorala.icu/)

The admin site can be accessed under [/admin/main](https://l.gorala.icu/admin/main)

See the credentials from above for auth.


## Contributing
Thanks for your interest! Do not hesitate to open an issue if you have a question, feedback or found something that is not working like it should.

Pull requests for improvements of the library or it's documentation are also highly appreciated.

## Licenses
This library and its content is released under the [MIT License](https://choosealicense.com/licenses/mit/).
