# Task Progress for Ecommerce App
A Flutter Mobile Application built for the 2024 A2SV Summer Internship in the mobile team.

## Table of contents
- [Screenshots](#screenshots)
- [Tasks](#tasks)

## Tasks

### Contents
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


## Task 14 - Task 14: Implement Local Data Source

### Task Objectives
- ✔️ Add SharedPreferences to project packages
 - ✔️ Implementing local datasource 
 - ✔️ Complete ProductLocalDatasourceImpl class
 - ✔️ Write test

  
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


## Screenshots

### Splash Screen
| Splash Screen | Sign up | Sign in | Logout |
|---------------|---------|---------|--------|
|![splash screen](https://github.com/user-attachments/assets/c77a45cd-1add-4d09-8d19-7ab07780d902)|![Sign Up Page](https://github.com/user-attachments/assets/33c0cb57-c28c-47b9-9ea3-c111599ab1e2)|![Sign in Page](https://github.com/user-attachments/assets/292d8c97-f99b-45c9-ac3e-8397f84547ff)|![Logout confirm](https://github.com/user-attachments/assets/a40f38b8-3007-4607-bac1-cf3249ea8143)|


| Home Light Default| Home Light with Avatar | Home Light With Product Added | Pick Avatar Light |
|-------------------|------------------------|-------------------------------|-------------------|
|![Home Page Light Mode Default](https://github.com/user-attachments/assets/425edad9-e8cf-4c67-8a70-4ba31a948aba)|![Home Page Light Mode After Avatar add](https://github.com/user-attachments/assets/38fb7507-132f-4c27-92c4-f509959ac715)|![Home Page Light Mode After Product Add](https://github.com/user-attachments/assets/9f4be1da-82d3-492c-9fb9-2c6c1855f7fe)|![Pick Avatar Light](https://github.com/user-attachments/assets/f139bb31-27cf-4a41-990f-fc870e52ca28)|


| Home Dark Default | Home Dark with Avatar | Home Dark with Product Added | Pick Avatar Dark |
|-------------------|-----------------------|------------------------------|------------------|
|![Home Page Dark Mode](https://github.com/user-attachments/assets/f31440c9-58f4-4d43-bd92-ebb88e880906)|![Home Page Dark Mode With Avatar](https://github.com/user-attachments/assets/c46d1582-7d2f-4a54-b3c3-e5811b3aba6f)|![Home Page Dark Mode After Product Add](https://github.com/user-attachments/assets/617d6404-ed1e-4cb3-a58d-b0238898fb5d)|![Pick Avatar Dark](https://github.com/user-attachments/assets/2903610d-297e-48b4-af89-8d0bc4c44e3c)|


| Details Light | Details Dark | Search Light | Search Dark |
|---------------|--------------|--------------|-------------|
|![Details Page Light Mode](https://github.com/user-attachments/assets/5db7f1d4-db8a-4dda-9fd4-0eee733c0598)|![Details Page Dark Mode](https://github.com/user-attachments/assets/08d28036-34ef-4235-a0c7-a84880889868)|![Search Light Mode](https://github.com/user-attachments/assets/29fb9b89-c2f8-4b22-a827-ae6cfe352d3a)|![Search Dark Mode](https://github.com/user-attachments/assets/5ef77f82-714a-445d-8f53-e6a085f13ad3)|


| Add Product Light Default | Add Product Dark Default | Add Product Light Filled | Add Product Dark Filled |
|---------------------------|--------------------------|--------------------------|-------------------------|
|![Add Product Light Mode Default](https://github.com/user-attachments/assets/0070c6b8-514b-424a-8b58-3f54f3cfdac8)|![Add Product Dark Mode Default](https://github.com/user-attachments/assets/85905491-4233-45d1-bb24-a67cdaef7bb7)|![Add Product Light Mode Filled](https://github.com/user-attachments/assets/46f44b84-a2ae-4498-a219-41975e7ad18e)| ![Add Product Dark Mode Filled](https://github.com/user-attachments/assets/c7087811-e86f-42d5-b0cc-665ddc45fa54)|


| Update Product Light | Update Product Dark | Chats List |
|----------------------|---------------------|------------|
|![Update Light Mode ](https://github.com/user-attachments/assets/c797ddcc-1bb7-4524-a163-f1aa861d6818)|![Update Dark Mode](https://github.com/user-attachments/assets/8d639190-0b5c-460a-a063-9c5aaf74278c)|![Chats](https://github.com/user-attachments/assets/c9e0720e-2938-4587-9e6a-ae23ec5a8641)|
