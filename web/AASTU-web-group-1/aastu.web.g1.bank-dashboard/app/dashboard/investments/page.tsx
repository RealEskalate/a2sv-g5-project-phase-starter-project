"use client";
import React,{useState, useEffect, use} from "react";
import Card from "./investment_components/Card";
import { investmentTypes } from "@/constants";
import { getInvestmentData } from "@/lib/investmentapies";
import Investment from "./investment_components/Investment";
import LineChartComp from "./investment_components/Chart1";
import Chart2 from "./investment_components/Chart2";
import Trending from "./investment_components/Trending";
import { useUser } from "@/contexts/UserContext";

const ShimmerEffect = ()=>{
  return(
    <>
      <div className="flex flex-col md:flex-row gap-3 md:gap-10 justify-center pt-4 w-[80%] my-10 mx-auto">
        {Array(3).fill("").map((_, index) => (<div key={index} className="w-full h-[80px] md:w-[30%] bg-gray-300 animate-pulse rounded-2xl"></div>))}
      </div>

      <div className="flex flex-col flex-wrap md:flex-row gap-10 my-4 justify-between w-[80%] mx-auto">
        <div className="w-[45%] mx-auto md:mx-0">
          <div className="h-6 w-3/4 bg-gray-300 animate-pulse rounded-md mb-3"></div>
          <div className="w-full h-64 bg-gray-300 animate-pulse rounded-md"></div>
        </div>
        <div className="w-[45%] mx-auto md:mx-0">
          <div className="h-6 w-3/4 bg-gray-300 animate-pulse rounded-md mb-3"></div>
          <div className="w-full h-64 bg-gray-300 animate-pulse rounded-md"></div>
        </div>
      </div>
    </>

  )
}

interface fetchDataType{
  "totalInvestment":number,
  "rateOfReturn":number,
  "yearlyTotalInvestment": [],
  "monthlyRevenue": []
}

const Investments = () => {
  const { isDarkMode } = useUser();
  const [data, setData] = useState<fetchDataType>();
  const [loading, setLoading] = useState(true);

  useEffect(()=>{
      const fetchData = async ()=>{
        try{
          const data = await getInvestmentData(5,5)
          if (data==null){
            throw new Error("Faild to fetch chart data!!!")
          }
          setData(data);
          setLoading(false);
        }catch(error){
          console.log("Error occured while fetching random invesment data : ",error);
      }
      }
      fetchData();
  },[])


  return (
    <div className={`flex flex-col justify-center ${isDarkMode ? "bg-gray-700" : "bg-[#F5F7FA]"} px-3`}>
      {loading ? (
         <ShimmerEffect/>
        ):(
          <>
          <div className="flex flex-col md:flex-wrap md:flex-row justify-evenly gap-3 md:gap-5 pt-4 w-[80%] mx-auto">
            {investmentTypes.map((item, index) => (
              ( data && <Card {...item} amount={data[item.id as keyof fetchDataType]} key={index} />)  
            ))}
          </div>
          <div className="flex flex-col md:flex-row md:flex-wrap gap-10 md:gap-5 my-4 w-[80%] justify-between mx-auto">
            <div className="w-full md:w-[35%] md:min-w-[600px] mx-auto md:mx-0">
              <h1 className={`my-3 font-[600] text-[16px] md:text-[22px] text-nowrap ${isDarkMode ? "text-white" : "text-[#333B69]"}`}>
                Yearly Total Investment
              </h1>
              <LineChartComp data = {data?.yearlyTotalInvestment}/>
            </div>
            <div className="w-full md:w-[35%] md:min-w-[600px] mx-auto md:mx-0">
              <h1 className={`my-3 font-[600] text-[16px] md:text-[22px] ${isDarkMode ? "text-white" : "text-[#333B69]"}`}>
                Monthly Revenue
              </h1>
              <Chart2 data ={data?.monthlyRevenue}/>
            </div>
          </div>
          </>
        )
      }
      <div className="flex flex-col md:flex-row gap-3 md:gap-10 my-2 w-[80%] justify-between mx-auto">
        <Investment />
        <Trending />
      </div>
    </div>
  );
};

export default Investments;
