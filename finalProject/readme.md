# Bookstore API

This project is a simple **Bookstore API** that provides functionality for managing books, authors, customers, orders, and book sales in an e-commerce setting. It is designed using Go, with a focus on an in-memory data store, and follows a clean architecture with separation of concerns.

## Features

The API allows you to perform CRUD operations on various resources like books, authors, customers, orders, and book sales.

- **Books**: Add, retrieve, update, and delete books from the store.
- **Authors**: Add, retrieve, update, and delete authors.
- **Customers**: Manage customer information (CRUD operations).
- **Orders**: Create, retrieve, update, and delete orders.
- **Book Sales**: Record and search for book sales.

### API Endpoints

The following sections describe the API endpoints, based on the Swagger documentation.

#### Books

- **POST /books**: Create a new book.
- **GET /books/{id}**: Retrieve a book by its ID.
- **PUT /books/{id}**: Update a book by its ID.
- **DELETE /books/{id}**: Delete a book by its ID.
- **GET /books**: Search for books by filters. All books are returned if no filters are provided in the JSON request.

#### Authors

- **POST /authors**: Create a new author.
- **GET /authors/{id}**: Retrieve an author by ID.
- **PUT /authors/{id}**: Update an author by ID.
- **DELETE /authors/{id}**: Delete an author by ID.
- **GET /authors**: Search for authors. All authors are returned if no filters are provided in the JSON request.

#### Customers

- **POST /customers**: Create a new customer.
- **GET /customers/{id}**: Retrieve a customer by ID.
- **PUT /customers/{id}**: Update a customer by ID.
- **DELETE /customers/{id}**: Delete a customer by ID.
- **GET /customers**: Get all customers.

#### Orders

- **POST /orders**: Create a new order.
- **GET /orders/{id}**: Retrieve an order by ID.
- **PUT /orders/{id}**: Update an order by ID.
- **DELETE /orders/{id}**: Delete an order by ID.
- **GET /orders**: Get all orders.


## Project Structure

The project is structured as follows:

```
/bookstore
  /handlers        # HTTP handlers for handling API requests
  /memory          # In-memory store for handling the data
  /models          # Data models representing the entities
  /repositories    # Interfaces for interacting with the data store
  /services        # Business logic layer for handling CRUD operations
  openapi.yml      # Swagger configuration
  main.go          # Entry point to run the application
```

### Directories and Files Breakdown

- **/handlers**: Contains the HTTP handler functions which process incoming requests, map them to the appropriate service methods, and return responses.
  - **BookHandler**: Manages HTTP operations related to books (e.g., `CreateBook`, `GetBookById`).
  - **AuthorHandler**: Handles requests for author-related operations.
  - **CustomerHandler**: Processes customer management requests.
  - **OrderHandler**: Deals with order processing.
  - **BookSaleHandler**: Manages operations related to book sales.

- **/memory**: Implements the in-memory data store using Go maps and sync mechanisms (mutexes). Singleton instances are used for managing resources like books, customers, and orders.
  - **BookMemoryStore**: A map-based storage for books, using Go’s `sync.Mutex` for thread-safe operations.
  - **AuthorMemoryStore**: Handles the in-memory storage for authors.
  - **CustomerMemoryStore**: Manages customer data in memory.
  - **OrderMemoryStore**: Stores order data, with methods for CRUD operations.
  - **BookSaleMemoryStore**: Tracks book sales in memory, enabling quick access and modifications.

- **/models**: Defines the data models that represent entities such as books, authors, orders, and book sales.
  - **Book**: Represents a book with attributes like `ID`, `Title`, `AuthorID`, `Price`, and `Stock`.
  - **Author**: Defines an author entity with fields such as `ID`, `Name`, and `Biography`.
  - **Customer**: Represents a customer, including details like `ID`, `Name`, `Email`, and `Address`.
  - **Order**: Captures the details of an order, including `ID`, `CustomerID`, `BookIDs`, and `Status`.
  - **BookSale**: Represents a sale transaction, including `ID`, `BookID`, `Quantity`, and `SaleDate`.

- **/repositories**: Contains interfaces for data access layers, such as methods for creating, retrieving, and deleting entities from the data store.
  - **BookRepository**: An interface defining methods for managing books in the data store (e.g., `FindByID`, `Save`, `Delete`).
  - **AuthorRepository**: Interface for CRUD operations on authors.
  - **CustomerRepository**: Interface for managing customer data.
  - **OrderRepository**: Defines methods for order management.
  - **BookSaleRepository**: Interface for handling book sales.

- **/services**: Handles the business logic and interacts with the repositories for CRUD operations and data management.
  - **BookService**: Contains methods like `CreateBook`, `UpdateBook`, `DeleteBook`, and `SearchBooks`.
  - **AuthorService**: Manages operations related to authors, such as creating and retrieving author information.
  - **CustomerService**: Handles customer-related business logic, including registration and updates.
  - **OrderService**: Manages the lifecycle of orders, ensuring proper validation and processing.
  - **BookSaleService**: Handles the business logic for recording and managing book sales.

- **openapi.yml**: Swagger/OpenAPI specification for the API.
  - **API Endpoints**: Descriptions of each endpoint, including HTTP methods, paths, and expected responses.
  - **Request and Response Models**: Details of the request payloads and response formats for each endpoint.
  - **Validation Rules**: Specifications for required fields, data types, and possible error responses.

- **main.go**: The main entry point for the application, where the server is set up, and routing is initialized.
  - **Initializing Services**: Sets up the service layer with dependencies like repositories.
  - **Setting Up Routes**: Configures the HTTP router with all API endpoints, linking handlers to paths.
  - **Starting the Server**: Launches the HTTP server, listening for incoming requests on the specified port.

## Example Requests

Refer to the Swagger file to explore different APIs and their examples.

## Technologies Used

- **Go**: The programming language used for building the API.
- **Swagger/OpenAPI**: API documentation and testing interface.
- **In-Memory Store**: Simple in-memory storage for entities.
- **Goroutines & Mutexes**: For managing concurrency and thread safety.

## Contributing

Feel free to fork the repository, open issues, and submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
