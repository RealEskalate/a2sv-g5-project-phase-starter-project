const API_BASE_URL = "https://bank-dashboard-6acc.onrender.com";
const token = 
  "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJheXV1IiwiaWF0IjoxNzI0MTQ5MzgyLCJleHAiOjE3MjQyMzU3ODJ9.ho0P9ZYtpOiDLT810v9r_YAMUwb865p4O4iXIWu0H5ujqjdxbLI_K6lH4m_YOxPm";
// GET /bank-services/{id}
export const getBankServiceById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/bank-services/${id}`, {
    method: "GET",
  });
  return response.json();
};
// PUT /bank-services/{id}
export const updateBankServiceById = async (
  id: string,
  updateData: {
    id: string;
    updateData: any;
  }
) => {
  const response = await fetch(`${API_BASE_URL}/bank-services/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(updateData),
  });
  return response.json();
};

// DELETE /bank-services/{id}
export const deleteBankServiceById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/bank-services/${id}`, {
    method: "DELETE",
  });
  return response.json();
};

// GET /bank-services
export const getAllBankServices = async () => {
  const response = await fetch(`${API_BASE_URL}/bank-services?page=${0}&size=${5}`, {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
     
    },

  });
  return response.json();
};

// POST /bank-services
export const createBankService = async (serviceData: any) => {
  const response = await fetch(`${API_BASE_URL}/bank-services`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(serviceData),
  });
  return response.json();
};

// GET /bank-services/search
export const searchBankServices = async (queryParams: any) => {
  const query = new URLSearchParams(queryParams).toString();
  const response = await fetch(
    `${API_BASE_URL}/bank-services/search?${query}`,
    {
      method: "GET",
    }
  );
  return response.json();
};
