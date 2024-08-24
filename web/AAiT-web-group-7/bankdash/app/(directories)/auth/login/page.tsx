"use client";
import Link from "next/link";
import { useForm } from "react-hook-form";
import { useUserLoginMutation } from "@/redux/api/authentication-controller";
import { signIn } from "next-auth/react";

interface FormValues {
  userName: string;
  password: string;
}

const Login = () => {
  const [login, { isLoading }] = useUserLoginMutation();
  const { register, handleSubmit, formState } = useForm<FormValues>();
  const { errors } = formState;
  const onSubmit = async (inputData: FormValues) => {
    console.log("inputData", inputData);

    try {
      const response = await signIn("credentials", {
        userName: inputData.userName,
        password: inputData.password,
        redirect: false,
        callbackUrl: "/",
      });
      console.log("response", response);
    } catch (error) {
      console.log("error", error);
    }
  };

  return (
    <div className="flex items-center justify-center h-screen">
      <div className="w-[70%] flex items-center justify-between gap-10">
      <div className="w-fit h-fit  flex flex-col gap-8">
        <h1 className="text-[32px] font-poppins font-black leading-[38.4px] text-center text-[#4640DE] w- h-[38px]">
          Welcome Back,
        </h1>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-[40px]"
        >
          <div className="w-[350px] h-[60px] flex flex-col gap-2">
            <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
              Username
            </label>
            <input
              className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-[6px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
              placeholder="Choose a username"
              type="text"
              {...register("userName", { required: "Username is required" })}
            />
            <p className="error text-[12px] text-center text-red-700">
              {errors.userName?.message}
            </p>
          </div>
          <div className="w-[350px] h-[60px] flex flex-col gap-2">
            <label className="text-[16px] font-epilogue font-semibold leading-[25.6px] text-[#515B6F]">
              Password
            </label>
            <input
              className="w-[350px] h-[60px] px-[16px] py-[12px] gap-[10px] rounded-[7px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
              placeholder="Enter your password"
              type="password"
              {...register("password", {
                required: "Password is required",
                minLength: {
                  value: 6,
                  message: "Password must be at least 6 characters long",
                },
              })}
            />
            <p className="text-center text-[12px] text-red-600">
              {errors.password?.message}
            </p>
          </div>

          <button
            className="bg-[#4640DE] text-white rounded-[80px] px-[24px] py-[12px] w-max-[408px] h-[50px]"
            type="submit"
          >
            Continue
          </button>
        </form>

        <div className="flex items-center gap-2 mx-auto">
          <p className="text-sm font-epilogue font-normal text-[#202430] opacity-[0.7]">
            Don&apos;t have an account?
          </p>
          <Link
            href="/auth/signup"
            className="hover:underline text-[#4640DE]"
          >
            Sign up
          </Link>
        </div>
      </div>
      <div className="w-full">
        <img src="/bankloin.jpg" alt="login" className="rounded-full" />
      </div>
      </div>
    </div>
  );
};

export default Login;
