"use client"

import { useState } from "react"
import { TrendingUp } from "lucide-react"
import { Bar, BarChart, CartesianGrid, XAxis, Cell } from "recharts"

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
  { month: "January", desktop: 186},
  { month: "February", desktop: 305},
  { month: "March", desktop: 237},
  { month: "April", desktop: 73},
  { month: "May", desktop: 209},
  { month: "June", desktop: 214},
]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#EDF0F7", // Default color for bars
  },

} satisfies ChartConfig

export default function DashboardBarChart() {
  const [activeIndex, setActiveIndex] = useState<number | null>(null)

  const handleMouseOver = (index: number) => {
    setActiveIndex(index)
  }

  const handleMouseOut = () => {
    setActiveIndex(null)
  }

  return (
    <Card className=" w-full lg:w-1/3 mr-2">
      <CardHeader>
        <CardTitle>My Expenses</CardTitle>
        <CardDescription>January - June 2024</CardDescription>
      </CardHeader>

      <CardContent>
          <ChartContainer config={chartConfig} className="max-h-48 w-full px-1">
            <BarChart accessibilityLayer data={chartData}>
              <CartesianGrid vertical={false} />
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
              <Bar dataKey="desktop" radius={8}>
                {chartData.map((entry, index) => (
                  <Cell
                    key={`cell-${index}`}
                    fill={
                      index === activeIndex
                        ? "#16DBCC" // Hover color
                        : chartConfig.desktop.color // Default color
                    }
                    onMouseOver={() => handleMouseOver(index)}
                    onMouseOut={handleMouseOut}
                  />
                ))}
              </Bar>
            </BarChart>
          </ChartContainer>
      </CardContent>
      <CardFooter className="flex-col items-start gap-2 text-sm">
        {/* <div className="flex gap-2 font-medium leading-none">
          Trending up by 5.2% this month <TrendingUp className="h-4 w-4" />
        </div>
        <div className="leading-none text-muted-foreground">
          Showing total visitors for the last 6 months
        </div> */}
      </CardFooter>
    </Card>
  )
}
