"use client";

import { CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis, ResponsiveContainer } from "recharts";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from "../../@/components/ui/card";

const monthlyRevenueData = [
  { month: "Jan", revenue: 4000 },
  { month: "Feb", revenue: 3500 },
  { month: "Mar", revenue: 4500 },
  { month: "Ap", revenue: 5000 },
  { month: "May", revenue: 4700 },
  { month: "Jun", revenue: 5200 },
  { month: "Jul", revenue: 4800 },
  { month: "Aug", revenue: 5100 },
  { month: "Sep", revenue: 5300 },
  { month: "Oct", revenue: 4900 },
  { month: "Nov", revenue: 5500 },
  { month: "Dec", revenue: 5800 },
];

export default function MonthlyRevenueChart() {
  return (
    <Card style={{ height: '100%' }}>
      <CardContent style={{ height: '100%' }}>
        <div className='pt-6' style={{ width: '100%', height: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <LineChart
              data={monthlyRevenueData}
              margin={{
                top: 5, right: 30, left: 20, bottom: 5,
              }}
            >
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="month" />
              <YAxis tickFormatter={(value) => `$${value}`} />
              <Tooltip formatter={(value) => `$${value}`} />
              <Line type="monotone" dataKey="revenue" stroke="#82ca9d" strokeWidth={3} dot={false} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}