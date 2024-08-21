"use client";
import React from "react";
import Image from "next/image";

const NotFound = () => {
  return (
    <section className="flex flex-col justify-center items-start gap-6 p-6">
      {/* Shimmer Effect for Card Placeholder */}
      <div className="flex flex-col gap-6">
        <div className="flex justify-between items-center mb-4">
          <div className="h-6 w-32 bg-gray-200 animate-pulse rounded"></div>
          <div className="h-6 w-24 bg-gray-200 animate-pulse rounded"></div>
        </div>
        <div className="flex gap-4 overflow-x-auto scrollbar-hide">
          <div className="w-64 h-32 bg-gray-200 animate-pulse rounded"></div>
          <div className="w-64 h-32 bg-gray-200 animate-pulse rounded"></div>
        </div>
      </div>

      {/* Shimmer Effect for Stats Placeholder */}
      <div className="flex flex-col gap-6">
        <div className="h-6 w-32 bg-gray-200 animate-pulse rounded mb-4"></div>
        <div className="flex flex-col p-4 bg-gray-200 animate-pulse rounded shadow-sm">
          <div className="h-4 w-48 bg-gray-300 mb-2 rounded"></div>
          <div className="h-4 w-32 bg-gray-300 mb-2 rounded"></div>
          <div className="h-4 w-20 bg-gray-300 rounded"></div>
        </div>
        <div className="h-40 w-full bg-gray-200 animate-pulse rounded"></div>
      </div>
    </section>
  );
};
// return (
//   <div className="w-full h-screen flex flex-col items-center justify-start pt-[20vh] gap-6">
//     <div>
//       <Image src="/assets/pageNotFound.svg" alt="" width={220} height={220} />
//     </div>
//     <h1 className="text-2xl text-gray-500">Page Not Found</h1>
//     {/* <div className="flex justify-center gap-8">
//       <div className="w-80 h-64 animate-pulse bg-gray-200 rounded-md"></div>
//       <div className="w-80 h-64 animate-pulse bg-gray-200 rounded-md"></div>
//     </div> */}
//   </div>
// );
// };

export default NotFound;
