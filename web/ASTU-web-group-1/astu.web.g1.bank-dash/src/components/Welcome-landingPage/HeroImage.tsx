import React from "react";
import Image from "next/image";

const HeroImage = () => {
  return (
    <div className="p-4 lg:mx-32 bg-gray-200 border-2 rounded-3xl mb-32">
      <div className="relative h-[100vh] rounded-xl">
        <Image
          src="/dashboard.png"
          alt="Bank Logo"
          layout="fill"
          objectFit="cover"
          objectPosition="left top"
          sizes="100vw, 100vh"
          className="rounded-3xl"
        />
      </div>
    </div>
  );
};

export default HeroImage;
