"use client";
import React from "react";
import { useGetAllServicesQuery } from "@/lib/services/services";
import { BankService } from "@/types";
import Image from "next/image";



const ServiceList = () => {
  const { data, isError, isLoading, error, isSuccess } = useGetAllServicesQuery(
    {
      accessToken:?
      page:? ,
      size: ?,
    }
  );

  if (isLoading) return <div>Loading...</div>;
  if (!isSuccess) return <div>Failed</div>;

  const bankServices = data?.data.content;

  return (
    <>
      {bankServices.map((service: BankService, index: number) => (
        <div
          key={index}
          className="grid  grid-cols-3 md:grid-cols-6 gap-3 bg-white items-center p-3 md:p-7 rounded-[20px] mb-6 "
        >
          <div className="flex items-center gap-3  col-span-2">
            <div
              className={`w-[45px] h-[45px] rounded-xl flex items-center justify-center   ${index % 2 === 0 ? 'bg-[#FFF5D9]' : 'bg-[#DCFAF8]'} `}
            >
              <Image
                src={service.icon}
                alt="icon"
                width={20}
                height={20}
                className="w-[20px] h-[20px]"
                unoptimized={true}
              />
            </div>

            <div>
              <p className="service_list_title">{service.name}</p>
              <p className="service_list_subtitle text-[#718EBF]">
                {service.details}
              </p>
            </div>
          </div>
          <div className="hidden md:block col-span-1">
            <p className="service_list_title">Type</p>
            <p className="service_list_subtitle text-[#718EBF]">
              {service.type}
            </p>
          </div>
          <div className="hidden md:block col-span-1">
            <p className="service_list_title">Status</p>
            <p className="service_list_subtitle text-[#718EBF]">
              {service.status}
            </p>
          </div>
          <div className="hidden md:block col-span-1">
            <p className="service_list_title">Number of users</p>
            <p className="service_list_subtitle text-[#718EBF] ">
              {service.numberOfUsers}
            </p>
          </div>
          <div className="col-span-1">
            <button className="service_list_Button">View Details</button>
          </div>
        </div>
      ))}
    </>
  );
};

export default ServiceList;
