"use client";
import React, { useEffect, useState } from "react";
import Image from "next/image";
import axios from "axios";
import TotalAmmount_img from "@/public/assests/icon/Investments/Group303.png";
import Number_img from "@/public/assests/icon/Investments/Group305.png";
import Rate_img from "@/public/assests/icon/Investments/Group307.png";
import ChartCard_Invest from "./ChartCard_Invest";
import MonthlyRevenueChart from './MonthlyRevenueChart'
import { tradingStockData, investmentsData } from "./mockData";
import { useSession } from "next-auth/react";

const Investments = () => {
  const [investmentOverview, setInvestmentOverview] = useState({
    totalAmount: 0,
    numberOfInvestments: 0,
    rateOfReturn: 0,
  });

  const [yearlyTotalInvestment, setYearlyTotalInvestment] = useState([]);
  const [monthlyRevenue, setMonthlyRevenue] = useState([]);
  const { data: session } = useSession();

  const token: string =  ` Bearer ${session?.user?.accessToken} `;
  useEffect(() => {
    
    const fetchInvestmentData = async () => {
      try {
        const response = await axios.get(
          "https://bank-dashboard-1tst.onrender.com/user/random-investment-data?years=3&months=5",
          {
            headers: {
              Authorization: token, // Make sure to add your token here
            },
          }
        );

        const { totalInvestment, rateOfReturn, yearlyTotalInvestment, monthlyRevenue } =
          response.data.data;
          console.log(response.data.data, "responce.data.data");

        setInvestmentOverview({
          totalAmount: totalInvestment,
          numberOfInvestments: yearlyTotalInvestment.length,
          rateOfReturn: rateOfReturn,
        });

        setYearlyTotalInvestment(yearlyTotalInvestment);
        setMonthlyRevenue(monthlyRevenue);
      } catch (error) {
        console.error("Error fetching investment data:", error);
      }
    };

    fetchInvestmentData();
  }, []);

  return (
    <div className="bg-[#F5F7FA] space-y-8 ml-[30px] pt-3    w-full overflow-hidden">
      {/* Row 1: Investment Overview */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={TotalAmmount_img} alt="balance" />
          <div>
            <p>Total Invested Amount</p>
            <p className="text-xl font-semibold">${investmentOverview.totalAmount.toFixed(2)}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Number_img} alt="balance" />
          <div>
            <p>Number of Investments</p>
            <p className="text-xl font-semibold">{investmentOverview.numberOfInvestments.toFixed(2)}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Rate_img} alt="balance" />
          <div>
            <p>Rate of Return</p>
            <p className="text-xl font-semibold">{investmentOverview.rateOfReturn.toFixed(2)}%</p>
          </div>
        </div>
      </div>

      {/* Row 2: Yearly Total Investment and Monthly Revenue */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div className="p-4 bg-gray-100 rounded-lg">
          <p>Yearly Total Investment</p>
          <div className="h-36 bg-white rounded mt-4" style={{ width: "100%", height: 329 }}>
            <ChartCard_Invest data={yearlyTotalInvestment} />
          </div>
        </div>
        <div className="p-4 bg-gray-100 rounded-lg">
          <p>Monthly Revenue</p>
          <div className="h-36 bg-white rounded mt-4" style={{ width: "100%", height: 329 }}>
            <MonthlyRevenueChart data={monthlyRevenue} />
          </div>
        </div>
      </div>

      {/* Row 3: Investments and Trading Stock */}
      <div className="flex flex-col md:flex-row gap-4">
        {/* Investments Section */}
        <div className="md:w-[58%] p-4 bg-gray-100 rounded-lg min-h-[345px]">
          <p className="text-lg font-semibold">My Investments</p>
          <div className="space-y-4 mt-4">
            {investmentsData.slice(0, 3).map((investment) => (
              <div
                key={investment.id}
                className="flex items-center space-x-4 p-2 bg-white rounded-lg shadow"
              >
                <Image
                  src={investment.image}
                  alt={investment.name}
                  width={44}
                  height={44}
                  className="rounded-full object-cover"
                />
                <div className="flex-1">
                  <p className="font-semibold">{investment.name}</p>
                  <p className="text-gray-500">{investment.service}</p>
                </div>
                <div>
                  <p className="text-sm font-semibold">{investment.value}</p>
                  <p className="text-xs text-gray-500">Investment value</p>
                </div>
                <div>
                  <div>
                    {investment.return < 0 ? (
                      <p className="text-red-500">{investment.return}%</p>
                    ) : (
                      <p className="text-green-500">{investment.return}%</p>
                    )}
                  </div>
                  <p className="text-xs text-gray-500">Return</p>
                </div>
              </div>
            ))}
          </div>
        </div>

        {/* Trading Stock Section */}
        <div className="md:w-[42%] p-4 bg-gray-100 rounded-lg min-h-[345px]">
          <p className="text-lg font-semibold">Trading Stock</p>
          <div className="mt-4">
            <table className="w-full bg-white rounded-lg shadow">
              <thead>
                <tr className="bg-gray-200">
                  <th className="p-2">Sl.No</th>
                  <th className="p-2">Name</th>
                  <th className="p-2">Price</th>
                  <th className="p-2">Return</th>
                </tr>
              </thead>
              <tbody>
                {tradingStockData.map((stock, index) => (
                  <tr key={stock.id}>
                    <td className="p-2">{index + 1}</td>
                    <td className="p-2">{stock.name}</td>
                    <td className="p-2">{stock.price}</td>
                    <td className="p-2">
                      {stock.return < 0 ? (
                        <p className="text-red-500">{stock.return}%</p>
                      ) : (
                        <p className="text-green-500">{stock.return}%</p>
                      )}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Investments;