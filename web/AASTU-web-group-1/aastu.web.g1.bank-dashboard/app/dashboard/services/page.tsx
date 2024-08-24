"use client";

import { servicesList } from "@/constants/index";
import { useUser } from "@/contexts/UserContext";
import { getSession } from "next-auth/react";
import { useEffect, useState } from "react";
import BenefitComp from "./serviceComponenet/BenefitComp";
import ServiceList from "./serviceComponenet/ServiceList";
import { Loading } from "../_components/Loading";

// Type definition for a single bank service
interface BankService {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

// Type definition for the response from the API
interface BankServicesResponse {
  success: boolean;
  message: string;
  data: {
    content: BankService[];
    totalPages: number;
  };
}

const Services = () => {
  const [bankServices, setBankServices] = useState<BankService[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [currentPage, setCurrentPage] = useState<number>(0); // New state for current page
  const [totalPages, setTotalPages] = useState<number>(0); // New state for total pages

  useEffect(() => {
    const fetchBankServices = async () => {
      const session = await getSession();
      const token = session?.user?.accessToken;

      try {
        const response: BankServicesResponse = await fetch(
          `${process.env.NEXT_PUBLIC_BASE_URL}/bank-services?page=${currentPage}&size=4`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        ).then((res) => res.json());

        console.log(response, "this is the data ");
        if (response.success) {
          setBankServices(response.data.content);
          setTotalPages(response.data.totalPages); // Set the total number of pages
        }
      } catch (error) {
        console.error("Error fetching bank services:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchBankServices();
  }, [currentPage]); // Effect triggers when currentPage changes

  const { isDarkMode } = useUser();

  const handlePageChange = (newPage: number) => {
    if (newPage >= 0 && newPage < totalPages) {
      setLoading(true); // Show loading spinner while fetching data
      setCurrentPage(newPage);
    }
  };
  const pages = [];
  for (let i = 0; i < totalPages; i++) {
    pages[i] = i;
  }

  return loading ? (
    <Loading />
  ) : (
    <div
      className={`p-4 flex flex-col w-full gap-5 lg:p-8 ${
        isDarkMode ? "bg-gray-800 text-gray-300" : "text-gray-900"
      }`}
    >
      <div className="flex gap-3 justify-start w-[100vw] md:w-full overflow-x-scroll scrollbar-hidden md:overflow-hidden lg:justify-start lg:gap-16 ">
        {servicesList.map((items, index) => (
          <BenefitComp items={items} key={index} />
        ))}
      </div>

      <div className="flex flex-col gap-[7px] md:gap-4">
        <h1
          className={`font-semibold text-lg ${
            isDarkMode ? "text-gray-300" : "text-gray-900"
          }`}
        >
          Bank Services List
        </h1>
        {bankServices.map((items, index) => (
          <ServiceList
            icon={items.icon}
            name={items.name}
            details={items.details}
            key={index}
          />
        ))}
      </div>

      {/* Pagination Controls */}
      <div className="flex justify-center gap-4 mt-4">
        <button
          onClick={() => handlePageChange(currentPage - 1)}
          disabled={currentPage === 0}
          className={`px-4 py-2 rounded ${
            isDarkMode
              ? "bg-gray-700 text-gray-300"
              : "bg-gray-200 text-gray-900"
          } ${currentPage === 0 && "opacity-50 cursor-not-allowed"}`}
        >
          Previous
        </button>
        <span className="self-center">
          Page {currentPage + 1} of {totalPages}
          {pages.map((page) => (
            <button
              key={page}
              onClick={() => handlePageChange(page)}
              className={`px-4 py-2 rounded ${
                isDarkMode
                  ? "bg-gray-700 text-gray-300"
                  : "bg-gray-200 text-gray-900"
              } ${currentPage === page - 1 && "opacity-50 cursor-not-allowed"}`}
            >
              {page}
            </button>
          ))}
        </span>
        <button
          onClick={() => handlePageChange(currentPage + 1)}
          disabled={currentPage === totalPages - 1}
          className={`px-4 py-2 rounded ${
            isDarkMode
              ? "bg-gray-700 text-gray-300"
              : "bg-gray-200 text-gray-900"
          } ${
            currentPage === totalPages - 1 && "opacity-50 cursor-not-allowed"
          }`}
        >
          Next
        </button>
      </div>
    </div>
  );
};

export default Services;
