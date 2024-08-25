// import { MonthlyRevenue } from '@/components/Charts/MonthlyRevenue';
export interface YearlyInvestment {
  time: string;
  value: number;
}

export interface InvestmentType {
  data: {
    totalInvestment: number;
    rateOfReturn: number;
    yearlyTotalInvestment: YearlyInvestment[];
    monthlyRevenue: YearlyInvestment[];
  };
}

export interface ApiResponse {
  success: boolean;
  message: string;
  data: InvestmentType;
}
