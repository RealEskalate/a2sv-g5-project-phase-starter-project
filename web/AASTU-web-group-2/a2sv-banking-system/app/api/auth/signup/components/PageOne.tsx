import React from "react";
import DatePicker from "react-datepicker";
import { useFormContext, Controller } from "react-hook-form";
import { BsExclamationCircle } from "react-icons/bs";
import "react-datepicker/dist/react-datepicker.css";

const PageOne = () => {
  const { control } = useFormContext();

  return (
    <div className="flex justify-center min-h-screen dark:bg-[#1e1e2e]">
      <div className="flex flex-col gap-10 w-1/2">
        {/* Welcome Text */}
        <h1 className="text-4xl font-black text-[#202430] dark:text-[#cdd6f4] mt-4 justify-center">
          Basics
        </h1>

        {/* Form Container */}
        <form className="flex flex-col gap-5">
          {/* Name Field */}
          <div className="mb-3 flex flex-col gap-2">
            <label
              htmlFor="name"
              className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]"
            >
              Name
            </label>
            <Controller
              name="name"
              control={control}
              defaultValue=""
              rules={{ required: "Name is required" }}
              render={({ field, fieldState: { error } }) => (
                <>
                  <input
                    {...field}
                    type="text"
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                    placeholder="Your Name"
                  />
                  {error && (
                    <div className="flex flex-row align-middle mt-2">
                      <BsExclamationCircle className="text-red-500 mr-2" />
                      <span className="text-red-500">{error.message}</span>
                    </div>
                  )}
                </>
              )}
            />
          </div>

          {/* Email Field */}
          <div className="mb-3 flex flex-col gap-2">
            <label
              htmlFor="email"
              className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]"
            >
              Email
            </label>
            <Controller
              name="email"
              control={control}
              defaultValue=""
              rules={{
                required: "Email is required",
                pattern: {
                  value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                  message: "Invalid email format",
                },
              }}
              render={({ field, fieldState: { error } }) => (
                <>
                  <input
                    {...field}
                    type="email"
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                    placeholder="john@example.com"
                  />
                  {error && (
                    <div className="flex flex-row align-middle mt-2">
                      <BsExclamationCircle className="text-red-500 mr-2" />
                      <span className="text-red-500">{error.message}</span>
                    </div>
                  )}
                </>
              )}
            />
          </div>

          {/* Date of Birth Field */}
          <div className="mb-3 flex flex-col gap-2">
            <label
              htmlFor="dateOfBirth"
              className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]"
            >
              Date of Birth
            </label>
            <Controller
              name="dateOfBirth"
              control={control}
              defaultValue={null}
              rules={{ required: "Date of Birth is required" }}
              render={({ field, fieldState: { error } }) => (
                <>
                  <DatePicker
                    selected={field.value}
                    onChange={(date) => field.onChange(date)}
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                  />
                  {error && (
                    <div className="flex flex-row align-middle mt-2">
                      <BsExclamationCircle className="text-red-500 mr-2" />
                      <span className="text-red-500">{error.message}</span>
                    </div>
                  )}
                </>
              )}
            />
          </div>

          {/* Username Field */}
          <div className="mb-3 flex flex-col gap-2">
            <label
              htmlFor="username"
              className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]"
            >
              Username
            </label>
            <Controller
              name="username"
              control={control}
              defaultValue=""
              rules={{
                required: "Username is required",
                minLength: {
                  value: 6,
                  message: "Username must be at least 6 characters",
                },
              }}
              render={({ field, fieldState: { error } }) => (
                <>
                  <input
                    {...field}
                    type="text"
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                    placeholder="Your Username"
                  />
                  {error && (
                    <div className="flex flex-row align-middle mt-2">
                      <BsExclamationCircle className="text-red-500 mr-2" />
                      <span className="text-red-500">{error.message}</span>
                    </div>
                  )}
                </>
              )}
            />
          </div>

          {/* Password Field */}
          <div className="mb-3 flex flex-col gap-2">
            <label
              htmlFor="password"
              className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]"
            >
              Password
            </label>
            <Controller
              name="password"
              control={control}
              defaultValue=""
              rules={{
                required: "Password is required",
                minLength: {
                  value: 6,
                  message: "Password must be at least 6 characters",
                },
              }}
              render={({ field, fieldState: { error } }) => (
                <>
                  <input
                    {...field}
                    type="password"
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                    placeholder="********"
                  />
                  {error && (
                    <div className="flex flex-row align-middle mt-2">
                      <BsExclamationCircle className="text-red-500 mr-2" />
                      <span className="text-red-500">{error.message}</span>
                    </div>
                  )}
                </>
              )}
            />
          </div>

          {/* Confirm Password Field */}
          <div className="mb-3 flex flex-col gap-2">
            <label
              htmlFor="confirmPassword"
              className="text-[#515B6F] font-semibold dark:text-[#cdd6f4]"
            >
              Confirm Password
            </label>
            <Controller
              name="confirmPassword"
              control={control}
              defaultValue=""
              rules={{
                required: "Confirm Password is required",
                validate: (value, { password }) =>
                  value === password || "Passwords must match",
              }}
              render={({ field, fieldState: { error } }) => (
                <>
                  <input
                    {...field}
                    type="password"
                    className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                    placeholder="********"
                  />
                  {error && (
                    <div className="flex flex-row align-middle mt-2">
                      <BsExclamationCircle className="text-red-500 mr-2" />
                      <span className="text-red-500">{error.message}</span>
                    </div>
                  )}
                </>
              )}
            />
          </div>
        </form>
      </div>
    </div>
  );
};

export default PageOne;
