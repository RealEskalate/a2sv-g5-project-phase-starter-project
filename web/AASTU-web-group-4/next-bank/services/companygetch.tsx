const API_BASE_URL = 'https://your-api-domain.com';

// GET /companies/{id}
export const getCompanyById = async (id:any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: 'GET',
  });
  return response.json();
};

// PUT /companies/{id}
export const updateCompanyById = async (id: any, updateData: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updateData),
  });
  return response.json();
};

// DELETE /companies/{id}
export const deleteCompanyById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/companies/${id}`, {
    method: 'DELETE',
  });
  return response.json();
};

// GET /companies
export const getAllCompanies = async () => {
  const response = await fetch(`${API_BASE_URL}/companies`, {
    method: 'GET',
  });
  return response.json();
};

// POST /companies
export const createCompany = async (companyData: any) => {
  const response = await fetch(`${API_BASE_URL}/companies`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(companyData),
  });
  return response.json();
};

// GET /companies/trending-companies
export const getTrendingCompanies = async () => {
  const response = await fetch(`${API_BASE_URL}/companies/trending-companies`, {
    method: 'GET',
  });
  return response.json();
};


