"use client";
import React, { useEffect, useState } from "react";
import DescriptionCard from "@/app/components/Card/DescriptionCard";
import ServicesCard from "@/app/components/Card/ServicesCard";
import { faX } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import axios from "axios";
import { useSession } from "next-auth/react";

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
  const accessToken = session?.accessToken as string;
  async function fetchData(accessToken: string) {
    try {
      const response = await axios.post(
        "https://bank-dashboard-1tst.onrender.com/bank-services",
        formData,
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
    <form
      className="flex flex-col space-y-4 p-2 rounded-lg max-h-[70vh] overflow-y-auto scrollbar-hide"
      onSubmit={handleSubmit(onSubmit)}
    >
      <div className="flex justify-between">
        <p className="text-base font-semibold dark:text-gray-200">
          Add Service
        </p>
        <button className="text-right" onClick={onClose}>
          <FontAwesomeIcon icon={faX} className="dark:text-white" />
        </button>
      </div>

      <div>
        <label
          htmlFor="name"
          className="block text-md font-medium text-gray-700 dark:text-gray-300"
        >
          Service Name
        </label>
        <input
          {...register("name", {
            required: "Receiver username is required",
          })}
          type="text"
          placeholder="Enter Service Name"
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm placeholder:text-xs"
        />
        {errors.name && (
          <span className="text-red-600 text-sm">{errors.name.message}</span>
        )}
      </div>

      <div>
        <label
          htmlFor="details"
          className="block text-md font-medium text-gray-700 dark:text-gray-300"
        >
          Details
        </label>
        <input
          {...register("details", { required: "Details is required" })}
          type="text"
          placeholder="Enter Details"
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm placeholder:text-xs"
        />
        {errors.details && (
          <span className="text-red-600 text-sm">{errors.details.message}</span>
        )}
      </div>

      <div className="flex gap-4">
        <div>
          <label
            htmlFor="numberOfUsers"
            className="block text-md font-medium text-gray-700 dark:text-gray-300"
          >
            Number of Users
          </label>
          <input
            {...register("numberOfUsers", {
              required: "Number of users is required",
            })}
            type="number"
            placeholder="Enter No. of Users"
            className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm placeholder:text-xs"
          />
          {errors.numberOfUsers && (
            <span className="text-red-600 text-sm">
              {errors.numberOfUsers.message}
            </span>
          )}
        </div>
        {/* </div> */}

        <div>
          <label
            htmlFor="status"
            className="block text-md font-medium text-gray-700 dark:text-gray-300"
          >
            Status
          </label>
          <select
            {...register("status", { required: "Status is required" })}
            className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          >
            {StatusOptions.map((item, index) => (
              <option key={index} value={item}>
                {item}
              </option>
            ))}
          </select>
          {errors.status && (
            <span className="text-red-600 text-sm">
              {errors.status.message}
            </span>
          )}
        </div>

        <div>
          <label
            htmlFor="type"
            className="block text-md font-medium text-gray-700 dark:text-gray-300"
          >
            Type
          </label>
          <select
            {...register("type", { required: "Type is required" })}
            className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          >
            {typeOptions.map((item, index) => (
              <option key={index} value={item}>
                {item}
              </option>
            ))}
          </select>
          {errors.type && (
            <span className="text-red-600 text-sm">{errors.type.message}</span>
          )}
        </div>
      </div>

      <div>
        <label
          htmlFor="icon"
          className="block text-md font-medium text-gray-700 dark:text-gray-300"
        >
          Icon
        </label>
        <select
          {...register("icon", { required: "Icon is required" })}
          className="mt-1 p-3 border block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
        >
          {iconOptions.map((item, index) => (
            <option key={index} value={item.value}>
              {item.label}
            </option>
          ))}
        </select>
        {errors.icon && (
          <span className="text-red-600 text-sm">{errors.icon.message}</span>
        )}
      </div>

      <button
        type="submit"
        className="w-full bg-blue-600 text-white py-2 px-4 rounded-lg shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
      >
        Add Service
      </button>
    </form>
  );
};

export default Services;
