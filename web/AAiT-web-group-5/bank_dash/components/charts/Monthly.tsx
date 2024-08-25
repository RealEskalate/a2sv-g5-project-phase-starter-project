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
} from "recharts";

const data = [
  { year: 2016, revenue: 3500 },
  { year: 2017, revenue: 6000 },
  { year: 2018, revenue: 2500 },
  { year: 2019, revenue: 7000 },
  { year: 2020, revenue: 4000 },
  { year: 2021, revenue: 8000 },
];

export default function Monthly() {
  return (
    <div className="bg-white shadow-lg rounded-lg p-4 w-full max-w-md mx-auto">
      <h2 className="text-lg  mb-4">Yearly Revenue</h2>
      <ResponsiveContainer width="100%" height={200}>
        <LineChart data={data}>
          <CartesianGrid strokeDasharray="3 3" stroke="#e5e7eb" />
          <XAxis dataKey="year" tick={{ fill: "#6b7280" }} />
          <YAxis tick={{ fill: "#6b7280" }} />
          <Tooltip />
          <Line
            type="monotone"
            dataKey="revenue"
            stroke="#10b981"
            strokeWidth={4}
            dot={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
