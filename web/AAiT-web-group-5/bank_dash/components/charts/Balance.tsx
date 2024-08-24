"use client";

import React from "react";
import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

const data = [
  { month: "Jul", balance: 200 },
  { month: "Aug", balance: 350 },
  { month: "Sep", balance: 275 },
  { month: "Oct", balance: 425 },
  { month: "Nov", balance: 700 },
  { month: "Dec", balance: 250 },
  { month: "Jan", balance: 500 },
];

export default function BalanceChart() {
  return (
    <div className="bg-white shadow-lg rounded-lg p-4 w-full max-w-md mx-auto">
      <h2 className="text-lg  mb-4">Balance</h2>
      <ResponsiveContainer width="100%" height={200}>
        <AreaChart data={data}>
          <defs>
            <linearGradient id="gradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="5%" stopColor="#6d28d9" stopOpacity={0.8} />
              <stop offset="95%" stopColor="#6d28d9" stopOpacity={0} />
            </linearGradient>
          </defs>
          <CartesianGrid strokeDasharray="3 3" stroke="#e5e7eb" />{" "}
          <XAxis dataKey="month" tick={{ fill: "#6b7280" }} />{" "}
          <YAxis tick={{ fill: "#6b7280" }} />
          <Tooltip />
          <Area
            type="monotone"
            dataKey="balance"
            stroke="#4f46e5"
            fill="url(#gradient)"
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
}
