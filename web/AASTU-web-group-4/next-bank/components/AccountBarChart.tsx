"use client"
import { Bar, BarChart, YAxis, Legend, XAxis } from "recharts"
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
  { day: "Sat", debit: 50, credit: 20 },
  { day: "Sun", debit: 30, credit: 10 },
  { day: "Mon", debit: 45, credit: 50 },
  { day: "Tue", debit: 33, credit: 15 },
  { day: "Wed", debit: 49, credit: 11 },
  { day: "Thu", debit: 33, credit: 44 },
  { day: "Fri", debit: 10, credit: 22 },
]

const chartConfig = {
  debit: {
    label: "Debit",
    color: "#1814F3",
  },
  credit: {
    label: "Credit",
    color: "#FC7900",
  },
} satisfies ChartConfig

export default function Component() {
  return (
    <Card className="flex flex-col w-full h-[350px]">
      <CardHeader>
        <CardTitle>Weekly Activity</CardTitle>
      </CardHeader>
      <CardContent className="flex-1">
        <ChartContainer config={chartConfig} className="p-0">
          <BarChart
            data={chartData}
            barCategoryGap="10%"
            barGap={5}
            barSize={30}
            // width="100%"        
            height={240}       // Set height to 100% of the container
            margin={{ top: 20, right: 20, left: 20, bottom: 20 }}  
          >
            <XAxis dataKey="day" axisLine={true} tickLine={false} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Legend verticalAlign="top" align="right" />
            <Bar dataKey="debit" fill={chartConfig.debit.color} radius={[5, 5, 0, 0]} />
            <Bar dataKey="credit" fill={chartConfig.credit.color} radius={[5, 5, 0, 0]} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}