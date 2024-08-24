"use client";
import React, { useEffect, useState } from "react";
import Link from "next/link";
import { getAllBankServices } from "@/services/bankseervice";
import BusinessLoansIcon from "@/public/icons/BusinessLoansIcon";
import CheckingAccountsIcon from "@/public/icons/CheckingAccountsIcon";
import SavingAccountsIcon from "@/public/icons/SavingAccountsIcon";
import DebitCreditIcon from "@/public/icons/DebitCreditIcon";
import SafetyIcon from "@/public/icons/SafetyIcon";
import Pagination from "./Pagination";
import { TbFileSad } from "react-icons/tb";
import { colors } from "@/constants";
import Cookie from "js-cookie";


const icons = [
  BusinessLoansIcon,
  CheckingAccountsIcon,
  SavingAccountsIcon,
  DebitCreditIcon,
  SafetyIcon,
];

const renderShimmer = (count: number) => {

  const shimmers = [];

  for (let i = 0; i < count; i++) {
    shimmers.push(
      <div key={i}>
        {/* Mobile View */}
        <div className="lg:hidden shadow-lg p-4 rounded-md flex items-center justify-between animate-pulse mb-4">
          <div className="flex items-center space-x-4">
            <div className="w-13 h-13 bg-gray-300 rounded-full"></div>
            <div>
              <div className="h-4 bg-gray-300 rounded w-32 mb-2"></div>
              <div className="h-3 bg-gray-200 rounded w-24"></div>
            </div>
          </div>
          <div className="h-3 bg-gray-200 rounded w-20"></div>
        </div>

        {/* Desktop View */}
        <div
          className="hidden lg:flex shadow-lg p-4 rounded-md items-center animate-pulse mb-4"
          style={{ width: "1110px", height: "90px" }}
        >
          <div className="w-13 h-13 bg-gray-300 rounded-full"></div>
          <div className="flex-1 ml-3">
            <div className="flex justify-between">
              <div>
                <div className="h-4 bg-gray-300 rounded w-32 mb-2"></div>
                <div className="h-3 bg-gray-200 rounded w-24"></div>
              </div>
              <div className="flex space-x-28">
                <div>
                  <div className="h-4 bg-gray-300 rounded w-20 mb-2"></div>
                  <div className="h-3 bg-gray-200 rounded w-16"></div>
                </div>
                <div>
                  <div className="h-4 bg-gray-300 rounded w-20 mb-2"></div>
                  <div className="h-3 bg-gray-200 rounded w-16"></div>
                </div>
                <div>
                  <div className="h-4 bg-gray-300 rounded w-20 mb-2"></div>
                  <div className="h-3 bg-gray-200 rounded w-16"></div>
                </div>
              </div>
            </div>
          </div>
          <div className="h-3 bg-gray-200 rounded w-28"></div>
        </div>
      </div>
    );
  }

  return shimmers;
};


const ITEMS_PER_PAGE = 10;

const BankservicesList: React.FC = () => {

  const [currentPage, setCurrentPage] = useState(0);
  const[filtered , setfiltered] = useState([])
  const [services , setservices] = useState([])
  const randomIcons = icons;
  const [totalPages, setTotalPages] = useState(0);
  // const token = Cookie.get("accessToken") || 'null'

  
  const [status, setStatus] = useState<"loading" | "error" | "success">(
    "loading"
  );
  const filter = (e: any) => {
    const keyword = e.target.value;
    if (keyword !== "") {
      const results = services.filter((service: any) => {
        return service.name.toLowerCase().startsWith(keyword.toLowerCase());
      });
      setfiltered(results);
    } else {
      setfiltered(services);}}

  useEffect(() => {
    const fetchData = async () => {
      setStatus("loading");

      try {
        const response= await getAllBankServices( currentPage, ITEMS_PER_PAGE);
        console.log("response:" , response)
        if (response.success === true) {
          setStatus("success");
          
        setservices(response.data.content || []);
        setfiltered(response.data.content || []);
        setTotalPages(response.data.totalPages);
        
      }
     } catch (error) {
        setStatus("error");
        console.error("Error fetching bank services:", error);
      }
    };
    fetchData();
  },[]);

  if (status === "loading") {
    return (
      <div className="max-w-[1110px] px-4 md:mx-auto">
        <h2 className="text-xl font-bold mb-4 dark:text-blue-500 animate-pulse">
          Bank Services List
        </h2>
        {renderShimmer(3)}
      </div>
    );
  } else if (status === "error") {
    return (
      <div className="max-w-[1110px] px-4 md:mx-auto">
        <h2 className="text-xl font-bold mb-4 dark:text-blue-500 animate-pulse">
          Bank Services List
        </h2>
        <div className="text-xl w-[100%] text-center gap-4 flex flex-col items-center  font-bold mb-4 text-red-500">
        <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
          <div> Failed to fetch the data</div>
        </div>
      </div>
    );
  } else if (status === "success") {
    return (
      <>
        {services.length == 0 ? (
          <div className="max-w-[1110px] px-4 md:mx-auto mt-4">
            <div className="shadow-lg p-4 rounded-md flex items-center justify-between bg-gray-100">
              <div className="flex items-center space-x-4">
                <div className="w-13 h-13 bg-gray-300 rounded-full"></div>
                <div>
                  <h3 className="text-[16px] font-semibold text-gray-700">
                    No Data Available
                  </h3>
                  <p className="text-[14px] text-gray-500">
                    There are no bank services to display at the moment.
                  </p>
                </div>
              </div>
            </div>
          </div>
        
      )
  
        : (
          <div className="max-w-[1110px] px-4 md:mx-auto">
            <h2 className="text-xl font-bold mb-4 dark:text-blue-500">
              Bank Services List
            </h2>
            <input type="text" onChange={filter} placeholder="search" />
            {filtered.map((service: any, index: any) => (
              <div key={index} className="mb-4">
             
                {/* Mobile View */}
                <div className="lg:hidden shadow-lg p-4 rounded-md flex items-center justify-between">
                  <div className="flex items-center space-x-4">
                    {icons[index % icons.length] &&
                      React.createElement(icons[index % icons.length], {
                        className: "w-13 h-13",
                        "aria-hidden": "true",
                      })}
                    <div>
                      <h3 className="text-[14px] font-semibold">
                        {service.name}
                      </h3>
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
                  {icons[index % icons.length] &&
                    React.createElement(icons[index % icons.length], {
                      className: "w-13 h-13",
                      "aria-hidden": "true",
                    })}
                  <div className="flex-1 ml-3">
                    <div className="flex justify-between">
                      <div>
                        <h3 className="text-[16px] font-semibold">
                          {service.name}
                        </h3>
                        <p className="text-[15px] text-gray-500">
                          {service.details}
                        </p>
                      </div>
                      <div className="flex space-x-28">
                        <div>
                          <h4 className="text-[14px] font-semibold">
                            {service.type}
                          </h4>
                          <p className="text-[12px] text-gray-500">type</p>
                        </div>
                        <div>
                          <h4 className="text-[14px] font-semibold">
                            {service.status}
                          </h4>
                          <p className="text-[12px] text-gray-500">status</p>
                        </div>
                        <div>
                          <h4 className="text-[14px] font-semibold">
                            {service.numberOfUsers}
                          </h4>
                          <p className="text-[12px] text-gray-500">
                            number of users
                          </p>
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
          </div>)
  }
   </>
  )
  }

    
  return null; // fallback, though it shouldn't reach here
};

export default BankservicesList;
