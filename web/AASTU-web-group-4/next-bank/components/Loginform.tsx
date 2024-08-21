"use client";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import Link from "next/link";
import Image from "next/image";
import {
  AtSymbolIcon,
  KeyIcon,
  ExclamationCircleIcon,
  ArrowPathIcon,
} from "@heroicons/react/24/outline";
import { creditcardstyles, colors ,logo } from "../constants/index";
import Cookie from 'js-cookie';
import { loginUser } from '@/services/authentication';


const LoginForm: React.FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();
  
  const [isLoading, setIsLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const onSubmit = async (data: any) => {
    setIsLoading(true);
      try {
        const loggedInUser = await loginUser(data);
        Cookie.set('accessToken', loggedInUser.data.access_token);
        Cookie.set('refreshToken', loggedInUser.data.refresh_token);
        console.log("successful login:", loggedInUser);
      } catch (error) {
        console.error('Login Error:', error);
      }
    };

  return (
    <div className="flex items-center justify-center h-[50vh]">
      <div className="flex-col items-center justify-center h-24 w-[50vh]">
        <form onSubmit={handleSubmit(onSubmit)} className="p-4 rounded-2xl">
          <div className="flex flex-col justify-center items-center my-4">
            <Image
              src={logo.icon}
              alt="next logo"
              width={150}
              height={50}
              className="h-30 w-40 m-auto my-4"
            />
            <h1 className="font-bold text-4xl text-gray-700 font-serif">NEXT BANK</h1>
          </div>

          <div className="my-10">
            <div>
              <label htmlFor="userName" className="block font-bold mb-2 text-gray-700">
                UserName
              </label>
              <input
                id="userName"
                type="text"
                placeholder="Username"
                {...register("userName", { required: "Username is required" })}
                className="w-full m-auto border-gray-200 border-2 rounded-lg shadow-sm focus:border-indigo-500 focus:ring-indigo-500 h-14 px-2.5"
              />
              {errors.userName && (
                <div className="flex gap-1">
                  <ExclamationCircleIcon className="h-5 w-5 text-red-500" />
                  <p className="text-red-500">{errors.userName.message as string}</p>
                </div>
              )}
            </div>
          </div>

          <div className="my-10">
            <div>
              <label htmlFor="password" className="block font-bold mb-2 text-gray-700">
                Password
              </label>
              <input
                id="password"
                type="password"
                placeholder="Password"
                {...register("password", { required: "Password is required" })}
                className="w-full m-auto border-gray-200 border-2 rounded-lg shadow-sm focus:border-indigo-500 focus:ring-indigo-500 h-14 px-2.5"
              />
              {errors.password && (
                <div className="flex gap-1">
                  <ExclamationCircleIcon className="h-5 w-5 text-red-500" />
                  <p className="text-red-500">{errors.password.message as string}</p>
                </div>
              )}
            </div>
          </div>

          {errorMessage && (
            <div className="text-red-500 text-center mb-4">
              {errorMessage}
            </div>
          )}

          <button
            type="submit"
            className={`${colors.blue} text-white px-4 py-2 mt-4 w-full rounded-3xl text-xl`}
          >
            {isLoading ? (
              <ArrowPathIcon className="h-5 w-5 animate-spin mr-2 text-white" />
            ) : (
              "Login"
            )}
          </button>

          <div className="my-14 flex flex-col items-center text-xl">
            <p>
              Don&apos;t have an account?{" "}
              <span className={`${colors.textblue} font-medium text-xl`}>
                <Link href="/signup">Sign Up</Link>
              </span>
            </p>
            <span className={`${colors.textblue} font-medium text-xl py-6`}>
              <Link href="/forgotpassword">Forgot password?</Link>
            </span>
          </div>
        </form>
      </div>
    </div>
  );
};

export default LoginForm;
