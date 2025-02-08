This is a simple movie search API built with Go and the OMDB API.

## Getting Started

To get started, you'll need to have Go installed on your machine. You can download the latest version from the official Go website: <https://go.dev/dl/>.

Once you have Go installed, you can clone this repository and navigate to the project directory:

```bash
git clone https://github.com/your-username/go-omdb-integration.git
cd go-omdb-integration
```

Next, you can build the project by running the following command:

```bash
go build
```

This will create an executable file named `go-omdb-integration` in the current directory.

## Running the Project

To run the project, you can use the following command:

```bash
./go-omdb-integration
```

This will start the server and listen on port 8080.

## Testing the API

You can test the API by making a GET request to the `/` endpoint with a `search` parameter:

```bash
curl http://localhost:8080/?search=the+matrix
```

This will return a JSON response with the search results.
