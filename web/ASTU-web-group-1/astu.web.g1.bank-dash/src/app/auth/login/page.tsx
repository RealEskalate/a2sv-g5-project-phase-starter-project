import LoginForm from "@/components/Form/AuthForm/LoginForm";
import Image from "next/image";
import Link from "next/link";

const Login = () => {
  return (
    <div className="flex min-h-screen max-w-7xl mx-auto">
      <div className="flex-1  hidden lg:flex relative">
        <Image src="/assets/images/login.png" fill alt="" />
      </div>
      <div className=" flex-1 flex flex-col justify-center items-center w-full gap-3 px-5 md:px-20 ">
        <div className="flex flex-col items-center justify-center gap-6 md:px-16 w-full">
          <h1 className="font-Lato font-black text-4xl text-black">
            Welcome Back,
          </h1>
          <div className="w-full flex">
            <div className="flex-1 flex w-full h-[1px] bg-gray-500"></div>
            <div className="flex flex-1"> </div>
            <div className="flex-1 flex w-full h-[1px] bg-gray-500"></div>
          </div>
        </div>
        <LoginForm />
        <div className="flex gap-2 justify-start items-center w-full   md:px-16">
          <p className="text-center">Don't have an account? </p>
          <Link
            href="/auth/signup"
            className="text-blue font-epilogue font-epilogue600"
          >
            Sign Up
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Login;
