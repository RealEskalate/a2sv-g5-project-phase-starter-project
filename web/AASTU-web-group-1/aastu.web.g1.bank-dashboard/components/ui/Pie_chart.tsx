"use client";
import {
  LabelList,
  Pie,
  PieChart,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import {
  Card,
  CardContent,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { getallTransactions } from "@/app/dashboard/transactions/component/getTransactions";
import { TransactionData } from "@/types";
import { useState,useEffect } from "react";




const categoryTotals ={
  Other:0,
  Investment:0,
  BillExpense:0,
  Entertainment:0
};


export function Pie_chart() {
  const [transactions, setTransactions] = useState<TransactionData[]>([]);
  const [chartData,setChartData] = useState<{ expense: string; value: number; fill: string; }[]>([]);
  useEffect(()=>{
    const fetchTransactions = async () => {
      const transactions = await getallTransactions(0, 100);
       setTransactions(transactions || []);}
       fetchTransactions();
  },[]);

  useEffect(()=>{
      transactions.map((transaction: TransactionData) => {
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
            categoryTotals.BillExpense += transaction.amount;
            break;
          default:
            categoryTotals.Other += transaction.amount;
            break;
        }
      });

      const totalSum =
        categoryTotals.Other +
        categoryTotals.Investment +
        categoryTotals.BillExpense +
        categoryTotals.Entertainment;

        const newChartData = [
          {
            expense: "Other",
            value: Math.round((categoryTotals.Other / totalSum) * 100),
            fill: "#1814F3",
          },
          {
            expense: "Investment",
            value: Math.round((categoryTotals.Investment / totalSum) * 100),
            fill: "#343C6A",
          },
          {
            expense: "Bill Expense",
            value: Math.round((categoryTotals.BillExpense / totalSum) * 100),
            fill: "#FC7900",
          },
          {
            expense: "Entertainment",
            value: Math.round((categoryTotals.Entertainment / totalSum) * 100),
            fill: "#FA00FF",
          },
        ];
        setChartData(newChartData);
  },[transactions]);
  
 const chartConfig = {
   value: {
     label: "Value",
   },
   other: {
     label: "Other",
     color: "#1814F3",
   },
   investment: {
     label: "Investment",
     color: "#343C6A",
   },
   billExpense: {
     label: "Bill Expense",
     color: "#FC7900",
   },
   entertainment: {
     label: "Entertainment",
     color: "#FA00FF",
   },
 } satisfies ChartConfig;

  return (
    <Card >
      <CardContent>
        <ChartContainer config={chartConfig} className="h-60 w-full ">
        
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
                  formatter={(value: number) =>
                    ` ${value}`
                  }
                  // stroke="#fff"
                  // className=" font-bold text-blue "
                />
              </Pie>
            </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
