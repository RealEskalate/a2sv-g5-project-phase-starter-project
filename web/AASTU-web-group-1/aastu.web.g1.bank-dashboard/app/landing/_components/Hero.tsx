import { Button } from "@/components/ui/button";
import { motion } from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import React from "react";

const Hero = () => {
  return (
    <div className="flex flex-col justify-center items-center text-center pt-20 bg-gray-50">
      <div className="w-[60%]">
        <h1 className="text-5xl pb-4">
          Banking Solutions Tailored for growing companies
        </h1>
        <p className="text-xl text-gray-600 pb-5">
          The all-in-one financial planning platform for employees, experience
          the well being benefit that is actually being used.
        </p>
        <Button className="rounded-full bg-[#1814F3] mb-5 animate-bounce">
          <Link href="/auth/sign-up">Get an Account</Link>
        </Button>
      </div>
      <img
        src="/icons/dashImage.png"
        alt="logo"
        className="w-[80%] rounded-3xl border-gray-300 p-2 border-b-0 -mb-5 border-t-0"
      />
    </div>
  );
};

export default Hero;
