import axios from "axios";
import Company from "../../types/companyInterface";
import { getServerSession } from "next-auth";

// Extend the user type to include accessToken
interface ExtendedUser {
  refresh_token: string;
  data: any; // Assuming `data` contains user information or other details
  accessToken?: string;
}

interface ExtendedSession {
  user?: ExtendedUser;
}

const fetchSession = async (): Promise<ExtendedSession> => {
  const session = await getServerSession();
  return session as ExtendedSession;
};

const getAccessToken = async (): Promise<string | undefined> => {
  const session = await fetchSession();
  return session?.user?.accessToken;
};

const baseUrl = "https://bank-dashboard-6acc.onrender.com";

export async function getCompanyById(id: string) {
  try {
    const accessToken = await getAccessToken();
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

export async function createCompany(company: Company) {
  try {
    const accessToken = await getAccessToken();
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

export async function updateCompany(id: string, company: Company) {
  try {
    const accessToken = await getAccessToken();
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

export async function deleteCompany(id: string) {
  try {
    const accessToken = await getAccessToken();
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

export async function getCompanies(page: number, size: number) {
  try {
    const accessToken = await getAccessToken();
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

export async function getTrendingCompanies() {
  try {
    const accessToken = await getAccessToken();
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
