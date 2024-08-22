import { useUser } from "@/contexts/UserContext";
import React from "react";
import Header from "./_components/Header";
import Hero from "./_components/Hero";

const Landing = () => {
  return (
    <div
      style={{ fontFamily: '"Poppins", sans-serif' }}
      className="absolute inset-0 -z-10 h-full w-full bg-white bg-[radial-gradient(#e5e7eb_1px,transparent_1px)] [background-size:16px_16px]"
    >
      <Header />
      <Hero />
    </div>
  );
};

export default Landing;
