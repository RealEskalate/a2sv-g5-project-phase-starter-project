"use client"

import { TrendingUp } from "lucide-react"
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts"

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

const chartData = [
  { month: "Saturday", desktop: 186, mobile: 80 },
  { month: "Sunday", desktop: 305, mobile: 200 },
  { month: "Monday", desktop: 237, mobile: 120 },
  { month: "Tuesday", desktop: 73, mobile: 190 },
  { month: "Wednesday", desktop: 209, mobile: 130 },
  { month: "Thursday", desktop: 214, mobile: 140 },
  { month: "Friday", desktop: 214, mobile: 140 },
]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#1814f3", // Color for desktop bars
  },
  mobile: {
    label: "Mobile",
    color: "#16dbcc", // Color for mobile bars
  },
} satisfies ChartConfig

export function WeeklyActivity() {
  return (
    <Card className="my-4 mx-4 rounded-3xl">
      <CardHeader>
        <CardTitle className="text-[#343C6A] font-bold text-xl">Weekly Activity</CardTitle>
        {/* <CardDescription>January - June 2024</CardDescription> */}
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart data={chartData}>
          <CartesianGrid 
                vertical={false} 
                strokeDasharray="none" // Remove dashed lines for solid lines
                stroke="#E0E0E0" // Lighter grey color for the grid lines
                strokeWidth={0.5} // Thinner lines for a lighter appearance
                />
            <YAxis
              tickCount={6}
              tickFormatter={(value) => value}
              domain={[0, 500]}
              interval={0}
              tickLine={false}
              axisLine={false}
              tickMargin={10}
            />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar 
              dataKey="desktop" 
              fill="#1814f3" // Updated color for desktop bars
              radius={4} 
            />
            <Bar 
              dataKey="mobile" 
              fill="#16dbcc" // Updated color for mobile bars
              radius={4} 
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
      {/* <CardFooter className="flex-col items-start gap-2 text-sm">
        <div className="flex gap-2 font-medium leading-none">
          Trending up by 5.2% this month <TrendingUp className="h-4 w-4" />
        </div>
        <div className="leading-none text-muted-foreground">
          Showing total visitors for the last 6 months
        </div>
      </CardFooter> */}
    </Card>
  )
}
