"use client";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import UserValue from "@/types/UserValue";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import AuthService from "@/app/Services/api/authService";
import ProgressComp from "../Box/ProgressComp";
import Image from "next/image";

// Define the schema using Zod
const schema = z
  .object({
    name: z.string().min(1, "Name is required"),
    email: z.string().email("Invalid email address"),
    dateOfBirth: z.string().min(1, "Date of Birth is required"),
    permanentAddress: z.string(),
    postalCode: z.string(),
    username: z.string().min(1, "Username is required"),
    password: z.string().min(6, "Password must be at least 6 characters"),
    confirmPassword: z.string().min(6, "Confirm Password is required"),
    presentAddress: z.string(),
    city: z.string(),
    country: z.string(),
    profilePicture: z.string().url("Invalid URL"),
    preference: z.object({
      currency: z.string(),
      sentOrReceiveDigitalCurrency: z.boolean().optional(),
      receiveMerchantOrder: z.boolean().optional(),
      accountRecommendations: z.boolean().optional(),
      timeZone: z.string(),
      twoFactorAuthentication: z.boolean().optional(),
    }),
  })
  .refine((data) => data.password === data.confirmPassword, {
    path: ["confirmPassword"],
    message: "Passwords do not match",
  });

type FormData = Omit<UserValue, "password"> & {
  password: string;
  confirmPassword: string;
};

const SignUpForm = () => {
  const [step, setStep] = useState(1);
  const {
    register,
    handleSubmit,
    trigger,
    formState: { errors },
    watch,
  } = useForm<FormData>({
    resolver: zodResolver(schema),
  });
  const confirmData = watch("password");
  const onSubmit = async (data: FormData) => {
    console.log(data);
    const { confirmPassword, ...userData } = data;
    console.log("Signup successful:", userData);
    try {
      const responseData = await AuthService.register(userData);
      console.log(responseData);
      if (responseData.success) {
        console.log("Signup successful:", responseData.message);
      } else {
        console.error("Signup failed:", responseData.message);
      }
    } catch (error) {
      console.error("An error occurred:", error);
    }
  };

  const nextStep = async () => {
    const fieldsToValidate: (
      | keyof FormData
      | `preference.${keyof FormData["preference"]}`
    )[] =
      step === 1
        ? ["name", "email", "username", "password", "confirmPassword"]
        : step === 2
        ? [
            "permanentAddress",
            "postalCode",
            "presentAddress",
            "city",
            "country",
          ]
        : ["profilePicture", "preference.currency", "preference.timeZone"];

    const isStepValid = await trigger(fieldsToValidate);
    if (isStepValid) {
      setStep((prev) => prev + 1);
    }
  };

  const prevStep = () => setStep((prev) => prev - 1);

  return (
    <>
      <div className="left flex flex-col gap-3">
        <h1 className="text-2xl flex flex-col text-center">SignUp</h1>
        <ProgressComp />
      </div>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-col justify-center items-center gap-3 w-[40%] min-w-[35%] p-6 py-8 bg-[#ffff] rounded-2xl border-[0.3px] border-solid border-[#1814f326] shadow-sm shadow-blue-100"
      >
        <div className="flex items-center circle p-4 pt-5 bg-blue-50 rounded-full">
          <Image src="/assets/logo-blue.svg" width={48} height={48} alt="" />
        </div>
        {step === 1 && (
          <div key={1} className="flex flex-col w-full px-6 gap-2">
            <h3 className="text-lg font-normal text-gray-400 text-center">
              Basic information
            </h3>
            {/* <ProgressComp /> */}

            <div className="flex flex-col">
              <label className="mb-1 text-slate-500" htmlFor="userName">
                Name
              </label>
              <input
                {...register("name", { required: "* FullName is required" })}
                placeholder="John Doe"
                id="userName"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
                type="text"
              />
              {errors.name && (
                <p className="text-[#1814F3]">{errors.name.message}</p>
              )}
            </div>
            <div className="flex flex-col">
              <label className="mb-1 text-slate-500" htmlFor="userName">
                Username
              </label>
              <input
                {...register("username", {
                  required: "* Username is required",
                })}
                placeholder="@John"
                id="userName"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
                type="text"
              />
              {errors.username && (
                <p className="text-[#1814F3]">{errors.username.message}</p>
              )}
            </div>
            <div className="flex flex-col">
              <label className="mb-1 text-slate-500" htmlFor="userName">
                Email
              </label>
              <input
                {...register("email", { required: "* Email is required" })}
                placeholder="example@gmail.com"
                id="userName"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]  "
                type="text"
              />
              {errors.email && (
                <p className="text-[#1814F3]">{errors.email.message}</p>
              )}
            </div>
            <div className="w-full password-box flex gap-1 flex-col">
              <label htmlFor="password" className="mb-1 text-slate-500">
                Password
              </label>
              <input
                type="password"
                className="password p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE] "
                placeholder="Enter password"
                {...register("password", {
                  required: "* Password is required",
                  pattern: {
                    value: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d@$!%*?&]{6,}$/,
                    message:
                      "* Password must contain be at least one letter and one number",
                  },
                  minLength: {
                    value: 6,
                    message: "* Password must be at least 6 characters long",
                  },
                })}
              />
              <p
                className="error text-[#1814F3]"
                style={{
                  display: errors.password?.message == null ? "none" : "flex",
                }}
              >
                {errors.password?.message}
              </p>
            </div>

            <div className="w-full password-box flex gap-1 flex-col">
              <label htmlFor="password" className="mb-1 text-slate-500">
                Confirm
              </label>
              <input
                type="password"
                className="password p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE] "
                placeholder="confirm"
                {...register("confirmPassword", {
                  required: "* Password is required",
                  pattern: {
                    value: /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d@$!%*?&]{6,}$/,
                    message:
                      "* Password must contain be at least one letter and one number",
                  },
                  minLength: {
                    value: 6,
                    message: "* Password must be at least 6 characters long",
                  },
                })}
              />
              <p
                className="error text-[#1814F3]"
                style={{
                  display:
                    errors.confirmPassword?.message == null ? "none" : "flex",
                }}
              >
                {errors.confirmPassword?.message}
              </p>
            </div>
          </div>
        )}

        {step === 2 && (
          <div key={2} className="flex flex-col gap-3 w-full px-6">
            <h3 className="text-xl font-semibold mb-4">Address Information</h3>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="permanentAddress">
                Permanent Address
              </label>
              <input
                {...register("permanentAddress")}
                id="permanentAddress"
                placeholder="123 Main St."
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.permanentAddress && (
                <p className="text-[#1814F3]">
                  {errors.permanentAddress.message}
                </p>
              )}
            </div>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="postalCode">
                Postal Code
              </label>
              <input
                {...register("postalCode")}
                id="postalCode"
                placeholder="12345"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.postalCode && (
                <p className="text-[#1814F3]">{errors.postalCode.message}</p>
              )}
            </div>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="presentAddress">
                Present Address
              </label>
              <input
                {...register("presentAddress")}
                id="presentAddress"
                placeholder="456 Elm St."
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.presentAddress && (
                <p className="text-[#1814F3]">
                  {errors.presentAddress.message}
                </p>
              )}
            </div>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="city">
                City
              </label>
              <input
                {...register("city")}
                id="city"
                placeholder="New York"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.city && (
                <p className="text-[#1814F3]">{errors.city.message}</p>
              )}
            </div>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="country">
                Country
              </label>
              <input
                {...register("country")}
                id="country"
                placeholder="USA"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.country && (
                <p className="text-[#1814F3]">{errors.country.message}</p>
              )}
            </div>
          </div>
        )}

        {step === 3 && (
          <div key={3} className="flex flex-col gap-3 w-full px-6">
            <h3 className="text-xl font-semibold mb-4">Personal Information</h3>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="profilePicture">
                Profile Picture URL
              </label>
              <input
                {...register("profilePicture")}
                id="profilePicture"
                placeholder="https://example.com/profile.jpg"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.profilePicture && (
                <p className="text-[#1814F3]">
                  {errors.profilePicture.message}
                </p>
              )}
            </div>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="dateOfBirth">
                Date of Birth
              </label>
              <input
                {...register("dateOfBirth")}
                id="dateOfBirth"
                type="date"
                className="p-3 border-2 border-gray-200 rounded-lg focus:outline-none focus:border-[#4640DE]"
              />
              {errors.dateOfBirth && (
                <p className="text-[#1814F3]">{errors.dateOfBirth.message}</p>
              )}
            </div>

            <h3 className="text-xl font-semibold mt-6">Preferences</h3>

            <div className="flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="currency">
                Currency
              </label>
              <input
                {...register("preference.currency")}
                id="currency"
                placeholder="USD"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.preference?.currency && (
                <p className="text-[#1814F3]">
                  {errors.preference.currency.message}
                </p>
              )}
            </div>

            <div className="flex gap-2">
              <input
                {...register("preference.sentOrReceiveDigitalCurrency")}
                id="sentOrReceiveDigitalCurrency"
                type="checkbox"
                className="h-5 w-5 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#4640DE]"
              />
              <label
                className="text-slate-500"
                htmlFor="sentOrReceiveDigitalCurrency"
              >
                Send or Receive Digital Currency
              </label>
              {errors.preference?.sentOrReceiveDigitalCurrency && (
                <p className="text-[#1814F3]">
                  {errors.preference.sentOrReceiveDigitalCurrency.message}
                </p>
              )}
            </div>

            <div className="flex gap-2">
              <input
                {...register("preference.receiveMerchantOrder")}
                id="receiveMerchantOrder"
                type="checkbox"
                className="h-5 w-5 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#4640DE]"
              />
              <label className="text-slate-500" htmlFor="receiveMerchantOrder">
                Receive Merchant Order
              </label>
              {errors.preference?.receiveMerchantOrder && (
                <p className="text-[#1814F3]">
                  {errors.preference.receiveMerchantOrder.message}
                </p>
              )}
            </div>

            <div className="flex gap-2">
              <input
                {...register("preference.accountRecommendations")}
                id="accountRecommendations"
                type="checkbox"
                className="h-5 w-5 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#4640DE]"
              />
              <label
                className="text-slate-500"
                htmlFor="accountRecommendations"
              >
                Account Recommendations
              </label>
              {errors.preference?.accountRecommendations && (
                <p className="text-[#1814F3]">
                  {errors.preference.accountRecommendations.message}
                </p>
              )}
            </div>

            <div className="mt-2 flex flex-col gap-2">
              <label className="mb-1 text-slate-500" htmlFor="timeZone">
                Time Zone
              </label>
              <input
                {...register("preference.timeZone")}
                id="timeZone"
                placeholder="UTC+0"
                className="p-3 border-2 border-gray-200 rounded-lg placeholder:text-slate-400 focus:outline-none focus:border-[#4640DE]"
                type="text"
              />
              {errors.preference?.timeZone && (
                <p className="text-[#1814F3]">
                  {errors.preference.timeZone.message}
                </p>
              )}
            </div>

            <div className="flex gap-2">
              <input
                {...register("preference.twoFactorAuthentication")}
                id="twoFactorAuthentication"
                type="checkbox"
                className="h-5 w-5 border-2 border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-[#4640DE]"
              />
              <label
                className="text-slate-500"
                htmlFor="twoFactorAuthentication"
              >
                Two-Factor Authentication
              </label>
              {errors.preference?.twoFactorAuthentication && (
                <p className="text-[#1814F3]">
                  {errors.preference.twoFactorAuthentication.message}
                </p>
              )}
            </div>
          </div>
        )}

        <div className="w-full p-6 flex justify-between mt-4">
          {step > 1 && (
            <button
              type="button"
              onClick={prevStep}
              className="px-4 py-2 rounded-xl border-2 border-[#1814F3] text-[#1814F3]"
            >
              Previous
            </button>
          )}
          {step < 3 ? (
            <button
              type="button"
              onClick={nextStep}
              className={`bg-[#1814F3] text-white px-6 py-3 rounded-xl ${
                step == 1 ? "grow" : ""
              }`}
            >
              Next
            </button>
          ) : (
            <button
              type="submit"
              className="bg-[#1814F3] text-white px-6 py-3 rounded-md"
            >
              Submit
            </button>
          )}
        </div>
      </form>
    </>
  );
};

export default SignUpForm;
