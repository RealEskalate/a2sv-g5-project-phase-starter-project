"use client";

import React, { useState, useEffect } from "react";
import axios from "axios";
import { Bar, BarChart, CartesianGrid, XAxis, Cell, LabelList } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import { ChartConfig, ChartContainer } from "@/components/ui/chart";
import { useSession } from "next-auth/react";

const chartConfig = {
  desktop: {
    label: "Expenses",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

const getAllMonths = () => {
  const months = [
    "Jan", "Feb", "Mar", "Apr", "May", "Jun",
    "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"
  ];
  return months;
};

export function BarChartComponent() {
  const [chartData, setChartData] = useState<{ month: string, desktop: number }[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [activeIndex, setActiveIndex] = useState<number | null>(null);
  interface ExtendedUser {
    name?: string;
    email?: string;
    image?: string;
    accessToken?: string;
    }
    const { data: session, status } = useSession();
    const user = session?.user as ExtendedUser;
    const accessToken = user?.accessToken;
  useEffect(() => {
    const fetchChartData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-1tst.onrender.com/transactions/expenses?page=0&size=10', {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        });
        if (response.data.success) {
          // Initialize data with zeros for each month
          const months = getAllMonths();
          const initialData = months.map(month => ({ month, desktop: 0 }));

          // Map API response to chart format
          response.data.data.forEach((item: { date: string, amount: number }) => {
            const month = new Date(item.date).toLocaleString('default', { month: 'short' });
            const index = months.indexOf(month);
            if (index !== -1) {
              initialData[index].desktop += item.amount;
            }
          });

          setChartData(initialData);
        } else {
          setError('Failed to fetch data');
        }
      } catch (err) {
        console.error('Error fetching chart data:', err);
        setError('Failed to fetch data');
      } finally {
        setLoading(false);
      }
    };

    fetchChartData();
  }, []);

  const handleMouseEnter = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseLeave = () => {
    setActiveIndex(null);
  };

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart data={chartData} onMouseLeave={handleMouseLeave}>
            <CartesianGrid vertical={false} horizontal={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <Bar dataKey="desktop" radius={10}>
              {chartData.map((entry, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={activeIndex === index ? "#12887E33" : "#EDF0F7"}
                  onMouseEnter={() => handleMouseEnter(index)}
                />
              ))}
              <LabelList
                dataKey="desktop"
                position="top"
                content={({ x, y, value, index }) =>
                  activeIndex === index ? (
                    <text
                      x={x}
                      y={y}
                      dy={-10}
                      fill="black"
                      fontSize={12}
                      textAnchor="top"
                    >
                      {value}
                    </text>
                  ) : null
                }
              />
            </Bar>
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}

export default BarChartComponent;
