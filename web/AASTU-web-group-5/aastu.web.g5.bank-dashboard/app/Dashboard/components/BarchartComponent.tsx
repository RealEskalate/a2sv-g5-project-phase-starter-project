"use client";

import { useState, useEffect } from "react";
import axios from "axios";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";

import { Card, CardContent } from "@/components/ui/card";
import { ChartConfig, ChartContainer } from "@/components/ui/chart";

// Define types for transactions and chart data
interface Transaction {
  date: string;
  type: 'deposit' | 'withdraw';
  amount: number;
}

interface ChartData {
  day: string;
  deposit: number;
  withdraw: number;
}

const chartConfig = {
  deposit: {
    label: "Deposit",
    color: "hsl(var(--chart-1))",
  },
  withdraw: {
    label: "Withdraw",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig;

// Helper function to aggregate data by day of the week
const aggregateDataByDay = (data: { date: string; amount: number; type: 'deposit' | 'withdraw'; }[]) => {
  const daysOfWeek = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
  const aggregatedData = daysOfWeek.map(day => ({ day, deposit: 0, withdraw: 0 }));

  data.forEach(({ date, amount, type }) => {
    const dayOfWeek = new Date(date).getDay(); // Get day of week (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
    if (type === 'deposit') {
      aggregatedData[dayOfWeek].deposit += amount;
    } else {
      aggregatedData[dayOfWeek].withdraw += amount;
    }
  });

  return aggregatedData;
};

export function BarchartComponent() {
  const [data, setData] = useState<ChartData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get<{ data: Transaction[] }>('https://bank-dashboard-6acc.onrender.com/transactions?page=0', {
          headers: {
            Authorization: `Bearer ${process.env.NEXT_PUBLIC_ACCESS_TOKEN}`,
          },
        });

        const transactions = response.data.data;

        // Process data to aggregate deposits and withdrawals by day of the week
        const processedData = aggregateDataByDay(transactions);

        setData(processedData);
      } catch (err) {
        setError('Failed to fetch data. Please check the console for more details.');
        console.error('Error fetching data:', err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <Card className="max-w-full overflow-hidden">
      <CardContent className="p-4 relative">
        <ChartContainer config={chartConfig}>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={data}>
              <CartesianGrid vertical={false} />
              <XAxis
                dataKey="day"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tickFormatter={(value) => value}
              />
              <YAxis
                tickCount={6}
                tickSize={5}
                tickFormatter={(value) => `${value / 1000}K`} // Format as thousands
                domain={[0, 'dataMax']} // Adjust if needed
                interval="preserveStartEnd"
                orientation="left"
              />
              <Tooltip />
              <Bar
                dataKey="deposit"
                fill="blue"
                radius={[10, 10, 10, 10]}
                barSize={10}
                name="Deposit"
              />
              <Bar
                dataKey="withdraw"
                fill="green"
                radius={[10, 10, 10, 10]}
                barSize={10}
                name="Withdraw"
              />
            </BarChart>
          </ResponsiveContainer>
        </ChartContainer>
        <div className="absolute top-0 right-4 p-2 bottom-4">
          <div className="flex flex-row items-center space-x-4">
            <div className="flex items-center">
              <div className="w-4 h-4 rounded-full mr-2" style={{ backgroundColor: "blue" }}></div>
              <span className="text-sm">Deposit</span>
            </div>
            <div className="flex items-center">
              <div className="w-4 h-4 mr-2 rounded-full" style={{ backgroundColor: "green" }}></div>
              <span className="text-sm">Withdraw</span>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
