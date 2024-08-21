import Image from "next/image";
import React from "react";

const Description = () => {
  return (
    <>
      <div
        className="flex flex-col md:flex-row justify-between p-5 h-auto md:h-2/4 space-y-5 md:space-y-0 md:space-x-10 w-full relative"
        id="why"
      >
        <Image
          src="/assets/images/online.jpg"
          alt="Bank Logo"
          className="object-cover opacity-20 rounded-lg"
          fill
        />
        <div className="flex w-full justify-center items-center p-4 md:w-2/5 h-64 md:h-auto">
          <p className="z-[5] inset-0 text-navy font-Inter text-5xl lg:text-5xl font-extrabold text-center">
            Why Choose Our BANK?
          </p>
        </div>
        <ul className="px-10 md:w-3/5 text-gray-700 md:pt-2 list-disc space-y-2 md:space-y-4">
          <li>
            <strong>Secure & Reliable:</strong> Your security is our priority.
            We use cutting-edge technology to safeguard your assets and personal
            information.
          </li>
          <li>
            <strong>Customer-Centric Services:</strong> Our services are
            designed with you in mind. Whether you’re saving for the future,
            managing daily expenses, or seeking financial growth, we’ve got you
            covered.
          </li>
          <li>
            <strong>Expert Guidance:</strong> Our team of financial experts is
            here to provide you with personalized advice, helping you make
            informed decisions that align with your financial goals.
          </li>
          <li>
            <strong>Convenient & Accessible:</strong> Bank from anywhere,
            anytime with our user-friendly mobile app and online banking
            platform.
          </li>
        </ul>
      </div>
    </>
  );
};

export default Description;
