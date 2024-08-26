"use client";
import { LabelList, Legend, Pie, PieChart } from "recharts";
import { useState, useEffect } from "react";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { TransactionContent } from "@/types";
import { useUser } from "@/contexts/UserContext";

const categoryTotals = {
  Other: 0,
  Investment: 0,
  BillExpense: 0,
  Entertainment: 0,
};

export function Pie_chart({
  transactions,
}: {
  transactions: TransactionContent[];
}) {
  const { isDarkMode } = useUser();
  const [chartData, setChartData] = useState<
    { expense: string; value: number; fill: string }[]
  >([]);

  useEffect(() => {
    transactions.forEach((transaction: TransactionContent) => {
      switch (transaction.type) {
        case "shopping":
          categoryTotals.Entertainment += transaction.amount;
          break;
        case "deposit":
          categoryTotals.Investment += transaction.amount;
          break;
        case "service":
          categoryTotals.BillExpense += transaction.amount;
          break;
        case "transfer":
          categoryTotals.Other += transaction.amount;
          break;
        default:
          categoryTotals.Other += transaction.amount;
          break;
      }
    });
    console.log(transactions)
    const totalSum =
      categoryTotals.Other +
      categoryTotals.Investment +
      categoryTotals.BillExpense +
      categoryTotals.Entertainment;

    const newChartData = [
      {
        expense: "Other",
        value: Math.round((categoryTotals.Other / totalSum) * 100),
        fill: isDarkMode ? "#FF4500" : "#1814F3", 
      },
      {
        expense: "Investment",
        value: Math.round((categoryTotals.Investment / totalSum) * 100),
        fill: isDarkMode ? "#8A2BE2" : "#343C6A", 
      },
      {
        expense: "Bill Expense",
        value: Math.round((categoryTotals.BillExpense / totalSum) * 100),
        fill: isDarkMode ? "#00CED1" : "#FC7900", 
      },
      {
        expense: "Entertainment",
        value: Math.round((categoryTotals.Entertainment / totalSum) * 100),
        fill: isDarkMode ? "#FF69B4" : "#FA00FF", 
      },
    ];
    setChartData(newChartData);
  }, [transactions, isDarkMode]);

  const chartConfig = {
    value: {
      label: "Value",
    },
    other: {
      label: "Other",
      color: isDarkMode ? "#FF4500" : "#1814F3",
    },
    investment: {
      label: "Investment",
      color: isDarkMode ? "#8A2BE2" : "#343C6A", 
    },
    billExpense: {
      label: "Bill Expense",
      color: isDarkMode ? "#00CED1" : "#FC7900", 
    },
    entertainment: {
      label: "Entertainment",
      color: isDarkMode ? "#FF69B4" : "#FA00FF", 
    },
  };

  return (
    <Card
      className={` ${
        isDarkMode ? "bg-gray-800  " : "bg-white "
      } py-5 border-none`}
    >
      <CardContent
        className="p-0 border-none"
        style={{
          backgroundColor: isDarkMode ? "#1f2937" : "#ffffff", 
          borderColor: isDarkMode ? "#333333" : "#dddddd", 
        }}
      >
        <ChartContainer
          config={chartConfig}
          className="h-64 w-full rounded-xl"
          style={{
            backgroundColor: isDarkMode ? "#1f2937" : "#ffffff", 
            borderColor: isDarkMode ? "#34495e" : "#dddddd", 
          }}
        >
          <PieChart>
            <ChartTooltip
              content={<ChartTooltipContent nameKey="expense" hideLabel />}
            />
            <Pie
              data={chartData}
              dataKey="value"
              nameKey="expense"
              outerRadius="80%"
              label
              paddingAngle={5}
            >
              <LabelList
                dataKey="expense"
                position="inside"
                fontSize={12}
                fontFamily="font-lustria"
                formatter={(value: number) => ` ${value}`}
              />
            </Pie>
            <Legend />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
