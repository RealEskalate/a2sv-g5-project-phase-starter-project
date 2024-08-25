'use client'
import React,{useState, useEffect} from "react";
import Card from "./investment_components/Card";
import { investmentTypes } from "@/constants";
import Investment from "./investment_components/Investment";
import LineChartComp from "./investment_components/Chart1";
import Chart2 from "./investment_components/Chart2";
import Trending from "./investment_components/Trending";
import { useUser } from "@/contexts/UserContext";

const Investments = () => {
  const { isDarkMode } = useUser();
  const [cardData, setCardData] = useState();
  return (
    <div className={`flex flex-col justify-center ${ isDarkMode ? "bg-gray-700" : "bg-[#F5F7FA]"} px-3`}>
      <div className="flex flex-col md:flex-row gap-3 md:gap-10 justify-center pt-4 w-full">
        {investmentTypes.map((item) => (
          <Card {...item} key={item.name} />
        ))}
      </div>
      <div className="flex flex-col md:flex-row flex-wrap md:gap-30 gap-10 my-4 w-full justify-center">
        <div className="w-[90%] md:w-[650px] mx-10 md:mx-0">
          <h1 className={`my-3 font-[600] text-[16px] md:text-[22px] text-nowrap ${isDarkMode ? "text-gray-200" : "text-[#333B69]"}`}>
            Yearly Total Investment
          </h1>
          <LineChartComp />
        </div>
        <div className="w-[90%] md:w-[650px] mx-auto md:mx-0">
          <h1 className={`my-3 font-[600] text-[16px] md:text-[22px] ${isDarkMode ? "text-gray-200" : "text-[#333B69]"}`}>
            Monthly Revenue
          </h1>
          <Chart2 />
        </div>
      </div>
      <div className="flex flex-col md:flex-row gap-3 md:gap-8 my-2 w-full justify-center">
        <Investment />
        <Trending />
      </div>
    </div>
  );
};

export default Investments;
