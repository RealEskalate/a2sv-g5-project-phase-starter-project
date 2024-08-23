import { Skeleton } from "@/components/ui/skeleton";
import React from "react";

const ActiveLoanSkeleton = () => {
  const TableData = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  return (
    <div className="flex flex-col gap-4">
      <Skeleton className="h-5 w-52 bg-slate-200" />
      <div className="relative overflow-x-auto bg-white px-4 md:px-6 pt-5 md:pt-6 rounded-2xl md:rounded-2xl">
        <table className="bg-white px-5 lg:px-11 w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
          <thead className=" text-12px md:text-15px font-Lato font-normal text-blue-steel bg-white border-b">
            <tr className="">
              <th scope="col" className="hidden md:table-cell pb-2">
                <Skeleton className="h-5 w-40" />
              </th>
              <th scope="col" className=" pb-2">
                <Skeleton className="h-5 w-40" />
              </th>
              <th scope="col" className=" pb-2">
                <Skeleton className="h-5 w-40" />
              </th>
              <th scope="col" className="hidden lg:table-cell pb-2">
                <Skeleton className="h-5 w-40" />
              </th>
              <th scope="col" className="hidden min-[900px]:table-cell pb-2">
                <Skeleton className="h-5 w-40" />
              </th>
              <th scope="col" className="hidden min-[900px]:table-cell pb-2">
                <Skeleton className="h-5 w-40" />
              </th>
              <th scope="col" className=" pb-2 w-fit">
                <Skeleton className="h-5 w-40" />
              </th>
            </tr>
          </thead>
          <tbody className="text-12px xl:text-15px text-gray-dark cursor-pointer  hover:bg-gray-100 dark:hover:bg-gray-700">
            {TableData.map((data, index) => (
              <tr
                key={index}
                className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-700"
              >
                <td className="hidden md:table-cell py-3">
                  <Skeleton className="h-5 w-28" />
                </td>
                <td className="py-3">
                  <Skeleton className="h-5 w-28" />
                </td>
                <td className="py-3">
                  <Skeleton className="h-5 w-28" />
                </td>
                <td className="hidden lg:table-cell py-3">
                  <Skeleton className="h-5 w-28" />
                </td>
                <td className="hidden min-[900px]:table-cell py-3">
                  <Skeleton className="h-5 w-28" />
                </td>
                <td className="hidden min-[900px]:table-cell py-3">
                  <Skeleton className="h-5 w-28" />
                </td>
                <td className="py-3 w-24 md:w-32 ">
                  <Skeleton className="h-6 w-16 px-6" />
                </td>
              </tr>
            ))}
            <tr className="bg-white align-bottom text-candyPink font-medium dark:bg-gray-800 dark:border-gray-700">
              <td className="hidden md:table-cell py-3 md:py-6">
                <Skeleton className="h-5 w-28" />
              </td>
              <td className="py-3 md:py-6 flex flex-col">
                <Skeleton className="h-5 w-28" />
              </td>
              <td className="py-3 md:py-6">
                <Skeleton className="h-5 w-28" />
              </td>
              <td className="hidden md:table-cell py-3 md:py-6"></td>
              <td className="hidden md:table-cell py-3 md:py-6"></td>
              <td className="hidden min-[900px]:table-cell py-3 md:py-6">
                <Skeleton className="h-5 w-28" />
              </td>
              <td className="py-3 md:py-6 whitespace-nowrap"></td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default ActiveLoanSkeleton;
