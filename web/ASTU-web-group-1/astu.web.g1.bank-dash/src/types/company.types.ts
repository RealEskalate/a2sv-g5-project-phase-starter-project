export interface CompanyDataType {
  id: string;
  companyName: string;
  type: string;
  icon: string;
}
export interface CompanyResponseType {
  success: boolean;
  message: string;
  data: CompanyDataType[];
}
export interface singleCompanyResponseType {
  success: boolean;
  message: string;
  data: CompanyDataType;
}
