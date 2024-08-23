"use client";

import Link from "next/link";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Inter } from "next/font/google";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
import notify from "@/utils/notify";
import { useState } from "react";

const schema = z.object({
  userName: z.string().min(1, { message: "User name field is required" }),
  password: z.string().min(1, { message: "Password field is required" }),
});

type FormData = z.infer<typeof schema>;
const inter = Inter({ subsets: ["latin"] });

const SignIn = ({ onClose }: { onClose: () => void }) => {
  const [loading, setLoading] = useState(false);
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({ resolver: zodResolver(schema) });

  const onSubmit = async (data: FormData) => {
    setLoading(true);
    const res = await signIn("credentials", {
      userName: data.userName,
      password: data.password,
      redirect: false,
    });
    if (!res?.ok) {
      notify.error("Invalid Credentials");
    } else {
      notify.success("Successfully logged in");
      onClose(); // Close the modal after successful login
      router.push("/dashboard");
    }
    setLoading(false);
  };

  return (
    <div 
      onClick={onClose}
      className="fixed inset-0 backdrop-blur w-full h-full flex justify-center items-center z-[1000]"
    >
      <div 
        className="relative w-[500px] rounded-3xl p-[2rem] bg-white shadow-md border-[3px]" 
        onClick={(e) => e.stopPropagation()} // Prevent clicks inside the modal from closing it
      >
        <div className={`${inter.className} flex justify-center pt-8 pb-10 rounded-3xl`}>
          <div className="flex flex-col gap-6 w-full">
            <div className="text-center text-[32px] font-black text-[#25324B]">
              Welcome Back,
            </div>
            <div className="flex justify-between items-center my-4">
              <div className="border-t border-[#D6DDEB] w-2/6"></div>
              <div className="border-t border-[#D6DDEB] w-2/6"></div>
            </div>
            <form
              onSubmit={handleSubmit(onSubmit)}
              className="flex flex-col gap-6 w-full"
              noValidate
            >
              <div>
                <label
                  htmlFor="userName"
                  className="font-semibold text-base text-[#515B6F]"
                >
                  User Name
                </label>
                <input
                  id="userName"
                  type="text"
                  placeholder="Enter your user name"
                  className="text-[#515B6F] border focus:border-[#083E9E] border-[#DFEAF2] px-4 py-3 rounded-lg focus:outline-none w-full mt-2"
                  {...register("userName")}
                  disabled={loading}
                  aria-invalid={!!errors.userName}
                  aria-describedby="userNameError"
                />
                {errors.userName && (
                  <p
                    id="userNameError"
                    className="text-red-700 text-sm mt-1"
                  >
                    {errors.userName.message}
                  </p>
                )}
              </div>
              <div>
                <label
                  htmlFor="password"
                  className="font-semibold text-base text-[#515B6F]"
                >
                  Password
                </label>
                <input
                  id="password"
                  type="password"
                  placeholder="Enter password"
                  className="text-[#515B6F] border focus:border-[#083E9E] border-[#DFEAF2] px-4 py-3 rounded-lg focus:outline-none w-full mt-2"
                  {...register("password")}
                  disabled={loading}
                  aria-invalid={!!errors.password}
                  aria-describedby="passwordError"
                />
                {errors.password && (
                  <p
                    id="passwordError"
                    className="text-red-700 text-sm mt-1"
                  >
                    {errors.password.message}
                  </p>
                )}
              </div>
              <button
                type="submit"
                className="font-bold text-white text-center bg-[#083E9E] px-6 py-3 rounded-3xl mt-4"
                disabled={loading}
                aria-busy={loading}
              >
                {loading ? (
                  <div className="w-8 h-8 border-4 border-dashed rounded-full animate-spin [animation-duration:3s] border-white mx-auto"></div>
                ) : (
                  "Login"
                )}
              </button>
            </form>
            <div className="text-[#515B6F] text-center text-sm mt-4">
              Don't have an account?{" "}
              <Link href="/signup" className="font-semibold text-[#083E9E]">
                Register
              </Link>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SignIn;
