## BANKING DASHBOARD

# table of content

1 About
2 Demo
3 Features
4 installation
5 usage
6 configuration
7 tests
8 contribution
9 License
10 Content

# ABout

The Banking Dashboard project focuses on building a user-friendly front end for an already established banking API. Using the provided Figma design and API documentation, our goal is to create a responsive and intuitive dashboard where users can manage their financial accounts seamlessly.
This project involves developing a front end with Next.js, Tailwind CSS, and TypeScript. By following a structured approach, we start by analyzing the Figma design to break down the UI into manageable components and review the API documentation to understand data flows and integration points.

# Demo

# Features

1.Real-time Financial Data Display: The app integrates with an existing API to display up-to-date financial information from multiple accounts.

2.Component-Based UI: The user interface is built using reusable components, such as Navbar, Account Summary, and Transaction List, providing a clean and organized layout that is easy to navigate.

3.Responsive Layout: The dashboard will have a responsive design, ensuring it works well on various devices, including desktops, tablets, and mobile phones.

4.Account Overview: Users can view a summary of their account details, including balances and recent transactions, allowing them to monitor their financial status at a glance.

5.Transaction List: A detailed list of transactions is provided, showing the user all their recent transactions. This list can be filtered and sorted for easier navigation.

6.API Integration: The app uses a service layer or custom hooks to fetch data from the API. This includes fetching transaction data, account details, and any other necessary financial information.

7.Error Handling and Loading States: The app includes robust error handling and displays appropriate loading states to provide a smooth user experience, even in cases where data might be delayed or not available.

8.Authentication Support: The app integrates with the provided API’s authentication mechanisms to ensure secure access to user data.

9.Testing Suite: The app includes unit and integration tests to ensure components work as expected and the user experience is intuitive and error-free.
.

# Installation

1 . clone the repository

```
git clone https://github.com/RealEskalate/a2sv-g5-project-phase-starter-project.git
```

2. change the branch to astu.web.g2.main
   `git checkout astu.web.g2.main`
   3.Navigate to the directory
   `cd a2sv-g5-project-phase-starter-project/web/ASTU-web-group-2/banking-dashboard`
   4.Install dependencies
   ` npm install ```5.Run the project `npm run dev`

# Usage

- To start the application locally, run the following command:
  `npm run dev`
  This command will start the development server, and you can access the dashboard by navigating to http://localhost:3000 in your web browser.
  -Signing Up
  Create an Account: On the homepage, click the "Sign Up" button to create a new account.
  -Fill in Your Details: Enter your personal information, such as your name, email address, password, and then submit the form.

Log In: After sign up, it redirect you to the dashboard.

-Accounts Page:
View recent transactions and account summaries.
Access detailed credit card information and financial overviews.
-Credit Cards Page:
Manage credit card details and settings.
View expense statistics and card lists.
-Investments Page:
Monitor yearly and monthly investment summaries.
Track investment performance and trending stocks.
-Loans Page:
View an overview of active loans and their statuses.
-Services Page:
Access and manage various banking services.
-Settings Page:
Edit user profile and account settings.

# Tests

-To run the test suite, use the following command:
`npx cypress open`
and chooose the option that you went to test

# Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any improvements or new features .

# License

This project is licensed under the MIT License.

# Contact

For questions or feedback, please reach out to one of the following. -`fasilhawultie19@gmail.com` -`kalebwondimu33@gmail.com` -`dureti104@gmail.com`

-`biniyam.negasa@a2sv.org`
