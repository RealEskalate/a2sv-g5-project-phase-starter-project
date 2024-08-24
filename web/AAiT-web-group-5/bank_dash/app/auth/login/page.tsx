"use client";

import Link from "next/link";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Inter } from "next/font/google";
import { signIn } from "next-auth/react";
import { useRouter } from "next/navigation";
import notify from "@/components/notify";
import { useState } from "react";
import { ToastContainer } from "react-toastify"; // Import ToastContainer

const schema = z.object({
  userName: z.string().min(1, { message: "User name field is required" }),
  password: z.string().min(1, { message: "Password field is required" }),
});

type FormData = z.infer<typeof schema>;
const inter = Inter({ subsets: ["latin"] });

interface SignInProps {
  onClose: () => void;
}

const SignIn: React.FC<SignInProps> = ({ onClose }) => {
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
      router.push("/dashboard");
    }
    setLoading(false);
  };

  return (
    <div
      onClick={onClose}
      className="fixed inset-0 bg-gray-800 bg-opacity-70 w-full h-full flex justify-center items-center z-[1000]"
      aria-modal="true"
      role="dialog"
    >
      <div
        className="relative w-full max-w-md rounded-lg p-8 bg-white shadow-lg"
        onClick={(e) => e.stopPropagation()}
      >
        <div className={`${inter.className} flex flex-col gap-6`}>
          <div className="text-center text-2xl font-bold text-gray-800">
            Welcome Back
          </div>
          <form
            onSubmit={handleSubmit(onSubmit)}
            className="flex flex-col gap-4"
            noValidate
          >
            <div>
              <label
                htmlFor="userName"
                className="font-medium text-gray-700"
              >
                User Name
              </label>
              <input
                id="userName"
                type="text"
                placeholder="Enter your user name"
                className="mt-2 w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:outline-none"
                {...register("userName")}
                disabled={loading}
                aria-invalid={!!errors.userName}
                aria-describedby="userNameError"
              />
              {errors.userName && (
                <p
                  id="userNameError"
                  className="text-red-600 text-sm mt-1"
                  aria-live="polite"
                >
                  {errors.userName.message}
                </p>
              )}
            </div>
            <div>
              <label
                htmlFor="password"
                className="font-medium text-gray-700"
              >
                Password
              </label>
              <input
                id="password"
                type="password"
                placeholder="Enter password"
                className="mt-2 w-full px-4 py-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500 focus:outline-none"
                {...register("password")}
                disabled={loading}
                aria-invalid={!!errors.password}
                aria-describedby="passwordError"
              />
              {errors.password && (
                <p
                  id="passwordError"
                  className="text-red-600 text-sm mt-1"
                  aria-live="polite"
                >
                  {errors.password.message}
                </p>
              )}
            </div>
            <button
              type="submit"
              className="mt-4 w-full py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors duration-200"
              disabled={loading}
              aria-busy={loading}
            >
              {loading ? (
                <div className="flex justify-center">
                  <span className="inline-flex items-center justify-center space-x-1">
                    <span className="w-2 h-2 bg-white rounded-full animate-bounce"></span>
                    <span className="w-2 h-2 bg-white rounded-full animate-bounce [animation-delay:0.1s]"></span>
                    <span className="w-2 h-2 bg-white rounded-full animate-bounce [animation-delay:0.2s]"></span>
                  </span>
                </div>
              ) : (
                "Login"
              )}
            </button>
          </form>
          <div className="text-center text-sm mt-4">
            {"Don't have an account? "}
            <Link href="/signup" className="font-medium text-blue-600 hover:underline">
              Register
            </Link>
          </div>
        </div>
        {/* Include the ToastContainer for displaying notifications */}
        <ToastContainer />
      </div>
    </div>
  );
};

export default SignIn;
