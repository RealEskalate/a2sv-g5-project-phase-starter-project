import axios from "axios";
const baseUrl = "https://a2svwallet.onrender.com";
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
