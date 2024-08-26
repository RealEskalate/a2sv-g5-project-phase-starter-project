"use client"
import React from 'react'
import { useUser } from '@/contexts/UserContext';

const TableShimmer = () => {
    const { isDarkMode } = useUser();
  return (
    <div className="flex flex-col justify-center rounded-2xl bg-transparent">
    {/* Placeholder for the table */}
    <div className={`w-[80%] md:w-full mx-auto rounded-2xl my-4 ${isDarkMode ? "bg-gray-950" : "bg-white"} animate-pulse`}>
      {/* Placeholder for table header */}
      <div className={`flex justify-around py-4 ${isDarkMode ? "border-gray-700" : ""}`}>
        <div className="hidden md:block h-6 w-12 bg-gray-300 rounded-md"></div>
        <div className="h-6 w-20 bg-gray-300 rounded-md"></div>
        <div className="h-6 w-24 bg-gray-300 rounded-md"></div>
        <div className="hidden md:block h-6 w-16 bg-gray-300 rounded-md"></div>
        <div className="hidden md:block h-6 w-20 bg-gray-300 rounded-md"></div>
        <div className="hidden md:block h-6 w-20 bg-gray-300 rounded-md"></div>
        <div className="h-6 w-16 bg-gray-300 rounded-md"></div>
      </div>
  
      {/* Placeholder for table rows */}
      {Array(5).fill("").map((_, idx) => (
        <div key={idx} className={`flex justify-around py-4 ${isDarkMode ? idx % 2 === 0 ? "bg-gray-900" : "bg-gray-800" : idx % 2 === 0 ? "bg-[#dfdfdf59]" : "bg-white"}`}>
          <div className="hidden md:block h-6 w-12 bg-gray-300 rounded-md"></div>
          <div className="h-6 w-20 bg-gray-300 rounded-md"></div>
          <div className="h-6 w-24 bg-gray-300 rounded-md"></div>
          <div className="hidden md:block h-6 w-16 bg-gray-300 rounded-md"></div>
          <div className="hidden md:block h-6 w-20 bg-gray-300 rounded-md"></div>
          <div className="hidden md:block h-6 w-20 bg-gray-300 rounded-md"></div>
          <div className="h-6 w-16 bg-gray-300 rounded-md"></div>
        </div>
      ))}
  
      {/* Placeholder for table footer */}
      <div className={`flex justify-around py-4 ${isDarkMode ? "bg-gray-950 border-gray-700 text-red-300" : "bg-white text-red-700"} rounded-b-2xl`}>
        <div className="hidden md:block h-6 w-12 bg-gray-300 rounded-md"></div>
        <div className="h-6 w-20 bg-gray-300 rounded-md"></div>
        <div className="h-6 w-24 bg-gray-300 rounded-md"></div>
        <div className="hidden md:block h-6 w-16 bg-gray-300 rounded-md"></div>
        <div className="hidden md:block h-6 w-20 bg-gray-300 rounded-md"></div>
        <div className="h-6 w-16 bg-gray-300 rounded-md"></div>
      </div>
    </div>
  
    {/* Placeholder for pagination */}
    <div className="flex justify-center mt-4">
      <div className="flex items-center gap-2">
        <div className={`${isDarkMode ? "bg-gray-950 text-gray-50" : "bg-gray-200"} h-8 w-8 rounded-full animate-pulse`}></div>
        <div className="flex gap-2">
          {Array(5).fill("").map((_, idx) => (
            <div key={idx} className={`${isDarkMode ? "bg-gray-950 text-gray-50" : "bg-gray-200"} h-8 w-8 rounded-full animate-pulse`}></div>
          ))}
        </div>
        <div className={`${isDarkMode ? "bg-gray-950 text-gray-50" : "bg-gray-200"} h-8 w-8 rounded-full animate-pulse`}></div>
      </div>
    </div>
  </div>
  
  )
}

export default TableShimmer
