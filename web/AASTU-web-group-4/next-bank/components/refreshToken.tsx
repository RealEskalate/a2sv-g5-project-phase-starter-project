import Cookies from "js-cookie";

const ACCESS_TOKEN_KEY = "accessToken";
const REFRESH_TOKEN_KEY = "refreshToken";

// Function to get the refresh token from cookies
export const getRefreshToken = () => {
  const token = localStorage.getItem(REFRESH_TOKEN_KEY);
  console.log("Current Refresh Token:", token);
  return token;
};

// Function to refresh the access token using the refresh token
export const refreshAccessToken = async () => {
  const refreshToken = getRefreshToken();
  
  if (!refreshToken) return;

  try {
    console.log("Refreshing access token...",refreshToken);
    const response = await fetch(" https://web-team-g4.onrender.com/auth/refresh_token", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
         Authorization: `Bearer ${refreshToken}`
      },
    });
     

    if (!response.ok) {
      throw new Error("Failed to refresh access token");
    }

    const data = await response.json();
    // Store the new access token in cookies
    Cookies.set(ACCESS_TOKEN_KEY, data.accessToken, { expires: 2 / 1440 });
    // Cookie expires in 15 minutes
    console.log("New Access Token:", data);
  } catch (error) {
    console.error("Error refreshing access token:", error);
  }
};

// Function to initialize the auto-refresh of the access token
export const initializeTokenRefresh = () => {
  console.log("Initializing token refresh every 30 seconds.");
  setInterval(() => {
    console.log("Attempting to refresh token...");
    refreshAccessToken(); // Call the refresh function
  }, 30 * 1000); // 30 seconds
};

