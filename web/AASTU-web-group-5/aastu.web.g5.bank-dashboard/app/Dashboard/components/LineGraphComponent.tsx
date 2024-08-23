"use client"

import { TrendingUp } from "lucide-react"
import { Area, AreaChart, CartesianGrid, XAxis, YAxis, ResponsiveContainer } from "recharts"

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
  { day: "01", desktop: 186 },
  { day: "05", desktop: 305 },
  { day: "10", desktop: 237 },
  { day: "15", desktop: 73 },
  { day: "20", desktop: 209 },
  { day: "25", desktop: 214 },
  { day: "30", desktop: 300 },
]

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig

export function LineGraphComponent() {
<<<<<<< HEAD
  const [chartData, setChartData] = useState<{ time: string, value: number }[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchChartData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-o9tl.onrender.com/transactions/balance-history', {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        });
        if (response.data.success) {
          setChartData(response.data.data);
        } else {
          setError('Failed to fetch data');
        }
      } catch (err) {
        console.error('Error fetching chart data:', err);
        setError('Failed to fetch data');
      } finally {
        setLoading(false);
      }
    };

    fetchChartData();
  }, []);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;

=======
>>>>>>> aastu.web.g5.yetnayet.transactions
  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <ResponsiveContainer width="100%" height={300}>
            <AreaChart
              data={chartData}
              margin={{
                top: 16,
                left: 4,
                right: 8,
                bottom: 4,
              }}
            >
              <CartesianGrid vertical={false} strokeDasharray="3 3" />
              <XAxis
                dataKey="day"
                tickLine={false}
                axisLine={false}
                tickMargin={8}
              />
              <YAxis
                domain={[0, 800]}
                tickCount={5}
                tickFormatter={(value) => `${value}`}
              />
              <ChartTooltip
                cursor={false}
                content={<ChartTooltipContent indicator="line" />}
              />
              <Area
                dataKey="desktop"
                type="natural"
                fill="blue"
                fillOpacity={0.2}
                stroke="blue"
              />
            </AreaChart>
          </ResponsiveContainer>
        </ChartContainer>
      </CardContent>
    </Card>
  )
}
