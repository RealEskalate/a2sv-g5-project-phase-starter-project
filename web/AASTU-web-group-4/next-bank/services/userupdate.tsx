import Cookie from "js-cookie";

const API_BASE_URL = "https://web-team-g4.onrender.com";
const token = Cookie.get("accessToken");

// Update User Details - PUT Request
export const updateUserDetails = async (userData: any) => {
  try {
    const response = await fetch(
      "https://web-team-g4.onrender.com/user/update",
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

// Update User Preferences - PUT Request
export const updatePreference = async (userData: any) => {
  try {
    const response = await fetch(
      "https://web-team-g4.onrender.com/user/update-preference",
      {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to update user preferences");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Fetch User Details - GET Request
export const fetchUserDetails = async (username: string) => {
  try {
    const response = await fetch(
      `https://web-team-g4.onrender.com/user/{username}`,
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

// Fetch Random Investment Data - GET Request
export const randomInvestmentData = async () => {
  try {
    const response = await fetch(
      `https://web-team-g4.onrender.com/user/random-investment-data?years=5&months=8`,
      {
        method: "GET",
        headers: {
          authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch investment data");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Fetch Current User - GET Request
export const currentuser = async () => {
  try {
    const response = await fetch(
      `https://web-team-g4.onrender.com/user/current`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      throw new Error("Failed to fetch current user details");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};
