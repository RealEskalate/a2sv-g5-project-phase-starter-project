"use client";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { useState } from "react";
import { useUserLoginMutation } from "@/redux/api/authentication-controller";
import { signIn } from "next-auth/react";
import ErrorMessage from "@/components/Message/ErrorMessage";

interface FormValues {
  userName: string;
  password: string;
}

const Login = () => {
  const router = useRouter();
  const [login, { isLoading }] = useUserLoginMutation();
  const { register, handleSubmit, formState } = useForm<FormValues>();
  const { errors } = formState;
  const [loginErrorMessage, setErrorMessage] = useState("");
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
      if (response?.ok) {
        setErrorMessage("");
        router.push("/");
      } else {
        setErrorMessage("Invalid email or password");
      }
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
            className="flex flex-col gap-9"
          >
            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[14px] font-epilogue font-semibold leading-[22px] text-[#515B6F]">
                Username
              </label>
              <input
                className="w-[350px] h-[45px] px-[14px] py-[10px] gap-[8px] rounded-xl border border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your username"
                type="text"
                {...register("userName", { required: "Username is required" })}
              />
              <ErrorMessage message={errors.userName?.message} />
            </div>
            <div className="w-[350px] h-[60px] flex flex-col gap-2">
              <label className="text-[16px] font-epilogue font-semibold leading-[25.6px] text-[#515B6F]">
                Password
              </label>
              <input
                className="w-[350px] h-[60px] px-[16px] py-[12px] gap-[10px] rounded-xl border border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
                placeholder="Enter your password"
                type="password"
                {...register("password", {
                  required: {
                    value: true,
                    message: "Password is required",
                  },
                  minLength: {
                    value: 6,
                    message: "Password must be at least 6 characters",
                  },
                })}
              />
              <div>
              <ErrorMessage message={errors.password?.message} />
              {!errors.password && loginErrorMessage && (
              <ErrorMessage message={loginErrorMessage} />

            )}
              </div>
              
            </div>
            <button
              className="bg-[#4640DE] text-white rounded-[80px] px-[24px] py-[12px] w-max-[408px] h-[50px] mt-3"
              type="submit"
            >
              Continue
            </button>
          </form>

          <div className="flex items-center -mt-4 gap-2 mx-auto">
            <p className="text-sm font-epilogue font-normal text-[#202430] opacity-[0.7]">
              Don't have an account?
            </p>
            <Link
              href="/auth/signup"
              className="hover:underline text-[#4640DE] font-medium"
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
