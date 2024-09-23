"use client";
import React from "react";
import { useDispatch } from "react-redux";
import { useForm } from "react-hook-form";

import { signIn, useSession } from "next-auth/react";
import Link from "next/link";
import { redirect } from "next/navigation";
import { AppDispatch } from "@/lib/store";

const Signup: React.FC = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { data: session } = useSession({
    required: false,
  });

  if (session) {
    redirect("/");
  }

  const {
    register,
    watch,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = async (data: any) => {
    try {
      const result = await signIn("akil-signup", {
        redirect: false,
        firstName: data.firstName,
        lastName: data.lastName,
        userName: data.userName,
        email: data.email,
        password: data.password,
        confirmPassword: data.confirmPassword,
        city: data.city,
        country: data.country,
      });

      if (!result)
        throw new Error("Result is undefined. Please try again later.");
      // TODO: handle error and set errorMessage

      if (!result.ok) {
        // TODO: Handle Error and set Error Message
      } else {
        redirect("/");
      }
    } catch (err) {
      // TODO : Handle Error and set Error Messages
    }
  };

  return (
    <div className="flex items-center justify-center my-2">
      <div className="w-full md:w-2/3 lg:w-1/2  px-4 py-10 md:py-12 rounded-lg">
        <h1 className="text-3xl font-bold text-gray-800 my-2 text-center">
          Sign Up Today!
        </h1>
        <form onSubmit={handleSubmit(onSubmit)} className="w-full">
          {/* first and last name */}
          <div className="my-3">
            <div className="flex flex-wrap -mx-3">
              {/* firstname */}
              <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                <label
                  htmlFor="firstName"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  First Name
                </label>
                <input
                  type="text"
                  id="firstName"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter your first name"
                  {...register("firstName", {
                    required: "first name is required",
                  })}
                />
                {errors.firstName && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.firstName.message?.toString()}
                  </p>
                )}
              </div>
              {/* lastname */}
              <div className="w-full md:w-1/2 px-3">
                <label
                  htmlFor="lastName"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  Last Name
                </label>
                <input
                  type="text"
                  id="lastName"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter your last name"
                  {...register("LastName", {
                    required: "last name is required",
                  })}
                />
                {errors.lastName && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.lastName.message?.toString()}
                  </p>
                )}
              </div>
            </div>
          </div>

          {/* username and email */}
          <div className="my-3">
            <div className="flex flex-wrap -mx-3">
              {/* username */}
              <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                <label
                  htmlFor="userName"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  Username
                </label>
                <input
                  type="userName"
                  id="userName"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter Username"
                  {...register("userName", {
                    required: "userName address is required",
                  })}
                />
                {errors.userName && "message" in errors.userName && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.userName.message?.toString()}
                  </p>
                )}
              </div>
              {/* email */}
              <div className="w-full md:w-1/2 px-3">
                <label
                  htmlFor="email"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  Email Address
                </label>
                <input
                  type="email"
                  id="email"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter email address"
                  {...register("email", {
                    required: "Email address is required",
                    pattern: {
                      value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$/i,
                      message: "Invalid email address",
                    },
                  })}
                />
                {errors.email && "message" in errors.email && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.email.message?.toString()}
                  </p>
                )}
              </div>
            </div>
          </div>

          {/* Passwords */}
          <div className="my-3">
            <div className="flex flex-wrap -mx-3">
              {/* Password */}
              <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                <label
                  htmlFor="password"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  Password
                </label>
                <input
                  type="password"
                  id="password"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter password"
                  {...register("password", {
                    required: "Password is required",
                    minLength: {
                      value: 6,
                      message: "Password must be at least 6 characters long",
                    },
                  })}
                />
                {errors.password && "message" in errors.password && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.password.message?.toString()}
                  </p>
                )}
              </div>

              {/* Confirm Password */}
              <div className="w-full md:w-1/2 px-3">
                <label
                  htmlFor="confirmPassword"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  Confirm Password
                </label>
                <input
                  type="password"
                  id="confirmPassword"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter password"
                  {...register("confirmPassword", {
                    required: "Please confirm your password",
                    validate: (value) =>
                      value === watch("password") || "Passwords do not match",
                  })}
                />
                {errors.confirmPassword && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.confirmPassword.message?.toString()}
                  </p>
                )}
              </div>
            </div>
          </div>

          {/*Location  */}
          <div className="my-3">
            <div className="flex flex-wrap -mx-3">
              {/* city */}
              <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                <label
                  htmlFor="city"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  City
                </label>
                <input
                  type="text"
                  id="city"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter your City"
                  {...register("city", {
                    required: "City is required",
                  })}
                />
                {errors.city && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.city.message?.toString()}
                  </p>
                )}
              </div>
              {/* country */}
              <div className="w-full md:w-1/2 px-3">
                <label
                  htmlFor="country"
                  className="block text-gray-700 text-sm font-medium mb-2"
                >
                  Country
                </label>
                <input
                  type="text"
                  id="country"
                  className="w-full px-4 py-3 border rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                  placeholder="Enter your Country"
                  {...register("country", {
                    required: "Country is required",
                  })}
                />
                {errors.country && (
                  <p className="text-red-500 text-xs italic mt-2">
                    {errors.country.message?.toString()}
                  </p>
                )}
              </div>
            </div>
          </div>

          {/* Continue */}
          <button
            style={{
              backgroundColor: "#4640DE",
            }}
            type="submit"
            className="hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-3xl focus:outline-none focus:shadow-outline w-full"
          >
            Continue
          </button>
        </form>
        <p className="mx-2 my-5 text-md text-gray-600">
          Already have an account?{" "}
          <Link
            href="/login"
            style={{
              color: "#4640DE",
            }}
            className="font-medium text-blue-500 hover:text-blue-700"
          >
            Login
          </Link>
        </p>
        <p className="text-center mt-4 text-sm text-gray-500">
          {`By clicking "Continue", you acknowledge that you have read and
          accepted our  `}
          <a href="#" className="text-customBlue hover:underline">
            Terms of Service
          </a>{" "}
          and{" "}
          <a
            href="#"
            style={{
              color: "#4640DE",
            }}
            className="text-blue-500 hover:underline"
          >
            Privacy Policy
          </a>
          .
        </p>
      </div>
    </div>
  );
};

export default Signup;
