import React from "react";
import { ListCardLoading } from "./ListCard";
const Loading = () => {
  return (
    <>
      <div className="flex flex-col h-full bg-[#F5F7FA] px-3 py-3 gap-5">
        <div className="flex flex-wrap gap-2">
          <ListCardLoading></ListCardLoading>
          <ListCardLoading></ListCardLoading>
          <ListCardLoading></ListCardLoading>
          <ListCardLoading></ListCardLoading>
        </div>

        <div className="flex flex-col md:flex-row gap-5">
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl bg-gray-300 font-semibold w-24 h-5 rounded-lg"></span>
            <div className="bg-white flex flex-col justify-between rounded-2xl animate-pulse p-4">
              <div className="w-full h-10 bg-gray-200 rounded mb-4"></div>
              <div className="w-full h-10 bg-gray-200 rounded mb-4"></div>
              <div className="w-full h-10 bg-gray-200 rounded mb-4"></div>
              <div className="w-full h-10 bg-gray-200 rounded"></div>
            </div>
          </div>

          <div className="md:w-1/2 gap-1 flex flex-col">
            <div className="flex justify-between mr-2">
              <span className="text-xl bg-gray-300 font-semibold w-24 h-5 rounded-lg"></span>
              <span className="text-xl bg-gray-300 font-semibold w-10 h-5 rounded-lg"></span>
            </div>
            <div className="border rounded-3xl my-4 mx-2 animate-pulse">
              <div className="relative w-full bg-gradient-to-b from-gray-200 to-gray-300 text-transparent rounded-3xl shadow-md h-[230px] min-w-[350px]">
                <div className="flex justify-between items-start px-6 pt-6">
                  <div>
                    <p className="text-xs font-semibold bg-gray-300 rounded w-16 h-4 mb-2"></p>
                    <p className="text-xl font-medium bg-gray-300 rounded w-24 h-6"></p>
                  </div>
                  <div className="w-8 h-8 bg-gray-300 rounded-full"></div>
                </div>

                <div className="flex justify-between gap-12 mt-4 px-6">
                  <div>
                    <p className="text-xs font-medium bg-gray-300 rounded w-16 h-4 mb-2"></p>
                    <p className="font-medium text-base bg-gray-300 rounded w-24 h-6"></p>
                  </div>
                  <div className="pr-8">
                    <p className="text-xs font-medium bg-gray-300 rounded w-16 h-4 mb-2"></p>
                    <p className="font-medium text-base md:text-lg bg-gray-300 rounded w-24 h-6"></p>
                  </div>
                </div>

                <div className="relative mt-8 flex justify-between py-4 items-center">
                  <div className="absolute inset-0 w-full h-full bg-gradient-to-b from-white/20 to-transparent z-0"></div>
                  <div className="ml-4 relative z-10 text-base font-medium px-6 bg-gray-300 rounded w-40 h-6"></div>
                  <div className="flex justify-end relative z-10 px-6">
                    <div className="w-10 h-10 bg-gray-300 rounded-full "></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div className="flex flex-col md:flex-row gap-5">
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl bg-gray-300 font-semibold w-24 h-5 rounded-lg"></span>
            <div className="w-full h-64 bg-gray-200 rounded-2xl animate-pulse"></div>
          </div>
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl bg-gray-300 font-semibold w-24 h-5 rounded-lg"></span>
            <div className="bg-white flex flex-col justify-between rounded-2xl animate-pulse p-4">
              <div className="w-full h-10 bg-gray-200 rounded mb-4"></div>
              <div className="w-full h-10 bg-gray-200 rounded mb-4"></div>
              <div className="w-full h-10 bg-gray-200 rounded mb-4"></div>
              <div className="w-full h-10 bg-gray-200 rounded"></div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Loading;
