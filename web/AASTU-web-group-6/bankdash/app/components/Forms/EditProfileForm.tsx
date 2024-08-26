"use client";
import React, { useEffect } from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPencilAlt } from "@fortawesome/free-solid-svg-icons";
import { useForm } from "react-hook-form";
import UserService from "@/app/Services/api/userService";
import { useAppSelector } from "@/app/Redux/store/store";
import UserValue from "@/types/UserValue"; // Make sure to import UserValue type

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
  const userData: UserValue | null = useAppSelector((state) => state.user.user); // Assuming state.user.user contains the UserValue type
  const { register, handleSubmit, reset } = useForm<FormData>({
    defaultValues: {
      name: "",
      email: "",
      username: "",
      password: "",
      confirmPassword: "",
      dateOfBirth: "",
      permanentAddress: "",
      postalCode: "",
      presentAddress: "",
      city: "",
      country: "",
      profilePicture: "",
    },
  });

  useEffect(() => {
    if (userData) {
      reset({
        name: userData.name,
        email: userData.email,
        username: userData.username,
        dateOfBirth: userData.dateOfBirth,
        permanentAddress: userData.permanentAddress,
        postalCode: userData.postalCode,
        presentAddress: userData.presentAddress,
        city: userData.city,
        country: userData.country,
        profilePicture: userData.profilePicture,
      });
    }
  }, [userData, reset]);

  const onSubmit = async (data: FormData) => {
    const { confirmPassword, ...updatedUserData } = data;
    console.log("Updating user profile:", updatedUserData);

    try {
      const responseData = await UserService.update(
        updatedUserData,
        "accessToken"
      ); // Call the update method
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
    <div className="w-full py-8 flex gap-3 xss:mb-2 sm:mb-0 xxs:flex-wrap xxs:justify-center sm:justify-normal md:flex-nowrap md:px-6 md:justify-around lg:justify-between ">
      <div className="flex justify-center xs:w-full md:w-fit md:pl-5">
        <div className="relative h-fit flex justify-center">
          <Image
            src={userData?.profilePicture || "/assets/profile-1.png"}
            width={132}
            height={130}
            alt="Profile"
            className="rounded-full mr-4 border-2 border-solid border-gray-600 dark:border-2 dark:border-gray-100"
          />
          <button className="absolute bottom-3 right-1 px-2 py-1 bg-[#1814F3] text-white rounded-full">
            <FontAwesomeIcon icon={faPencilAlt} className="text-lg" />
          </button>
        </div>
      </div>

      <form
        className="w-flex flex-col xxs:justify-center xxs:px-3 sm:px-0 sm:justify-normal md:justify-around lg:w-[800px] py-2   "
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="name"
            >
              Your Name
            </label>
            <input
              type="text"
              id="name"
              {...register("name")}
              className="p-3 border-2 dark:bg-gray-200 dark:border-gray-300  border-gray-200 rounded-lg   placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="Charlene Reed"
            />
          </div>
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="username"
            >
              User Name
            </label>
            <input
              type="text"
              id="username"
              {...register("username")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700 dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="Charlene Reed"
            />
          </div>
        </div>

        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="email"
            >
              Email
            </label>
            <input
              type="email"
              id="email"
              {...register("email")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700  dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="charlenereed@gmail.com"
            />
          </div>
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="password"
            >
              Password
            </label>
            <input
              type="password"
              id="password"
              {...register("password")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700  dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="********"
            />
          </div>
        </div>

        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="dob"
            >
              Date of Birth
            </label>
            <input
              type="date"
              id="dob"
              {...register("dateOfBirth")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700  dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="25 January 1990"
            />
          </div>
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="present-address"
            >
              Present Address
            </label>
            <input
              type="text"
              id="present-address"
              {...register("presentAddress")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700  dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="San Jose, California, USA"
            />
          </div>
        </div>

        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="permanent-address"
            >
              Permanent Address
            </label>
            <input
              type="text"
              id="permanent-address"
              {...register("permanentAddress")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700  dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="San Jose, California, USA"
            />
          </div>
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="city"
            >
              City
            </label>
            <input
              type="text"
              id="city"
              {...register("city")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700  dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="San Jose"
            />
          </div>
        </div>

        <div className="gap-y-3 flex gap-x-2 justify-between flex-wrap">
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="postal-code"
            >
              Postal Code
            </label>
            <input
              type="text"
              id="postal-code"
              {...register("postalCode")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-600 dark:border-gray-600 placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="45962"
            />
          </div>
          <div className="xxs:w-full sm:w-[250px] lg:w-[280px] xl:w-[380px] flex flex-col">
            <label
              className="mb-1 text-slate-800 dark:text-gray-300"
              htmlFor="country"
            >
              Country
            </label>
            <input
              type="text"
              id="country"
              {...register("country")}
              className="p-3 border-2 border-gray-200 rounded-lg dark:bg-gray-700   dark:border-gray-600 placeholder:text

-slate-400 focus:outline-none focus:border-[#4640DE]  "
              placeholder="USA"
            />
          </div>
        </div>

        <div className="mt-6 flex justify-end">
          <button
            className="px-4 py-2 bg-[#1814F3] text-white rounded-md hover:bg-[#0702db] transition-all duration-300"
            type="submit"
          >
            Save Changes
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfileForm;
