'use client';
import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';
import { useSession } from 'next-auth/react';

const Page = () => {
  const [services, setServices] = useState([]);
  const { data: session } = useSession();
  const token: string = `Bearer ${session?.user?.accessToken}`;

  useEffect(() => {
    const fetchServices = async () => {
      try {
        const response = await axios.get(
          "https://bank-dashboard-1tst.onrender.com/bank-services?page=0&size=7",
          {
            headers: {
              Authorization: token,
            },
          }
        );
        setServices(response.data.data.content);
      } catch (error) {
        console.error("Error fetching services:", error);
      }
    };

    fetchServices();
  }, [token]);

  const getImageProps = (serviceName: string) => {
    switch (serviceName) {
      case 'Business Loans':
        return { src: '/images/loan 1.png', alt: `${serviceName} Icon`, bgClass: 'bg-pink-100' };
      case 'Checking Accounts':
        return { src: '/images/orange-image.png', alt: `${serviceName} Icon`, bgClass: 'bg-orange-100' };
      case 'Savings Accounts':
        return { src: '/images/pink-image.png', alt: `${serviceName} Icon`, bgClass: 'bg-pink-100' };
      case 'Debit and credit cards':
        return { src: '/images/blue-image.png', alt: `${serviceName} Icon`, bgClass: 'bg-blue-100' };
      case 'Life Insurance':
        return { src: '/images/green-image.png', alt: `${serviceName} Icon`, bgClass: 'bg-green-100' };
      default:
        return { src: '/images/default-image.png', alt: `${serviceName} Icon`, bgClass: 'bg-gray-100' };
    }
  };

  return (
    <div className="bg-gray-100">
      {/* Life Insurance Section */}
      <div className="w-full flex items-center bg-gray-100 p-6 gap-6 pr-10 pl-10">
        <div className="mr-5 lg:mr-0 flex gap-4 overflow-x-auto lg:overflow-x-visible [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none] lg:pl-10 lg:pt-10 w-full">
          <div className="bg-white h-[90px] rounded-xl flex-grow flex justify-center items-center gap-4">
            <div className="flex items-center justify-center bg-blue-100 rounded-full w-14 h-14">
              <Image
                src="/images/life-insurance filled 1.png"
                alt="heart Icon"
                objectFit="cover"
                width={25}
                height={25}
              />
            </div>
            <div className="flex flex-col">
              <h3 className="font-inter text-black text-lg font-medium leading-[16.94px] text-left">Life Insurance</h3>
              <p className="font-inter text-blue-900 opacity-60 text-base">Unlimited Protection</p>
            </div>
          </div>

          {/* Shopping */}
          <div className="bg-white h-[90px] rounded-xl flex-grow flex justify-center items-center gap-4">
            <div className="flex items-center justify-center bg-orange-100 rounded-full w-14 h-14">
              <Image
                src="/images/shopping.png"
                alt="Shopping Icon"
                objectFit="cover"
                width={20}
                height={20}
              />
            </div>
            <div className="flex flex-col">
              <h3 className="font-inter text-black text-lg font-medium leading-[16.94px] text-left">Shopping</h3>
              <p className="font-inter text-blue-900 opacity-60 text-base">Buy. Think. Grow.</p>
            </div>
          </div>

          {/* Safety */}
          <div className="bg-white h-[90px] rounded-xl flex-grow flex justify-center items-center gap-4">
            <div className="flex items-center justify-center bg-green-100 rounded-full w-14 h-14">
              <Image
                src="/images/shield 1.png"
                alt="Safety Icon"
                objectFit="cover"
                width={50}
                height={50}
              />
            </div>
            <div className="flex flex-col">
              <h3 className="font-inter text-black text-lg font-medium leading-[16.94px] text-left">Safety</h3>
              <p className="font-inter text-blue-900 opacity-60 text-base">We are all your allies</p>
            </div>
          </div>
        </div>
      </div>

      {/* Title Section */}
      <div className="p-2 sm:p-1">
        <p className="ml-8 text-xl sm:text-2xl font-inter text-blue-950 font-semibold">
          Bank Service List
        </p>
      </div>

      {/* Bank Service List Section */}
      <div className="p-5 m-4 space-y-3">
        {services.map((service) => {
          const { src, alt, bgClass } = getImageProps(service.name);
          return (
            <div key={service.id} className="w-full flex items-center justify-center">
              <div className="bg-white h-[70px] rounded-xl flex-grow items-center justify-center sm:shadow-md">
                <div className="flex items-center justify-center p-1 m-0">
                  <div className="flex items-center gap-3 w-full">
                    <div className={`flex items-center justify-center rounded-xl w-12 h-12 ${bgClass}`}>
                      <Image
                        src={src}
                        alt={alt}
                        width={25}
                        height={25}
                      />
                    </div>
                    <div className="flex-1 sm:flex-none">
                      <h3 className="font-inter text-black text-base font-medium leading-[16.94px] text-left whitespace-nowrap">
                        {service.name}
                      </h3>
                      <p className="sm:font-inter sm:text-xs sm:font-normal sm:leading-[14.52px] sm:text-left text-blue-900 opacity-60 whitespace-nowrap">
                        {service.details.slice(0, 25)}
                      </p>
                    </div>
                  </div>

                  <div className="flex sm:flex-row items-center justify-between w-full sm:w-auto mt-3 sm:mt-0 sm:ml-auto sm:pr-6">
                    <div className="hidden sm:flex flex-row gap-4 sm:gap-12 items-start pr-6">
                      <div className="flex-col items-center gap-3 w-48">
                        <p className="font-semibold">Number of Users</p>
                        <p className="font-inter text-blue-900 opacity-60">{service.numberOfUsers}</p>
                      </div>
                      <div className="flex-col items-center w-20">
                        <p className="font-semibold">Status</p>
                        <p className="font-inter text-blue-900 opacity-60">{service.status}</p>
                      </div>
                      <div className="flex-col items-center w-20">
                        <p className="font-semibold">Type</p>
                        <p className="font-inter text-blue-900 opacity-60">{service.type}</p>
                      </div>
                    </div>
                  </div>

                  <button className="text-purple-600 mr-1 hover:bg-purple-50 rounded-full">
                    View Details
                  </button>
                </div>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default Page;
