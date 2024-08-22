# Task Progress for Ecommerce App
A Flutter Mobile Application built for the 2024 A2SV Summer Internship in the mobile team.

## Table of contents
- [Task 9: Entities, Use Cases, and Repositories](#task-9-entities-use-cases-and-repositories)
- [Task 10: Data Overview Layer](#task-10-data-overview-layer)
- [Task 11: Contracts of Data Sources](#task-11-contracts-of-data-sources)
- [Task 12: Implement Repository](#task-12-implement-repository)
- [Task 13: Implement Network Info](#task-13-implement-network-info)
- [Task 14: Implement Local Data Source](#task-14-implement-local-data-source)
- [Task 15: Implement Remote Data Source](#task-15-implement-remote-data-source)
- [Task 16: Improve Code Organization and Reusability](#task-16-improve-code-organization-and-reusability)
- [Task 17: Implement Bloc](#task-17-implement-bloc)
- [Task 18: Dependency Injection](#task-18-dependency-injection)
- [Task 19: Implement User Interface](#task-19-implement-user-interface)
- [Task 20: Consume Bloc for eCommerce](#task-20-consume-bloc-for-ecommerce)
  
## Task 9: Entities, Use Cases, and Repositories

### Task Objectives
 - ✔️ Create entities for the products in the eCommerce Mobile App.
 - ✔️ Define use cases for inserting, updating, deleting, and getting a product.
 - ✔️ Implement repositories to handle data operations for the products.

### Testing the usecase
![Screenshot 2024-08-07 120634](https://github.com/user-attachments/assets/7aacc7a7-908d-4c72-9a8d-830d9f5178a7)


## Task 10: Data Overview Layer

### Task Objectives
- ✔️ Step 1: Folder Setup
- ✔️ Step 2: Implement Models
- ✔️ Step 3: Documentation

### Testing model layer
![Screenshot 2024-08-07 170342](https://github.com/user-attachments/assets/0821d2db-31b0-4b33-a9d0-b7ba6641779c)


## Task 11: Contracts of Data Sources 

### Task Objectives
- ✔️ Contract and Repository
  - ✔️ Implementation of a contract defining repository methods
  - ✔️ Interfaces or abstract classes for repository dependencies
  - ✔️ Basic structure of the repository

## Task 12: Implement Repository

### Task Objectives
- ✔️ Contract and Repository:
 - ✔️ Use local datasource when network is unavailable
 - ✔️ Use remote datasource when network is available

### Commit links
- https://github.com/AryamEzra/2024-internship-mobile-tasks/commit/7e7a750c32afa4debaed16a9f723fe4f265c9738


## Task 13: Implement Network Info

### Task Objectives
- ✔️ Step 1: Create the NetworkInfo Class 
 - ✔️ Class Definition
    - Create a new Dart file named network_info.dart and correctly define the NetworkInfo abstract class along with the NetworkInfoImpl class that implements it, using the provided starting point.
 - ✔️ Constructor and Implementation
   -  Implement the NetworkInfoImpl class constructor that takes an InternetConnectionChecker instance.
   -  Implement the isConnected getter method using the provided resource, ensuring proper usage of the connectionChecker instance.

- ✔️ Step 2: Use NetworkInfo in the Repository 
 - ✔️ Import and Dependency Injection 
 - ✔️ Network Connectivity Usage 
 - ✔️ Error Handling
   
### Commit links
- https://github.com/AryamEzra/2024-internship-mobile-tasks/commit/cefb3da59eee3f01999d23b3770624447a0e06cb


## Task 14 - Task 14: Implement Local Data Source

### Task Objectives
- ✔️ Add SharedPreferences to project packages
 - ✔️ Implementing local datasource 
 - ✔️ Complete ProductLocalDatasourceImpl class
 - ✔️ Write test

### Commit links:
- https://github.com/AryamEzra/2024-internship-mobile-tasks/commit/be47c42af14482aa3635884044f2147073073fc2

  
## Task 15: Implement Remote Data Source

### Task Objectives
- ✔️ Implementing remote datasource
 - ✔️ Complete ProductRemoteDatasourceImpl class
 - ✔️ Write test


## Task 16: Improve Code Organization and Reusability

### Task Objectives
- ✔️ Code Organization
- ✔️ Reusability 
- ✔️ Integration and Funcionality
 
### Commit links:
- https://github.com/AryamEzra/2024-internship-mobile-tasks/commit/bff2d6de75a8de51acc141b576069a56ded25cc8

### Testing the app so far
![Screenshot 2024-08-12 114807](https://github.com/user-attachments/assets/32f8b16e-e1ff-4f04-af9f-699b5ea3ea63)


## Task 17: Implement Bloc

### Task Objectives
- ✔️ Task 17.1: Create Event Classes
- ✔️ Task 17.2: Create State Classes
- ✔️ Task 17.3: Create ProductBloc (5 points)
  

## Task 18: Dependency Injection

### Task Objectives
- ✔️ Setup and configuration
- ✔️ Creating the Injection Container

  
## Task 19: Implement User Interface

### Task Objectives
- ✔️ Presentation Logic Holder
- ✔️ Building the UI Structure
- ✔️ Displaying Different States
- ✔️ Handling UI states (Empty, Loading, Error, Loaded) using BlocBuilder.
- ✔️ Creating custom widgets for displaying messages and loading indicators.
- ✔️ Handling User Input
- ✔️ Refining UI Components


## Task 20: Consume Bloc for eCommerce

### Task Objectives
- ✔️ Design a screen that allows users to input product details and create a new product.
- ✔️ Consume the appropriate bloc method to handle the product creation process.
- ✔️ Retrieve All Products Page
- ✔️ Product Detail Page [Get/Delete]
- ✔️ Update Product Page
- ✔️ Deleting Product 
- ✔️ Navigation and Routing


This project is a starting point for a Flutter application.

A few resources to get you started if this is your first Flutter project:

- [Lab: Write your first Flutter app](https://docs.flutter.dev/get-started/codelab)
- [Cookbook: Useful Flutter samples](https://docs.flutter.dev/cookbook)

For help getting started with Flutter development, view the
[online documentation](https://docs.flutter.dev/), which offers tutorials,
samples, guidance on mobile development, and a full API reference.
