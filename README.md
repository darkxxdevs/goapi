# API - GOlANG

<p align="left">
  <img src="./goimg/go.png" alt="Go Logo">
</p>

Welcome to the API developed by darkxxdevs . This package provides a simple API for managing and retrieving information about courses. It is built using Go and the Gorilla Mux router.

## Getting Started

Before you begin, make sure you have Go installed on your system.

1. Clone this repository.
2. Install the required dependencies by running:

   ```sh
   go get -u github.com/gorilla/mux

   ```

3. Build and run api

   ```
   go run main.go

   ```

- api will be available at [http://localhost:8080](http://localhost:8080).

## API Endpoints

### Home

- **GET** `/`
  - Description: Displays a welcome message.
  - Example: [http://localhost:8080](http://localhost:8080)

### Get All Courses

- **GET** `/courses`
  - Description: Retrieve a list of all courses.
  - Example: [http://localhost:8080/courses](http://localhost:8080/courses)

### Get One Course

- **GET** `/courses/{courseid}`
  - Description: Retrieve a single course by its ID.
  - Example: [http://localhost:8080/courses/1](http://localhost:8080/courses/1)

### Create One Course

- **POST** `/courses/post`
  - Description: Create a new course.
  - Example: [POST Request with JSON Data](#sample-post-request)

### Update One Course

- **PUT** `/courses/update/{courseid}`
  - Description: Update an existing course.
  - Example: [PUT Request with JSON Data](#sample-put-request)

### Delete One Course

- **DELETE** `/courses/delete/{courseid}`
  - Description: Delete a course by its ID.
  - Example: [DELETE Request](#sample-delete-request)

### Delete All Courses

- **DELETE** `/courses/deleteall`
  - Description: Delete all courses.
  - Example: [DELETE Request](#sample-delete-all-request)

## Sample POST Request

To create a new course, you can send a POST request with the following JSON data:

```json
{
  "courseid": "2",
  "name": "New Course",
  "price": "2000",
  "author": {
    "author_name": "John Doe",
    "auhtor_websitr": "https://example.com"
  }
}
```

## Sample GET Request

To retrieve a single course by its ID, send a GET request to the following endpoint with the appropriate `courseid` in the URL.

Example: Retrieve a course with `courseid` equal to 1

```http
GET /courses/1
```

## Sample PUT Request

To update an existing course, send a PUT request with the updated JSON data for the course, including the `courseid` of the course you want to update.

Example JSON data for updating a course:

```json
{
  "courseid": "2",
  "name": "Updated Course",
  "price": "2500",
  "author": {
    "author_name": "Jane Smith",
    "auhtor_websitr": "https://updated-example.com"
  }
}
```

Make sure to include the courseid in the URL to specify which course to update, such as /courses/update/2.

## Sample DELETE Request

To delete a specific course by its ID, send a DELETE request to the following endpoint with the appropriate `courseid` in the URL.

Example: Delete a course with `courseid` equal to 1

```http
DELETE /courses/delete/1
```

The API will remove the specified course, and the response will indicate the success of the operation.

## Sample DELETE All Request

To delete all courses, send a DELETE request to the following endpoint.

Example: Delete all courses

```http
DELETE /courses/deleteall
```

## Developers

- **Developer Name**: [Arpit Yadav]
- **GitHub Profile**: [GitHub Profile](https://github.com/darkxxdevs)

Feel free to connect with the developer and contribute to this project. If you have any questions or need assistance, please reach out to the developer for support.

Thank you for using this API package!
