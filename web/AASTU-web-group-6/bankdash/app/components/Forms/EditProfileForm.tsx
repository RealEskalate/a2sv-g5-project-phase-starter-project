import React from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPencilAlt } from "@fortawesome/free-solid-svg-icons";

const EditProfileForm = () => {
  return (
    <div className="content-center w-full flex justify-center gap-x-12 py-4 flex-wrap sm:pt-4 md:pt-8 lg:pt-12 bg-white rounded-lg shadow-md">
      <div className="flex justify-center  gap-3 w-fit sm:w-full md:w-auto">
        <div className="relative h-fit flex justify-center">
          <Image
            src="/assets/profile-1.png"
            width={132}
            height={130}
            alt="Profile"
            className="rounded-full mr-4"
          />
          <button className="absolute bottom-3 right-1 px-2 py-1 bg-[#1814F3] text-white rounded-full">
            <FontAwesomeIcon icon={faPencilAlt} className="text-lg" />
          </button>
        </div>
      </div>

      <form className="mt-8 lg:space-y-6 flex flex-col sm:flex-wrap">
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="name" className="">
              Your Name
            </label>
            <input
              type="text"
              id="name"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="Charlene Reed"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="username" className="">
              User Name
            </label>
            <input
              type="text"
              id="username"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="Charlene Reed"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="email" className="">
              Email
            </label>
            <input
              type="email"
              id="email"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="charlenereed@gmail.com"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="password" className="">
              Password
            </label>
            <input
              type="password"
              id="password"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="********"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full mr-3 md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="dob" className="">
              Date of Birth
            </label>
            <input
              type="date"
              id="dob"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="25 January 1990"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="present-address" className="">
              Present Address
            </label>
            <input
              type="text"
              id="present-address"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="San Jose, California, USA"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="permanent-address" className="">
              Permanent Address
            </label>
            <input
              type="text"
              id="permanent-address"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="San Jose, California, USA"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="city" className="">
              City
            </label>
            <input
              type="text"
              id="city"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="San Jose"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="postal-code" className="">
              Postal Code
            </label>
            <input
              type="text"
              id="postal-code"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="45962"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="country" className="">
              Country
            </label>
            <input
              type="text"
              id="country"
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="USA"
            />
          </div>
        </div>

        <div className="flex lg:justify-end mt-3 sm:w-full sm:justify-center">
          <button
            type="submit"
            className="w-[192px] py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfileForm;
