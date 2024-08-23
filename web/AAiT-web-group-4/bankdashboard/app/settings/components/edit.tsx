"use client";

import React from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";

interface FormData {
  name: string;
  email: string;
  dateofbirth: string;
  permanentAddress: string;
  postalCode: string;
  userName: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
}

const EditProfile = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>();

  const [activeButton, setActiveButton] = React.useState("edit");

  // Handle form submission
  const onSubmit = (data: FormData) => {
    console.log("Data to be sent:", data);
    alert("Profile updated successfully!");
  };

  return (
    <div className="w-[1191px] h-[924px] pt-8 pl-[41px]">
      <div className="flex flex-col w-[1110px] h-[717px] pt-[37px] pl-[30px] gap-[52px] bg-white">
        <div className="flex flex-row border-b w-[1050px] h-[30px] text-[#718EBF] gap-12">
          <button
            className={`w-[114px] h-[30px] items-center ${
              activeButton === "edit"
                ? "border-b-2 text-[#1814F3] border-[#1814F3]"
                : ""
            }`}
            onClick={() => setActiveButton("edit")}
          >
            Edit Profile
          </button>
          <button
            className={`w-[114px] h-[30px] items-center ${
              activeButton === "preferences"
                ? "border-b-2 text-[#1814F3] border-[#1814F3]"
                : ""
            }`}
            onClick={() => setActiveButton("preferences")}
          >
            Preferences
          </button>
          <button
            className={`w-[114px] h-[30px] items-center ${
              activeButton === "Security"
                ? "border-b-2 text-[#1814F3] border-[#1814F3]"
                : ""
            }`}
            onClick={() => setActiveButton("Security")}
          >
            Security
          </button>
        </div>

        <form onSubmit={handleSubmit(onSubmit)} id="editProfileForm">
          <div className="flex flex-row w-fit h-fit gap-[53px]">
            <div className="w-[142px] h-[140px] relative">
              <Image
                src="/theGirl.png"
                width={110}
                height={113}
                className="rounded-full"
                alt="Profile"
              />
              <button
                type="button"
                className="bg-[#1814F3] w-[30px] h-[30px] flex justify-center items-center rounded-full absolute bottom-[36px] right-8"
              >
                <Image
                  src="/profilePen.png"
                  width={25}
                  height={25}
                  alt="Edit"
                />
              </button>
            </div>

            <div className="flex flex-col gap-[22px]">
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Your Name</label>
                <input
                  {...register("name", { required: "Full Name is required" })}
                  placeholder="Full Name"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.name && (
                  <span className="text-red-500">{errors.name.message}</span>
                )}
              </div>
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Email</label>
                <input
                  {...register("email", {
                    required: "Email is required",
                    pattern: {
                      value: /\S+@\S+\.\S+/,
                      message: "Email is invalid",
                    },
                  })}
                  placeholder="Email"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.email && (
                  <span className="text-red-500">{errors.email.message}</span>
                )}
              </div>

              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Date of Birth</label>
                <input
                  type="date"
                  {...register("dateofbirth", {
                    required: "Date of Birth is required",
                  })}
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] px-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.dateofbirth && (
                  <span className="text-red-500">
                    {errors.dateofbirth.message}
                  </span>
                )}
              </div>
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Permanent Address</label>
                <input
                  {...register("permanentAddress", {
                    required: "Permanent Address is required",
                  })}
                  placeholder="Address"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.permanentAddress && (
                  <span className="text-red-500">
                    {errors.permanentAddress.message}
                  </span>
                )}
              </div>

              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Postal Code</label>
                <input
                  {...register("postalCode", {
                    required: "Postal Code is required",
                  })}
                  placeholder="Code"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.postalCode && (
                  <span className="text-red-500">
                    {errors.postalCode.message}
                  </span>
                )}
              </div>
            </div>

            <div className="flex flex-col gap-[22px] -ml-6">
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">User Name</label>
                <input
                  {...register("userName", {
                    required: "User Name is required",
                  })}
                  placeholder="Name"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.userName && (
                  <span className="text-red-500">
                    {errors.userName.message}
                  </span>
                )}
              </div>
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Password</label>
                <input
                  type="password"
                  {...register("password", {
                    required: "Password is required",
                  })}
                  placeholder="Password"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.password && (
                  <span className="text-red-500">
                    {errors.password.message}
                  </span>
                )}
              </div>

              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Present Address</label>
                <input
                  {...register("presentAddress", {
                    required: "Present Address is required",
                  })}
                  placeholder="Address"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.presentAddress && (
                  <span className="text-red-500">
                    {errors.presentAddress.message}
                  </span>
                )}
              </div>
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">City</label>
                <input
                  {...register("city", { required: "City is required" })}
                  placeholder="City"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.city && (
                  <span className="text-red-500">{errors.city.message}</span>
                )}
              </div>
              <div className="w-[418px] h-[80px]">
                <label className="text-[#232323]">Country</label>
                <input
                  {...register("country", { required: "Country is required" })}
                  placeholder="Country"
                  className="w-[418px] h-[50px] rounded-[15px] mt-[11px] pl-5 border border-[#DFEAF2] text-[#718EBF]"
                />
                {errors.country && (
                  <span className="text-red-500">{errors.country.message}</span>
                )}
              </div>
            </div>
          </div>
          <div className=" flex justify-end pt-[60px] pb-[53px] px-12">
            <button
              type="submit"
              className="bg-[#1814F3] w-[147px] h-[55px] text-white text-[20px] rounded-[15px]"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default EditProfile;
