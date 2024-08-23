import Image from "next/image";
import React from "react";
import TableButton from "../TableButton/TableButton";
export const BankServices = [
  {
    icon: "/assets/icons/business-loans.svg",
    firstcol: ["Business Loans", "it is a long established"],
    secondcol: ["lorem Ipsum", "Many Publishing"],
    thirdcol: ["lorem Ipsum", "Many Publishing"],
    fourthcol: ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/checking-account.svg",
    firstcol: ["Checking Accounts", "it is a long established"],
    secondcol: ["lorem Ipsum", "Many Publishing"],
    thirdcol: ["lorem Ipsum", "Many Publishing"],
    fourthcol: ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/saving-accounts.svg",
    firstcol: ["Saving Accounts", "it is a long established"],
    secondcol: ["lorem Ipsum", "Many Publishing"],
    thirdcol: ["lorem Ipsum", "Many Publishing"],
    fourthcol: ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/debit-and-credit-card.svg",
    firstcol: ["Debit and Credit Cards", "it is a long established"],
    secondcol: ["lorem Ipsum", "Many Publishing"],
    thirdcol: ["lorem Ipsum", "Many Publishing"],
    fourthcol: ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/life-insurance.svg",
    firstcol: ["Life Insurance", "it is a long established"],
    secondcol: ["lorem Ipsum", "Many Publishing"],
    thirdcol: ["lorem Ipsum", "Many Publishing"],
    fourthcol: ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/business-loans.svg",
    firstcol: ["Business Loans", "it is a long established"],
    secondcol: ["lorem Ipsum", "Many Publishing"],
    thirdcol: ["lorem Ipsum", "Many Publishing"],
    fourthcol: ["lorem Ipsum", "Many Publishing"],
  },
];

const BankServicesList = () => {
  return (
    <div className="flex flex-col gap-4">
      <p className="text-[#333B69] pb-3 font-semibold">Bank Services List</p>
      <div className="flex flex-col gap-3">
        {BankServices.map((items, index) => (
          <div
            key={index}
            className="grid grid-cols-7 md:grid-cols-7 xl:grid-cols-11 items-center bg-white p-3 rounded-xl"
          >
            <div className="flex col-span-5 md:col-span-3 items-center gap-3 lg:gap-4">
              <div className="flex relative w-9 h-9 lg:w-12 lg:h-12 items-center">
                <Image
                  src={items.icon}
                  alt=""
                  fill
                  data-testid={`icon-${index}`}
                />
              </div>
              <div className="flex flex-col">
                <span
                  data-testid={`service-name-${index}`}
                  className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark"
                >
                  {items.firstcol[0]}
                </span>
                <span
                  data-testid={`service-description-${index}`}
                  className="text-12px lg:text-15px text-blue-steel"
                >
                  {items.firstcol[1]}
                </span>
              </div>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden">
              <span
                data-testid={`second-col-name-${index}`}
                className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark"
              >
                {items.secondcol[0]}
              </span>
              <span
                data-testid={`second-col-description-${index}`}
                className="text-12px lg:text-15px text-blue-steel"
              >
                {items.secondcol[1]}
              </span>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden">
              <span
                data-testid={`third-col-name-${index}`}
                className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark"
              >
                {items.thirdcol[0]}
              </span>
              <span
                data-testid={`third-col-description-${index}`}
                className="text-12px lg:text-15px text-blue-steel"
              >
                {items.thirdcol[1]}
              </span>
            </div>
            <div className="flex-col  md:flex col-span-2 hidden">
              <span
                data-testid={`fourth-col-name-${index}`}
                className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark"
              >
                {items.fourthcol[0]}
              </span>
              <span
                data-testid={`fourth-col-description-${index}`}
                className="text-12px lg:text-15px text-blue-steel"
              >
                {items.fourthcol[1]}
              </span>
            </div>
            <div className="flex-col  flex col-span-2 items-end md:items-start">
              <div className="">
                <TableButton
                  data-testid={`view-details-button-${index}`}
                  text="View Details"
                  classname="hidden md:flex px-6 lg:text-[14px]"
                />
              </div>
              <span
                data-testid={`mobile-view-details-${index}`}
                className="font-medium text-12px lg:text-15px text-blue-bright md:hidden"
              >
                View Details
              </span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BankServicesList;
