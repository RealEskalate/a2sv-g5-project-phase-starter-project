import React from "react";
import Image from "next/image";

const page = () => {
  return (
    <div className="w-full h-screen flex flex-col items-center justify-start pt-[20vh] gap-6">
      <div>
        <Image src="/assets/pageNotFound.svg" alt="" width={220} height={220} />
      </div>
      <h1 className="text-2xl text-gray-500">Page Not Found</h1>
    </div>
  );
};

export default page;
