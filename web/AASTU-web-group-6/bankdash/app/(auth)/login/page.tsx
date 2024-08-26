import React from "react";
import LoginForm from "@/app/components/Forms/LoginForm";
import Image from "next/image";
import imgUrl from "../../../public/images/vackground-com-SfDofjXtxHE-unsplash.jpg";
const Login = () => {
  return (
    <div className="py-10  h-screen flex justify-center items-center content-center">
      <LoginForm />
    </div>
  );
};

export default Login;
