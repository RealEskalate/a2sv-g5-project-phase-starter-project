"use client";
import { useForm } from "react-hook-form";
import { DevTool } from "@hookform/devtools";
import { BsExclamationCircle } from "react-icons/bs";
import { FaGoogle, FaGithub } from "react-icons/fa";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
import { useState } from "react";

type FormValues = {
  username: string;
  password: string;
};

const Contact = () => {
  const [loading, setLoading] = useState(false);
  const form = useForm<FormValues>();
  const { register, control, handleSubmit, formState } = form;
  const { errors } = formState;
  const router = useRouter();

  const onSubmit = async (data: FormValues) => {
    setLoading(true);
    try {
      await signIn("credentials", {
        redirect: true,
        userName: data.username,
        password: data.password,
      });
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  const route = () => {
    router.push("/api/auth/signup");
  };

  return (
    <div className="flex justify-center min-h-screen dark:bg-[#1e1e2e]">
      <div className="flex flex-col items-center justify-center w-1/2 gap-10">
        <div className="flex justify-center">
          <span className="text-4xl mt-4 font-black text-[#202430] dark:text-[#cdd6f4]">
            Welcome
          </span>
        </div>
        <div className="flex justify-between w-2/3">
          <span className="border-b border-gray-300 text-gray-300 w-1/3 dark:text-gray-600"></span>
          <span className="border-b border-gray-300 text-gray-300 w-1/3 dark:text-gray-600"></span>
        </div>

        <div className="card-body w-2/3">
          <div className="flex flex-col justify-center gap-10 w-full">
            <form onSubmit={handleSubmit(onSubmit)} noValidate className="flex flex-col gap-5">
              <div className="mb-3 flex flex-col gap-2">
                <label htmlFor="username" className="text-[#515B6F] font-semibold">
                  Username
                </label>
                <input
                  type="text"
                  className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                  id="username"
                  placeholder="john@example.com"
                  {...register("username")}
                />
                {errors.username && (
                  <div className="flex flex-row align-middle mt-2">
                    <BsExclamationCircle className="text-red-500 mt-1.5 small" />
                    <p className="text-red-500 px-3">{errors.username.message}</p>
                  </div>
                )}
              </div>
              <div className="mb-3 flex flex-col gap-2">
                <label
                  htmlFor="password"
                  className="text-[#515B6F] font-semibold"
                >
                  Password
                </label>
                <input
                  type="password"
                  className="border border-gray-400 rounded-lg py-2 px-5 w-full dark:border-gray-600 dark:bg-[#313244] dark:text-[#cdd6f4] dark:focus:bg-[#313244] dark:focus:border-[#4640DE] dark:focus:text-[#cdd6f4]"
                  id="password"
                  placeholder="********"
                  {...register("password", {
                    required: "Password Field is Required",
                  })}
                />
                {errors.password && (
                  <div className="flex flex-row align-middle mt-2">
                    <BsExclamationCircle className="text-red-500 mt-1.5 small" />
                    <p className="text-red-500 px-3">
                      {errors.password.message}
                    </p>
                  </div>
                )}
              </div>
              <div className="">
                <button
                  type="submit"
                  className="border border-gray-400 bg-[#4640DE] hover:bg-[#2721dd] text-white font-bold rounded-3xl py-2 px-5 w-full dark:border-gray-600 dark:bg-[#5857ED] dark:hover:bg-[#4640DE] dark:text-white flex justify-center items-center"
                  disabled={loading}
                >
                  {loading ? (
                    <svg
                      className="animate-spin h-5 w-5 mr-3"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        className="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        strokeWidth="4"
                      ></circle>
                      <path
                        className="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8v8z"
                      ></path>
                    </svg>
                  ) : (
                    "Login"
                  )}
                </button>
              </div>
            </form>
          </div>
          <div className="py-2 px-1">
            <span className="text-[#515B6F]">Don’t have an account?</span>
            <button onClick={route} className="text-[#4640DE] font-bold px-2">
              Sign Up
            </button>
          </div>
          <DevTool control={control}></DevTool>
        </div>
      </div>
    </div>
  );
};

export default Contact;
