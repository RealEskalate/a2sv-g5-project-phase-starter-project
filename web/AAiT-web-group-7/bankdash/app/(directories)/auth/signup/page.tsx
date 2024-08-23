"use client";
import React from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";

interface PreferenceValues {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

interface FormValues {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  preference: PreferenceValues;
}

const Signup = () => {
  const { register, setValue, handleSubmit, formState, watch } =
    useForm<FormValues>();
  const { errors } = formState;

  // onSubmit function
  const onSubmit = async (formData: FormValues) => {
    console.log("formData", formData);
  };

  const handleProfilePictureChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result as string;
        setValue("profilePicture", base64String);
      };
      reader.readAsDataURL(file);
    }
  };

  return (
    <div className="container  flex justify-between w-full  p-8">
      <div className="signup-container ml-6 w-fit h-fit flex flex-col gap-4">
        <div className="w-full flex flex-col gap-4">
          <h1 className="text-[28px] font-poppins font-black leading-[34px] text-center text-[#25324B]">
            Sign Up Today!
          </h1>
        </div>

        <form
          className="w-fit h-fit flex justify-between"
          onSubmit={handleSubmit(onSubmit)}
        >
          <div className="flex flex-col gap-5">
            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Full Name
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your full name"
                type="text"
                {...register("name", {
                  required: "Full Name is required",
                  minLength: {
                    value: 2,
                    message: "Full Name must be at least 2 characters long",
                  },
                })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.name?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Date of Birth
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                type="date"
                {...register("dateOfBirth", {
                  required: "Date of Birth is required",
                })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.dateOfBirth?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Permanent Address
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your permanent address"
                type="text"
                {...register("permanentAddress", {
                  required: "Permanent Address is required",
                })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.permanentAddress?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Postal Code
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your postal code"
                type="text"
                {...register("postalCode", {
                  required: "Postal Code is required",
                })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.postalCode?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Username
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Choose a username"
                type="text"
                {...register("username", { required: "Username is required" })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.username?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Password
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your password"
                type="password"
                {...register("password", {
                  required: "Password is required",
                  minLength: {
                    value: 6,
                    message: "Password must be at least 6 characters long",
                  },
                })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.password?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Present Address
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your present address"
                type="text"
                {...register("presentAddress", {
                  required: "Present Address is required",
                })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.presentAddress?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                City
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your city"
                type="text"
                {...register("city", { required: "City is required" })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.city?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Country
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your city"
                type="text"
                {...register("city", { required: "City is required" })}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.city?.message}
              </p>
            </div>

            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Profile Picture
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                type="file"
                accept="image/*"
                onChange={handleProfilePictureChange}
              />
              <p className="error text-[12px] text-center text-red-700">
                {errors.profilePicture?.message}
              </p>
            </div>
          </div>
          <div className="container h-[100vh] flex justify-between w-full p-8">
            <div className="signup-container ml-6 w-fit h-fit flex flex-col gap-6">
              <h1 className="text-[28px] font-poppins font-black leading-[34px] text-center text-[#25324B]">
                Set Your Preferences
              </h1>

              <div className="w-[350px] h-[60px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Currency
                </label>
                <select
                  className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  {...register("preference.currency", {
                    required: "Currency is required",
                  })}
                >
                  <option value="">Select your currency</option>
                  <option value="USD">USD - US Dollar</option>
                  <option value="EUR">EUR - Euro</option>
                  <option value="GBP">GBP - British Pound</option>
                  <option value="JPY">JPY - Japanese Yen</option>
                  <option value="AUD">AUD - Australian Dollar</option>
                  <option value="CAD">CAD - Canadian Dollar</option>
                  <option value="INR">INR - Indian Rupee</option>
                </select>
                <p className="error text-[12px] text-center text-red-700">
                  {errors.preference?.currency?.message}
                </p>
              </div>

              <div className="w-[350px] h-[60px] flex flex-col gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Time Zone
                </label>
                <select
                  className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                  {...register("preference.timeZone", {
                    required: "Time Zone is required",
                  })}
                >
                  <option value="">Select your time zone</option>
                  <option value="GMT">GMT - Greenwich Mean Time</option>
                  <option value="UTC">UTC - Coordinated Universal Time</option>
                  <option value="EST">EST - Eastern Standard Time</option>
                  <option value="PST">PST - Pacific Standard Time</option>
                  <option value="CST">CST - Central Standard Time</option>
                  <option value="IST">IST - Indian Standard Time</option>
                </select>
                <p className="error text-[12px] text-center text-red-700">
                  {errors.preference?.timeZone?.message}
                </p>
              </div>

              <div className=" flex flex-row-reverse  gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Send or Receive Digital Currency
                </label>
                <input
                  className=" rounded-[6px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                  type="checkbox"
                  {...register("preference.sentOrReceiveDigitalCurrency")}
                />
              </div>

              <div className=" flex flex-row-reverse gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Receive Merchant Order
                </label>
                <input
                  className="rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                  type="checkbox"
                  {...register("preference.receiveMerchantOrder")}
                />
              </div>

              <div className=" flex flex-row-reverse gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Account Recommendations
                </label>
                <input
                  className=" rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                  type="checkbox"
                  {...register("preference.accountRecommendations")}
                />
              </div>

              <div className=" flex flex-row-reverse gap-2">
                <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                  Two-Factor Authentication
                </label>
                <input
                  className=" rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px]]"
                  type="checkbox"
                  {...register("preference.twoFactorAuthentication")}
                />
              </div>

              <button
                type="submit"
                className="w-[350px] h-[45px] bg-[#4640DE] text-white rounded-[6px] text-center text-[16px] font-semibold"
              >
                Submit Preferences
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Signup;
