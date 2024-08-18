import Image from "next/image";
import React from "react";
import TableButton from "../TableButton/TableButton";
const BankServices = [
  {
    icon: "/assets/icons/business-loans.svg",
    "first-col": ["Business Loans", "it is a long established"],
    "second-col": ["lorem Ipsum", "Many Publishing"],
    "third-col": ["lorem Ipsum", "Many Publishing"],
    "fourth-col": ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/checking-account.svg",
    "first-col": ["Checking Accounts", "it is a long established"],
    "second-col": ["lorem Ipsum", "Many Publishing"],
    "third-col": ["lorem Ipsum", "Many Publishing"],
    "fourth-col": ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/saving-accounts.svg",
    "first-col": ["Saving Accounts", "it is a long established"],
    "second-col": ["lorem Ipsum", "Many Publishing"],
    "third-col": ["lorem Ipsum", "Many Publishing"],
    "fourth-col": ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/debit-and-credit-card.svg",
    "first-col": ["Debit and Credit Cards", "it is a long established"],
    "second-col": ["lorem Ipsum", "Many Publishing"],
    "third-col": ["lorem Ipsum", "Many Publishing"],
    "fourth-col": ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/life-insurance.svg",
    "first-col": ["Life Insurance", "it is a long established"],
    "second-col": ["lorem Ipsum", "Many Publishing"],
    "third-col": ["lorem Ipsum", "Many Publishing"],
    "fourth-col": ["lorem Ipsum", "Many Publishing"],
  },
  {
    icon: "/assets/icons/business-loans.svg",
    "first-col": ["Business Loans", "it is a long established"],
    "second-col": ["lorem Ipsum", "Many Publishing"],
    "third-col": ["lorem Ipsum", "Many Publishing"],
    "fourth-col": ["lorem Ipsum", "Many Publishing"],
  },
];

const BankServicesList = () => {
  return (
    <div className="flex flex-col gap-4">
      <p className='text-[#333B69] pb-3 font-semibold'>Bank Services List</p>
      <div className="flex flex-col gap-3">
        {BankServices.map((items, index) => (
          <div key={index} className="grid grid-cols-7 md:grid-cols-7 xl:grid-cols-11 items-center bg-white p-3 rounded-xl">
            <div className="flex col-span-5 md:col-span-3 items-center gap-3 lg:gap-4">
              <div className="flex relative w-9 h-9 lg:w-12 lg:h-12 items-center">
                <Image src={items.icon} alt="" fill />
              </div>
              <div className="flex flex-col">
                <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                  {items["first-col"][0]}
                </span>
                <span className="text-12px lg:text-15px text-blue-steel">
                  {items["first-col"][1]}
                </span>
              </div>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden">
              <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                {items["second-col"][0]}
              </span>
              <span className="text-12px lg:text-15px text-blue-steel">
                {items["second-col"][1]}
              </span>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden">
              <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                {items["third-col"][0]}
              </span>
              <span className="text-12px lg:text-15px text-blue-steel">
                {items["third-col"][1]}
              </span>
            </div>
            <div className="flex-col  md:flex col-span-2 hidden">
              <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                {items["fourth-col"][0]}
              </span>
              <span className="text-12px lg:text-15px text-blue-steel">
                {items["fourth-col"][0]}
              </span>
            </div>
            <div className="flex-col  flex col-span-2 items-end md:items-start">
              <div className="">
                <TableButton
                  text="View Details"
                  classname="hidden md:flex px-6 lg:text-[14px]"
                />
              </div>
              <span className="font-medium text-12px lg:text-15px text-blue-bright md:hidden">
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
