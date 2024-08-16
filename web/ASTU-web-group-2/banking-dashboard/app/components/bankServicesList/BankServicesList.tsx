import React from "react";
import BankService from "../bankService/BankService";
import BankServiceMobile from "../bankService/BankServiceMobile";

const BankServicesList = () => {
  const data = [
    {
      category: "Business Loans",
      description: "It is a long established",
      logo: "/assets/bankService/businessLoans.svg",
      details: [
        {
          title: "Loan Option 1",
          subtitle: "Many publishing",
        },
        {
          title: "Loan Option 2",
          subtitle: "Many publishing",
        },
        {
          title: "Loan Option 3",
          subtitle: "Many publishing",
        },
        {
          title: "Loan Option 4",
          subtitle: "Many publishing",
        },
      ],
      action: {
        label: "View Details",
        link: "/details/business-loans",
      },
    },
    {
      category: "Checking Accounts",
      description: "It is a long established",
      logo: "/assets/bankService/checkAccounts.svg",
      details: [
        {
          title: "Account Option A",
          subtitle: "Many publishing",
        },
        {
          title: "Account Option B",
          subtitle: "Many publishing",
        },
        {
          title: "Account Option C",
          subtitle: "Many publishing",
        },
        {
          title: "Account Option D",
          subtitle: "Many publishing",
        },
      ],
      action: {
        label: "View Details",
        link: "/details/checking-accounts",
      },
    },
    {
      category: "Savings Accounts",
      description: "It is a long established",
      logo: "/assets/bankService/savingAccounts.svg",
      details: [
        {
          title: "Savings Plan 1",
          subtitle: "Many publishing",
        },
        {
          title: "Savings Plan 2",
          subtitle: "Many publishing",
        },
        {
          title: "Savings Plan 3",
          subtitle: "Many publishing",
        },
        {
          title: "Savings Plan 4",
          subtitle: "Many publishing",
        },
      ],
      action: {
        label: "View Details",
        link: "/details/savings-accounts",
      },
    },
    {
      category: "Debit and Credit Cards",
      description: "It is a long established",
      logo: "/assets/bankService/debitCredit.svg",
      details: [
        {
          title: "Card Plan 1",
          subtitle: "Earn rewards",
        },
        {
          title: "Card Plan 2",
          subtitle: "Earn rewards",
        },
        {
          title: "Card Plan 3",
          subtitle: "Earn rewards",
        },
        {
          title: "Card Plan 4",
          subtitle: "Earn rewards",
        },
      ],
      action: {
        label: "View Details",
        link: "/details/debit-credit-cards",
      },
    },
    {
      category: "Life Insurance",
      description: "It is a long established",
      logo: "/assets/bankService/lifeInsurance.svg",
      details: [
        {
          title: "Insurance Plan 1",
          subtitle: "Many publishing",
        },
        {
          title: "Insurance Plan 2",
          subtitle: "Many publishing",
        },
        {
          title: "Insurance Plan 3",
          subtitle: "Many publishing",
        },
        {
          title: "Insurance Plan 4",
          subtitle: "Many publishing",
        },
      ],
      action: {
        label: "View Details",
        link: "/details/life-insurance",
      },
    },
    {
      category: "Business Loans",
      description: "It is a long established",
      logo: "/assets/bankService/businessLoans.svg",
      details: [
        {
          title: "Loan Option 1",
          subtitle: "Many publishing",
        },
        {
          title: "Loan Option 2",
          subtitle: "Many publishing",
        },
        {
          title: "Loan Option 3",
          subtitle: "Many publishing",
        },
        {
          title: "Loan Option 4",
          subtitle: "Many publishing",
        },
      ],
      action: {
        label: "View Details",
        link: "/details/business-loans",
      },
    },
  ];

  return (
    <div>
      <div className="flex flex-col gap-5 max-md:hidden">
        {data.map((bankService) => (
          <BankService
            logoLink={bankService.logo}
            category={bankService.category}
            description={bankService.description}
            details={bankService.details}
            action={bankService.action}
          />
        ))}
      </div>

      {/* Mobile view */}
      <div className="flex flex-col gap-5 md:hidden">
        {data.map((bankService) => (
          <BankServiceMobile
            logoLink={bankService.logo}
            category={bankService.category}
            description={bankService.description}
            details={bankService.details}
            action={bankService.action}
          />
        ))}
      </div>
    </div>
  );
};

export default BankServicesList;
