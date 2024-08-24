"use client";

import { NormalTextColor } from "@/constants/transactions/colors";
import React, { useState } from "react";
import {
    Bar,
    BarChart,
    CartesianGrid,
    LabelList,
    XAxis,
} from "recharts";

export function ExpenseChart() {
    const isDarkMode = false;
    const [activeIndex, setActiveIndex] = useState<number | null>(null);

    const chartData = [
        { month: "Jan", expense: 2000 },
        { month: "Feb", expense: 1500 },
        { month: "Mar", expense: 1000 },
        { month: "Apr", expense: 2500 },
        { month: "May", expense: 2000 },
        { month: "Jun", expense: 700 },
    ];

    const barSize = 30; // Store barSize in a variable

    return (
        <BarChart
            className="z-50"
            width={350}
            height={200}
        
            data={chartData}
            margin={{
                top: 10,
                right: 10,
                left: 10,
                bottom: 10, 
            }}
        >
            <XAxis
                dataKey="month"
                axisLine={false} // Hide the axis line
                tickLine={false} // Hide the tick lines
                tick={{ fill: isDarkMode ? "white" : "#718EBF", fontSize: 14 }} // Customize the tick labels
            />
            <Bar
                dataKey="expense"
                fill={isDarkMode ? "#1e3a8a" : "#EDF0F7"} // Bar color
                radius={[10, 10, 10, 10]} // Rounded corners on top
                barSize={barSize}
                style={{ transition: "fill 0.3s ease" }}
                className="hover:fill-[#16DBCC] pl-2"
                onMouseEnter={(data, index) => setActiveIndex(index)}
                onMouseLeave={() => setActiveIndex(null)}
            >
                {/* Expense labels on hover */}
                <LabelList
                    dataKey="expense"
                    position="centerTop"
                    fill={"#343C6A"} // Label color
                    fontSize={14}
                    className="z-50"
                    content={(props) => {
                        const { x, y, value, index } = props;
                        return index === activeIndex ? (
                            <text
                                x={x}
                                y={y}
                                dy={-6}
                                fill={"#343C6A"}
                                fontSize={14}
                                textAnchor="middle"
                                className={"font-medium " + NormalTextColor}
                            >
                                {`$${value}`}
                            </text>
                        ) : null;
                    }}
                />
            </Bar>

        </BarChart>
    );
}
