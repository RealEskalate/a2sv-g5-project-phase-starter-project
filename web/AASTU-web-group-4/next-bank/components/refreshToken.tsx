import Cookies from "js-cookie";

const ACCESS_TOKEN_KEY = "accessToken";
const REFRESH_TOKEN_KEY = "refreshToken";

// Function to get the refresh token from cookies
export const getRefreshToken = () => {
  return Cookies.get(REFRESH_TOKEN_KEY);
};

// Function to refresh the access token using the refresh token
export const refreshAccessToken = async () => {
  const refreshToken = getRefreshToken();
  
  if (!refreshToken) return;

  try {
    const response = await fetch("https://your-api-domain.com/refresh-token", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ refreshToken }),
    });

    if (!response.ok) {
      throw new Error("Failed to refresh access token");
    }

    const data = await response.json();
    // Store the new access token in cookies
    Cookies.set(ACCESS_TOKEN_KEY, data.accessToken, { expires: 1/96 }); // Cookie expires in 15 minutes
    console.log("New Access Token:", data.accessToken);
  } catch (error) {
    console.error("Error refreshing access token:", error);
  }
};

// Function to initialize the auto-refresh of the access token
export const initializeTokenRefresh = () => {
  // Set an interval to refresh the token every 15 minutes
  setInterval(() => {
    refreshAccessToken();
  }, 15 * 60 * 1000); // 15 minutes
};
