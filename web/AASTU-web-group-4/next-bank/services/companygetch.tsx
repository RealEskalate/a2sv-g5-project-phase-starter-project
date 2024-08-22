const API_BASE_URL = 'https://bank-dashboard-o9tl.onrender.com
';
const token = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJxd2VyIiwiaWF0IjoxNzI0MTQwNzE1LCJleHAiOjE3MjQyMjcxMTV9.90gS2PauXlM2v4Dv8LlEG2r2Dr4ZnlWS19A7cDRf-OA0SpWxwanSEDW8ddH_vn9E"
// GET /companies/{id}
export const getCompanyById = async (id:any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
  }});
  return response.json();
};

// PUT /companies/{id}
export const updateCompanyById = async (id: any, updateData: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(updateData),
  });
  return response.json();
};

// DELETE /companies/{id}
export const deleteCompanyById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });
  return response.json();
};

// GET /companies
export const getAllCompanies = async () => {
  const response = await fetch(`${API_BASE_URL}/companies`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });
  return response.json();
};

// POST /companies
export const createCompany = async (companyData: any) => {
  const response = await fetch(`${API_BASE_URL}/companies`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify(companyData),
  });
  return response.json();
};

// GET /companies/trending-companies
export const getTrendingCompanies = async () => {
  const response = await fetch(`${API_BASE_URL}/companies/trending-companies`, {
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });
  return response.json();
};


