"use client";
import React, { useState, useEffect, Suspense } from "react";
import ServiceList from "./serviceComponenet/ServiceList";
import BenefitComp from "./serviceComponenet/BenefitComp";
import { bankServices } from "@/constants/index";
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
          "https://bank-dashboard-6acc.onrender.com/bank-services?page=0&size=10",
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
      className={`p-4 flex  flex-col   ${
        isDarkMode ? "bg-gray-700" : "bg-gray-200"
      } w-full h-full gap-5 lg:p-8`}
    >
      <div className=" flex gap-3 justify-start w-[100vw] md:w-full overflow-x-scroll scrollbar-hidden md:overflow-hidden lg:justify-between">
        {servicesList.map((items, index) => (
          <BenefitComp items={items} key={index} />
        ))}
      </div>

      <div className="flex flex-col gap-[7px]">
        <h1 className="font-semibold text-lg text-[#343C6A]">
          Bank Services List
        </h1>
        {loading ? (
          <div>Loading...</div>
        ) : (
          bankServices.map((items, index) => (
            <ServiceList
              icon={items.icon} // Use the icon provided in the data
              name={items.name}
              key={index}
            />
          ))
        )}
      </div>
    </div>
  );
};

export default Services;
