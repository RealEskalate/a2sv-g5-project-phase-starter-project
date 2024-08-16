"use client"

import { TrendingUp } from "lucide-react"
import { Area, AreaChart, CartesianGrid, XAxis, YAxis , Line } from "recharts"

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
  { Year: "2016", investment: 5000 },
  { Year: "2017", investment: 25000 },
  { Year: "2018", investment: 18000 },
  { Year: "2019", investment: 38000 },
  { Year: "2020", investment: 20000 },
  { Year: "2021", investment: 30000 },
]

const chartConfig = {
  investment: {
    label: "investment",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig

export function YearlyInvest() {
  return (
    <Card className="rounded-3xl py-5 shadow-lg border-gray-300">
      <CardContent>
        <ChartContainer config={chartConfig}>
          <AreaChart
            accessibilityLayer
            data={chartData}
            margin={{
              left: 12,
              right: 12,
              
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="Year"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              padding={{left:20}}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              ticks={[0, 10000, 20000, 30000, 40000]}
              tickFormatter={(value) => value.toLocaleString()}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" hideLabel />}
            />
            <Area
              dataKey="investment"
              type="linear"
              fill="white"
              fillOpacity={0.4}
              stroke="#EDA10D"
              strokeWidth={3}
              dot={{
                fill: "white",
                r:6
              }}
              activeDot={{
                r: 6,
              }}
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
      
    </Card>
  )
}
