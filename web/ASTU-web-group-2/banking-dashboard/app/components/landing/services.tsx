"use client";
import { useRouter } from "next/navigation";
import React, { useEffect, useRef, useState } from "react";
import { motion as m } from 'framer-motion';

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

const sectionVariants = {
  hidden: { opacity: 0, y: 60 },
  visible: { opacity: 1, y: 0 },
  exit: { opacity: 0, y: -60 },
};

const itemVariants = {
  hidden: { opacity: 0, scale: 0.9 },
  visible: { opacity: 1, scale: 1 },
};

const Services: React.FC = () => {
  const router = useRouter();
  const [isInView, setIsInView] = useState(false);
  const sectionRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsInView(true);
        }
      },
      { threshold: 0.1 }
    );

    if (sectionRef.current) {
      observer.observe(sectionRef.current);
    }

    return () => {
      if (sectionRef.current) {
        observer.unobserve(sectionRef.current);
      }
    };
  }, []);

  return (
    <m.section
      ref={sectionRef}
      variants={sectionVariants}
      initial="hidden"
      animate={isInView ? "visible" : "hidden"}
      exit="exit"
      transition={{ duration: 0.85, ease: "easeOut" }}
      id="services"
      className="w-full py-16 mt-24 pt-[80px]" 
    >
      <div className="text-center mb-12">
        <m.h1
          variants={sectionVariants}
          initial="hidden"
          animate={isInView ? "visible" : "hidden"}
          exit="exit"
          transition={{ duration: 0.85, ease: "easeOut" }}
          className="text-[#083E9E] text-3xl sm:text-4xl lg:text-5xl mt-10 font-extrabold"
        >
          Our Services
        </m.h1>
      </div>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 px-4 md:px-8 lg:px-12">
        {services.map((item, index) => (
          <m.div
            key={index}
            variants={itemVariants}
            initial="hidden"
            animate={isInView ? "visible" : "hidden"}
            exit="hidden"
            transition={{ duration: 0.6, delay: index * 0.2, ease: "easeOut" }}
            className="flex flex-col items-center p-8 rounded-3xl border border-dashed border-gray-400 bg-white shadow-md hover:shadow-lg transform hover:scale-105 transition-all duration-300"
          >
            <m.img
              src={item.icon}
              alt={item.title}
              className="w-24 h-24 mb-4"
              whileHover={{ scale: 1.1 }}
              transition={{ duration: 0.3 }}
            />
            <h2 className="text-xl lg:text-2xl text-[#0b1739] mb-2 text-center">{item.title}</h2>
            <p className="text-sm lg:text-md text-[#6F6969] mb-4 text-center">
              {item.description}
            </p>
            <m.button
              onClick={() => router.push('/login')}
              className="rounded-full px-4 py-2 border-[1px] border-dashed border-gray-400 bg-[#083E9E] text-white hover:bg-[#065B8F] transition-colors duration-300"
              whileHover={{ scale: 1.05 }}
              transition={{ duration: 0.3 }}
            >
              Apply now
            </m.button>
          </m.div>
        ))}
      </div>
      <div className="flex justify-center mt-12">
        <m.button
          whileHover={{ scale: 1.05, backgroundColor: "#065B8F" }}
          whileTap={{ scale: 0.95 }}
          className="bg-[#083E9E] w-60 text-white rounded-full px-6 py-3 border border-transparent hover:bg-[#065B8F] transition-colors duration-300"
        >
          View more
        </m.button>
      </div>
    </m.section>
  );
};

export default Services;
