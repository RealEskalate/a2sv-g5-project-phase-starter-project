"use client";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { User } from "@/types/index";
import { creditcardstyles, colors, logo } from "../../../constants/index";
import Image from "next/image";
import Link from "next/link";
import { registerUser } from '@/services/authentication';
import Cookie from "js-cookie";
import router from 'next/router';
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const SignupForm = () => {
  const [step, setStep] = useState(1);
  
  const {
    register,
    handleSubmit,
    formState: { errors, isValid },
    trigger,
  } = useForm<User>({
    mode: "onChange"
  });

  const handleRegister = async (data: User) => {
    try {
      const registeredUser = await registerUser(data);
      console.log('Registered User:', registeredUser);
      Cookie.set('accessToken', registeredUser.data.access_token);
      Cookie.set('refreshToken', registeredUser.data.refresh_token);
      router.push('/signin'); // Correctly navigate to the sign-in page
    } catch (error) {
      console.error('Registration Error:', error);
      toast.error('Registration failed. Please try again.');
    }
  };

  const onSubmit = async (data: User) => {
    console.log("User Data:", data);
    handleRegister(data);
  };

  const handleNextStep = async () => {
    const result = await trigger(); // Trigger validation for the current step
    if (result) {
      setStep((prevStep) => prevStep + 1);
    } else {
      toast.error('Please fill in all required fields.');
    }
  };

  const handlePreviousStep = () => setStep((prevStep) => prevStep - 1);

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-dark dark:text-white">
      <form onSubmit={handleSubmit(onSubmit)} className="max-w-3xl mx-auto p-3 bg-white dark:bg-gray-900 shadow-md rounded-lg">
        <div className="flex justify-center items-center m-2">
          <Image src={logo.icon} alt="Logo" height={60} width={60} />
          <h1 className="font-bold text-3xl text-gray-700 font-serif p-2">
            <p className="text-gray-600">NEXT BANK</p>
          </h1>
        </div>
        <h2 className="text-blue-600 font-semibold text-2xl text-center mb-4">Sign Up</h2>
        <p className="text-center text-gray-600 mb-8">Please enter your details</p>
  
        {step === 1 && (
          <div className="grid gap-6 sm:grid-cols-1 md:grid-cols-2">
            {/* Step 1 Fields */}
            <div>
              <label htmlFor="name" className="block font-medium text-gray-700">Full Name</label>
              <input
                type="text"
                id="name"
                placeholder="ex: John"
                {...register("name", { required: "Name is required" })}
                className="border border-gray-300 dark:text-white dark:bg-black rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
              />
              {errors.name && <p className="text-red-500 mt-1">{errors.name.message}</p>}
            </div>
  
            <div>
              <label htmlFor="email" className="block font-medium text-gray-700">Email</label>
              <input
                type="email"
                id="email"
                placeholder="ex: example@gmail.com"
                {...register("email", {
                  required: "Email is required",
                  pattern: {
                    value: /^\S+@\S+\.\S+$/,
                    message: "Invalid email format",
                  },
                })}
                className="border dark:text-white dark:bg-black border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
              />
              {errors.email && <p className="text-red-500 mt-1">{errors.email.message}</p>}
            </div>
  
            <div>
              <label htmlFor="dateOfBirth" className="block font-medium text-gray-700">Date of Birth</label>
              <input
                type="date"
                id="dateOfBirth"
                {...register("dateOfBirth", { required: "Date of Birth is required" })}
                className="border dark:text-white dark:bg-black border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
              />
              {errors.dateOfBirth && <p className="text-red-500 mt-1">{errors.dateOfBirth.message}</p>}
            </div>
  
            <div>
              <label htmlFor="username" className="block font-medium text-gray-700">Username</label>
              <input
                type="text"
                id="username"
                placeholder="ex: Star"
                {...register("username", { 
                  required: "Username is required",  
                  minLength: {
                    value: 4,
                    message: "Username must be at least 4 characters",
                  },
                  maxLength: {
                    value: 12,
                    message: "Username must not exceed 12 characters",
                  }
                })}
                className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
              />
              {errors.username && <p className="text-red-500 mt-1">{errors.username.message}</p>}
            </div>
  

            <div>
            <label htmlFor="password" className="block font-medium text-gray-700">Password</label>
            <input
              type="password"
              id="password"
              placeholder="ex: password"
              {...register("password", {
                required: "Password is required",
                minLength: {
                  value: 8,
                  message: "Password must be at least 8 characters",
                },
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.password && <p className="text-red-500 mt-1">{errors.password.message}</p>}
          </div>

          <div className="md:row-span-2">
            <label htmlFor="permanentAddress" className="block font-medium text-gray-700">Permanent Address</label>
            <textarea
              id="permanentAddress"
              placeholder="ex: A.A"
              {...register("permanentAddress", { required: "Permanent Address is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.permanentAddress && <p className="text-red-500 mt-1">{errors.permanentAddress.message}</p>}
          </div>

          <div>
            <label htmlFor="postalCode" className="block font-medium text-gray-700">Postal Code</label>
            <input
              type="text"
              id="postalCode"
              placeholder="ex: 1000"
              {...register("postalCode", { required: "Postal Code is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.postalCode && <p className="text-red-500 mt-1">{errors.postalCode.message}</p>}
          </div>
        </div>
      )}

      {step === 2 && (
        <div className="grid gap-6 sm:grid-cols-1 mb-6 md:py-14 md:grid-cols-2">
          {/* Step 2 Fields */}
          <div>
            <label htmlFor="presentAddress" className="block font-medium text-gray-700">Present Address</label>
            <input
              type="text"
              id="presentAddress"
              placeholder="ex: Bola"
              {...register("presentAddress", { required: "Present Address is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.presentAddress && <p className="text-red-500 mt-1">{errors.presentAddress.message}</p>}
          </div>

          <div>
            <label htmlFor="city" className="block font-medium text-gray-700">City</label>
            <input
              type="text"
              id="city"
              placeholder="ex: A.A"
              {...register("city", { required: "City is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.city && <p className="text-red-500 mt-1">{errors.city.message}</p>}
          </div>

          <div>
            <label htmlFor="country" className="block font-medium text-gray-700">Country</label>
            <input
              type="text"
              id="country"
              placeholder="ex: Ethiopia"
              {...register("country", { required: "Country is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.country && <p className="text-red-500 mt-1">{errors.country.message}</p>}
          </div>


          <div>
            <label htmlFor="profilePicture" className="block font-medium text-gray-700">Profile Picture</label>
            <input
              type="file"
              id="profilePicture"
              {...register("profilePicture")}
              className="p-2 w-full border border-gray-300 rounded-lg mt-1 focus:outline-none focus:border-blue-500"
            />
            {errors.profilePicture && <p className="text-red-500 mt-1">{errors.profilePicture.message}</p>}
          </div>

          <div>
            <label htmlFor="currency" className="block font-medium text-gray-700">Preferred Currency</label>
            <select
              id="currency"
              {...register("currency", { required: "Preferred Currency is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            >
              <option value="">Select Currency</option>
              <option value="USD">USD</option>
              <option value="ETB">ETB</option>
              <option value="EUR">EUR</option>
              <option value="GBP">GBP</option>
              <option value="JPY">JPY</option>
              <option value="CAD">CAD</option>
            </select>
            {errors.currency && <p className="text-red-500 mt-1">{errors.currency.message}</p>}
          </div>

            <div>
              <label htmlFor="timeZone" className="block font-medium text-gray-700">Time Zone</label>
              <select
                id="timeZone"
                {...register("preference.timeZone", { required: "Time Zone is required" })}
                className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
              >
                <option value="">Select Time Zone</option>
                <option value="UTC">UTC</option>
                {/* Add more time zones as needed */}
              </select>
              {errors.preference?.timeZone && <p className="text-red-500 mt-1">{errors.preference.timeZone.message}</p>}
            </div>

          <div  className="flex flex-row p-3">
            <input
              type="checkbox"
              id="sentOrReceiveDigitalCurrency"
              {...register("preference.sentOrReceiveDigitalCurrency")}
              className="mr-1"
            />
              <label htmlFor="sentOrReceiveDigitalCurrency" className="block font-medium text-gray-700">Send or Receive Digital Currency?</label>
            </div>

            <div  className="flex flex-row p-3">
              <input
                type="checkbox"
                id="receiveMerchantOrder"
                {...register("preference.receiveMerchantOrder")}
                className="mr-1"
              />
              <label htmlFor="receiveMerchantOrder" className="block font-medium text-gray-700">Receive Merchant Order Notifications?</label>
            </div>

            <div className="flex flex-row p-3">
              <input
                type="checkbox"
                id="accountRecommendations"
                {...register("preference.accountRecommendations")}
                className="mr-1"
              />
              <label htmlFor="accountRecommendations" className="block font-medium text-gray-700">Receive Account Recommendations?</label>
            </div>


            <div  className="flex flex-row p-3">
              <input
                type="checkbox"
                id="twoFactorAuthentication"
                {...register("preference.twoFactorAuthentication")}
                className="mr-1"
              />
              <label htmlFor="twoFactorAuthentication" className="block font-medium text-gray-700">Enable Two-Factor Authentication?</label>
            </div>
          </div>
        )}

      <div className="flex justify-between mt-8">
        {step > 1 && (
          <button
            type="button"
            onClick={handlePreviousStep}
            className="bg-gray-700 text-white py-2 px-4 rounded-md hover:bg-gray-600"
          >
            Previous
          </button>
        )}
        {step < 2 ? (
        <div className="flex flex-col w-full items-center">
        <button
          type="button"
          onClick={handleNextStep}
          className="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600 mb-4"
        >
          Continue
        </button>
        <div className="flex justify-center">
          <p className="mx-2">Already have an account?</p>
          <Link href="/signin" className="text-blue-900">Login</Link>
        </div>
      </div>
        ) : (
          <button
            type="submit"
            className="bg-blue-700 text-white py-2 px-4 rounded-md hover:bg-green-600"
          >
            Sign Up
          </button>
        )}
      </div>
      <ToastContainer />
    </form>
  </div>
);
};

export default SignupForm;