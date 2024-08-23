import { getSession } from "next-auth/react";

export  async function getLoansAll() {
    try {
        const session = await getSession();
        const accessToken = session?.user.accessToken;
        const response = await fetch(`${process.env.BASE_URL}/active-loans/all`, {
            method: "GET",
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
            body: null,
        }).then((res)=>(res.json()));

        if (response.success){
            // console.log(response.data);
            return response.data;
        }else{
            throw new Error("failed to get data");
        }
    } catch (error) {
      console.error("An error occurred on card:", error);
    }
}