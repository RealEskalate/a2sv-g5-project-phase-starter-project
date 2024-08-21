"use client";
import React, { useEffect, useState } from "react";
import DescriptionCard from "@/app/components/Card/DescriptionCard";
import ServicesCard from "@/app/components/Card/ServicesCard";
import axios from "axios";
import { useSession } from "next-auth/react";
import ModalService from "@/app/components/Card/ModalService";

interface BankService {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
  colors: string;
}

const Services = () => {
  const { data: session } = useSession();
  const colors = [
    "bg-orange-100",
    "bg-pink-100",
    "bg-blue-100",
    "bg-green-100",
    "bg-pink-100",
  ];
  const [services, setServices] = useState<BankService[]>([]);
  const [pageNumber, setPageNumber] = useState(1);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleModalToggle = () => {
    setIsModalOpen(!isModalOpen);
  };

  const accessToken = session?.accessToken;
  async function fetchData(accessToken: string) {
    try {
      const response = await axios.get(
        `https://bank-dashboard-1tst.onrender.com/bank-services?page=0&size=50`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );
      setServices(response.data.data.content);
      console.log(services);
    } catch (error) {
      console.error("There was a problem with the axios request:", error);
    }
  }

  useEffect(() => {
    fetchData(accessToken);
  }, []);

  return (
    <div className="w-[96%] xxs:pt-4 xs:pt-20 md:pt-5 lg:pt-0">
      <div className="ml-5 lg:ml-0 ">
        <div className="mr-5 lg:mr-0 flex gap-4 overflow-x-auto lg:overflow-x-visible [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none] lg:pl-10 lg:pt-10 w-full">
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/lifeInsurance.svg"
            title="Life Insurance"
            desc="Unlimited Protection"
          />
          {/* </div> */}
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/shoppingBag.svg"
            title="Shopping"
            desc="Buy. Think. Grow"
          />
          {/* </div> */}
          {/* <div className="w-[100%] lg:w-[350px] flex-shrink-0"> */}
          <ServicesCard
            img="/assets/safety.svg"
            title="Safety"
            desc="We are your allies"
          />
        </div>
        {/* </div> */}

        <div>
          <div className="flex justify-between">
            <p className="font-semibold text-[22px] text-[#343C6A] pt-5 pb-5 lg:p-10 ">
              Bank Services List
            </p>
            <div
              className={`flex items-center text-base text-[#718EBF] dark:bg-gray-700 dark:text-gray-400 rounded-[50px] py-1 pl-6 grow justify-end ${
                isModalOpen ? "blur-sm" : ""
              }`}
            >
              <button
                onClick={handleModalToggle}
                className="bg-blue-600 text-white p-3 rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                type="button"
              >
                Add
              </button>
            </div>
            {isModalOpen && (
              <div
                className="fixed inset-0 z-50 flex justify-center items-center bg-black bg-opacity-50 backdrop-blur-sm"
                onClick={handleModalToggle}
              >
                <div
                  className="relative bg-white p-6 rounded-lg shadow-lg max-w-lg w-full"
                  onClick={(e) => e.stopPropagation()} // Prevent modal from closing when clicking inside it
                >
                  <ModalService
                    isOpen={isModalOpen}
                    onClose={handleModalToggle}
                  />
                </div>
              </div>
            )}
          </div>
          <div className="w-full flex flex-col grow items-start px-4">
            {services.length > 0 ? (
              services.map((service, index) => (
                <DescriptionCard
                  key={service.id}
                  img={service.icon}
                  title={service.name}
                  desc={service.details}
                  colOne="Number of Users"
                  descOne={service.numberOfUsers}
                  colTwo="Status"
                  descTwo={service.status}
                  colThree="Type"
                  descThree={service.type}
                  btn="View Details"
                  color={colors[index]}
                />
              ))
            ) : (
              <p className="pl-10">No services available</p>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Services;
