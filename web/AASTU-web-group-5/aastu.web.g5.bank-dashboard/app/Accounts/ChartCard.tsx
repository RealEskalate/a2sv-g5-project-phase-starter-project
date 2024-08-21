"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "../../@/components/ui/card";

const chartData = [
  { day: "Mon", debited: 150, credited: 100 },
  { day: "Tue", debited: 200, credited: 50 },
  { day: "Wed", debited: 120, credited: 80 },
  { day: "Thu", debited: 75, credited: 70 },
  { day: "Fri", debited: 100, credited: 40 },
  { day: "Sat", debited: 80, credited: 30 },
  { day: "Sun", debited: 50, credited: 50 },
];

const chartConfig = {
  debited: {
    label: "Debited",
    color: "blue",  // Red for debited
  },
  credited: {
    label: "Credited",
    color: "orange",  // Green for credited
  },
};

export default function Component() {
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
            Total debited: $775, Total credited: $420 this week <TrendingUp className="h-4 w-4" />
          </div>
        </div>

        {/* Bar Chart */}
        <div className="flex-1 h-[calc(100%-2rem)] pb-0">
          <ResponsiveContainer width="100%" height="80%">
            <BarChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis
                dataKey="day"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tick={{ fontSize: 12 }} // Adjust font size for better visibility
              />
              <YAxis />
              <Tooltip />
              <Bar dataKey="debited" fill={chartConfig.debited.color} radius={4} />
              <Bar dataKey="credited" fill={chartConfig.credited.color} radius={4} />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}