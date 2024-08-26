import { getSession } from "next-auth/react";

export async function getLoansAll(page: number, size: number) {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/active-loans/all?page=${page}&size=${size}`,
      {
        method: "GET",
        cache: "reload",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    ).then((res) => res.json());

    if (response.success) {
      return response.data;
    } else {
      throw new Error("failed to get data");
    }
  } catch (error) {
    console.error("An error occurred on table:", error);
  }
}

export async function getDetailData() {
  try {
    const session = await getSession();
    const accessToken = session?.user.accessToken;
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BASE_URL}/active-loans/detail-data`,
      {
        method: "GET",
        cache: "force-cache",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    ).then((res) => res.json());

    if (response.success) {
      console.log(response.data);
      return response.data;
    } else {
      throw new Error(`failed to fetch the data: ${response}`);
    }
  } catch (error) {
    console.error(error);
  }
}
