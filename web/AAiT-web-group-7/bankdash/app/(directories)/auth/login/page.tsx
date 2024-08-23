"use client";
import Link from "next/link";
import { useForm } from "react-hook-form";

interface FormValues {
  email: string;
  password: string;
}

const Login = () => {
  const { register, handleSubmit, formState } = useForm<FormValues>();
  const { errors } = formState;
  const onSubmit = async (inputData: FormValues) => {
    console.log("inputData", inputData);
  };

  return (
    <div className="flex  gap-10 p-11 ">
      <div className="w-fit h-fit  flex flex-col gap-[24px]">
        <h1 className="text-[32px] font-poppins font-black leading-[38.4px] text-center text-[#202430] w- h-[38px]">
          Welcome Back,
        </h1>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-[40px]"
        >
          <div className="w-[408px] h-[68px] flex flex-col gap-2">
            <label className="text-[16px] font-epilogue font-semibold leading-[25.6px] text-[#515B6F]">
              Email Address
            </label>
            <input
              className="w-[408px] h-[50px] px-[16px] py-[12px] gap-[10px] rounded-[7px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
              placeholder="Enter your email address"
              type="email"
              {...register("email", {
                required: "Email Address is required",
                pattern: {
                  value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                  message: "Invalid email format",
                },
              })}
            />
            <p className="text-center text-[12px] text-red-600">
              {errors.email?.message}
            </p>
          </div>
          <div className="w-[408px] h-[68px] flex flex-col gap-2">
            <label className="text-[16px] font-epilogue font-semibold leading-[25.6px] text-[#515B6F]">
              Password
            </label>
            <input
              className="w-[408px] h-[50px] px-[16px] py-[12px] gap-[10px] rounded-[7px] border-[1px] border-solid border-[#D6DDEB] focus:outline-none focus:ring-[1px] focus:ring-[#4640DE]"
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

        <div className="w-fit h-fit flex gap-3">
          <p className="text-[16px] font-epilogue font-normal leading-[25.6px] text-[#202430] w-[203px] h-[26px] opacity-[0.7]">
            Don't have an account?
          </p>
          <Link
            href="/auth/signup"
            className="text-[16px] font-inter font-semibold leading-[24px] text-[#4640DE] h-[24px]"
          >
            Sign up
          </Link>
        </div>
      </div>
      <div className="w-full">
        <img src="/bankloin.jpg" alt="login" className="object-cover" />
      </div>
    </div>
  );
};

export default Login;
