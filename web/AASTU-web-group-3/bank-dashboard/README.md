# BankDash

Welcome to the **BankDash**! This project manages various banking operations through a web application. It includes features like user authentication, account management, loans, investments, transactions, and more. The project is developed using Next.js, TypeScript, Tailwind CSS, and Redux for state management.
![Screenshot](../Screenshot%202024-08-26%20200823.png)

## Table of Contents

- [Project Structure](#project-structure)
- [Setup Instructions](#setup-instructions)
- [Features](#features)
- [Technologies Used](#technologies-used)


## Project Structure

This project is structured into several branches, each managed by different teams. Each branch corresponds to a specific feature set or module of the banking system. Below are the main branches:



## Setup Instructions

To set up the project locally, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/RealEskalate/a2sv-g5-starter-project
2. **Pull from branch**:

   ```bash
   git pull origin aastu.web.g3.main
3. **Navigate to the project directory**:

   ```bash
   cd bank-dashboard
4. **Install dependencies**:

   ```bash
   npm install
4. **Run the development server**:

   ```bash
   npm run dev

5. **Open application**
    ```bash
    via localhost:300 on browser

## Features

### Authentication (Group 1)

- **Login Page**: Secure login with email and password.
- **Sign-up Page**: Register a new account with multi-step forms.

### Dashboard 

- Overview of user’s financial status, including account balances, recent transactions, and alerts.

### Account Management 

- View and edit account details.
- Manage multiple bank accounts under one user profile.

### Investments 

- Browse available investment options.
- Manage investment portfolios.

### Loans 

- Apply for new loans.
- Manage existing loans, view repayment schedules, and make payments.

### Transactions 

- View all transaction history.
- Transfer money between accounts or to external accounts.

### Services 

- Access various banking services like bill payments and subscriptions.

### Settings 
- Update personal information and account preferences.
- Manage notification settings and security options.

## Technologies Used

- **Next.js 14**: Framework for building React applications with server-side rendering and static site generation.
- **TypeScript**: Type-safe JavaScript for better development experience.
- **Tailwind CSS**: Utility-first CSS framework for rapid UI development.
- **Redux Toolkit (RTK Query)**: State management and data fetching.
- **NextAuth.js**: Authentication library for handling login and sign-up.
