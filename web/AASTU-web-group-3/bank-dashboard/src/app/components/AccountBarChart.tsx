'use client';
import { useGetExpenseTransactionsQuery, useGetIncomeTransactionsQuery } from "@/lib/redux/api/transactionsApi";
import { IncomeResponse, MyExpenseResponse } from "@/lib/redux/types/transactions";
import React, { useEffect, useState } from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

interface ChartData {
  name: string;
  deposit: number;
  withdraw: number;
}

const Example = () => {
  const [Expence, setExpence] = useState<MyExpenseResponse>();
  const [Income, setIncome] = useState<IncomeResponse>();
  
  const { data: expense, isLoading: isExpenseLoading, isError: isExpenseError } = useGetExpenseTransactionsQuery({ page: 0, size: 8 });
  const { data: income, isLoading: isIncomeLoading, isError: isIncomeError } = useGetIncomeTransactionsQuery({ page: 0, size: 8 });

  useEffect(() => {
    if (expense) {
      setExpence(expense);
    }
  }, [expense]);
  
  useEffect(() => {
    if (income) {
      setIncome(income);
    }
  }, [income]);

  const data: ChartData[] = [];

  useEffect(() => {
    if (Income && Expence) {
      const mappedData = Income.data.content.map((inc, index: number) => ({
        name: `Entry ${index + 1}`,
        deposit: inc.amount,
        withdraw: Expence.data.content[index]?.amount || 0,
      }));
      data.push(...mappedData);
    }
  }, [Income, Expence]);

  console.log("income and expence", Income, Expence);
  
  if (isIncomeLoading || isExpenseLoading || Expence === undefined || Income === undefined) {
    return <div>Loading...</div>;
  }

  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart
        width={500}
        height={200}
        data={data}
        margin={{ top: 5, right: 30, left: 20, bottom: 5 }}
        barCategoryGap="20%"
        barGap={8}
      >
        <CartesianGrid strokeDasharray="3 3" />
        <XAxis dataKey="name" />
        <YAxis domain={[100, 500]} ticks={[100, 200, 300, 400, 500]} /> {/* Removed 0 from ticks */}
        <Tooltip />
        <Legend verticalAlign="top" align="right" />
        <Bar
          dataKey="withdraw"
          fill="#1814F3"
          radius={[10, 10, 10, 10]}
          barSize={20}
        />
        <Bar
          dataKey="deposit"
          fill="#16DBCC"
          radius={[10, 10, 10, 10]}
          barSize={20}
        />
      </BarChart>
    </ResponsiveContainer>
  );
};

export default Example;
