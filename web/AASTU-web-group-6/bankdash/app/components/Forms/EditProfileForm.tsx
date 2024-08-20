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
    <div className="w-full py-8 flex justify-around xs:flex-wrap md:flex-nowrap">
      <div className="flex justify-center xs:w-full md:w-fit md:pl-5">
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

      <form className="w-flex flex-col md:justify-around lg:w-[800px] gap-y-5 py-2 xs:px-2 sm:px-0 " onSubmit={handleSubmit(onSubmit)}>
        

        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="name">Your Name</label>
            <input
              type="text"
              id="name"
              {...register("name")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="Charlene Reed"
            />
          </div>
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="username">User Name</label>
            <input
              type="text"
              id="username"
              {...register("username")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="Charlene Reed"
            />
          </div>
        </div>


        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="email">Email</label>
            <input
              type="email"
              id="email"
              {...register("email")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="charlenereed@gmail.com"
            />
          </div>
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="password">Password</label>
            <input
              type="password"
              id="password"
              {...register("password")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="********"
            />
          </div>
        </div>



        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="dob">Date of Birth</label>
            <input
              type="date"
              id="dob"
              {...register("dateOfBirth")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="25 January 1990"
            />
          </div>
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="present-address">Present Address</label>
            <input
              type="text"
              id="present-address"
              {...register("presentAddress")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="San Jose, California, USA"
            />
          </div>
        </div>



        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="permanent-address">Permanent Address</label>
            <input
              type="text"
              id="permanent-address"
              {...register("permanentAddress")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="San Jose, California, USA"
            />
          </div>
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="city">City</label>
            <input
              type="text"
              id="city"
              {...register("city")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="San Jose"
            />
          </div>
        </div>



        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="postal-code">Postal Code</label>
            <input
              type="text"
              id="postal-code"
              {...register("postalCode")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="45962"
            />
          </div>
          <div className="xs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label className="mb-1 text-slate-800" htmlFor="country">Country</label>
            <input
              type="text"
              id="country"
              {...register("country")}
              className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
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
