"use client"
import React, { PureComponent } from "react";
import {
  BarChart,
  Bar,
  Rectangle,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  LabelList,
} from "recharts";


const data = [
  {
    name: "Sat",
    deposit: 250,
    withdraw: 490
  },
  {
    name: "Sun",
    deposit: 110,
    withdraw: 350
  },
  {
    name: "Mon",
    deposit: 110,
    withdraw: 350
  },
  {
    name: "Tue",
    deposit: 390,
    withdraw: 490
  },
  {
    name: "Wed",
    deposit: 150,
    withdraw: 120
  },
  {
    name: "Thu",
    deposit: 250,
    withdraw: 410
  },
  {
    name: "Fri",
    deposit: 310,
    withdraw: 400
  },
];

export default class Example extends PureComponent {
  static demoUrl = "https://codesandbox.io/p/sandbox/simple-bar-chart-72d7y5";

  render() {
    return (
      <ResponsiveContainer width="100%" height={300}>
        <BarChart
          width={500}
          height={200}
          data={data}
          margin={{
            top: 5,
            right: 30,
            left: 20,
            bottom: 5,
          }}
          barCategoryGap="20%"
          barGap={8}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="name" />
          <YAxis domain={[100, 500]} ticks={[0, 100, 200, 300, 400, 500]} />
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
  }
}