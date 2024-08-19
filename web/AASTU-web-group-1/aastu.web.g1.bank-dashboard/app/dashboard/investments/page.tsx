import React from "react";
import Card from "./investment_components/Card";
import { investmentTypes } from "@/constants";
import Investment from "./investment_components/Investment";
import  LineChartComp  from "./investment_components/Chart1";
import Chart2 from "./investment_components/Chart2";
import Trending from "./investment_components/Trending";
const Investments = () => {

  return(
    <div className="flex flex-col justify-center">
      <div className=" flex flex-col md:flex-row gap-3 md:gap-10 justify-center pt-4 w-full">
        {investmentTypes.map((item) => (
          <Card {...item} key="item.name"/>
        ))}
      </div>
      <div className="flex flex-col md:flex-row md:gap-40 md:ml-20 md:justify-start gap-10 my-4 w-full justify-center " >
        <div className="w-[90%] md:w-[35%] mx-auto md:mx-0">
          <h1 className="my-3 font-[600] text-[16px] text-[#333B69] md:text-[22px]" >Yearly Total Investment</h1>
          <LineChartComp/>
        </div>
        <div className="w-[90%] md:w-[35%] mx-auto md:mx-0">
          <h1 className="my-3 font-[600] text-[16px] text-[#333B69] md:text-[22px]" >Monthly Revenue</h1>
          <Chart2/>
        </div>
      </div>
      <div className=" flex flex-col md:flex-row gap-3 md:gap-10 my-2 w-full justify-center md:justify-start md:ml-20">
        <Investment/>
        <Trending/>
      </div>
    </div>
  )
};

export default Investments;
