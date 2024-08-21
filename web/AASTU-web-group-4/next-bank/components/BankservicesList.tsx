'use client'
import React, { useEffect, useState } from "react";
import Link from "next/link";
import { getAllBankServices } from "@/services/bankseervice";
import { FaIceCream } from "react-icons/fa";
import BusinessLoansIcon from "@/public/icons/BusinessLoansIcon";
import CheckingAccountsIcon from "@/public/icons/CheckingAccountsIcon";
import SavingAccountsIcon from "@/public/icons/SavingAccountsIcon";
import DebitCreditIcon from "@/public/icons/DebitCreditIcon";
import SafetyIcon from "@/public/icons/SafetyIcon";

const icons = [
  BusinessLoansIcon,
  CheckingAccountsIcon,
  SavingAccountsIcon,
  DebitCreditIcon,
  SafetyIcon,
];


const BankservicesList: React.FC = () => {
  const [services , setservices] = useState([])
  const randomIcons = icons;
  useEffect(() => {
    const fetch = async () => {
   
      try {
        const response= await getAllBankServices();
        console.log("response:" , response)
        setservices(response.data.content || [])
        
      } catch (error) {
        console.error(' Error feetching bank services:', error);
      }
    };
    fetch();
  }, [])
  

  return (
    <div className="max-w-[1110px] mx-auto">
      <h2 className="text-xl font-bold mb-4">Bank Services List</h2>
      {services.map((service: any, index: any) => (
        <div key={index} className="mb-4">
          {/* Mobile View */}
          <div className="lg:hidden shadow-lg p-4 rounded-md flex items-center justify-between">
            <div className="flex items-center space-x-4">
            {randomIcons[index % randomIcons.length] && (
                  React.createElement(randomIcons[index % randomIcons.length], {
                    className: "w-13 h-13",
                    "aria-hidden": "true"
                  })
                )}
              <div>
                <h3 className="text-[14px] font-semibold">{service.name}</h3>
                <p className="text-[12px] text-gray-500">
                  {service.details}
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
            {randomIcons[index % randomIcons.length] && (
                  React.createElement(randomIcons[index % randomIcons.length], {
                    className: "w-13 h-13",
                    "aria-hidden": "true"
                  })
                )}
            <div className="flex-1 ml-3">
              <div className="flex justify-between">
                <div>
                  <h3 className="text-[16px] font-semibold">{service.name}</h3>
                  <p className="text-[15px] text-gray-500">
                    {service.details}
                  </p>
                </div>
                <div className="flex space-x-28">
                  <div>
                    <h4 className="text-[14px] font-semibold">{service.type}</h4>
                    <p className="text-[12px] text-gray-500">type</p>
                  </div>
                  <div>
                    <h4 className="text-[14px] font-semibold">{service.status}</h4>
                    <p className="text-[12px] text-gray-500">status</p>
                  </div>
                  <div>
                    <h4 className="text-[14px] font-semibold">{service.numberOfUsers}</h4>
                    <p className="text-[12px] text-gray-500">number of users</p>
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
