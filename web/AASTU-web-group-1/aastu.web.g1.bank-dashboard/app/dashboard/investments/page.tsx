import React from "react";
import Card from "./investment_components/Card";
import { investmentTypes } from "@/constants";
import Investment from "./investment_components/Investment";
const Investments = () => {

  return(
    <div>
      <div className=" flex flex-col md:flex-row gap-3 md:gap-10 justify-center pt-4 w-full">
        {investmentTypes.map((item) => (
          <Card {...item} key="item.name"/>
        ))}
      </div>
      <div>
        Some graphs
      </div>
      <div className=" flex md:flex-row gap-3 md:gap-10 pt-4 w-full justify-between">
        <Investment/>
      </div>
    </div>
  )
};

export default Investments;
