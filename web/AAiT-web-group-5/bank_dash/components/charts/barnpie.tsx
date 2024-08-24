import React from "react";
import BarStat from "./bar";
import PieStat from "./pie";

const BarNPie = () => {
  return (
    <div className="flex gap-3">
      <div className="w-2/3">
        <BarStat />
      </div>
      <div className="w-1/3">
        <PieStat />
      </div>
    </div>
  );
};

export default BarNPie;
