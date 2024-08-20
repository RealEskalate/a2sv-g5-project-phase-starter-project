"use client";

import { LabelList, Pie, PieChart } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { browser: "Services", visitors: 15, fill: "#FC7900", outerRadius: 130 },
  { browser: "Transfer", visitors: 30, fill: "#343C6A", outerRadius: 140 },
  { browser: "Shopping", visitors: 20, fill: "#FA00FF", outerRadius: 150 },
  { browser: "Other", visitors: 35, fill: "#1814F3", outerRadius: 160 },
];

const totalVisitors = 100;

const chartConfig = {
  Services: {
    label: `Services`,
    color: "#FC7900",
  },
  Transfer: {
    label: `Transfer`,
    color: "#343C6A",
  },
  Shopping: {
    label: `Shopping`,
    color: "#FA00FF",
  },
  Other: {
    label: `Other`,
    color: "#1814F3",
  },
} satisfies ChartConfig;

export function PieComp() {
  return (
    <div className="flex w-full it gap-6 p-8 bg-white dark:bg-[#232328]  rounded-3xl">
      <Card className="w-full flex flex-col border-0 shadow-none bg-transparent">
        <CardContent className="flex pb-0 p-0">
          <ChartContainer
            config={chartConfig}
            className="sm:w-full lg:w-[330px] lg:h-64"
          >
            <PieChart className=" ">
              <ChartTooltip
                content={<ChartTooltipContent nameKey="visitors" hideLabel />}
              />
              <Pie
                data={chartData}
                dataKey="visitors"
                innerRadius={2} // Adjust the inner radius of the pie
                paddingAngle={5} // Add padding between each section
                outerRadius={110} // Set a fixed maximum outer radius
              >
                {}
                <LabelList
                  dataKey="browser"
                  fill="#FFFFFF"
                  className="font-bold text-[13px] text-wrap "
                  stroke="none"
                  // fontSize={12}
                  formatter={(value: keyof typeof chartConfig) =>
                    chartConfig[value]?.label
                  }
                />
              </Pie>
            </PieChart>
          </ChartContainer>
        </CardContent>
      </Card>
    </div>
  );
}
