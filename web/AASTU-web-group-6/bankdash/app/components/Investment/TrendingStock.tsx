import React from "react";
interface props {
  data: string[][];
}

const TrendingStock = ({ data }: props) => {
  return (
    <div className="border border-solid rounded-3xl overflow-hidden w-[100%]  bg-white dark:bg-[#232328] p-4">
      <table className="w-full">
        <thead>
          <tr className="text-left ">
            <th className="lg:px-3 xl:px-6 px:1 sm:px-3 py-6 font-medium lg:text-[12px] xl:text-base text-[12px] text-[#718EBF] dark:text-gray-300">
              SL No
            </th>
            <th className="lg:px-3 xl:px-6 px:1 sm:px-3 py-6 font-medium lg:text-[12px] xl:text-base text-[12px] text-[#718EBF] dark:text-gray-300">
              Name
            </th>
            <th className="lg:px-3 xl:px-6 px:1 sm:px-3 py-6 font-medium lg:text-[12px] xl:text-base text-[12px] text-[#718EBF] dark:text-gray-300">
              Price
            </th>
            <th className="lg:px-3 xl:px-6 px:1 sm:px-3 py-6 font-medium lg:text-[12px] xl:text-base text-[12px] text-[#718EBF] dark:text-gray-300">
              Return
            </th>
          </tr>
        </thead>
        <tbody>
          {data.map((row, rowIndex) => (
            <tr key={rowIndex}>
              {row.map((cell, cellIndex) => {
                const isLastCell = cellIndex === row.length - 1;
                const cellColor =
                  isLastCell && typeof cell === "string" && cell.startsWith("-")
                    ? "text-[#FE5C73]"
                    : "text-[#16DBAA]";
                const cellClasses = isLastCell
                  ? `lg:px-3 xl:px-5 px:1 sm:px-3 py-4 whitespace-nowrap font-inter font-normal lg:text-[12px] xl:text-base text-[12px] ${cellColor}`
                  : "lg:px-3 xl:px-5 px:1 sm:px-3 py-4 whitespace-nowrap font-inter font-normal lg:text-[12px] xl:text-base text-[12px] text-[#232323] dark:text-gray-300";

                return (
                  <td key={cellIndex} className={cellClasses}>
                    {cell}
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TrendingStock;
