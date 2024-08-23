"use client"
import { useRouter } from "next/navigation";
import React from "react";

const services = [
  {
    title: "Personal loan",
    description:
      "Personal loans provide borrowers with flexibility in how they use the funds.",
    icon: "assets/landing/icon1.svg",
  },
  {
    title: "Business loan",
    description:
      "Business Loan Services provide financial assistance to businesses for various purposes.",
    icon: "assets/landing/icon2.svg",
  },
  {
    title: "Auto loan",
    description:
      "Auto Loan Services provide financing options for individuals and businesses to purchase a vehicle.",
    icon: "assets/landing/icon3.svg",
  },
];

const Services = () => {
  const router = useRouter()
  return (
    <section id="services" className="w-full  py-10 mt-20">
      <div className="text-center mb-10">
        <h1 className="text-[#083E9E] text-3xl sm:text-4xl lg:text-5xl font-extrabold">
          Our Services
        </h1>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 pl-5 pr-5 lg:grid-cols-3 gap-6 px-4 md:px-10">
        {services.map((item, index) => (
          <div
            key={index}
            className="flex flex-col items-center p-6 rounded-3xl border border-dashed border-gray-400 hover:shadow-lg"
          >
            <img src={item.icon} alt={item.title} className="w-24 h-24 mb-4" />
            <h2 className="text-2xl text-[#0b1739] mb-2">{item.title}</h2>
            <p className="text-md text-[#6F6969] mb-4 text-center">
              {item.description}
            </p>
            <button 
            onClick={()=>router.push('/login')}
            className="rounded-full px-4 py-2 border-[1px] border-dashed  border-gray-400  hover:bg-[#083E9E] hover:text-white">
              Apply now
            </button>
          </div>
        ))}
      </div>
      <div className="flex justify-center mt-10">
        <button className="bg-[#083E9E] w-60 text-white rounded-full px-6 py-3 border border-transparent hover:bg-[#083E9E]">
          View more
        </button>
      </div>
    </section>
  );
};

export default Services;