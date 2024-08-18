import React from "react";
import Link from "next/link";
import BusinessLoansIcon from "@/public/icons/BusinessLoansIcon";
import CheckingAccountsIcon from "@/public/icons/CheckingAccountsIcon";
import SavingAccountsIcon from "@/public/icons/SavingAccountsIcon";
import DebitCreditIcon from "@/public/icons/DebitCreditIcon";
import SafetyIcon from "@/public/icons/SafetyIcon";

const bankServices = [
  {
    title: "Business Loans",
    description: "It is a long established",
    icon: BusinessLoansIcon,
  },
  {
    title: "Checking Accounts",
    description: "It is a long established",
    icon: CheckingAccountsIcon,
  },
  {
    title: "Saving Accounts",
    description: "It is a long established",
    icon: SavingAccountsIcon,
  },
  {
    title: "Debit and Credit Cards",
    description: "It is a long established",
    icon: DebitCreditIcon,
  },
  {
    title: "Life Insurance",
    description: "It is a long established",
    icon: SafetyIcon,
  },
  {
    title: "Business Loans",
    description: "It is a long established",
    icon: BusinessLoansIcon,
  },
];

const BankservicesList: React.FC = () => {
  return (
    <div className="max-w-[1110px] mx-auto">
      <h2 className="text-xl font-bold mb-4">Bank Services List</h2>
      {bankServices.map((service, index) => (
        <div key={index} className="mb-4">
          {/* Mobile View */}
          <div className="lg:hidden shadow-lg p-4 rounded-md flex items-center justify-between">
            <div className="flex items-center space-x-4">
              <service.icon className="w-13 h-13" aria-hidden="true" />
              <div>
                <h3 className="text-[14px] font-semibold">{service.title}</h3>
                <p className="text-[12px] text-gray-500">
                  {service.description}
                </p>
              </div>
            </div>
            <Link href="/details" className="text-[12px] text-blue-600">
              View Details
            </Link>
          </div>

          {/* Larger Screens */}
          <div
            className="hidden lg:flex shadow-lg p-4 rounded-md items-center"
            style={{ width: "1110px", height: "90px" }}
          >
            <service.icon
              className="w-130 h-130 flex-shrink-0"
              aria-hidden="true"
            />
            <div className="flex-1 ml-3">
              <div className="flex justify-between">
                <div>
                  <h3 className="text-[16px] font-semibold">{service.title}</h3>
                  <p className="text-[15px] text-gray-500">
                    {service.description}
                  </p>
                </div>
                <div className="flex space-x-28">
                  <div>
                    <h4 className="text-[14px] font-semibold">Lorem Ipsum</h4>
                    <p className="text-[12px] text-gray-500">Many publishing</p>
                  </div>
                  <div>
                    <h4 className="text-[14px] font-semibold">Lorem Ipsum</h4>
                    <p className="text-[12px] text-gray-500">Many publishing</p>
                  </div>
                  <div>
                    <h4 className="text-[14px] font-semibold">Lorem Ipsum</h4>
                    <p className="text-[12px] text-gray-500">Many publishing</p>
                  </div>
                </div>
              </div>
            </div>
            <Link
              href="/details"
              className="text-[15px] text-blue-600 ml-28 border border-blue-600 px-2 py-1 rounded-full"
            >
              View Details
            </Link>
          </div>
        </div>
      ))}
    </div>
  );
};

export default BankservicesList;
