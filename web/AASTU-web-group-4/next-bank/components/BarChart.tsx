"use client"
import { Bar, BarChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts"
import {
  Card,
  CardContent,
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
  { month: "January", desktop: 186, mobile: 80 },
  { month: "February", desktop: 305, mobile: 200 },
  { month: "March", desktop: 237, mobile: 120 },
  { month: "April", desktop: 73, mobile: 190 },
  { month: "May", desktop: 209, mobile: 130 },
  { month: "June", desktop: 214, mobile: 140 },
]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#1814F3",
  },
  mobile: {
    label: "Mobile",
    color: "#16DBCC",
  },
} satisfies ChartConfig

export default function Component() {
  return (
    <Card className="w-[100%] flex items-end justify-start ">
      <CardContent className=" p-0 w-[100%] flex items-end justify-start " >
        <ChartContainer className="w-[100%] " config={chartConfig}>
              
              <BarChart className="w-[100%]" accessibilityLayer barSize={10} data={chartData} >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={true}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis axisLine={true} tickLine={false} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Legend verticalAlign="top" align="right"  />

            <Bar dataKey="desktop" fill={chartConfig.desktop.color} radius={[10,10,10,10]} />
            <Bar dataKey="mobile" fill={chartConfig.mobile.color} radius={[10,10,10,10]} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
