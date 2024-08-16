import React from "react";
import { Label, Pie, PieChart, Sector } from "recharts";
import { PieSectorDataItem } from "recharts/types/polar/Pie";

import { Card, CardContent } from "@/components/ui/card";
import {
	ChartConfig,
	ChartContainer,
	ChartTooltip,
	ChartTooltipContent,
} from "@/components/ui/chart";

interface DonutChartProps {
	data: {
		browser: string;
		visitors: number;
		fill: string;
	}[];
}

const chartConfig = {
	visitors: {
		label: "Visitors",
	},
	chrome: {
		label: "DBL Bank",
		color: "hsl(var(--chart-1))",
	},
	safari: {
		label: "ABM Bank",
		color: "hsl(var(--chart-2))",
	},
	firefox: {
		label: "BRC Bank",
		color: "hsl(var(--chart-3))",
	},
	edge: {
		label: "MCP Bank",
		color: "hsl(var(--chart-4))",
	},
} satisfies ChartConfig;

export default function DonutChart({ data }: DonutChartProps) {
	return (
		<Card className="flex flex-col border-none">
			<CardContent className="flex-1 pb-0">
				<ChartContainer
					config={chartConfig}
					className="mx-auto aspect-square max-h-[200px] w-full"
				>
					<PieChart width={250} height={250}>
						<ChartTooltip
							cursor={true}
							content={<ChartTooltipContent hideLabel />}
						/>
						<Pie
							data={data}
							dataKey="visitors"
							nameKey="browser"
							innerRadius={35}
							strokeWidth={5}
							activeIndex={0}
							activeShape={({
								outerRadius = 0,
								...props
							}: PieSectorDataItem) => (
								<Sector {...props} outerRadius={outerRadius + 10} />
							)}
						/>
					</PieChart>
				</ChartContainer>
			</CardContent>
			<div className="flex justify-center px-5 py-2 gap-2">
				<div>
					{data.slice(0, 2).map((item, index) => (
						<div key={index} className="flex items-center gap-1 pb-2">
							<div
								className="w-3.5 h-3.5 rounded-full"
								style={{ backgroundColor: item.fill }}
							></div>
							<div className="text-blue-900 opacity-60">{item.browser}</div>
						</div>
					))}
				</div>
				<div>
					{data.slice(2).map((item, index) => (
						<div key={index} className="flex items-center gap-1 pb-2">
							<div
								className="w-3.5 h-3.5 rounded-full"
								style={{ backgroundColor: item.fill }}
							></div>
							<div className="text-blue-900 opacity-60">{item.browser}</div>
						</div>
					))}
				</div>
			</div>
		</Card>
	);
}
