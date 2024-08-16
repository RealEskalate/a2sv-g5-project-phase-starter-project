"use client"
import { TrendingUp } from "lucide-react"
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts"
import { 
  Card, 
  CardContent, 
  CardDescription, 
  CardFooter, 
  CardHeader, 
  CardTitle 
} from "@/components/ui/card"
import { 
  ChartConfig, 
  ChartContainer, 
  ChartTooltip, 
  ChartTooltipContent 
} from "@/components/ui/chart"

const chartData = [
  { month: "July", desktop: 95 },
  { month: "August", desktop: 215 },
  { month: "September", desktop: 415 },
  { month: "October", desktop: 730 },
  { month: "November", desktop: 200 },
  { month: "December", desktop: 530 },
  { month: "January", desktop: 230  },
]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig

export function BalanceHistory() {
  return (
    <Card className="my-4 mx-4 rounded-3xl">
      <CardHeader>
        <CardTitle className="text-[#343C6A] font-bold  text-xl">Balance History</CardTitle>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <AreaChart
            accessibilityLayer
            data={chartData}
            margin={{
              left: 0,  // Removed padding
              right: 0, // Removed padding
              top: 10,  // Added some margin at the top
              bottom: 0, // Added some margin at the bottom
            }}
          >
            <CartesianGrid 
              strokeDasharray="3 3"  // Dotted lines
              stroke="rgba(0, 0, 0, 0.5)"  // More visible lines
              vertical={true}   // Enable vertical lines
              horizontal={true} // Enable horizontal lines
            />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              interval={0}
              ticks={[0, 200, 400, 600, 800]}
              domain={[0, 800]}  // Ensure the Y axis ends at 800
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="line" />}
            />
            <Area
              dataKey="desktop"
              type="natural"
              fill="rgba(0, 0, 255, 0.2)"
              stroke="blue"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
      <CardFooter>
        {/* <div className="flex w-full items-start gap-2 text-sm">
          <div className="grid gap-2">
            <div className="flex items-center gap-2 font-medium leading-none">
              Trending up by 5.2% this month <TrendingUp className="h-4 w-4" />
            </div>
            <div className="flex items-center gap-2 leading-none text-muted-foreground">
              January - June 2024
            </div>
          </div>
        </div> */}
      </CardFooter>
    </Card>
  )
}
