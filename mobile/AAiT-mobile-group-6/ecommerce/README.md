# Ecommerce App

This repository contains the codebase for an ecommerce mobile application built using Flutter. The application follows a clean architecture pattern and implements best practices for maintainable and scalable code.


## Overview

The Ecommerce app is a mobile application designed to provide a product management experinece. Users can browse products, view detailed information, Manage products.

## Architecture

The app is built using **Clean Architecture**, which emphasizes the separation of concerns and makes the codebase more modular, testable, and maintainable. The architecture consists of three main layers:

1. **Domain Layer**: This is the core of the application, containing business logic, use cases, and entity definitions.
2. **Data Layer**: Responsible for managing data sources, including remote APIs and local databases. It contains repository implementations that communicate with the domain layer.
3. **Presentation Layer**: This layer handles UI and user interaction. It includes widgets, state management, and view models.

### Domain Layer

- **Entities**: Core business objects that are used across the app.
- **Use Cases**: Application-specific business rules. Each use case represents a single piece of business logic.
- **Repositories (Interfaces)**: Abstract definitions that the domain layer uses to interact with data sources.

### Data Layer

- **Repositories (Implementations)**: Concrete implementations of the repositories defined in the domain layer.
- **Data Sources**: Manage data retrieval and storage, including API calls and database interactions.
- **Models**:Are entities with additional functionality, particularly the ability to be serialized and deserialized from JSON. 

### Presentation Layer

- **Screens**: Flutter widgets that render the UI.
- **Widgets**: Reusable components to enhance the performance .
- **Bloc**: Manages the state of the UI and handles user inputs.


## Data Flow

The data flow in the application follows a unidirectional pattern:

1. **User Interaction**: The user interacts with the UI (e.g., tapping a button).
2. **ViewModel**: The UI sends the user action to a ViewModel.
3. **Use Case**: The ViewModel invokes a use case from the domain layer.
4. **Repository**: The use case interacts with the repository to fetch or update data.
5. **Data Source**: The repository communicates with the data source (e.g., API, database) to perform the required operation.
6. **UI Update**: The result is passed back through the layers to the ViewModel, which updates the UI accordingly.

## Features

- Product browsing
- Product details
- Product Management


## Dependencies

- [Flutter](https://flutter.dev/) - UI framework
- [Dartz](https://pub.dev/packages/dartz) - Functional programming in Dart
- [Mockito](https://pub.dev/packages/mockito) - Mocking framework for unit tests
- [Provider](https://pub.dev/packages/flutter_bloc) - State management

## Getting Started

To get started with the project, follow these steps:

1. **Clone the repository:**

   ```bash
   git clone https://github.com/HEran-ite/2024-internship-mobile-tasks.git
   ```

2. **Navigate to the project directory:**

   ```bash
   cd mobile/heran_eshetu/ecommerce
   ```

3. **Install dependencies:**

   ```bash
   flutter pub get
   ```

4. **Run the app:**

   ```bash
   flutter run
   ```

## Project Structure

```plaintext
lib/
├── core/                         # Core utilities, themes, and shared code
├── features/                     # Feature-specific code
│   ├── feature_1/                # Example feature module
│   │   ├── data/                 # Data layer
│   │   ├── domain/               # Domain layer
│   │   └── presentation/         # Presentation layer
├── main.dart                     # Main entry point of the application
test/                             # Unit and widget tests
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a Pull Request.

