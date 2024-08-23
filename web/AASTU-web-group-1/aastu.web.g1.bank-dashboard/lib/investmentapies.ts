import { getSession } from "next-auth/react";

const baseUrl = `https://bank-dashboard-o9tl.onrender.com`

export  async function getInvestmentData(year:number, months:number) {
    try {
        const session = await getSession();
        const accessToken = session?.user.accessToken;
        const response = await fetch(`${baseUrl}/random-investment-data?years=${year}&months=${months}`, {
            method: "GET",
            cache: "reload",
            headers: {
                Authorization: `Bearer ${accessToken}`,
            },
            body: null,
        }).then((res)=>(res.json()));

        if (response.success){
            console.log(response.data);
            return response.data;
        }else{
            throw new Error("failed to get data");
        }
    } catch (error) {
      console.error("An error occurred on card:", error);
    }
}