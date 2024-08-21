"use client"
import { Pie, PieChart } from "recharts"
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
import { useEffect, useState } from "react"

const chartData = [
  { category: "Entertainment", amount: 375, fill: "#3C4264 " },  // Dim Blue
  { category: "Shopping", amount: 200, fill: "#FF8900 " },       // Dim Green
  { category: "Groceries", amount: 387, fill: "#FF00FF " },       // Dim Orange
  { category: "Bills", amount: 173, fill: "#001BFF " },             // Dim Red
]



export default function Component() {
  const [pierad , setpierad] = useState(130)
  useEffect(()=>{

    const fun= () =>{
      if(window.innerWidth < 1024){
        setpierad(20)
        alert('less than 1024')
    
      }
      else{
        setpierad(100)
      }
    fun()
    window.addEventListener('resize' , fun)
    return()=>{
      window.removeEventListener('resize' , fun)
    }
    }
    },[])
  return (
    <Card className=" w-[100%] md:py-14 ">

      <CardContent className=" w-[100%] ">
        <ChartContainer
          className="mx-auto my- max-h-[90%] "
          config={{}}
        >
          <PieChart>
            <ChartTooltip
              content={<ChartTooltipContent nameKey="category" />}
            />
            <Pie
              data={chartData}
              dataKey="amount"
              nameKey="category"
              cx="50%"
              cy="50%"
              outerRadius={80}
              style={{stroke: "none"}}
              // className="md:"
              labelLine={false}
              label={({ cx, cy, midAngle, innerRadius, outerRadius, percent, index }) => {
                const RADIAN = Math.PI / 180;
                const radius = innerRadius + (outerRadius - innerRadius) * 0.5;
                const x = cx + radius * Math.cos(-midAngle * RADIAN);
                const y = cy + radius * Math.sin(-midAngle * RADIAN);
                return (
                  <text
                    x={x}
                    y={y}
                    fill="white"
                    textAnchor="middle"
                    dominantBaseline="central"
                    fontSize={12}
                  >
                    {`${(percent * 100).toFixed(0)}%`}
                  </text>
                );
              }}
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
