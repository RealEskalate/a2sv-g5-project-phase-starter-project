"use client"
import { Area, AreaChart, CartesianGrid, Legend, XAxis, YAxis,  } from "recharts"
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
import { TbBackground } from "react-icons/tb"

const chartData = [
  { month: "January", balance: 186 },
  { month: "February", balance: 305 },
  { month: "March", balance: 237 },
  { month: "April", balance: 73 },
  { month: "May", balance: 209 },
  { month: "June", balance: 214 },
]

const chartConfig = {
  balance: {
    color: "#1814F3",
  },
} satisfies ChartConfig

export default function LineChart() {
  return (
    <Card className=" w-[100%] rounded-2xl">
      <CardContent>
        <ChartContainer 
          config={chartConfig}
          className="w-[100%]"
        >
          <AreaChart
            accessibilityLayer
            data={chartData}
            
          >
            <defs>
              <linearGradient id="colorBalance" x1="0" y1="0" x2="0" y2="1">
                <stop offset="0%" stopColor={chartConfig.balance.color} stopOpacity={0.2} />
                <stop offset="100%" stopColor={chartConfig.balance.color} stopOpacity={0.05} />
              </linearGradient>
            </defs>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={true}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis dataKey="balance" axisLine={true} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="line" />}
            />
            <Area
              dataKey="balance"
              type="natural"
              fill="url(#colorBalance)"
              stroke={chartConfig.balance.color}
              fillOpacity={1}
              strokeWidth={3}
            />
            <Legend verticalAlign="top"  height={36} />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
