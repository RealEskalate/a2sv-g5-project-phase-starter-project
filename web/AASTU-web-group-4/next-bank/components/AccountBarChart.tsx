import { Bar, BarChart, XAxis } from "recharts";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { day: "Sat", debit: 50, credit: 20 },
  { day: "Sun", debit: 30, credit: 10 },
  { day: "Mon", debit: 45, credit: 50 },
  { day: "Tue", debit: 33, credit: 15 },
  { day: "Wed", debit: 49, credit: 11 },
  { day: "Thu", debit: 33, credit: 44 },
  { day: "Fri", debit: 10, credit: 22 },
];

// Calculate total debit and credit
const totalDebit = chartData.reduce((sum, data) => sum + data.debit, 0);
const totalCredit = chartData.reduce((sum, data) => sum + data.credit, 0);

const chartConfig = {
  debit: {
    label: "Debit",
    color: "#1814F3",
  },
  credit: {
    label: "Credit",
    color: "#FC7900",
  },
} satisfies ChartConfig;

export default function Component() {
  return (
    <Card className="flex flex-col w-full h-[350px]">
      <CardHeader className="flex justify-between">
        <div className="flex flex-row justify-between space-x-4">
          <div className="hidden md:flex text-sm font-normal">
            <span className="font-bold">${totalDebit}</span> Debited &{" "}
            <span className="font-bold"> ${totalCredit}</span> Credited in this
            Week
          </div>
          <div className="flex px-3 text-right">
            <span
              className="w-4 h-4 rounded-xl"
              style={{ backgroundColor: chartConfig.debit.color }}
            ></span>
            <span className="text-sm font-normal pl-2 pr-4">Debit</span>

            <span
              className="w-4 h-4 rounded-xl"
              style={{ backgroundColor: chartConfig.credit.color }}
            ></span>
            <span className="text-sm font-normal pl-2">Credit</span>
          </div>
        </div>
      </CardHeader>
      <CardContent className="flex-1">
        <ChartContainer config={chartConfig} className="p-0">
          <BarChart
            data={chartData}
            barCategoryGap="10%"
            barGap={5}
            barSize={20} // Adjusted bar size
            height={200} // Adjusted height
            margin={{ top: 10, right: 20, left: 20, bottom: 90 }} // Reduced top margin
          >
            <XAxis dataKey="day" axisLine={true} tickLine={false} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar
              dataKey="debit"
              fill={chartConfig.debit.color}
              radius={[5, 5, 0, 0]}
            />
            <Bar
              dataKey="credit"
              fill={chartConfig.credit.color}
              radius={[5, 5, 0, 0]}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
