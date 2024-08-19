"use client";
import { useForm } from "react-hook-form";
import { DevTool } from "@hookform/devtools";
import { BsExclamationCircle } from "react-icons/bs";
import { FaGoogle, FaGithub } from "react-icons/fa";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
type FormValues = {
  username: string;
  password: string;
};

const Contact = () => {
  const form = useForm<FormValues>();
  const { register, control, handleSubmit, formState } = form;
  const { errors } = formState;
  const router = useRouter()
  
  const onSubmit = async (data: FormValues) => {
    await signIn("credentials", {
      redirect: true,
      callbackUrl: "/dashboard",
      userName: data.username,
      password: data.password,
    });
  };


  const route = () => {
    router.push("/api/auth/signup")
  }

  return (
    <div className="flex justify-center">
      <div className="flex flex-col items-center justify-center w-1/2 gap-10">
        <div className="flex justify-center">
          <span className="text-4xl mt-4 font-black text-[#202430]">
            Welcome Back,
          </span>
        </div>
        <div className="flex justify-between w-2/3">
          <span className="border-b border-gray-300 text-gray-300 w-1/3"></span>
          <span className="border-b border-gray-300 text-gray-300 w-1/3"></span>
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
                  className="border border-gray-400 rounded-lg py-2 px-5 w-full"
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
                  className="border border-gray-400 rounded-lg py-2 px-5 w-full"
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
                  className="border border-gray-400 bg-[#4640DE] hover:bg-[#2721dd] text-white font-bold rounded-3xl py-2 px-5 w-full"
                >
                  Login
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
