import ky from "ky";
import { useSession } from "next-auth/react";

interface Credentials {
  userName: string;
  password: string;
}

interface SignInResponse {
  success: boolean;
  message: string;
  data: {
    id: string;
    name: string;
    email: string;
    access_token: string;
    refresh_token: string;
  };
}

export async function signInWithCredentials(
  credentials: Credentials
): Promise<SignInResponse> {
  return await ky
    .post(`${process.env.NEXT_PUBLIC_BASE_URL}/auth/login`, {
      json: credentials,
    })
    .json<SignInResponse>();
}

// export async function changePassword(credentials: any){
//     console.log(accessToken);
//     const formData = {
//       password: credentials.currentPassword,
//       newPassword: credentials.newPassword,
//     };

//     console.log(accessToken);

//     if (!accessToken) {
//       throw new Error("No access token found");
//     }

//     const res = await ky
//       .post("https://bank-dashboard-o9tl.onrender.com/auth/change_password", {
//         json: formData,
//         headers: {
//           Authorization: `Bearer ${accessToken}`,
//         },
//       })
//       .json();

//     return res;
// }

interface returnData {
  access_token: string;
  refresh_token: string;
  data: null;
}

export async function refreshAccessToken(
  refresh_token: string
): Promise<returnData | null> {
  console.log("Refresh Token Entering", refresh_token);
  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/auth/refresh_token`,
      {
        method: "POST",
        headers: {
          Authorization: `Bearer ${refresh_token}`,
        },
      }
    );

    if (response.status === 200) {
      const data = await response.json();
      console.log("Fetched Data from auth.ts", data);

      // Ensure that `data.data` contains both `access_token` and `refresh_token`
      if (data.data && data.data.access_token && data.data.refresh_token) {
        return data.data; // Ensure the returnData structure is as expected
      } else {
        throw new Error("Missing access or refresh token in response data");
      }
    } else {
      throw new Error(
        `Failed to refresh token. Status code: ${response.status}`
      );
    }
  } catch (error) {
    console.error("Error refreshing token:", error);
    return null;
  }
}
