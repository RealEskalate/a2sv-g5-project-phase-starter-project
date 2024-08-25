"use client";
import React from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import Link from "next/link";
import logo from "../../../public/images/Logo.svg";
import Image from "next/image";
import { signIn } from "next-auth/react";

interface formtype {
  userName: string;
  password: string;
}

const Page = () => {
  const router = useRouter();
  const form = useForm<formtype>();
  const { register, formState, handleSubmit } = form;
  const { errors } = formState;


  const onSubmit = async (data: formtype) => {
    const result = await signIn("credentials", {
      redirect: false,
      userName: data.userName,
      password: data.password,
    });

    if (result?.ok) {
      router.push(`/`)
    } else {
      alert('Invalid Information')
    }
  };

  return (
    <div className="flex justify-center">
      {/* using react-hook-form to handle signup form */}

      <form
        onSubmit={handleSubmit(onSubmit)}
        className=" px-6 space-y-5 w-4/12 mt-16"
        action=""
      >
        <div className="flex mx-2 scale-150 mb-10 flex-col space-y-6 items-center">
          <div className=" md:block hidden">
            <Image src={logo} className="ml-1" alt="LOGO" />
          </div>
        </div>

        <div className="flex justify-center w-full">
          <h1 className=" text-[#202430] font-poppins text-3xl font-black">
            Welcome Back,
          </h1>
        </div>
        <hr />

        <div className="space-y-4">
          <div className="flex space-y-2 flex-col">
            <label
              className="text-[#515B6F] ml-1 epi text-sm font-semibold"
              htmlFor="email"
            >
              User Name
            </label>
            <input
              id="a"
              className="bg-white rounded-lg border p-3"
              type="text"
              placeholder="Enter email address"
              {...register("userName", { required: "Email is required" })}
            />
          </div>
          {errors.userName ? (
            <p className="text-red-500 text-sm">{errors.userName.message}</p>
          ) : null}

          <div className="flex space-y-2 flex-col">
            <label
              className="text-[#515B6F] ml-1 epi text-sm font-semibold"
              htmlFor="email"
            >
              Password
            </label>
            <input
              id="b"
              className="bg-white rounded-lg border p-3"
              type="password"
              placeholder="Enter password"
              {...register("password", { required: "Password is required" })}
            />
          </div>
          {errors.password ? (
            <p className="text-red-500 text-sm">{errors.password.message}</p>
          ) : null}
        </div>
        <button
          onClick={handleSubmit(onSubmit)}
          className="w-full py-4 text-white epi rounded-full bg-indigo-900"
        >
          Login
        </button>

        <h2 className="text-base epi">
          Don&apos;t have an account?{" "}
          <Link href={`/auth/signup-first`}>
            <span className="font-semibold ml-1 text-base text-blue-900">
               Sign Up
            </span>
          </Link>
        </h2>
      </form>
    </div>
  );
};

export default Page;
