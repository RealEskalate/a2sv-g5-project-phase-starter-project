import Image from "next/image";
import React from "react";
import { motion as m } from 'framer-motion';

const LandingHome = () => {
  return (
    <div
      id="home"
      className="bg-[#083E9E] w-full h-[200px] md:h-[400px] lg:h-[600px] relative text-white flex"
    >
      <m.div 
        initial={{ opacity: 0, y: "40%" }}
        animate={{ opacity: 1, y: "0%" }}
        transition={{ duration: 0.75, ease: "easeOut" }}
        className="pt-12 sm:pt-20 pl-12 md:pl-20 lg:pt-32 lg:pl-32 flex flex-col md:w-[40%] z-50"
      >
        <m.span
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.75, ease: "easeOut", delay: 0.2 }}
          className="font-normal md:font-bold lg:font-extrabold text-[12px] md:text-[25px] lg:text-[35px]"
        >
          Start building your financial future today. Collect and manage your savings effortlessly
        </m.span>
        <m.span
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.75, ease: "easeOut", delay: 0.4 }}
          className="font-thin md:font:normal text-[5px] md:text-[15px]"
        >
          Join millions of agents saving funds for people who don't have access
          to their bank details and earn returns.
        </m.span>
      </m.div>

      <Image
        src={"assets/landing/net.svg"}
        width={0}
        height={0}
        alt="net"
        style={{ width: "100%", height: "100%" }}
        className="absolute bottom-0 right-0 z-0"
      />

      <m.section
        initial={{ opacity: 0, y: "40%", scale: 0.9 }}
        animate={{ opacity: 1, y: "15%", scale: 1 }}
        transition={{ duration: 0.85, ease: "easeOut" }}
        className="w-[50rem] flex justify-center"
      >
        <m.div
          initial={{ opacity: 0, rotate: -10, y: "10%" }} // Adjusted y property for additional downward movement
          animate={{ opacity: 1, rotate: 0, y: "0%" }}
          transition={{ duration: 0.75, ease: "easeOut", delay: 0.2 }}
          className="ml-10 z-10"
        >
          <Image
            src={"assets/landing/phone_pc.svg"}
            width={0}
            height={0}
            style={{ width: "100%" }}
            alt="phone and pc"
            className="z-10"
          />
        </m.div>
      </m.section>
    </div>
  );
};

export default LandingHome;
