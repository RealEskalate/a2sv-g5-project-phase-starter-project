import React from "react";
import InfoboxForInvestementPage from "../components/infobox/InfoboxForInvestementPage";
import MyExpenseChart from "../components/charts/MyExpenseChart";
import Card from "../components/card/Card";

const InvestmentsPage = () => {
  return (
    <div className="flex flex-col gap-2">
      <InfoboxForInvestementPage />
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="My Expense"
          className="flex flex-col lg:w-[730px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254]"
        >
          <MyExpenseChart />
        </Card>
      </div>
    </div>
  );
};

export default InvestmentsPage;
