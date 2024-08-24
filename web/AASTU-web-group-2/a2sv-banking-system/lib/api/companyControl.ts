import axios from "axios";
import Company from "../../types/companyInterface";

const baseUrl = "https://bank-dashboard-mih0.onrender.com";

export async function getCompanyById(id: string, accessToken: string) {
  try {
    const response = await axios.get(baseUrl + `/companies/${id}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching company by ID:", error);
    throw error;
  }
}

export async function createCompany(company: Company, accessToken: string) {
  try {
    const response = await axios.post(baseUrl + `/companies`, company, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error creating company:", error);
    throw error;
  }
}

export async function updateCompany(
  id: string,
  company: Company,
  accessToken: string
) {
  try {
    const response = await axios.put(baseUrl + `/companies/${id}`, company, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error updating company:", error);
    throw error;
  }
}

export async function deleteCompany(id: string, accessToken: string) {
  try {
    const response = await axios.delete(baseUrl + `/companies/${id}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error deleting company:", error);
    throw error;
  }
}

export async function getCompanies(
  page: number,
  size: number,
  accessToken: string
) {
  try {
    const response = await axios.get(
      baseUrl + `/companies/?page=${page}&size=${size}`,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          "Content-Type": "application/json",
        },
      }
    );
    return response.data;
  } catch (error) {
    console.error("Error fetching companies:", error);
    throw error;
  }
}

export async function getTrendingCompanies(accessToken: string) {
  try {
    const response = await axios.get(
      baseUrl + `/companies/trending-companies`,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          "Content-Type": "application/json",
        },
      }
    );
    return response.data;
  } catch (error) {
    console.error("Error fetching trending companies:", error);
    throw error;
  }
}
