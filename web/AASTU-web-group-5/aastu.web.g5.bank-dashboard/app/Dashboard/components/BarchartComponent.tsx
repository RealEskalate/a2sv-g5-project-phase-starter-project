"use client"

import { Bar, BarChart, CartesianGrid, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts"

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart"

// Dummy data for the bar chart
const chartData = [
  { day: "Sat", deposit: 500, withdraw: 250 },
  { day: "Sun", deposit: 305, withdraw: 200 },
  { day: "Mon", deposit: 237, withdraw: 120 },
  { day: "Tue", deposit: 73, withdraw: 190 },
  { day: "Wed", deposit: 209, withdraw: 130 },
  { day: "Thu", deposit: 214, withdraw: 140 },
  { day: "Fri", deposit: 214, withdraw: 140 },
]

const chartConfig = {
  deposit: {
    label: "Deposit",
    color: "hsl(var(--chart-1))",
  },
  withdraw: {
    label: "Withdraw",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig

export function BarchartComponent() {
  return (
    <Card className="max-w-full overflow-hidden">
      <CardContent className="p-4 relative">
        <ChartContainer config={chartConfig}>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={chartData}>
              <CartesianGrid vertical={false} />
              <XAxis
                dataKey="day"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tickFormatter={(value) => value.slice(0, 3)}
              />
              <YAxis
                tickCount={6}
                tickSize={5}
                tickFormatter={(value) => `${value}`}
                domain={[0, 'dataMax']}
                interval="preserveStartEnd"
                orientation="left"
              />
              <Tooltip />
              <Bar
                dataKey="deposit"
                fill="blue"
                radius={[10, 10, 10, 10]} // Rounded corners (top-left, top-right, bottom-right, bottom-left)
                barSize={10} // Width of the bars
                name="Deposit"
              />
              <Bar
                dataKey="withdraw"
                fill="green"
                radius={[10, 10, 10, 10]} // Rounded corners (top-left, top-right, bottom-right, bottom-left)
                barSize={10} // Width of the bars
                name="Withdraw"
              />
            </BarChart>
          </ResponsiveContainer>
        </ChartContainer>
        <div className="absolute top-4 right-4 p-2">
  <div className="flex flex-row items-center space-x-4">
    <div className="flex items-center">
      <div className="w-4 h-4 rounded-full mr-2" style={{ backgroundColor: "blue" }}></div>
      <span className="text-sm">Deposit</span>
    </div>
    <div className="flex items-center">
      <div className="w-4 h-4 mr-2 rounded-full" style={{ backgroundColor: "green" }}></div>
      <span className="text-sm">Withdraw</span>
    </div>
  </div>
</div>

      </CardContent>
    </Card>
  )
}
