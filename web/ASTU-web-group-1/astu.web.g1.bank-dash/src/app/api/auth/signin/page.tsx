import LoginForm from "@/components/Form/AuthForm/LoginForm";
import React from "react";
import Image from "next/image";
import AuthProvider from "@/components/Formx/AuthProvider";

export default function Page() {
  return (
    <>
      <div className="flex justify-around items-center min-h-screen">
        <div className="hidden md:w-1/2 minrelative lg:flex items-center justify-center">
          <Image
            src="/assets/images/welcome-page.png"
            alt="hello"
            width={500}
            height={500}
            className="object-cover"
          />
        </div>
        <div className="min-h-[100vh] lg:bg-slate-200 w-full px-10 sm:w-1/2 lg:p-6 flex items-center justify-center ">
          <AuthProvider>
            <LoginForm />
          </AuthProvider>
        </div>
      </div>
    </>
  );
}
