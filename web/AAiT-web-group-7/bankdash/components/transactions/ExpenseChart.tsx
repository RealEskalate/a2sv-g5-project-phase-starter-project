import React from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const data = [
  { name: "Aug", Expense: 1000 },
  { name: "Sep", Expense: 2000 },
  { name: "Oct", Expense: 3000 },
  { name: "Nov", Expense: 4000 },
  { name: "Dec", Expense: 5000 },
  { name: "Jan", Expense: 600 }
];

const ExpenseChart = () => {
  return (
    <div className="w-[350px] mx-auto md:mx-0 md:w-[330px] rounded-3xl shadow-md bg-white">
      <ResponsiveContainer width="100%" height={190}>
        <BarChart
          data={data}
          margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
          barCategoryGap="10%"  
        >
          <XAxis dataKey="name" />
          <YAxis hide={true} />
          <Tooltip />
          <Bar
            dataKey="Expense"
            fill="#16DBCC"
            radius={[10, 10, 0, 0]}
            barSize={25}  
          />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default ExpenseChart;
