"use client";
import React, { useState, useEffect } from "react";
import LineChartNoBg from "@/components/LineChartNoBg";
import LineChartStright from "@/components/LineChartStraight";
import { colors } from "@/constants/index";
import { FaSackDollar } from "react-icons/fa6";
import { AiOutlineRetweet } from "react-icons/ai";
import { GiTakeMyMoney } from "react-icons/gi";
import MyInvestment from "@/components/MyInvestment";
import TrendingStock from "@/components/TrendingStock";
import { getTrendingCompanies } from "@/services/companygetch";
import { randomInvestmentData } from "@/services/userupdate";
const data = [
  {
    icon: "/icons/apple_store.png",
    color: "bg-red-100 ",
    colortext: colors.textblack,
    category: "E-commerce, marketplace",
    categorycolor: colors.textgray,
    name: "Apple Store",
    amount: "54000",
    percentage: "1.6%",
  },
  {
    icon: "/icons/Google_store.png",
    color: "bg-blue-100",
    colortext: colors.textblack,
    category: "E-commerce, marketplace",
    categorycolor: colors.textgray,
    name: "Google Store",
    amount: "25000",
    percentage: "2.23%",
  },
  {
    icon: "/icons/tesla.png",
    color: "bg-yellow-100",
    colortext: colors.textblack,
    category: "E-commerce, marketplace",
    categorycolor: colors.textgray,
    name: "Tesla Store",
    amount: "95000",
    percentage: "2.23%",
  },
];
const trendingdata = [
  { slNo: "01.", name: "Nokia", price: "$940", return: "+2%" },
  { slNo: "02.", name: "Apple", price: "$1500", return: "+5%" },
  { slNo: "03.", name: "Google", price: "$2500", return: "-3%" },
  { slNo: "04.", name: "Amazon", price: "$3000", return: "+4%" },
  { slNo: "05.", name: "Microsoft", price: "$2000", return: "-6%" },
];

interface chartData {
  time: string;
  value: number;
}
interface InvestmentData {
  totalInvestment: string;
  rateOfReturn: string;
  yearlyTotalInvestment: chartData[];
  monthlyRevenue: chartData[];
  // Add other properties as needed
}

const Investments = () => {
  // const fetch = async () => {
  //   try {
  //     const trendingcomp = await getTrendingCompanies();
  //     return trendingcomp;
  //   } catch (error) {
  //     console.error("Login Error:", error);
  //   }
  // };

  // const Trendingcomp = fetch();

  const [investment, setInvestment] = useState<InvestmentData>();
  const [status, setStatus] = useState<"loading" | "error" | "success">(
    "loading"
  );

  useEffect(() => {
    const fetchInvestmentData = async () => {
      setStatus("loading");
      try {
        const data = await randomInvestmentData();
        console.log(data, "sanhiubk");
        if (data.success) {
          setInvestment(data.data);
          setStatus("success");
        }
      } catch (error) {
        console.error("Error fetching investment data:", error);
        setStatus("error");
      }
    };
    fetchInvestmentData();
  }, []);

  if (status === "loading") {
    return (
      <div
        className={` ${colors.graybg} flex flex-col lg:gap-5 lg:ml-64 lg:pr-6 xl:pr-10 dark:bg-dark text-gray-900 dark:text-white`}
      >
        <div className="flex flex-col items-center px-6 pt-10 gap-4 lg:flex-row dark:bg-dark text-gray-900 dark:text-white">
          <div className="flex gap-3 w-[80%] bg-gray-200 justify-center items-center py-3 rounded-xl animate-pulse">
            <div className="bg-cyan-100 w-[50px] h-[50px] flex items-center justify-center rounded-full animate-pulse"></div>
            <div>
              <div className="bg-gray-300 h-[12px] w-[150px] rounded mb-2 animate-pulse"></div>
              <div className="bg-gray-300 h-[16px] w-[100px] rounded animate-pulse"></div>
            </div>
          </div>

          <div className="flex gap-3 w-[80%] bg-gray-200 justify-center items-center py-3 rounded-xl animate-pulse">
            <div className="bg-pink-100 w-[50px] h-[50px] flex items-center justify-center rounded-full animate-pulse"></div>
            <div>
              <div className="bg-gray-300 h-[12px] w-[150px] rounded mb-2 animate-pulse"></div>
              <div className="bg-gray-300 h-[16px] w-[100px] rounded animate-pulse"></div>
            </div>
          </div>

          <div className="flex gap-3 w-[80%] bg-gray-200 justify-center items-center py-3 rounded-xl animate-pulse">
            <div className="bg-indigo-100 w-[50px] h-[50px] flex items-center justify-center rounded-full animate-pulse"></div>
            <div>
              <div className="bg-gray-300 h-[12px] w-[150px] rounded mb-2 animate-pulse"></div>
              <div className="bg-gray-300 h-[16px] w-[100px] rounded animate-pulse"></div>
            </div>
          </div>
        </div>

        <div className="flex flex-col py-5 px-6 gap-14 lg:grid lg:grid-cols-2 lg:gap-6">
          <div className="flex flex-col gap-3 lg:gap-4 xl:gap-5">
            <div className="bg-gray-300 h-[22px] w-[200px] rounded mb-4 animate-pulse"></div>
            <div className="bg-gray-300 h-[250px] rounded animate-pulse"></div>
          </div>
          <div className="flex flex-col gap-3 lg:gap-4 xl:gap-5">
            <div className="bg-gray-300 h-[22px] w-[200px] rounded mb-4 animate-pulse"></div>
            <div className="bg-gray-300 h-[250px] rounded animate-pulse"></div>
          </div>
        </div>

        <div className="flex flex-col lg:grid lg:grid-cols-5">
          <div className="px-6 lg:col-span-3 flex flex-col gap-5">
            <div className="bg-gray-300 h-[22px] w-[200px] rounded mb-4 animate-pulse"></div>
            <div className="bg-gray-300 h-[150px] rounded mb-4 animate-pulse"></div>
            <div className="bg-gray-300 h-[150px] rounded mb-4 animate-pulse"></div>
            <div className="bg-gray-300 h-[150px] rounded mb-4 animate-pulse"></div>
          </div>
          <div className="lg:col-span-2">
            <div className="p-6 lg:p-0">
              <div className="bg-gray-300 h-[22px] w-[200px] rounded mb-4 animate-pulse"></div>
              <div className="bg-gray-300 h-[250px] rounded animate-pulse"></div>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div
      className={` ${colors.graybg}   flex  flex-col  lg:gap-5 lg:ml-64 lg:pr-6 xl:pr-10  dark:bg-dark text-gray-900 dark:text-white`}
    >
      <div className="flex flex-col items-center px-6 pt-10 gap-4 lg:flex-row dark:bg-dark text-gray-900 dark:text-white">
        <div className="flex gap-3 w-[80%] bg-white justify-center items-center py-3  rounded-xl  dark:bg-dark text-gray-900 dark:text-white">
          <div className="bg-cyan-100 w-[50px] h-[50px] flex items-center justify-center rounded-full  ">
            <FaSackDollar className="text-cyan-500 h-[25px] w-[20px] " />
          </div>
          <div>
            <p
              className={`${colors.textgray} font-normal text-[12px] lg:text-[18px] dark:text-white`}
            >
              Total Invested Amount
            </p>
            <p
              className={`${colors.textblack} font-semibold text-[16px] dark:text-white`}
            >
              {investment?.totalInvestment ?? "no data to display"}
            </p>
          </div>
        </div>

        <div className="flex gap-3 w-[80%] bg-white justify-center items-center py-3 rounded-xl  dark:bg-dark text-gray-900 dark:text-white">
          <div className="bg-pink-100 w-[50px] h-[50px] flex items-center justify-center rounded-full  ">
            <GiTakeMyMoney className="text-pink-500 h-[25px] w-[20px] " />
          </div>
          <div>
            <p
              className={`${colors.textgray} font-normal text-[12px] lg:text-[18px] dark:text-white`}
            >
              Number of Investments
            </p>
            <p
              className={`${colors.textblack} font-semibold text-[16px] dark:text-white`}
            >
              1809
            </p>
          </div>
        </div>

        <div className="flex gap-3 w-[80%] bg-white justify-center items-center py-3 rounded-xl dark:bg-dark text-gray-900 dark:text-white">
          <div className="bg-indigo-100 w-[50px] h-[50px] flex items-center justify-center rounded-full  ">
            <AiOutlineRetweet className="text-indigo-500 h-[25px] w-[20px] " />
          </div>
          <div>
            <p
              className={`${colors.textgray} font-normal text-[12px] w-[132px] lg:text-[18px] dark:text-white`}
            >
              Rate of Return
            </p>
            <p
              className={`${colors.textblack} font-semibold text-[16px] dark:text-white`}
            >
              {investment?.rateOfReturn ?? "no data to display"}
            </p>
          </div>
        </div>
      </div>

      <div className="  flex flex-col py-5 px-6 gap-14 lg:grid lg:grid-cols-2 lg:gap-6  ">
        <div className="flex flex-col gap-3 lg:gap-4 xl:gap-5 ">
          <h2
            className={`font-semibold text-[22px] ${colors.navbartext} dark:text-blue-500`}
          >
            Yearly Total Investments
          </h2>
          <LineChartStright
            yearlyData={
              investment?.yearlyTotalInvestment ?? [{ time: "", value: 0 }]
            }
          />
        </div>
        <div className="flex  flex-col gap-3 lg:gap-4 xl:gap-5 ">
          <h2
            className={`font-semibold text-[22px] ${colors.navbartext} dark:text-blue-500`}
          >
            Monthly Revenue
          </h2>
          <LineChartNoBg
            monthlyData={investment?.monthlyRevenue ?? [{ time: "", value: 0 }]}
          />
        </div>
      </div>

      <div className="flex flex-col lg:grid lg:grid-cols-5 ">
        <div className="px-6  lg:col-span-3 flex flex-col gap-5">
          <h2
            className={`font-semibold text-[22px] ${colors.navbartext} dark:text-blue-500`}
          >
            My Investment
          </h2>
          {data.map((item, index) => (
            <MyInvestment
              key={index}
              icon={item.icon}
              color={item.color}
              colortext={item.colortext}
              category={item.category}
              categorycolor={item.categorycolor}
              name={item.name}
              amount={item.amount}
              percentage={item.percentage}
            />
          ))}
        </div>
        <div className="lg:col-span-2">
          <div className="p-6 lg:p-0">
            <h2
              className={`font-semibold text-[22px] ${colors.navbartext} dark:text-blue-500`}
            >
              Trending Stock
            </h2>

            <TrendingStock items={trendingdata} />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Investments;
