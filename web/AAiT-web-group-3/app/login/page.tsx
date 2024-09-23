"use client";
import React from "react";
import { useForm } from "react-hook-form";
import { redirect } from "next/navigation";
import { signIn, useSession } from "next-auth/react";
import { useDispatch } from "react-redux";
import { AppDispatch, useAppSelector } from "@/lib/store";
import Link from "next/link";
import Image from "next/image";
import {
  updateIsLoading,
  updateErrorMessage,
  updateAccessToken,
} from "@/lib/features/auth/authSlice";

const Login: React.FC = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const dispatch = useDispatch<AppDispatch>();

  const errorMessage = useAppSelector(
    (state) => state.authReducer.value.errorMessage
  );
  const isLoading = useAppSelector(
    (state) => state.authReducer.value.isLoading
  );

  const { data: session } = useSession({
    required: false,
  });

  if (session) {
    redirect("/");
  }

  const onSubmit = async (data: any) => {
    dispatch(updateIsLoading(true));
    dispatch(updateErrorMessage(null));

    try {
      if (data.email && data.password) {
        const result = await signIn("akil-login", {
          redirect: false,
          email: data.email,
          password: data.password,
        });

        if (!result) throw new Error("Undefined result from server");

        if (result.error) {
          if (result.status === 401 || result.status === 404) {
            dispatch(updateErrorMessage("Invalid Credentials"));
          }
        } else {
          redirect("/");
        }
      } else {
        dispatch(updateErrorMessage("Invalid Credentials"));
      }
    } catch (err) {
      dispatch(
        updateErrorMessage(
          "Something went Wrong, Invalid response from server."
        )
      );
    } finally {
      dispatch(updateIsLoading(false));
    }
  };

  return (
    <div className="flex items-center justify-center">
      <div className="w-full flex items-center justify-center">
        <Image
          src="https://dododoyo.github.io/image/next-auth/loginPage.svg"
          alt="Loin Image"
          className="hidden lg:block lg:w-1/2 m-30 p-28"
          width={20}
          height={20}
        />

        <div className="w-full lg:w-1/2 md:w-1/2 sm:w-3/5 px-8 py-12 mx-20 rounded-lg">
          <h2 className="text-4xl font-extrabold text-gray-800 mb-8 text-center">
            Welcome Back,
          </h2>
          {errorMessage && (
            <p className="text-red-500 text-xl text-center italic mb-4">
              {errorMessage}
            </p>
          )}
          <form onSubmit={handleSubmit(onSubmit)} className="w-full">
            <div className="mb-6">
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
                  {errors?.email?.message?.toString()}
                </p>
              )}
            </div>
            <div className="mb-8">
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
                  {errors?.password?.message?.toString()}
                </p>
              )}
            </div>
            <button
              type="submit"
              style={{
                backgroundColor: "#4640DE",
              }}
              disabled={isLoading}
              className={`w-full hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-3xl focus:outline-none focus:shadow-outline 
            `}
            >
              {isLoading ? "Loading . . ." : "Login"}
            </button>
          </form>

          {errors.email && "message" in errors.email && (
            <p className="text-red-500 text-xs italic mt-2">
              {errors?.email?.message?.toString()}
            </p>
          )}

          <p className="mt-8 text-gray-600">
            {`Don't have an account ?   `}
            <Link
              href="/signup"
              className="font-medium text-customBlue hover:text-blue-700"
            >
              Sign Up
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
};

export default Login;
