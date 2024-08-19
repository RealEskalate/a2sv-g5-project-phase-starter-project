"use client";

import Link from "next/link";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Inter } from "next/font/google";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
import notify from "@/utils/notify";

const schema = z.object({
  userName: z.string().min(1, { message: "User name field is required" }),
  password: z.string().min(1, { message: "Password field is required" }),
});

type FormData = z.infer<typeof schema>;
const inter = Inter({ subsets: ["latin"] });

const SignIn = () => {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ resolver: zodResolver(schema) });

  const onSubmit = async (data: any) => {
    const res = await signIn("credentials", {
      userName: data.userName,
      password: data.password,
      redirect: false,
    });
    if (!res?.ok) {
      notify.error("Invalid Credentials");
      router.push("/logIn");
    } else {
      console.log("response: ", res);
      notify.success("Successfully logged in");
      router.push("/");
    }
  };
  return (
    <div
      className={`${inter.className} flex justify-center pt-[34px] pb-[50px] rounded-3xl bg-white`}
    >
      <div className="flex flex-col gap-6">
        <div className="flex flex-col gap-6">
          <div className={`text-center text-[32px] font-black text-[#25324B]`}>
            Welcome Back,
          </div>
        </div>
        <div className="flex justify-between items-center">
          <div className="border-[1px] w-2/6 h-0 border-[#D6DDEB]"></div>

          <div className="border-[1px] w-2/6 h-0 border-[#D6DDEB]"></div>
        </div>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col gap-[22px]"
          noValidate
        >
          <label
            htmlFor="userName"
            className={`font-semibold text-base text-[#515B6F]`}
          >
            User Name
          </label>
          <input
            type="text"
            placeholder="Enter your user name"
            className="border-[1px] border-[#DFEAF2] px-4 py-3 rounded-lg focus:outline-none"
            {...register("userName")}
          />
          {errors.userName && (
            <p className={` font-semibold text-base text-red-700`}>
              {errors.userName.message}
            </p>
          )}

          <label
            htmlFor="password"
            className={` font-semibold text-base text-[#515B6F]`}
          >
            Password
          </label>
          <input
            type="password"
            placeholder="Enter password"
            className="border-[1px] border-[#DFEAF2] px-4 py-3 rounded-lg focus:outline-none"
            {...register("password")}
          />
          {errors.password && (
            <p className={` font-semibold text-base text-red-700`}>
              {errors.password.message}
            </p>
          )}
          <button
            type="submit"
            className={` font-bold text-white text-center bg-[#1814F3] px-6 py-3 rounded-3xl`}
          >
            Login
          </button>
        </form>
        <div className={` font-normal`}>
          {`Don't have an account? `}
          <Link href="/SignUp" className="font-semibold text-[#1814F3]">
            Register
          </Link>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
