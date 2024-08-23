"use client";

import React, { useState, useEffect } from "react";
import ServiceList from "./serviceComponenet/ServiceList";
import BenefitComp from "./serviceComponenet/BenefitComp";
import { servicesList } from "@/constants/index";
import { getSession } from "next-auth/react";
import { useUser } from "@/contexts/UserContext";

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
  data: BankService[];
}

const Services = () => {
  const [bankServices, setBankServices] = useState<BankService[]>([]);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const fetchBankServices = async () => {
      const session = await getSession();
      const token = session?.user?.accessToken;
      console.log(token);
      try {
        const response = await fetch(

          `${process.env.BASE_URL}/bank-services?page=0&size=10`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        ).then((res) => res.json());

        console.log(response, "this is the data ");
        if (response.success) {
          setBankServices(response.data.content);
        }
      } catch (error) {
        console.error("Error fetching bank services:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchBankServices();
  }, []);

  const { isDarkMode } = useUser();

  return (
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
        {loading ? (
          <div>Loading...</div>
        ) : (
          bankServices.map((items, index) => (
            <ServiceList
              icon={items.icon}
              name={items.name}
              details={items.details}
              key={index}
            />
          ))
        )}
      </div>
    </div>
  );
};

export default Services;
