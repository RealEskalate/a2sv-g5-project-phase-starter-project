import Card from "../../components/Accounts/account";
import React from "react";
import { YearlyInvest } from "../../components/Investment/YearlyInvest";
import { MonthlyRev } from "../../components/Investment/MonthlyRev";
import MyInvestment from "../../components/Investment/MyInvestment";
import TrendingStock from "../../components/Investment/TrendingStock";
const page = () => {
  const data = [
    ["01.", "Trivago", "$520", "+5%"],
    [" 02.", "Canon", "$480", " +10%"],
    ["03.", "Uber Food", " $350", "-3%"],
    ["04.", " Nokia", " $940", "+2%"],
    ["05.", "Tiktok", "$670", "-12%"],
  ];

  return (
    <>
      <div className="flex flex-col lg:flex-row gap-6">
        <Card
          title="Total Invested Amount"
          amount="$150,000"
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
          amount="+5.80%"
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
          <YearlyInvest />
        </div>

        <div className="w-full lg:w-[48%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Monthly Revenue
          </p>
          <MonthlyRev />
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
            icon="/assets/Google.svg"
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
          <TrendingStock data={data} />
        </div>
      </div>
    </>
  );
};

export default page;
