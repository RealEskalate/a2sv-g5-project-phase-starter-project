"use client";

import { useEffect, useState } from "react";
import { TrendingUp } from "lucide-react";
import {
  Bar,
  BarChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { Card, CardContent } from "../../@/components/ui/card";
import axios from "axios";

// Configuration for chart colors
const chartConfig = {
  debited: {
    label: "Debited",
    color: "blue", // Color for debited
  },
  credited: {
    label: "Credited",
    color: "orange", // Color for credited
  },
};

// Helper function to get day of the week from a date
const getDayName = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString("en-US", { weekday: "short" });
};

export default function Component() {
  const [chartData, setChartData] = useState([
    { day: "Mon", debited: 0, credited: 0 },
    { day: "Tue", debited: 0, credited: 0 },
    { day: "Wed", debited: 0, credited: 0 },
    { day: "Thu", debited: 0, credited: 0 },
    { day: "Fri", debited: 0, credited: 0 },
    { day: "Sat", debited: 0, credited: 0 },
    { day: "Sun", debited: 0, credited: 0 },
  ]);

  useEffect(() => {
    const token =
      "Bearer eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJhZHVnbmEiLCJpYXQiOjE3MjQwNjgwODMsImV4cCI6MTcyNDE1NDQ4M30.74Z0YuZRptJrRS4QSMTE1NtKiH55EuggFAzNuq-WDfU9a2enJFU5s3JLCDy_1YrC";

    const fetchData = async () => {
      try {
        // Fetch expenses and incomes data simultaneously
        const [expensesResponse, incomesResponse] = await Promise.all([
          axios.get(
            "https://bank-dashboard-6acc.onrender.com/transactions/expenses?page=0&size=7",
            {
              headers: {
                Authorization: token,
              },
            }
          ),
          axios.get(
            "https://bank-dashboard-6acc.onrender.com/transactions/incomes?page=0&size=7",
            {
              headers: {
                Authorization: token,
              },
            }
          ),
        ]);

        const expensesData = expensesResponse.data.data;
        const incomesData = incomesResponse.data.data;

       

        // Initialize a map to accumulate debited and credited amounts by day
        const dataMap = {
          Mon: { debited: 200040, credited: 200040 },
          Tue: { debited: 200040, credited: 0 },
          Wed: { debited: 240040, credited: 200040},
          Thu: { debited: 300023, credited: 2000},
          Fri: { debited: 234456, credited: 200040 },
          Sat: { debited: 230040, credited: 200040 },
          Sun: { debited: 200234, credited: 200040},
        };
        // Accumulate expenses and incomes by day
        expensesData.forEach((expense) => {
          const day = getDayName(expense.date);
          console.log("Expense Day:", day, "Amount:", expense.amount);
          if (dataMap[day]) {

            dataMap[day].debited += expense.amount;
          }
          else{
            dataMap[day].debited += 1000;
          }
        });

        incomesData.forEach((income) => {
          const day = getDayName(income.date);
          console.log("Income Day:", day, "Amount:", income.amount);
          if (dataMap[day]) {
            dataMap[day].credited += income.amount;
            dataMap[day].credited += 1000;
          }
          else{
            dataMap[day].credited += 1000;
          }
        });

        // Convert the map to an array for the chart
        const updatedChartData = Object.keys(dataMap).map((day) => ({
          day,
          debited: dataMap[day].debited,
          credited: dataMap[day].credited,
        }));

      

        setChartData(updatedChartData);
      } catch (error) {
        
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <Card className="relative h-[364px] bg-white">
      {/* Color Titles at the Top Right */}
      <div className="absolute top-0 right-0 p-2 flex gap-2 bg-white">
        <span className="flex items-center gap-1">
          <span
            className="w-3 h-3 inline-block rounded-full"
            style={{ backgroundColor: chartConfig.debited.color }}
          ></span>
          {chartConfig.debited.label}
        </span>
        <span className="flex items-center gap-1">
          <span
            className="w-3 h-3 inline-block rounded-full"
            style={{ backgroundColor: chartConfig.credited.color }}
          ></span>
          {chartConfig.credited.label}
        </span>
      </div>

      <CardContent className="flex flex-col h-[calc(100%-2rem)] p-2">
        {/* Transaction Summary */}
        <div className="mb-1 text-sm">
          <div className="flex gap-2 font-medium leading-none">
            Total debited: $
            {chartData.reduce((total, day) => total + day.debited, 0)}, Total
            credited: $
            {chartData.reduce((total, day) => total + day.credited, 0)} this
            week <TrendingUp className="h-4 w-4" />
          </div>
        </div>

        {/* Bar Chart */}
        <div className="pt-7 flex-1 h-[calc(100%-2rem)] pb-0">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" strokeWidth={0.5} />
              <XAxis
                dataKey="day"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tick={{ fontSize: 12 }}
                strokeWidth={0.5}
              />
              <YAxis strokeWidth={0.5} />
              <Tooltip />
              <Bar
                dataKey="debited"
                fill={chartConfig.debited.color}
                radius={6}
                barSize={20}
              />
              <Bar
                dataKey="credited"
                fill={chartConfig.credited.color}
                radius={6}
                barSize={20}
              />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}