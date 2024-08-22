import Image from "next/image";
import React from "react";

const LandingHome = () => {
  return (
    <section
      id="home"
      className="bg-[#083E9E] w-full h-[200px] md:h-[400px] lg:h-[600px] relative text-white flex"
    >
      <div className="pt-12 sm:pt-20 pl-12 md:pl-20 lg:pt-32 lg:pl-32 flex flex-col  md:w-[40%] z-50">
        <span className="font-normal md:font-bold lg:font-extrabold text-[12px] md:text-[25px] lg:text-[35px]">
          Collect Thrift Savings for people around you. & Earn Returns
        </span>
        <span className="font-thin md:font:normal text-[5px] md:text-[15px]">
          Join Millions of agents saving funds for people who dont have access
          to their bank details and earn in returns.
        </span>
      </div>

      <Image
        src={"assets/landing/net.svg"}
        width={0}
        height={0}
        alt="net"
        style={{ width: "100%", height: "100%" }}
        className="absolute bottom-0 right-0 z-0"
      />

      <Image
        src={"assets/landing/phone_pc.svg"}
        width={0}
        height={0}
        style={{ width: "50%" }}
        alt="phone and pc"
        className="ml-10 z-10 mb-[-70px] sm:mb-[-130px]"
      />
    </section>
  );
};

export default LandingHome;
