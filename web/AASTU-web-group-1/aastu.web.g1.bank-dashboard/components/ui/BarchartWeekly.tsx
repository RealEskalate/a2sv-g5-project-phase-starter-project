"use client"

import { TrendingUp } from "lucide-react"
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts"

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
  { days: "Sunday", desktop: 186, mobile: 80 },
  { days: "Monday", desktop: 305, mobile: 200 },
  { days: "Tuesday", desktop: 237, mobile: 120 },
  { days: "Wednesday", desktop: 73, mobile: 190 },
  { days: "Thursday", desktop: 209, mobile: 130 },
  { days: "Friday", desktop: 214, mobile: 140 },
  { days: "Saturday", desktop: 214, mobile: 140 },
]

const chartConfig = {
  desktop: {
    label: "Debit",
    color: "#1A16F3",
  },
  mobile: {
    label: "Credit",
    color: "#FCAA0B",
  },
} satisfies ChartConfig

export function ChartWeekly() {
  return (
    <Card>
      <CardHeader>
        <CardDescription className="hidden md:block"> $7,560 <span className="text-[#718EBF]">Debited</span>  & $5,420 <span className="text-[#718EBF]">Credited</span> in this Week </CardDescription>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="lg:w-[800px] sm:w-[600px] w-[300px] h-60 " >
          <BarChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="days"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="desktop" fill="var(--color-desktop)" radius={4} />
            <Bar dataKey="mobile" fill="var(--color-mobile)" radius={4} />
          </BarChart>
        </ChartContainer>
      </CardContent>
      
    </Card>
  )
}
