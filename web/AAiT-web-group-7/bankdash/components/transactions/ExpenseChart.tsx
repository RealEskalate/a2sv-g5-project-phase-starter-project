import React from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const data = [
  { name: "Aug", Expense: 1000 },
  { name: "Sep", Expense: 2000 },
  { name: "Oct", Expense: 3000 },
  { name: "Nov", Expense: 4000 },
  { name: "Dec", Expense: 5000 },
  { name: "Jan", Expense: 600 },
  { name: "Feb", Expense: 700 },
  { name: "Mar", Expense: 8000 },
  { name: "Apr", Expense: 9000 },
  { name: "May", Expense: 100 },
  { name: "Jun", Expense: 11000 },
  { name: "Jul", Expense: 1000 },
];

const ExpenseChart = () => {
  return (
    <div className="w-[350px] h-fit mx-auto md:mx-0 md:w-[330px] rounded-3xl px-5 py-3 space-y-5 shadow-md bg-white">
      <ResponsiveContainer width="100%" height={225}>
        <BarChart
          data={data}
          margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="name" />
          <YAxis hide={true} />
          <Tooltip />
          <Bar
            dataKey="Expense"
            fill="#16DBCC"
            radius={[10, 10, 0, 0]}
            barSize={37}
          />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default ExpenseChart;
