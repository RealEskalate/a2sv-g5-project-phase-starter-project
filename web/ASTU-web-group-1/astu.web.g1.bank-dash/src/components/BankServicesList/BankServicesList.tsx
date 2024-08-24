"use client";
import Image from "next/image";
import React from "react";
import TableButton from "../TableButton/TableButton";
import {
  useGetBankServiceByIdQuery,
  useGetBankServiceQuery,
} from "@/lib/redux/slices/bankService";

const BankServicesList = () => {
  const page = 0;
  const size = 6;
  const { data, error, isLoading } = useGetBankServiceQuery({ page, size });
  console.log("data is ", data, error);

  if (isLoading) {
    return <p>Loading...</p>; // Add a loading indicator
  }

  if (error || !data || !data.data) {
    return <p>An error occurred</p>; // Add error handling
  }
  const contents = data.data.content;

  console.log("the contents", contents);
  console.log("the icon path", contents);

  return (
    <div className="flex flex-col gap-4">
      <p className="text-[#333B69] pb-3 font-semibold">Bank Services List</p>
      <div className="flex flex-col gap-3">
        {contents?.map((items, index) => (
          <div
            key={index}
            className="grid grid-cols-7 md:grid-cols-7 xl:grid-cols-11 items-center bg-white p-3 rounded-xl"
          >
            <div className="flex col-span-5 md:col-span-3 items-center gap-3 lg:gap-4">
              <div className="flex relative w-9 h-9 lg:w-12 lg:h-12 items-center">
                <Image
                  src={`/assets/icons/services${(index % 5) + 1}.svg`}
                  alt=""
                  fill
                />
              </div>
              <div className="flex flex-col">
                <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                  {items.name}
                </span>
                <span className="text-12px lg:text-15px text-blue-steel">
                  {items.details}
                </span>
              </div>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden">
              <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                Number of users
              </span>
              <span className="text-12px lg:text-15px text-blue-steel">
                {items.numberOfUsers}
              </span>
            </div>
            <div className="flex-col  xl:flex col-span-2 hidden">
              <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                Service Status
              </span>
              <span className="text-12px lg:text-15px text-blue-steel">
                {items.status}
              </span>
            </div>
            <div className="flex-col  md:flex col-span-2 hidden">
              <span className="text-13px md:text-13px lg:text-16px font-medium text-gray-dark">
                Service Type
              </span>
              <span className="text-12px lg:text-15px text-blue-steel">
                {items.type}
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
