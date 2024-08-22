# E_commerce_app_Clean

This project is a clean architecture-based e-commerce application that runs on Flutter. The application is made with a clear division of responsibilities across many layers, making it scalable, maintainable, and tested.

## table_of_contents
1. [Project Details](#project-overview)
2. [Features](#features)
   - [CRUD operations on Product](#CRUD-operations-on-Product)
       - [Domain Layer](#domain-layer)
       - [Data Layer](#data-layer)
       - [Presentation Layer](#presentation-layer)
3. [Testing](#testing)
4. [Getting Started](#getting-started)
6. [License](#license)

## Project Overview

This project adheres to test-driven development and clean architecture, and it is designed to manage products.

## Features 
 - ## CRUD operations on Product
    - ### Domain Layer
       - **Use cases**:  the user can get all products, get a single product, add a product, update a product and delete a product.
       - **Entities**: it has an entity that has and id, name, description, rating, price, category, and an image parameters to represent the product.
    - ### Data Layer
      **Models**: it has a model that handles Json serialization and deserialization.
         - `fromJson` handles the mapping of json files to the corresponding product entity.
         - `toJson` handles the transformation of the product entity to json format.
## Testing

since it is a test driven development, currently it has a unit test written to it for the domain layer usecases and for the data model

## Getting Started

1. Clone the repository.
2. Run `flutter pub get` to install dependencies.
3. Use `flutter run` to start the app.

## License

This project is open-source and available under the MIT License.
   
