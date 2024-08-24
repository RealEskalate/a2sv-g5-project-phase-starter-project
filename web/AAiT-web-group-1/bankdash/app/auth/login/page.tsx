"use client";
import React from "react";
import { useForm } from "react-hook-form";
import { useSignInMutation } from "@/lib/redux/api/authApi";
import { useRouter } from "next/navigation";
import Link from "next/link";
import logo  from '../../../public/images/Logo.svg'
import { useDispatch } from "react-redux";
import Image from "next/image";
// import { authorize } from "../service/loginSlice";

interface formtype {
  email: string;
  password: string;
}

const Page = () => {
  const dispatch = useDispatch();

  const router = useRouter();
  const form = useForm<formtype>();
  const { control, register, formState, handleSubmit } = form;
  const { errors } = formState;

  // rtk query hook
  const [signIn, { data, isError, isLoading }] = useSignInMutation();

  if (isError) {
    return (
      <h1 className="text-center text-lg mt-72 text-red-700">
        There seems to be an error while logging in your information
      </h1>
    );
  }
  if (isLoading) {
    return <h1 className="text-center text-lg mt-72">Loading ....</h1>;
  }

  // handling the signin process and storing access token for further use
  const onSubmit = async (user: formtype) => {
    try {
      const res = await signIn(user);
      const { data } = res;
      console.log(data)
      console.log(res)
      // localStorage.setItem("accessToken", data.data.accessToken);
      // dispatch(authorize());
      // router.push(`/`);
    } catch (err) {
      console.log("error");
      // router.push(`/signup`);
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
          <h1 className=" text-[#202430] font-poppins text-4xl font-black">
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
              Email Address
            </label>
            <input
              id="a"
              className="bg-white rounded-lg border p-3"
              type="text"
              placeholder="Enter email address"
              {...register("email", { required: "Email is required" })}
            />
          </div>
          {errors.email ? (
            <p className="text-red-500 text-sm">{errors.email.message}</p>
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
          <Link href={`/signup`}>
            {" "}
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
