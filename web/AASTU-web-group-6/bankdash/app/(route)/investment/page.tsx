"use client"
import Card from "../../components/Accounts/account";
import React from "react";
import { YearlyInvest } from "../../components/Investment/YearlyInvest";
import { MonthlyRev } from "../../components/Investment/MonthlyRev";
import MyInvestment from "../../components/Investment/MyInvestment";
import TrendingStock from "../../components/Investment/TrendingStock";
import { useState , useEffect } from "react";
import InvestmentService from "@/app/Services/api/investmentApi";
interface YearlyInvestment {
  time: string; 
  value: number;
}

interface MonthlyRevenue {
  time: string; 
  value: number;
}

interface InvestmentData {
  totalInvestment: number;
  rateOfReturn: number;
  yearlyTotalInvestment: YearlyInvestment[];
  monthlyRevenue: MonthlyRevenue[];
}
const page = () => {
  const data1 = [
    ["01.", "Trivago", "$520", "+5%"],
    [" 02.", "Canon", "$480", " +10%"],
    ["03.", "Uber Food", " $350", "-3%"],
    ["04.", " Nokia", " $940", "+2%"],
    ["05.", "Tiktok", "$670", "-12%"],
  ];
  const [data, setData] = useState<InvestmentData>();
  const [error, setError] = useState<string | null>(null);
  const accessToken = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJtaWhyZXQiLCJpYXQiOjE3MjQwNTY5MDEsImV4cCI6MTcyNDE0MzMwMX0.06ogiDUHZipaLn7gIoJDxGz4Bw_zFmsA72Zp99eKGkKVjOFRXy3MUvh55dspPaib";


  useEffect(() => {
    const getData = async () => {
      try {
        const result = await InvestmentService.getInvestmentData(accessToken);
        setData(result);
        console.log("data" , data);
      } catch (error) {
        setError("error");
        alert("error");
      }
    };

    getData();
  }, []);
  const formattedAmount = `$${data?.totalInvestment.toLocaleString()}`;
  const formattedReturn = `${data?.rateOfReturn.toFixed(2).toLocaleString()}%`
  return (
    <div className="w-[96%] flex flex-col grow gap-6 p-8 pt-6">
      <div className="flex flex-col lg:flex-row gap-6">
        <Card
          title="Total Invested Amount"
          amount={formattedAmount}
          color="#DCFAF8"
          icon="/assets/money-bag-of-dollars.svg"
          width="w-[32%]"
        ></Card>
        <Card
          title="Number of Investments"
          amount="1,250"
          color="#FFE0EB"
          icon="/assets/pie-chart.svg"
          width="w-[32%]"
        ></Card>
        <Card
          title="Rate of Return"
          amount={formattedReturn}
          color="#E7EDFF"
          icon="/assets/repeat.svg"
          width="w-[32%]"
        ></Card>
      </div>
      <div className=" flex flex-col lg:flex-row justify-between my-5">
        <div className=" w-full lg:w-[48%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Yearly Total Investment
          </p>
          <YearlyInvest data = {data}/>
        </div>

        <div className="w-full lg:w-[48%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Monthly Revenue
          </p>
          <MonthlyRev data = {data}/>
        </div>
      </div>
      <div className="flex flex-col lg:flex-row justify-between my-5">
        <div className=" w-full lg:w-[60%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            My Investment
          </p>
          <MyInvestment
            title="Apple Store"
            color="#FFE0EB"
            icon="/assets/applee.svg"
            amount="$54,000"
            Envestment="Envestment Value"
            returnValue="+16%"
            returnRe="Return Value"
            titleRe="E-commerce, Marketplace"
          />
          <MyInvestment
            title="Samsung Mobile"
            color="#E7EDFF"
            icon="/assets/google.svg"
            amount="$25,300"
            Envestment="Envestment Value"
            returnValue="-4%"
            returnRe="Return Value"
            titleRe="E-commerce, Marketplace"
          />
          <MyInvestment
            title="Tesla Motors"
            color="#FFF5D9"
            icon="/assets/Tesla.svg"
            amount="$8,200"
            Envestment="Envestment Value"
            returnValue="+25%"
            returnRe="Return Value"
            titleRe="Electric Vehicles"
          />
        </div>
        <div className=" lg:w-[35%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Trending Stock
          </p>
          <TrendingStock data={data1} />
        </div>
      </div>
    </div>
  );
};

export default page;
