"use client";

import React from "react";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
  Dot,
} from "recharts";

const data = [
  { year: 2016, revenue: 3500 },
  { year: 2017, revenue: 6000 },
  { year: 2018, revenue: 2500 },
  { year: 2019, revenue: 7000 },
  { year: 2020, revenue: 4000 },
  { year: 2021, revenue: 8000 },
];

// A function to add break lines for clarity
const BreakLine = ({ points }: { points: number[] }) => {
  return (
    <>
      {points.map((point: number, index: number) => (
        <Line
          key={index}
          type="linear"
          dataKey="revenue"
          stroke="#f59e0b"
          strokeWidth={4}
          dot={<Dot r={5} stroke="#f59e0b" strokeWidth={2} fill="none" />}
          data={[point]}
        />
      ))}
    </>
  );
};

export default function YearlyRevenueChart() {
  const points = data.map((item, index) => ({
    ...item,
    // Add condition to create break points. This is a simple approach for the demo.
    // In production, you may need more sophisticated logic to handle breaks dynamically.
    revenue: index % 2 === 0 ? item.revenue : null,
  }));

  return (
    <div className="bg-white shadow-lg rounded-lg p-4 w-full max-w-md mx-auto">
      <h2 className="text-lg font-semibold mb-4">Yearly Revenue</h2>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={points}>
          <CartesianGrid strokeDasharray="3 3" stroke="#e5e7eb" />{" "}
          {/* Tailwind gray-200 */}
          <XAxis dataKey="year" tick={{ fill: "#6b7280" }} />{" "}
          {/* Tailwind gray-500 */}
          <YAxis tick={{ fill: "#6b7280" }} /> {/* Tailwind gray-500 */}
          <Tooltip />
          {/* Render the line with breaks */}
          <BreakLine points={points} />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
