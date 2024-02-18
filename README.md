# Go Fiber CRM RESTful API 🚀

This project is a simple RESTful API built using GoLang with Fiber framework and SQLite database.

## Getting Started 🛠️

Follow these steps to run the project locally:

1. **Set Up Database:**

    Make sure you have SQLite installed. The database file is located at `/database/leads.db`.

2. **Set Up Dependencies:**

    Run `go mod tidy` to install the necessary dependencies.

3. **Run the Application:**

    ```bash
    go run main.go
    ```

The server will start at http://localhost:3000.

## API Endpoints 🚪

### Get Leads

-   **Endpoint:** `GET /api/v1/lead`
-   **Description:** Retrieve a list of all leads.

### Get Lead by ID

-   **Endpoint:** `GET /api/v1/lead/:id`
-   **Description:** Retrieve a specific lead by ID.

### Create Lead

-   **Endpoint:** `POST /api/v1/lead`
-   **Description:** Create a new lead.

### Delete Lead by ID

-   **Endpoint:** `DELETE /api/v1/lead/:id`
-   **Description:** Delete a specific lead by ID.

## Closing Notes📝

If you find any issues or have suggestions for improvement, please feel free to open an issue or submit a pull request.

Happy coding!🚀👨‍💻
