// Register User - POST Request
export const registerUser = async (userData: any) => {
  try {
    const response = await fetch(
      "https://bank-dashboard-1tst.onrender.com/auth/register",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
      }
    );

    if (!response.ok) {
    const responseData = await response.json();
    console.log("response", responseData.data);
      throw new Error("Failed to register user");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.log("Failed to register user", error);
    throw error;
  }
};

// Refresh Token - POST Request
export const refreshToken = async (tokenData: any) => {
  try {
    const response = await fetch(
      "https://bank-dashboard-1tst.onrender.com/auth/refresh_token",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(tokenData),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to refresh token");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Login User - POST Request
export const loginUser = async (loginData: any) => {
  try {
    const response = await fetch(
      "https://bank-dashboard-1tst.onrender.com/auth/login",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to login");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Change Password - POST Request
export const changePassword = async (passwordData: any) => {
  try {
    const response = await fetch(
      "https://bank-dashboard-1tst.onrender.com/auth/change_password",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(passwordData),
      }
    );

    if (!response.ok) {
      throw new Error("Failed to change password");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// You can export all functions from this file
