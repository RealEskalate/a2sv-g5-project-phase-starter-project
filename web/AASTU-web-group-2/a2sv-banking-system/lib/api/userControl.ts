import axios from "axios";
import User, { Preference } from "../../types/userInterface";
// Extend the user type to include accessToken
const baseUrl = "https://a2svwallets.onrender.com";

export async function userUpdate(user: User, accessToken: string) {
  try {
    const response = await axios.put(baseUrl + "/user/update", user, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data.data;
  } catch (error) {
    console.error("Error updating user:", error);
    throw error;
  }
}

export async function userUpdatePreference(
  preference: Preference,
  accessToken: string
) {
  try {
    const response = await axios.put(
      baseUrl + "/user/update-preference",
      preference,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          "Content-Type": "application/json",
        },
      }
    );
    return response.data.data;
  } catch (error) {
    console.error("Error updating user preference:", error);
    throw error;
  }
}

export async function getUserByUsername(username: string, accessToken: string) {
  try {
    const response = await axios.get(baseUrl + `/user/${username}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data.data;
  } catch (error) {
    console.error("Error fetching user by username:", error);
    throw error;
  }
}

export async function getRandomInvestementData(
  months: number,
  years: number,
  accessToken: string
) {
  try {
    const response = await axios.get(
      baseUrl + `/user/random-investment-data?months=${months}&years=${years}`,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          "Content-Type": "application/json",
        },
      }
    );
    return response.data.data;
  } catch (error) {
    console.error("Error fetching random investment data:", error);
    throw error;
  }
}

export async function getCurrentUser(accessToken: string) {
  try {
    const response = await axios.get(baseUrl + `/user/current`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data.data;
  } catch (error) {
    console.error("Error fetching current user:", error);
    throw error;
  }
}
