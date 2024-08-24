// Register User - POST Request
export const registerUser = async (userData: any) => {
  try {
    const response = await fetch('https://web-team-g4.onrender.com/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      const responseData = await response.json();
      console.error("Error registering user:", responseData);
      throw new Error('Failed to register user');
    }

    const data = await response.json();
    console.log("User registered successfully:", data);
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Refresh Token - POST Request
export const refreshToken = async (tokenData: any) => {
  try {
    const response = await fetch('https://web-team-g4.onrender.com/auth/refresh_token', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(tokenData),
    });

    if (!response.ok) {
      const responseData = await response.json();
      console.error("Error refreshing token:", responseData.message);
      throw new Error('Failed to refresh token');
    }

    const data = await response.json();
    console.log("Token refreshed successfully:", data);
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Login User - POST Request
export const loginUser = async (loginData: any) => {
  try {
    const response = await fetch('https://web-team-g4.onrender.com/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(loginData),
    });

    if (!response.ok) {
      const responseData = await response.json();
      console.error("Error logging in:", responseData.message);
      throw new Error('Failed to log in');
    }

    const data = await response.json();
    console.log("User logged in successfully:", data);
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Change Password - POST Request
export const changePassword = async (passwordData: any) => {
  try {
    const response = await fetch('https://web-team-g4.onrender.com/auth/change_password', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(passwordData),
    });

    if (!response.ok) {
      const responseData = await response.json();
      console.error("Error changing password:", responseData.message);
      throw new Error('Failed to change password');
    }

    const data = await response.json();
    console.log("Password changed successfully:", data);
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};
