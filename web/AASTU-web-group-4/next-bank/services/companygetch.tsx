import Cookies from "js-cookie";

const API_BASE_URL = "https://next-bank.onrender.com";
const token = Cookies.get("accessToken");
// GET /companies/{id}
export const getCompanyById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.json();
};

// PUT /companies/{id}
export const updateCompanyById = async (id: any, updateData: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(updateData),
  });
  return response.json();
};

// DELETE /companies/{id}
export const deleteCompanyById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.json();
};

// GET /companies
export const getAllCompanies = async () => {
  const response = await fetch(`${API_BASE_URL}/companies`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.json();
};

// POST /companies
export const createCompany = async (companyData: any) => {
  const response = await fetch(`${API_BASE_URL}/companies`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(companyData),
  });
  return response.json();
};

// GET /companies/trending-companies
export const getTrendingCompanies = async () => {
  const response = await fetch(`${API_BASE_URL}/companies/trending-companies`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.json();
};
