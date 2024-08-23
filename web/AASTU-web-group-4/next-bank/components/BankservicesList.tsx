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
import Pagination from "./Pagination";


const icons = [
  BusinessLoansIcon,
  CheckingAccountsIcon,
  SavingAccountsIcon,
  DebitCreditIcon,
  SafetyIcon,
];



const ITEMS_PER_PAGE = 5;

const BankservicesList: React.FC = () => {

  const [currentPage, setCurrentPage] = useState(1);
  const[filtered , setfiltered] = useState([])
  const [services , setservices] = useState([])
  const randomIcons = icons;
  const [totalPages, setTotalPages] = useState(0);
  useEffect(() => {
    const fetch = async () => {
   
      try {
        const response= await getAllBankServices( currentPage, ITEMS_PER_PAGE);
        console.log("response:" , response)
        setservices(response.data.content || []);
        setfiltered(response.data.content || []);
        setTotalPages(response.data.totalPages);
        
      } catch (error) {
        console.error(' Error fetching bank services:', error);
      }
    };
    fetch();
  }, [currentPage]);
  
  const Filter = (event:any) =>{
    setfiltered(services.filter((service:any) => service.name.toLowerCase().includes(event.target.value.toLowerCase())))
  }

  return (
    <div className="max-w-[1110px] px-4 md:mx-auto ">
      <h2 className="text-xl font-bold mb-4  dark:text-blue-500">Bank Services List</h2>
      <input type="text" className="form-control py-4 border-none" onChange={Filter} placeholder="search" />
      {filtered.map((service: any, index: any ) => (
        <div key={index} className="mb-4">
          {/* Mobile View */}
          <div className="lg:hidden">
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
          <div>
         
          </div>
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
       <Pagination
             currentPage={currentPage}
             totalPages={totalPages}
             onPageChange={setCurrentPage}
          />
    </div>
  );
};

export default BankservicesList;
