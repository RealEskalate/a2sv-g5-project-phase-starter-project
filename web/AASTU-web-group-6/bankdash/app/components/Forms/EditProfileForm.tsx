"use client";
import React from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPencilAlt } from "@fortawesome/free-solid-svg-icons";
import { useForm } from "react-hook-form"
import UserService from "@/app/Services/api/userService";

type FormData = {
  name: string;
  email: string;
  username: string;
  password: string;
  confirmPassword: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
};

const EditProfileForm = () => {
  const { register, handleSubmit } = useForm<FormData>();

  const onSubmit = async (data: FormData) => {
    const { confirmPassword, ...userData } = data; 
    console.log("Updating user profile:", userData);

    try {
      const responseData = await UserService.update(userData,"accessToken"); // Call the update method
      if (responseData.success) {
        console.log("Profile update successful:", responseData.message);
      } else {
        console.error("Profile update failed:", responseData.message);
      }
    } catch (error) {
      console.error("An error occurred during profile update:", error);
    }
  };

  return (
    <div className="content-center w-full flex justify-center gap-x-12 py-4 flex-wrap sm:pt-4 md:pt-8 lg:pt-12 bg-white rounded-lg shadow-md">
      <div className="flex justify-center  gap-3 w-fit xd:w-full sm:w-fit md:w-fit">
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

      <form className="mt-8 lg:space-y-6 flex flex-col sm:flex-wrap" onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="name">Your Name</label>
            <input
              type="text"
              id="name"
              {...register("name")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="Charlene Reed"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="username">User Name</label>
            <input
              type="text"
              id="username"
              {...register("username")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="Charlene Reed"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              {...register("email")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="charlenereed@gmail.com"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              {...register("password")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="********"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full mr-3 md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="dob">Date of Birth</label>
            <input
              type="date"
              id="dob"
              {...register("dateOfBirth")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="25 January 1990"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="present-address">Present Address</label>
            <input
              type="text"
              id="present-address"
              {...register("presentAddress")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="San Jose, California, USA"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="permanent-address">Permanent Address</label>
            <input
              type="text"
              id="permanent-address"
              {...register("permanentAddress")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="San Jose, California, USA"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="city">City</label>
            <input
              type="text"
              id="city"
              {...register("city")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="San Jose"
            />
          </div>
        </div>
        <div className="flex flex-wrap md:space-x-2 md:mb-2 lg:space-x-8">
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="postal-code">Postal Code</label>
            <input
              type="text"
              id="postal-code"
              {...register("postalCode")}
              className="sm:w-full md:w-[256px] lg:w-[418px] border-2 border-gray-300 px-5 py-4 rounded-xl"
              placeholder="45962"
            />
          </div>
          <div className="sm:w-full md:w-[256px] lg:w-[418px] flex flex-col lg:space-y-3">
            <label htmlFor="country">Country</label>
            <input
              type="text"
              id="country"
              {...register("country")}
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
