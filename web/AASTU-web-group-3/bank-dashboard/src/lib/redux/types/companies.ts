export interface Company {
    id: string;
    companyName: string;
    type: string;
    icon: string;
  }
  
  export interface CompanyResponse {
    success: boolean;
    message: string;
    data: Company;
  }
  
  export interface CompaniesResponse {
    success: boolean;
    message: string;
    data: {
      content: Company[];
      totalPages: number;
    };
  }
  
  export interface TrendingCompaniesResponse {
    success: boolean;
    message: string;
    data: Company[];
  }
  