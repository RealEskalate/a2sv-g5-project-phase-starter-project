"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
import LastTransService from "@/app/Services/api/lastTransService";
import { useState, useEffect } from "react";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartConfig = {
  debit: {
    label: "Debit",
    color: "hsl(var(--chart-1))",
  },
  credit: {
    label: "Credit",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig;
interface chartData {
  day: string;
  debit: number;
  credit: number;
}
function isDateInLast7Days(dateString: string): boolean {
  const currentDate = new Date();
  const sevenDaysAgo = new Date(currentDate);
  sevenDaysAgo.setDate(currentDate.getDate() - 7);
  sevenDaysAgo.setHours(0, 0, 0, 0);

  const transactionDate = new Date(dateString);
  return transactionDate >= sevenDaysAgo && transactionDate <= currentDate;
}
function getDayOfWeek(dateString: string): string {
  const date = new Date(dateString);
  const daysOfWeek = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ];
  return daysOfWeek[date.getDay()];
}
export function DebitCreditOver() {
  const [data, setData] = useState<chartData[]>([]);
  const [totalIncome, setTotalIncome] = useState(0);  
  const [totalExpense, setTotalExpense] = useState(0);
  useEffect(() => {
    const getData = async () => {
      try {
         const accessToken = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJtaWhyZXQiLCJpYXQiOjE3MjQwNTY5MDEsImV4cCI6MTcyNDE0MzMwMX0.06ogiDUHZipaLn7gIoJDxGz4Bw_zFmsA72Zp99eKGkKVjOFRXy3MUvh55dspPaib"

        const expense = await LastTransService.getExpenseData(accessToken);
        const income = await LastTransService.getIncomeData(accessToken);
        const chartData: { [day: string]: { debit: number; credit: number } } =
          {
            Sunday: { debit: 0, credit: 0 },
            Monday: { debit: 0, credit: 0 },
            Tuesday: { debit: 0, credit: 0 },
            Wednesday: { debit: 0, credit: 0 },
            Thursday: { debit: 0, credit: 0 },
            Friday: { debit: 0, credit: 0 },
            Saturday: { debit: 0, credit: 0 },
          };
          let incomeSum = 0;
          let expenseSum = 0;
        income.forEach((transaction) => {
          if (isDateInLast7Days(transaction.date)) {
            incomeSum+=transaction.amount;
            const dayOfWeek = getDayOfWeek(transaction.date);
            chartData[dayOfWeek].credit += transaction.amount;
          }
        });

        expense.forEach((transaction) => {
          if (isDateInLast7Days(transaction.date)) {
            expenseSum+=transaction.amount;
            const dayOfWeek = getDayOfWeek(transaction.date);
            chartData[dayOfWeek].debit += transaction.amount;
          }
        });

        const formattedChartData = Object.keys(chartData).map((day) => ({
          day: day,
          debit: chartData[day].debit,
          credit: chartData[day].credit,
        }));
        const currentDayIndex = new Date().getDay();
        const rotatedChartData = [
          ...formattedChartData.slice(currentDayIndex + 1),
          ...formattedChartData.slice(0, currentDayIndex + 1),]
        setData(rotatedChartData);
        setTotalExpense(expenseSum);
        setTotalIncome(incomeSum)
      } catch (error) {
        alert("Error Fetching data ");
      }
    };
    getData();
  }, []);
  return (
    <Card className="rounded-3xl shadow-lg border-gray-300   ">
      <CardHeader>
        <div className="flex justify-between ">
          <CardTitle className=" hidden lg:block text-base font-normal font-inter text-[#718EBF]">
            <span className="text-black">${totalExpense}</span> Debited &{" "}
            <span className="text-black">${totalIncome}</span> Credited in this Week
          </CardTitle>
          <div className="flex gap-5">
            <div className="flex items-center">
              <div className="border border-[#4C78FF] w-[15px] h-[15px] rounded-sm bg-[#4C78FF]"></div>
              <p className="font-inter font-normal text-base text-[#718EBF]">
                Debit
              </p>
            </div>
            <div className="flex items-center">
              <div className="border border-[#FCAA0B] w-[15px] h-[15px] rounded-sm bg-[#FCAA0B]"></div>
              <p className="font-inter font-normal text-base text-[#718EBF]">
                Credit
              </p>
            </div>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="h-[350px] w-[100%]">
          <BarChart accessibilityLayer data={data}>
            <CartesianGrid vertical={false} className="h-[70%]" />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="debit" fill="#1A16F3" radius={10} />
            <Bar dataKey="credit" fill="#FCAA0B" radius={10} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
