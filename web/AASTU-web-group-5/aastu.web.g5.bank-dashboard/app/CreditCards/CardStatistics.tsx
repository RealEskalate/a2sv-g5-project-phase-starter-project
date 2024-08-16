"use client";

import React, { useState, useEffect } from "react";
import DonutChart from "./DonutChart";

const CardStatistics = () => {
	const [chartData, setChartData] = useState([
		{ browser: "DBL Bank", visitors: 275, fill: "#4C78FF" },
		{ browser: "ABM Bank", visitors: 200, fill: "#16DBCC" },
		{ browser: "BRC Bank", visitors: 187, fill: "#FF82AC" },
		{ browser: "MCP Bank", visitors: 173, fill: "#FFBB38" },
	]);

	useEffect(() => {
		// Example: Fetch data from an API or compute data dynamically
		// fetch("/api/chartData")
		// 	.then((response) => response.json())
		// 	.then((data) => setChartData(data))
		// 	.catch((error) => console.error("Error fetching chart data:", error));
		// Or update chartData dynamically based on your application logic
	}, []);

	return (
		<div className="statics ">
			<div className="p-4 pl-5 font-semibold text-blue-900">
				Card Expense Statistics
			</div>
			<div className="pl-4">
				<div className="w-80 bg-white rounded-2xl pb-8 ">
					<div className="flex justify-center ">
						<DonutChart data={chartData} />
					</div>
				</div>
			</div>
		</div>
	);
};

export default CardStatistics;
