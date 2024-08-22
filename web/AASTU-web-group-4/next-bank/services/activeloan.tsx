import Cookies from "js-cookie";

const API_BASE_URL = "https://bank-dashboard-o9tl.onrender.com";
const token = Cookies.get("accessToken");

// POST /active-loans
export const createActiveLoan = async (loanData: any) => {
  const response = await fetch(`${API_BASE_URL}/active-loans`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(loanData),
  });
  return response.json();
};

// POST /active-loans/{id}/reject
export const rejectActiveLoan = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/active-loans/${id}/reject`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response.json();
};

// POST /active-loans/{id}/approve
export const approveActiveLoan = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/active-loans/${id}/approve`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
  });
  return response.json();
};

// GET /active-loans/{id}
export const getActiveLoanById = async (id: any) => {
  const response = await fetch(`${API_BASE_URL}/active-loans/${id}`, {
    method: "GET",
  });
  return response.json();
};

// GET /active-loans/my-loans
export const getMyLoans = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/active-loans/my-loans`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (!response.ok) {
      console.log(response);
      throw new Error("Failed to fetch ");
    }
    return response.json();
  } catch (error) {
    console.error("Error: ", error);
  }
};

// GET /active-loans/detail-data
export const getLoanDetailData = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/active-loans/detail-data`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (!response.ok) {
      console.log(response);
      throw new Error("Failed to fetch");
    }
    return response.json();
  } catch (error) {
    console.error("Error: ", error);
  }
};

// GET /active-loans/all
export const getAllActiveLoans = async () => {
  const response = await fetch(`${API_BASE_URL}/active-loans/all`, {
    method: "GET",
  });
  return response.json();
};
