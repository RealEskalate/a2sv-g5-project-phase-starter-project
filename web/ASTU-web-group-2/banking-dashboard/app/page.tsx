import React from "react";
import WeeklyActivityChart from "./components/charts/WeeklyActivityChart";
import Card from "./components/card/Card";
import CardForCreditCards from "./components/card/CardForCreditCards";

const page = () => {
  return (
    <div className="grid grid-cols-2">
      <CardForCreditCards title="Weekly Activity" button="See All" link="/report">
        <WeeklyActivityChart />
      </CardForCreditCards>
    </div>
  );
};

export default page;
