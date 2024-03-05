# GoLang HTTP Server with Request Counter

This GoLang program creates an HTTP server that responds to requests with a counter of the total number of requests received during the previous 60 seconds (moving window). The server persists the data to a file, ensuring that the count is maintained even after restarting.

## Prerequisites

Make sure you have Go installed on your machine. You can download and install it from [here](https://golang.org/dl/).

## Running the Server

1. Clone or download this repository to your local machine.
   
2. Navigate to the directory containing the files.

3. Open a terminal window and run the following command to build the executable:

    ```bash
    go build
    ```

    This will create an executable file named `main` in the same directory.

4. Run the executable by executing the following command:

    ```bash
    ./main
    ```

    This will start the HTTP server, which will listen on port 8080.

## Accessing the Server

Once the server is running, you can access it by opening a web browser and navigating to [http://localhost:8080](http://localhost:8080) or by using tools like `curl` or `wget` in the terminal:

```bash
curl http://localhost:8080/numberOfRequests
