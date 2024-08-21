import Cookie from "js-cookie"

const API_BASE_URL = "https://bank-dashboard-6acc.onrender.com";
const token = Cookie.get("accessToken")
// Update User Details - PUT Request
export const updateUserDetails = async (userData: any) => {
  try {
    const response = await fetch(
      "https://bank-dashboard-6acc.onrender.com/user/update",
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to update user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

export const updatePreference = async (userData: any) => {
  try {
    const response = await fetch(
      "https://bank-dashboard-6acc.onrender.com/user/update-preference",
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to update user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};
// Example of another API function - GET Request
export const fetchUserDetails = async (userId: string) => {
  try {
    const response = await fetch(
      `https://bank-dashboard-6acc.onrender.com/user/{username}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

export const randominvestmentdata = async (userId: string) => {
  try {
    const response = await fetch(
      `https://bank-dashboard-6acc.onrender.com/user/random-investment-data`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

export const currentuser = async () => {
  try {
    const response = await fetch(
      `https://bank-dashboard-6acc.onrender.com/user/current`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      console.log("current user fetch :", response)
      throw new Error("Failed to fetch user details");
    }

    const data = await response.json();
    console.log("succesful current user response:", data)

    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};
// You can add more API functions similarly...
