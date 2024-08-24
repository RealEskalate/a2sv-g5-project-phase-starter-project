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
import { Card, CardContent } from "@/components/ui/card";
import axios from "axios";
import { useSession } from "next-auth/react";

// Configuration for chart colors
const chartConfig = {
  debited: {
    label: "Deposit",
    color: "#1814F3", // Color for debited
  },
  credited: {
    label: "Withdraw",
    color: "#16DBCC", // Color for credited
  },
};

// Helper function to get day of the week from a date
const getDayName = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString("en-US", { weekday: "short" });
};

export default function Component() {
  const { data: session } = useSession();
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
    const token = `Bearer ${session?.user?.accessToken}`;

    const fetchData = async () => {
      try {
        // Fetch expenses and incomes data simultaneously
        const [expensesResponse, incomesResponse] = await Promise.all([
          axios.get(
            "https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=7",
            {
              headers: {
                Authorization: token,
              },
            }
          ),
          axios.get(
            "https://bank-dashboard-rsf1.onrender.com/transactions/incomes?page=0&size=7",
            {
              headers: {
                Authorization: token,
              },
            }
          ),
        ]);

        const expensesData = expensesResponse.data.data.content;
        const incomesData = incomesResponse.data.data.content;

        // Initialize a map to accumulate debited and credited amounts by day
        const dataMap = {
          Mon: { debited: 12000, credited: 10000},
          Tue: { debited: 15000, credited: 10000},
          Wed: { debited: 2344, credited: 7000 },
          Thu: { debited: 3345, credited: 9000},
          Fri: { debited: 12340, credited: 1000 },
          Sat: { debited: 8000, credited: 5000 },
          Sun: { debited:5000, credited: 6000},
        };

        // Accumulate expenses and incomes by day
        expensesData.forEach((expense) => {
          const day = getDayName(expense.date);
          if (dataMap[day]) {
            dataMap[day].debited += expense.amount;
          }
        });

        incomesData.forEach((income) => {
          const day = getDayName(income.date);
          if (dataMap[day]) {
            dataMap[day].credited += income.amount;
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
  }, [session]);


  return (
    <Card className="relative h-[364px] bg-white w-full">
      {/* Color Titles at the Top Right */}
      <div className="absolute top-0 right-0 p-2 flex gap-2 bg-white">
        <span className="flex text-sm items-center gap-1">
          <span
            className="w-3 h-3 inline-block rounded-full"
            style={{ backgroundColor: chartConfig.debited.color }}
          ></span>
          {chartConfig.debited.label}
        </span>
        <span className="flex text-sm items-center gap-1">
          <span
            className="w-3 h-3 inline-block rounded-full"
            style={{ backgroundColor: chartConfig.credited.color }}
          ></span>
          {chartConfig.credited.label}
        </span>
      </div>

      <CardContent className="flex  flex-col h-[calc(100%-2rem)] w-full">
        {/* Bar Chart */}
        <div className="pt-7 w-full flex-1 h-[calc(100%-2rem)] pb-0">
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
              <YAxis
                width={70} // Increased width for better visibility of large numbers
                tickMargin={10}
                tick={{ fontSize: 9 }} // Adjust font size if necessary
                strokeWidth={0.5}
                // padding={{ right: 10 }} // Adding padding for better visibility
              />
              {/* <Tooltip /> */}
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
