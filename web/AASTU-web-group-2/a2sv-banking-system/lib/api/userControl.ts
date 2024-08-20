import axios from "axios";
import User, { Preference } from "../../types/userInterface";
import { getServerSession } from "next-auth";

// Extend the user type to include accessToken
interface ExtendedUser {
  refresh_token: string;
  data: any; // Assuming `data` contains user information or other details
  accessToken?: string;
}

interface ExtendedSession {
  user?: ExtendedUser;
}

const fetchSession = async (): Promise<ExtendedSession> => {
  const session = await getServerSession();
  return session as ExtendedSession;
};

const getAccessToken = async (): Promise<string | undefined> => {
  const session = await fetchSession();
  return session?.user?.accessToken;
};

const baseUrl = "https://bank-dashboard-6acc.onrender.com";

export async function userUpdate(user: User) {
  try {
    const accessToken = await getAccessToken();
    const response = await axios.put(baseUrl + "/user/update", user, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error updating user:", error);
    throw error;
  }
}

export async function userUpdatePreference(preference: Preference) {
  try {
    const accessToken = await getAccessToken();
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
    return response.data;
  } catch (error) {
    console.error("Error updating user preference:", error);
    throw error;
  }
}

export async function getUserByUsername(username: string) {
  try {
    const accessToken = await getAccessToken();
    const response = await axios.get(baseUrl + `/user/${username}`, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching user by username:", error);
    throw error;
  }
}

export async function getRandomInvestementData(months: number, years: number) {
  try {
    const accessToken = await getAccessToken();
    const response = await axios.get(
      baseUrl + `/user/random-investment-data?months=${months}&years=${years}`,
      {
        headers: {
          Authorization: `Bearer ${accessToken}`,
          "Content-Type": "application/json",
        },
      }
    );
    return response.data;
  } catch (error) {
    console.error("Error fetching random investment data:", error);
    throw error;
  }
}

export async function getCurrentUser() {
  try {
    const accessToken = await getAccessToken();
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
