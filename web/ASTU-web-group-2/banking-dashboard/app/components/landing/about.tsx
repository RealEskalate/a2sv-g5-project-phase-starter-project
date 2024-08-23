"use client";
import Image from "next/image";
import React, { useEffect, useRef, useState } from "react";
import { motion as m } from 'framer-motion';

const sectionVariants = {
  hidden: { opacity: 0, y: 50 },
  visible: { opacity: 1, y: 0 },
  exit: { opacity: 0, y: -50 },
};

const itemVariants = {
  hidden: { opacity: 0, y: 30 },
  visible: { opacity: 1, y: 0 },
};

const About: React.FC = () => {
  const [isInView, setIsInView] = useState(false);
  const sectionRef = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setIsInView(true);
        } else {
          setIsInView(false);
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
      transition={{ duration: 0.75, ease: "easeOut" }}
      id="about"
      className="flex flex-col sm:flex-row justify-center items-start p-8 gap-32 pt-[80px] min-h-screen" // Adjust pt to ensure the section starts below the fixed navbar
    >
      <m.div
        variants={itemVariants}
        initial="hidden"
        animate={isInView ? "visible" : "hidden"}
        transition={{ duration: 0.75, ease: "easeOut" }}
        className="flex flex-col items-center justify-center w-full md:w-1/3 rounded-lg p-6"
      >
        <h1 className="text-blue-800 text-center mb-4 text-2xl md:text-3xl font-extrabold">
          About Us
        </h1>
        <div className="mb-4">
          <Image
            src="assets/landing/about.svg"
            width={0}
            height={0}
            alt="about"
            className="w-full h-auto max-w-xs md:max-w-md"
          />
        </div>
        <p className="text-md text-gray-700 text-center">
          Bank - Your trusted financial partner for loans. Quick approvals,
          competitive rates, and personalized solutions to meet your unique
          needs. Empowering you to achieve your financial goals. Apply online
          today!
        </p>
      </m.div>

      <m.div
        variants={itemVariants}
        initial="hidden"
        animate={isInView ? "visible" : "hidden"}
        transition={{ duration: 0.75, ease: "easeOut" }}
        className="flex flex-col items-center justify-center w-full sm:w-1/3 rounded-lg p-6 m-4"
      >
        <h1 className="text-blue-800 text-center mb-4 text-2xl md:text-3xl font-extrabold">
          Contact Us
        </h1>
        <input
          className="w-full text-center p-4 h-12 mb-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          name="name"
          placeholder="Full Name"
          type="text"
        />
        <input
          className="w-full text-center p-4 h-12 mb-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          name="email"
          placeholder="Email Address"
          type="email"
        />
        <textarea
          className="w-full h-48 text-center p-4 mb-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          name="message"
          placeholder="Message"
        />
        <button className="text-center h-12 rounded-md text-white w-1/3 bg-blue-800 hover:bg-blue-700 transition-colors duration-300">
          Send
        </button>
      </m.div>
    </m.section>
  );
};

export default About;
