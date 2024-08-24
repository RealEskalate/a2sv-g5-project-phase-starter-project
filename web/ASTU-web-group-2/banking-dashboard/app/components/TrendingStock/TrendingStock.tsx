"use client"
import React from "react";
import TableSkeleton from "./TableSkeleton";

const TrendingStock = () => {
  // if (true){
  //   return (
  //     <TableSkeleton />
  //   )
  // }
  return (
    <div className="bg-white rounded-3xl">
      <table className="w-full border-collapse">
        <thead>
          <tr>
            <th className="font-medium text-lg text-gray-500 p-4">SL No</th>
            <th className="font-medium text-lg text-gray-500 p-4">Name</th>
            <th className="font-medium text-lg text-gray-500 p-4">Price</th>
            <th className="font-medium text-lg text-gray-500 p-4">Return</th>
          </tr>
          <tr className="bg-gray-100 h-px">
            <td colSpan={4}></td>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td className="p-4 text-gray-700">01.</td>
            <td className="p-4 text-gray-700">Trivago</td>
            <td className="p-4 text-gray-700">$520</td>
            <td className="p-4 text-red-600">-5%</td>
          </tr>
          <tr>
            <td className="p-4 text-gray-700">02.</td>
            <td className="p-4 text-gray-700">Expedia</td>
            <td className="p-4 text-gray-700">$450</td>
            <td className="p-4 text-[#16DBAA]">+3%</td>
          </tr>
          <tr>
            <td className="p-4 text-gray-700">03.</td>
            <td className="p-4 text-gray-700">Airbnb</td>
            <td className="p-4 text-gray-700">$670</td>
            <td className="p-4 text-[#16DBAA]">+7%</td>
          </tr>
          <tr>
            <td className="p-4 text-gray-700">04.</td>
            <td className="p-4 text-gray-700">Booking</td>
            <td className="p-4 text-gray-700">$490</td>
            <td className="p-4 text-[#16DBAA]">+4%</td>
          </tr>
          <tr>
            <td className="p-4 text-gray-700">05.</td>
            <td className="p-4 text-gray-700">Kayak</td>
            <td className="p-4 text-gray-700">$380</td>
            <td className="p-4 text-red-700">-2%</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default TrendingStock;
