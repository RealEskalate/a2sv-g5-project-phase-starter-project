import Cookies from "js-cookie";

<<<<<<< HEAD
const API_BASE_URL = "https://web-team-g4.onrender.com/";
const token = Cookies.get("accessToken");
=======
const API_BASE_URL = 'https://web-team-g4.onrender.com';
const token = Cookies.get('accessToken')
>>>>>>> 4c18eaee10a18bbefdc03fa57c6e957d354b34ad

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
<<<<<<< HEAD
    const response = await fetch(`${API_BASE_URL}/active-loans/my-loans`, {
=======
    const response = await fetch(`${API_BASE_URL}/active-loans/my-loans?page=0&size=5`, {
>>>>>>> 4c18eaee10a18bbefdc03fa57c6e957d354b34ad
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
<<<<<<< HEAD
=======

>>>>>>> 4c18eaee10a18bbefdc03fa57c6e957d354b34ad
  });
  return response.json();
};
