"use client"
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts"
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
    <Card >
      <CardHeader>
        
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} 
            style={{height:'300px' , width:'600px'}}
        >
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
              dataKey="month"
              tickLine={false}
              axisLine={true}
              tickMargin={8}
              // style={{height:'50px' , width:'200px'}}
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
              fill={chartConfig.balance.color}
              fillOpacity={0.4}
              stroke={chartConfig.balance.color}
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
