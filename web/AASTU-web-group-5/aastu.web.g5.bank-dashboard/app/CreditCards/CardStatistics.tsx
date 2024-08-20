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

	return (
		<div className="statics  ">
			<div className="p-3 font-semibold text-blue-900">
				Card Expense Statistics
			</div>
			<div>
				<div className="w-full  bg-white rounded-2xl">
					<div className="flex justify-center p-6">
						<DonutChart data={chartData} />
					</div>
				</div>
			</div>
		</div>
	);
};

export default CardStatistics;
