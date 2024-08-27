# A2SV Banking System - Group 2

## Project Overview

The **A2SV Banking System** is a modern, full-featured banking application designed for both web and mobile platforms. It enables users to perform various banking operations such as managing accounts, making transactions, applying for loans, and more, all from a user-friendly interface. This project is developed by **AASTU Web Group 2** as part of the A2SV project start_up. 

### Key Features:
- **User Authentication**: Sign In and Sign Up functionality for secure user access.
- **Account Management**: Create, view, and manage multiple accounts.
- **Transactions**: Real-time deposits, withdrawals, transfers, and transaction history.
- **Investments & Credit Cards**: Management of investment portfolios and credit card applications.
- **Loans**: Apply for and manage loans.
- **Services**: Explore additional banking services.
- **Settings**: Manage user settings and preferences.


## Technology Stack

- **Frontend and Backend**: Next.js (Fullstack React framework)
- **Styling**: Tailwind CSS 
- **Mobile**: Responsiveness via tailwind
- **Testing**: Jest for unit testing
- **Deployment**: Netlify (for seamless frontend and backend deployment)


## Getting Started

### Prerequisites

- Node.js and npm installed (version 14 or higher)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/RealEskalate/a2sv-g5-project-phase-starter-project.git
    cd a2sv-g5-project-phase-starter-project/web/AASTU-web-group-2/a2sv-banking-system
    ```

2. Install the dependencies:

    ```bash
    npm install
    ```

3. Setup environment variables:
    - Copy `.env.example` and rename it to `.env`.

4. Start the application:

    ```bash
    npm run dev
    ```

5. The app should be running at `http://localhost:3000`.


## Usage

1. **Sign In / Sign Up**: Users can create accounts or log in to access banking features.
2. **Dashboard**: Overview of account balances, recent transactions, and quick access to services.
3. **Transactions**: Manage transactions such as deposits, withdrawals, and transfers.
4. **Investments**: Manage and track investment portfolios.
5. **Credit Cards**: Apply for and manage credit card accounts.
6. **Loans**: Apply for loans and manage existing loans.
7. **Services**: Access additional banking services like insurance, financial advisory, etc.
8. **Settings**: Customize account preferences, security options, and more.

## Project Structure

- **backend/**: Server-side code, handling API requests.
- **web/**: Next.js client-side code for the web interface.

## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/NewFeature`).
3. Commit your changes (`git commit -m 'Add NewFeature'`).
4. Push to the branch (`git push origin feature/NewFeature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License.

## Contact

For questions, contact [Contributors](https://github.com/RealEskalate/a2sv-g5-project-phase-starter-project).

---

## Functionalities

### Sign In Page:
- **Purpose**: Allows users to authenticate and access their banking accounts securely.
- **Functionality**: 
  - Accepts username and password.
  - Authenticates using the backend API (JWT tokens).
  - Redirects authenticated users to the dashboard.

### Sign Up Page:
- **Purpose**: Provides users with the ability to create a new banking account.
- **Functionality**:
  - Collects user information (name, email, password, etc.).
  - Validates input fields and interacts with the backend to create a new user in the database.
  - On success, the user is redirected to the Sign In page.

### Dashboard:
- **Purpose**: Acts as the central hub for the user's banking activities.
- **Functionality**:
  - Displays account balances.
  - Shows recent transactions.
  - Provides quick access links to various banking services such as accounts, transactions, and loans.

### Transactions Page:
- **Purpose**: Handles all user transactions including deposits, withdrawals, and transfers.
- **Functionality**:
  - Provides a detailed history of the user's transactions.
  - Allows users to filter transactions by date, type, and amount.
  - Includes forms to initiate new deposits, withdrawals, or transfers.

### Accounts Page:
- **Purpose**: Allows users to manage their bank accounts.
- **Functionality**:
  - Users can open new accounts, view existing accounts, and delete accounts.
  - Displays detailed information on each account such as account number, balance, and transaction history.

### Investments Page:
- **Purpose**: Provides a platform for users to manage and track their investment portfolios.
- **Functionality**:
  - Displays current investment holdings.
  - Allows users to buy and sell investment products like stocks or bonds.
  - Shows real-time performance data for their investments.

### Credit Cards Page:
- **Purpose**: Enables users to apply for and manage credit cards.
- **Functionality**:
  - Users can apply for new credit cards by providing necessary details.
  - View credit card balance, transaction history, and available credit.
  - Manage credit card payments and billing cycles.

### Loans Page:
- **Purpose**: Manages loan applications and payments.
- **Functionality**:
  - Users can apply for personal or business loans by submitting required documentation.
  - View current loans, remaining balance, and due dates.
  - Make loan repayments through linked accounts.

### Services Page:
- **Purpose**: Provides access to additional banking services such as insurance, financial advisory, and retirement plans.
- **Functionality**:
  - Explore various service offerings.
  - Allows users to sign up for new services or contact a service representative for more information.

### Settings Page:
- **Purpose**: Allows users to customize their account settings.
- **Functionality**:
  - Change personal information (e.g., name, email, password).
  - Manage security options such as enabling two-factor authentication (2FA).
  - Set notification preferences for transactions, account updates, etc.
